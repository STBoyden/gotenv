// gotenv is a general-purpose package to load environment variables from files.
package gotenv

import (
	"bufio"
	"bytes"
	"io"
	"io/fs"
	"os"
	"strings"
)

// LoadOptions is a struct that contains options for the LoadEnv and
// LoadEnvFromFS functions.
type LoadOptions struct {
	OverrideExistingVars bool
	FileName             string
}

var defaultLoadEnvOptions = LoadOptions{
	OverrideExistingVars: true,
	FileName:             ".env",
}

// DefaultLoadOptions returns the default options for the LoadEnv and
// LoadEnvFromFS functions. The default options are:
//
//	OverrideExistingVars: true
//	FileName: ".env"
func DefaultLoadOptions() LoadOptions {
	return defaultLoadEnvOptions
}

// LoadEnv loads environment variables from a given file name specified in the
// given optional parameter opts. Default values are given by
// DefaultLoadOptions.
//
// This function returns a map of the environment variables that were added by
// the environment file, and an error if one occurred. Environment variables can
// also be accessed normally using os.Getenv.
func LoadEnv(opts ...LoadOptions) (map[string]string, error) {
	var overrideExistingVars bool = defaultLoadEnvOptions.OverrideExistingVars
	var fileName string = defaultLoadEnvOptions.FileName

	if opts != nil {
		overrideExistingVars = opts[0].OverrideExistingVars
		if opts[0].FileName != "" {
			fileName = opts[0].FileName
		}
	}

	envFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer envFile.Close()

	return parseEnvFile(envFile, overrideExistingVars)
}

// LoadEnvFromFS loads environment variables from a given file name specified in
// the given optional parameter opts, using the given file system fsys. Default
// values for opts are given by DefaultLoadOptions.
//
// This function returns a map of the environment variables that were added by
// the environment file, and an error if one occurred. Environment variables can
// also be accessed normally using os.Getenv.
func LoadEnvFromFS(fsys fs.FS, opts ...LoadOptions) (map[string]string, error) {
	var overrideExistingVars bool = defaultLoadEnvOptions.OverrideExistingVars
	var fileName string = defaultLoadEnvOptions.FileName

	if opts != nil {
		overrideExistingVars = opts[0].OverrideExistingVars
		if opts[0].FileName != "" {
			fileName = opts[0].FileName
		}
	}

	b, err := fs.ReadFile(fsys, fileName)
	if err != nil {
		return nil, err
	}

	return parseEnvFile(bytes.NewReader(b), overrideExistingVars)
}

// LoadEnvFromReaderOptions is a struct that contains options for the
// LoadEnvFromReader function.
type LoadEnvFromReaderOptions struct {
	OverrideExistingVars bool
}

var defaultLoadEnvFromReaderOptions = LoadEnvFromReaderOptions{
	OverrideExistingVars: true,
}

// DefaultLoadEnvFromReaderOptions returns the default options for the
// LoadEnvFromReader function. The default options are:
//
//	OverrideExistingVars: true
func DefaultLoadEnvFromReaderOptions() LoadEnvFromReaderOptions {
	return defaultLoadEnvFromReaderOptions
}

// LoadEnvFromReader loads environment variables from a given io.Reader. Default
// values for opts are given by DefaultLoadEnvFromReaderOptions.
//
// This function returns a map of the environment variables that were added by
// the environment file, and an error if one occurred. Environment variables can
// also be accessed normally using os.Getenv.
func LoadEnvFromReader(reader io.Reader, opts ...LoadEnvFromReaderOptions) (map[string]string, error) {
	var overrideExistingVars bool = defaultLoadEnvFromReaderOptions.OverrideExistingVars

	if opts != nil {
		overrideExistingVars = opts[0].OverrideExistingVars
	}

	return parseEnvFile(reader, overrideExistingVars)
}

func parseEnvFile(reader io.Reader, overrideExistingVars bool) (map[string]string, error) {
	envMap := make(map[string]string)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		// if line is empty or starts with a # then skip the line
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		firstEqualsIndex := strings.Index(line, "=")
		if firstEqualsIndex == -1 {
			continue
		}

		key := line[:firstEqualsIndex]
		value := line[firstEqualsIndex+1:]

		// we want to trim potential leading and tailing double quotes, as we assume
		// those are just used to denote a string environment variable that contains
		// characters that could potentially break parsing in other environments.

		value = strings.TrimPrefix(value, "\"")

		// if string ends with \" then we want to keep that as that could be there
		// for a reason.
		if !strings.HasSuffix(value, "\\\"") {
			value = strings.TrimSuffix(value, "\"")
		}

		envMap[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	for key, value := range envMap {
		if !overrideExistingVars {
			if _, present := os.LookupEnv(key); present {
				continue
			}
		}

		if err := os.Setenv(key, value); err != nil {
			return nil, err
		}
	}

	return envMap, nil
}

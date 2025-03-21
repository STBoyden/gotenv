package gotenv

import (
	"embed"
	"os"
	"testing"
)

//go:embed .env
var envFile embed.FS

func TestLoadEnv(t *testing.T) {
	_, err := LoadEnv(LoadOptions{OverrideExistingVars: false})
	if err != nil {
		t.Fatalf("an error occurred when loading environment variables: %s", err.Error())
	}

	if v, present := os.LookupEnv("SOME_ENV"); !present && v != "hello, world!" {
		var presentStr string
		if present {
			presentStr = "yes"
		} else {
			presentStr = "no"
		}

		t.Fatalf("output of LookupEnv was not as expected. Variable present?: %s. Value of environment variable: %s", presentStr, v)
	}
}

func TestLoadEnvFromFS(t *testing.T) {
	_, err := LoadEnvFromFS(envFile, LoadOptions{OverrideExistingVars: false})
	if err != nil {
		t.Fatalf("an error occurred when loading environment variables: %s", err.Error())
	}

	if v, present := os.LookupEnv("SOME_ENV"); !present && v != "hello, world!" {
		var presentStr string
		if present {
			presentStr = "yes"
		} else {
			presentStr = "no"
		}

		t.Fatalf("output of LookupEnv was not as expected. Variable present?: %s. Value of environment variable: %s", presentStr, v)
	}
}

func TestLoadEnvFromReader(t *testing.T) {
	file, _ := os.Open("exampleenv")

	_, err := LoadEnvFromReader(file, LoadEnvFromReaderOptions{OverrideExistingVars: false})
	if err != nil {
		t.Fatalf("an error occurred when loading environment variables: %s", err.Error())
	}

	if v, present := os.LookupEnv("SOME_ENV"); !present && v != "hello, world!" {
		var presentStr string
		if present {
			presentStr = "yes"
		} else {
			presentStr = "no"
		}

		t.Fatalf("output of LookupEnv was not as expected. Variable present?: %s. Value of environment variable: %s", presentStr, v)
	}
}

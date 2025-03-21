package tests

import (
	"embed"
	"os"
	"testing"

	"github.com/STBoyden/gotenv/v2"
)

//go:embed .env
var envFile embed.FS

func TestLoadEnvV2(t *testing.T) {
	_, err := gotenv.LoadEnv(gotenv.LoadOptions{OverrideExistingVars: false})
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

func TestLoadEnvFromFSV2(t *testing.T) {
	_, err := gotenv.LoadEnvFromFS(envFile, gotenv.LoadOptions{OverrideExistingVars: false})
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

func TestLoadEnvFromReaderV2(t *testing.T) {
	file, _ := os.Open("exampleenv")

	_, err := gotenv.LoadEnvFromReader(file, gotenv.LoadEnvFromReaderOptions{OverrideExistingVars: false})
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

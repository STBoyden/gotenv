package gotenv

import (
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	_, err := LoadEnv(false)
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

func TestLoadEnvFromFile(t *testing.T) {
	_, err := LoadEnvFromFile("exampleenv", false, false)
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

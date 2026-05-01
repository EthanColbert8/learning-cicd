package auth

import (
	"errors"
	"net/http"
	"testing"
)

type apiKeyTest struct {
	in       string
	expected string
}

func TestGetApiKey(t *testing.T) {
	testCases := []apiKeyTest{
		apiKeyTest{in: "ApiKey abcdefgh", expected: "abcdefgh"},
		apiKeyTest{in: "ApiKey asdf", expected: "asdf"},
		apiKeyTest{in: "ApiKey asdf abcdefgh", expected: "asdf"},
		apiKeyTest{in: "ApiKey abcdefgh asdf", expected: "abcdefgh"},
		apiKeyTest{in: "abcdefgh", expected: ""},
		apiKeyTest{in: "abcdefgh ApiKey", expected: ""},
	}

	headers := make(http.Header)
	headers.Set("Accept", "*/*")

	for i, testCase := range testCases {
		headers.Set("Authorization", testCase.in)

		result, err := GetAPIKey(headers)

		if result != testCase.expected {
			t.Errorf("test %d: expected %q, got %q", i, testCase.expected, result)
		}

		if testCase.expected == "" {
			if err == nil {
				t.Errorf("test %d: expected error, got none", i)
			}
		}
	}
}

func TestGetApiKeyNull(t *testing.T) {
	headers := make(http.Header)
	headers.Set("Accept", "*/*")

	result, err := GetAPIKey(headers)

	if result != "" {
		t.Errorf("non-empty result, got %s", result)
	}

	if err == nil {
		t.Error("expected error, got none")
	} else if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Error("malformed error")
	}
}

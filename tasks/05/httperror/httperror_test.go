package httperror

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestNewHTTPError(t *testing.T) {
	err := fmt.Errorf("some error happened: %w", NewHTTPError(http.StatusBadRequest))

	var httpErr httpError
	if !errors.As(err, &httpErr) {
		t.Errorf("Expected err to be httpError, got %T", err)
	}

	if code := GetStatusCode(err); code != http.StatusBadRequest {
		t.Errorf("Expected statusCode to be %d, got %d", http.StatusBadRequest, code)
	}

	if code := GetStatusCode(errors.New("non http error")); code != http.StatusInternalServerError {
		t.Errorf("Expected statusCode to be %d, got %d", http.StatusInternalServerError, code)
	}
}

func TestWrapHTTPError(t *testing.T) {
	originalErr := errors.New("original error")
	err := fmt.Errorf("some error happened: %w", WrapWithHTTPError(originalErr, http.StatusBadRequest))

	var httpErr httpError
	if !errors.As(err, &httpErr) {
		t.Errorf("Expected err to be httpError, got %T", err)
	}

	if code := GetStatusCode(err); code != http.StatusBadRequest {
		t.Errorf("Expected statusCode to be %d, got %d", http.StatusBadRequest, code)
	}

	if !errors.Is(err, originalErr) {
		t.Errorf("Expected underlying error to be 'original error'")
	}

	if code := GetStatusCode(errors.New("non http error")); code != http.StatusInternalServerError {
		t.Errorf("Expected statusCode to be %d, got %d", http.StatusInternalServerError, code)
	}
}

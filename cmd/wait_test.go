package cmd

import (
	"reflect"
	"testing"
)

// TestValidateHosts calls wait.validateHosts with hosts,
// checking for a valid return value.
func TestValidateHosts(t *testing.T) {
	// Arrange
	hosts := []string{"invalidHost", "http://validHost"}
	want := []string{"http://validHost"}

	// Act
	got := validateHosts(hosts)

	// Assert
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

// TestWaitTimeout calls wait.waitTimeout with invalid host and timeout,
// checking for timeout.
func TestWaitTimeout(t *testing.T) {
	// Arrange
	hosts := []string{"http://validHost"}
	timeout := 2
	want := false

	// Act
	got := waitTimeout(hosts, timeout)

	// Assert
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// TestExecuteEntrypoint calls wait.executeEntrypoint with an entrypoint
func TestExecuteEntrypoint(t *testing.T) {
	// Arrange
	entrypoint := "echo Hello World"
	want := true

	// Act + Assert
	got := executeEntrypoint(entrypoint)

	// Assert
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

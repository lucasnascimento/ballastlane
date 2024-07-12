package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestReadCmdParameters tests the ReadCmdParameters function for both success and failure scenarios.
func TestReadCmdParameters(t *testing.T) {
	// Test case for successful parameter reading
	successArgs := []string{"program", "param1"}
	expectedParam := "param1"
	param, err := ReadCmdParameters(successArgs)
	assert.NoError(t, err, "ReadCmdParameters() with args %v should not return an error", successArgs)
	assert.Equal(t, expectedParam, param, "ReadCmdParameters() = %v, want %v", param, expectedParam)

	// Test case for not enough arguments error
	failureArgs := []string{"program"}
	_, err = ReadCmdParameters(failureArgs)
	assert.Error(t, err, "ReadCmdParameters() with args %v was expected to return an error", failureArgs)
	assert.EqualError(t, err, ErrNotEnoughArguments.Error(), "ReadCmdParameters() returned an unexpected error: got %v, want %v", err, ErrNotEnoughArguments)
}

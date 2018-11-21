package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun(t *testing.T) {

	logrus.SetLevel(logrus.DebugLevel)

	// Use case command without arguments
	result, err := Run("ls")
	assert.NoError(t, err)
	assert.Equal(t, 0, result.ExitCode)
	assert.NotEqual(t, "", result.Stdout)

	// Use case command with argument
	result, err = Run("ls -al")
	assert.NoError(t, err)
	assert.Equal(t, 0, result.ExitCode)
	assert.NotEqual(t, "", result.Stdout)

	// Use case when command not exist
	result, err = Run("fake-cmd")
	assert.Error(t, err)
}

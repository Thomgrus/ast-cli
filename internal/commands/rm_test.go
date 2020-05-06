package commands

import (
	"testing"

	"gotest.tools/assert/cmp"

	"gotest.tools/assert"
)

func TestSastResourcesScansCommand(t *testing.T) {
	err := executeASTTestCommand("sast-resources", "scans", "--format", "pretty")
	assert.NilError(t, err)
	err = executeASTTestCommand("sr", "scans")
	assert.NilError(t, err)
	err = executeASTTestCommand("sr", "scans", "--format", "table")
	assert.NilError(t, err)
	err = executeASTTestCommand("sast-resources", "scans", "-h")
	assert.NilError(t, err)
}

func TestSastResourcesEnginesCommand(t *testing.T) {
	err := executeASTTestCommand("sast-resources", "engines", "--format", "table")
	assert.NilError(t, err)
	err = executeASTTestCommand("sr", "engines")
	assert.NilError(t, err)
	err = executeASTTestCommand("sr", "engines", "--format", "pretty")
	assert.NilError(t, err)
	err = executeASTTestCommand("sr", "engines", "-h")
	assert.NilError(t, err)
}

func TestSastResourcesStatsCommand(t *testing.T) {
	err := executeASTTestCommand("sr", "stats", "-v", "-r", "hour", "-m", "engine-running")
	assert.NilError(t, err)
	err = executeASTTestCommand("sr", "stats", "-v", "--resolution", "hour", "-m", "scan-running")
	assert.NilError(t, err)
	err = executeASTTestCommand("sr", "stats", "-v", "-r", "day", "--metric", "scan-pending")
	assert.NilError(t, err)
	err = executeASTTestCommand("sr", "stats", "-v", "--resolution", "week")
	assert.NilError(t, err)
	err = executeASTTestCommand("sr", "stats", "-v", "--metric", "engine-waiting")
	assert.NilError(t, err)
	err = executeASTTestCommand("sast-resources", "-v", "stats", "--format", "table")
	assert.NilError(t, err)
	err = executeASTTestCommand("sr", "stats", "-v", "--format", "pretty")
	assert.NilError(t, err)
	err = executeASTTestCommand("sr", "stats", "-h")
	assert.NilError(t, err)
	err = executeASTTestCommand("sr", "stats", "-r", "day", "-m", "lala")
	assert.Assert(t, cmp.Equal(err.Error(), "unknown metric lala"))

	err = executeASTTestCommand("sr", "stats", "-r", "sdfsd", "-m", "engine-waiting")
	assert.Assert(t, cmp.Equal(err.Error(), "unknown resolution sdfsd"))
}

func executeASTTestCommand(args ...string) error {
	cmd := createASTTestCommand()
	return executeTestCommand(cmd, args...)
}

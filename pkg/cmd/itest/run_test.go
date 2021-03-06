package cmd_test

import (
	"testing"
)

func TestAbortedTestShouldFailLocal(t *testing.T) {
	err := runSingle(t,
		&terminateOpts{
			runner:  "local:exec",
			builder: "exec:go",
		},
		"run",
		"single",
		"--builder",
		"exec:go",
		"--runner",
		"local:exec",
		"--instances",
		"1",
		"--plan",
		"placebo",
		"--testcase",
		"abort",
		"--wait",
	)

	if err == nil {
		t.Fatalf("should exit with an error")
	}
}

func TestAbortedTestShouldFailDocker(t *testing.T) {
	err := runSingle(t,
		&terminateOpts{
			runner:  "local:docker",
			builder: "docker:go",
		},
		"run",
		"single",
		"--builder",
		"docker:go",
		"--runner",
		"local:docker",
		"--instances",
		"1",
		"placebo:abort",
		"--wait",
	)

	if err == nil {
		t.Fatalf("should exit with an error")
	}
}

func TestIncompatibleRun(t *testing.T) {
	err := runSingle(t,
		&terminateOpts{
			runner:  "local:docker",
			builder: "exec:go",
		},
		"run",
		"single",
		"--builder",
		"exec:go",
		"--runner",
		"local:docker",
		"--instances",
		"1",
		"--plan",
		"placebo",
		"--testcase",
		"ok",
		"--wait",
	)

	if err == nil {
		t.Fatalf("should exit with an error")
	}
}

func TestCompatibleRunLocal(t *testing.T) {
	err := runSingle(t,
		&terminateOpts{
			runner:  "local:exec",
			builder: "exec:go",
		},
		"run",
		"single",
		"--builder",
		"exec:go",
		"--runner",
		"local:exec",
		"--instances",
		"1",
		"--plan",
		"placebo",
		"--testcase",
		"ok",
		"--wait",
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCompatibleRunDocker(t *testing.T) {
	err := runSingle(t,
		&terminateOpts{
			runner:  "local:docker",
			builder: "docker:go",
		},
		"run",
		"single",
		"--builder",
		"docker:go",
		"--runner",
		"local:docker",
		"--instances",
		"1",
		"--plan",
		"placebo",
		"--testcase",
		"ok",
		"--wait",
	)

	if err != nil {
		t.Fatal(err)
	}
}

// example:artifact *only* works with docker:generic
func TestDockerGenericCustomizesImage(t *testing.T) {
	err := runSingle(t,
		&terminateOpts{
			runner: "local:docker",
		},
		"run",
		"single",
		"--builder",
		"docker:generic",
		"--runner",
		"local:docker",
		"--instances",
		"1",
		"--plan",
		"example",
		"--testcase",
		"artifact",
		"--wait",
	)

	if err != nil {
		t.Fatal(err)
	}
}

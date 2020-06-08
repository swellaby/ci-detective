package functional

import (
	"os/exec"
	"runtime"
)

var expCIEngineEnvVars = [...]string{
	"CI",
	"CONTINUOUS_INTEGRATION",
	"BUILD_NUMBER",
	"APPVEYOR",
	"bamboo_planKey",
	"BITBUCKET_COMMIT",
	"BITRISE_IO",
	"BUDDY_WORKSPACE_ID",
	"BUILD_ID",
	"BUILDKITE",
	"BUILDKITE_PULL_REQUEST",
	"CIRCLECI",
	"CIRCLE_PULL_REQUEST",
	"CIRRUS_CI",
	"CODEBUILD",
	"CI_NAME",
	"DRONE",
	"DRONE_BUILD_EVENT",
	"DSARI",
	"GITHUB_ACTIONS",
	"GITHUB_EVENT_NAME",
	"GITLAB_CI",
	"GO_PIPELINE_LABEL",
	"HUDSON_URL",
	"JENKINS_URL",
	"MAGNUM",
	"pull_request",
	"SAILCI",
	"SEMAPHORE",
	"SHIPPABLE",
	"STRIDER",
	"TASK_ID",
	"TDDIUM",
	"TEAMCITY_VERSION",
	"TF_BUILD",
	"SYSTEM_TEAMFOUNDATIONCOLLECTIONURI",
	"TRAVIS",
	"TRAVIS_PULL_REQUEST",
}

func buildCommandWithEnv(envVar string) *exec.Cmd {
	var runnerArgs []string
	runner := ""
	if runtime.GOOS == "windows" {
		runner = "cmd.exe"
		script := "set " + envVar + "=\"\" && ci-detective"
		runnerArgs = []string{"/V", "/C", script}
	} else {
		runner = "sh"
		script := "env " + envVar + "=\"\" ci-detective"
		runnerArgs = []string{"-c", script}
	}
	return exec.Command(runner, runnerArgs...)
}

func buildCommandWithoutEnv() *exec.Cmd {
	runner := ""
	runnerArg := ""
	ciDetectiveCmd := "ci-detective"
	if runtime.GOOS == "windows" {
		runner = "cmd.exe"
		runnerArg = "/C"
	} else {
		runner = "sh"
		runnerArg = "-c"
	}
	cmd := exec.Command(runner, runnerArg, ciDetectiveCmd)
	cmd.Env = []string{"user=foo"}
	return cmd
}

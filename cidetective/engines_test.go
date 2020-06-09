package cidetective

import (
	"github.com/stretchr/testify/assert"
	"testing"
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

func TestEnginesCheckedInCorrectOrder(t *testing.T) {
	for i, actVar := range ciEngineEnvVars {
		assert.Equal(t, expCIEngineEnvVars[i], actVar)
	}
}

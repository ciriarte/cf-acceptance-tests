package helpers

import (
	"github.com/cloudfoundry-incubator/cf-test-helpers/commandstarter"
	"github.com/cloudfoundry-incubator/cf-test-helpers/helpers/internal"
	"github.com/onsi/gomega/gexec"
)

var SkipSSLValidation bool

func Curl(args ...string) *gexec.Session {
	cmdStarter := runner.NewCommandStarter()
	return helpersinternal.Curl(cmdStarter, args...)
}

func CurlSkipSSL(skip bool, args ...string) *gexec.Session {
	cmdStarter := runner.NewCommandStarter()
	return helpersinternal.CurlSkipSSL(cmdStarter, skip, args...)
}

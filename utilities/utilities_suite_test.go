package utilities_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	"github.com/jarcoal/httpmock"
	log "github.com/Sirupsen/logrus"
	"github.com/enkhalifapro/pgen/config"
	"github.com/spf13/viper"
)

func TestUtilities(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Utilities Suite")
}

var _ = BeforeSuite(func() {
	if err := config.Load("test", "../etc"); err != nil {
		panic(err.Error())
	}
	log.Infof("loaded config: %v", viper.ConfigFileUsed())
	// block all HTTP requests
	httpmock.Activate()
	// mock Get technologies API
	httpmock.RegisterResponder("GET", "https://api.similartech.com/v1/site/jet.com/technologies/all/pages?userkey=apiKey&format=json",
		httpmock.NewStringResponder(200, `{
    "site": "jet.com",
    "found": true,
    "technologies": [
        {
            "pages": [
                "/help-center/jetcares",
                "/jetcares"
            ],
            "coverage": 0.02,
            "id": 1779,
            "name": "Typeform",
            "categories": [
                "Online Forms"
            ],
            "paying": "maybe"
        }
    ]
}`))
})

var _ = AfterSuite(func() {
	httpmock.DeactivateAndReset()
})

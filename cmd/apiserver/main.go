


package main

import (
	// Make sure dep tools picks up these dependencies
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "github.com/go-openapi/loads"

	"sigs.k8s.io/apiserver-builder-alpha/pkg/cmd/server"
	_ "k8s.io/client-go/plugin/pkg/client/auth" // Enable cloud provider auth

	"gopath/src/pet-project/pkg/apis"
	"gopath/src/pet-project/pkg/openapi"
	_ "gopath/src/pet-project/plugin/admission/install"
)

func main() {
	version := "v0"

	err := server.StartApiServerWithOptions(&server.StartOptions{
		EtcdPath:         "/registry/shevchenkovitalii.com",
		Apis:             apis.GetAllApiBuilders(),
		Openapidefs:      openapi.GetOpenAPIDefinitions,
		Title:            "Api",
		Version:          version,

		// TweakConfigFuncs []func(apiServer *apiserver.Config) error
		// FlagConfigFuncs []func(*cobra.Command) error
	})
	if err != nil {
		panic(err)
	}
}

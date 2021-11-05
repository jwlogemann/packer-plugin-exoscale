package exoscaleimport

import (
	"context"
	"fmt"

	egoscale "github.com/exoscale/egoscale/v2"
	exoapi "github.com/exoscale/egoscale/v2/api"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
)

const BuilderId = "packer.post-processor.exoscale-import"

type Artifact struct {
	StateData map[string]interface{}

	postProcessor *PostProcessor
	state         *multistep.BasicStateBag
	template      *egoscale.Template
}

func (a *Artifact) BuilderId() string {
	return BuilderId
}

func (a *Artifact) Id() string {
	return *a.template.ID
}

func (a *Artifact) Files() []string {
	return nil
}

func (a *Artifact) String() string {
	return fmt.Sprintf(
		"%s @ %s (%s)",
		*a.template.Name,
		*a.template.Zone,
		*a.template.ID,
	)
}

func (a *Artifact) State(name string) interface{} {
	return a.StateData[name]
}

func (a *Artifact) Destroy() error {
	ctx := exoapi.WithEndpoint(
		context.Background(),
		exoapi.NewReqEndpoint(a.postProcessor.config.APIEnvironment, *a.template.Zone),
	)

	return a.postProcessor.exo.DeleteTemplate(ctx, *a.template.Zone, a.template)
}

func (a *Artifact) Template() *egoscale.Template {
	return a.template
}

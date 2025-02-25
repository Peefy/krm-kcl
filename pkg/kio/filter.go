package kio

import (
	"kcl-lang.io/krm-kcl/pkg/config"
	"kcl-lang.io/krm-kcl/pkg/edit"

	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

// Filter implements kio.Filter
type Filter struct {
	rw *kio.ByteReadWriter
}

// Filter checks each input and ensures that all containers have cpu and memory
// reservations set, otherwise it returns an error.
func (f Filter) Filter(in []*yaml.RNode) ([]*yaml.RNode, error) {
	c, err := f.parseConfig()
	if err != nil {
		return nil, err
	}
	st := &edit.SimpleTransformer{
		Name:           config.DefaultProgramName,
		Source:         c.Spec.Source,
		FunctionConfig: f.rw.FunctionConfig,
	}
	return st.Transform(in)
}

// parseConfig parses the functionConfig into an API struct.
func (f *Filter) parseConfig() (*config.KCLRun, error) {
	// Parse the input function config.
	var config config.KCLRun
	if err := yaml.Unmarshal([]byte(f.rw.FunctionConfig.MustString()), &config); err != nil {
		return nil, err
	}
	return &config, nil
}

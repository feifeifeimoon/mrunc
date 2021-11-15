package util

import (
	"encoding/json"
	"fmt"
	"github.com/opencontainers/runtime-spec/specs-go"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	SpecConfigFile = "config.json"
)

// ValidateSpec validates that the spec is compatible with mrunc.
func ValidateSpec(spec *specs.Spec) error {
	if spec.Process == nil {
		return fmt.Errorf("error spec.Process undefined, %+v", spec)
	}

	if spec.Root == nil || len(spec.Root.Path) == 0 {
		return fmt.Errorf("error spec.Root undefined, %+v", spec)
	}

	if spec.Solaris != nil || spec.Windows != nil {
		return fmt.Errorf("solaris or windows is not supported, %+v", spec)
	}

	return nil
}

// ReadSpecFromBundle read config.json from bundle dir
func ReadSpecFromBundle(bundleDir string) (*specs.Spec, error) {
	specFile, err := os.Open(filepath.Join(bundleDir, SpecConfigFile))
	if err != nil {
		return nil, fmt.Errorf("error opening spec file %q: %v", filepath.Join(bundleDir, "config.json"), err)
	}
	defer specFile.Close()

	specBytes, err := ioutil.ReadAll(specFile)
	if err != nil {
		return nil, fmt.Errorf("error reading spec from file %q: %v", specFile.Name(), err)
	}

	var spec specs.Spec
	if err := json.Unmarshal(specBytes, &spec); err != nil {
		return nil, fmt.Errorf("error unmarshaling spec from file %q: %v\n %s", specFile.Name(), err, string(specBytes))
	}

	return &spec, nil
}

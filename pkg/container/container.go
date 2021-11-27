package container

import (
	"github.com/feifeifeimoon/mrunc/pkg/process"
	"github.com/feifeifeimoon/mrunc/pkg/util"
	"github.com/opencontainers/runtime-spec/specs-go"
	log "github.com/sirupsen/logrus"
)

type Container struct {
	ID string

	Spec *specs.Spec
}

func NewContainer(id, bundleDir string) (*Container, error) {
	log.Infof("new container [%s] [%s]", id, bundleDir)

	var container Container

	spec, err := util.ReadSpecFromBundle(bundleDir)
	if err != nil {
		log.Errorf("read spec err, %v", err)
		return nil, err
	}

	container.Spec = spec

	initProcess := process.NewInitProcess()
	initProcess.Create(spec)

	return &container, nil
}

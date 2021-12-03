package container

import (
	"encoding/json"
	"github.com/feifeifeimoon/mrunc/pkg/process"
	"github.com/feifeifeimoon/mrunc/pkg/util"
	"github.com/opencontainers/runtime-spec/specs-go"
	log "github.com/sirupsen/logrus"
	"os"
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

	pr, pw, err := os.Pipe()
	if err != nil {
		log.Errorf("create pipe err, %v", err)
		return nil, err
	}

	initProcess := process.NewInitProcess()
	initProcess.Create(spec, pr)

	// 将spc通过pipe发送给init进程
	if err := json.NewEncoder(pw).Encode(spec); err != nil {
		log.Errorf("write spec to init pipe err, %v", err)
		return nil, err
	}

	// 记录init进程的Pid

	// 为init进程设置cgroup

	return &container, nil
}

//+build linux

package process

import (
	"context"
	"github.com/opencontainers/runtime-spec/specs-go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	"os"
	"os/exec"
	"syscall"
)

type InitProcess struct {
	InitPath string
	InitArgs []string
}

func NewInitProcess() *InitProcess {

	return &InitProcess{
		InitPath: "/proc/self/exe",
		InitArgs: []string{"Init"},
	}
}

func (i *InitProcess) Create(spec *specs.Spec) {
	cmd := exec.CommandContext(context.Background(), i.InitPath, i.InitArgs...)

	//case ns to unix flag
	flag := uintptr(0)
	for _, ns := range spec.Linux.Namespaces {
		switch ns.Type {
		case specs.PIDNamespace:
			flag |= unix.CLONE_NEWPID
		case specs.NetworkNamespace:
			flag |= unix.CLONE_NEWNET
		case specs.MountNamespace:
			flag |= unix.CLONE_NEWNS
		case specs.IPCNamespace:
			flag |= unix.CLONE_NEWIPC
		case specs.UTSNamespace:
			flag |= unix.CLONE_NEWUTS
		case specs.UserNamespace:
			flag |= unix.CLONE_NEWUSER
		case specs.CgroupNamespace:
			flag |= unix.CLONE_NEWCGROUP
		default:
			log.Warning("no supported namespaces ", ns.Type)
		}

	}

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: flag,
	}

	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

	if err := cmd.Start(); err != nil {
		log.Errorf("start process err %v", err)
		return
	}

	if err := cmd.Wait(); err != nil {
		log.Errorf("wait process err %v", err)
		return
	}

	//defer func() {
	//	if err != nil {
	//		_ = cmd.Process.Kill()
	//	}
	//}()

}

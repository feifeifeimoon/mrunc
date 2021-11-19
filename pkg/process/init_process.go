package process

import (
	"context"
	"github.com/opencontainers/runtime-spec/specs-go"
	"os"
	"os/exec"
)

type InitProcess struct {
	InitPath string
	InitArgs []string
}

func NewInitProcess() *InitProcess {

	return &InitProcess{
		InitPath: "/proc/self/exec",
		InitArgs: []string{os.Args[0], "Init"},
	}
}

func (i *InitProcess) Create(spec *specs.Spec) {
	cmd := exec.CommandContext(context.Background(), i.InitPath, i.InitArgs...)

	// case ns to unix flag
	//flag := uintptr(0)
	//for _, ns := range spec.Linux.Namespaces {
	//	switch ns.Type {
	//	case specs.PIDNamespace:
	//		flag |= unix.CLONE_NEWPID
	//	case specs.NetworkNamespace:
	//		flag |= unix.CLONE_NEWNET
	//	case specs.MountNamespace:
	//		flag |= unix.CLONE_NEWNS
	//	case specs.IPCNamespace:
	//		flag |= unix.CLONE_NEWIPC
	//	case specs.UTSNamespace:
	//		flag |= unix.CLONE_NEWUTS
	//	case specs.UserNamespace:
	//		flag |= unix.CLONE_NEWUSER
	//	case specs.CgroupNamespace:
	//		flag |= unix.CLONE_NEWCGROUP
	//	default:
	//		log.Warning("no supported namespaces ", ns.Type)
	//	}
	//
	//}

	//cmd.SysProcAttr = &syscall.SysProcAttr{
	//	Cloneflags: flag,
	//}

	cmd.Run()

}

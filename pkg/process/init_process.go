//+build linux

package process

import (
	"context"
	"encoding/json"
	"github.com/opencontainers/runtime-spec/specs-go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	"os"
	"os/exec"
	"strconv"
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

func (i *InitProcess) Create(spec *specs.Spec, pw *os.File) {
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

	cmd.ExtraFiles = append(cmd.ExtraFiles, pw)
	// 3代表stdin、stdout、stderr占用了三个
	cmd.Env = append(cmd.Env, "INIT_PIPE_PD="+strconv.Itoa(3+len(cmd.ExtraFiles)-1))

	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

	if err := cmd.Start(); err != nil {
		log.Errorf("start process err %v", err)
		return
	}

	//if err := cmd.Wait(); err != nil {
	//	log.Errorf("wait process err %v", err)
	//	return
	//}

	//defer func() {
	//	if err != nil {
	//		_ = cmd.Process.Kill()
	//	}
	//}()

}

func (i *InitProcess) StartInitialization() error {
	pipeFd, err := strconv.Atoi(os.Getenv("INIT_PIPE_PD"))
	if err != nil {
		//panic(err)
		return err
	}
	log.Info("get init pipe fd ", pipeFd)

	pipe := os.NewFile(uintptr(pipeFd), "pipe")
	defer pipe.Close()

	var spec *specs.Spec
	if err := json.NewDecoder(pipe).Decode(&spec); err != nil {
		log.Errorf("read spec from init pipe err, %v", err)
		return err
	}

	log.Debugf("read spec %+v", spec)

	// no return
	return nil

}

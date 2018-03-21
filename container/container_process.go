package container

import (
	"syscall"
	"os/exec"
	"os"
	"github.com/gpmgo/gopm/modules/log"
)

func NewParentProcess(tty bool, command string) *exec.Cmd {
	log.Info("-----------------------NewParentProcess")
	args := []string{"init", command}
	log.Info("newparentprocess args=%s",args)

	cmd := exec.Command("/proc/self/exe", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}
	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd
}
package container

import (
	"os"
	"syscall"
	"github.com/Sirupsen/logrus"
)

func RunContainerInitProcess(command string, args []string) error {
	logrus.Infof("RunContainerInitProcess command %s", command)

        syscall.Sethostname([]byte("newhost"))

        syscall.Chroot("/root/go/src/hans/test/busybox")
        os.Chdir("/")
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")

	argv := []string{command}
	logrus.Infof("RunContainerInitProcess command=%s argv=%s",command,argv)
	if err := syscall.Exec(command, argv, os.Environ()); err != nil {
		logrus.Errorf(err.Error())
	}
	return nil
}

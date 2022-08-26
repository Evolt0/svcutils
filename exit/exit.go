package exit

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func SetupSignalHandler() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	select {
	case sig := <-sc:
		logrus.Infof("Got signal [%s] to exit.", sig)

		os.Exit(0)
	}
}

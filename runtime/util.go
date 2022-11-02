package runtime

import (
	"os"
	"os/signal"
	"syscall"
)

func ListenForSignals() os.Signal {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	sig := <-signals
	return sig
}

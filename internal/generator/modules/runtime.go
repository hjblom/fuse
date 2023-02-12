package modules

import (
	"fmt"

	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util"
)

var RuntimeGenerator = &runtimeGenerator{file: util.File}

type runtimeGenerator struct {
	file util.FileReadWriter
}

func (g *runtimeGenerator) Name() string {
	return "Main Generator"
}

func (g *runtimeGenerator) Description() string {
	return "Generate the main.go file."
}

func (g *runtimeGenerator) Plugins() map[string]string {
	return map[string]string{}
}

func (g *runtimeGenerator) Generate(mod *config.Module) error {
	path := fmt.Sprintf("internal/%s", "runtime.go")
	if g.file.Exists(path) {
		return nil
	}

	err := g.file.WriteFile(path, []byte(runtimeSrc))
	if err != nil {
		return fmt.Errorf("failed to write file %s: %w", path, err)
	}

	return nil
}

var runtimeSrc = `package internal

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

var logger = logrus.WithField("ctx", "runtime")

type Service interface {
	Start() error
	Stop() error
}

type Runtime struct {
	services []Service
}

func NewRuntime(services []Service) *Runtime {
	return &Runtime{
		services: services,
	}
}

func (r *Runtime) Start() error {
	logger.Info("starting services")
	wg := sync.WaitGroup{}
	var err error

	// Setup channels
	startErr := make(chan error, 1)
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)

	// Start app
	wg.Add(1)
	go func(startErr chan error) {
		defer wg.Done()
		err := r.start()
		if err != nil {
			startErr <- err
		}
	}(startErr)

	// Wait for service failure or kill signal
	select {
	case err = <-startErr:
		logger.Error("failed to start services", err)
		break
	case sig := <-osSignal:
		logger.Info(fmt.Sprintf("%s signal received", sig.String()))
		break
	}

	// Attempt to stop services gracefully
	logger.Info("attempting to stop services gracefully")
	var done = make(chan bool, 1)
	go func() {
		r.stop()
		wg.Wait()
		done <- true
	}()

	// Wait for app to stop or kill after graceful shutdown timeout
	select {
	case <-done:
		logger.Info("stopped all services")
		break
	case <-time.After(5 * time.Second):
		logger.Info("app stopped forcefully")
		break
	}

	return err
}

func (r *Runtime) start() error {
	wg := sync.WaitGroup{}
	var e = make(chan error, 1)
	var done = make(chan bool, 1)

	// Start all services
	for _, service := range r.services {
		wg.Add(1)
		go func(svc Service, e chan error) {
			defer wg.Done()
			err := svc.Start()
			if err != nil {
				e <- err
			}
		}(service, e)
	}

	// Wait for WaitGroup to finish
	go func() {
		wg.Wait()
		done <- true
	}()

		select {
		// If all services are done, return without error
		case <-done:
			return nil
		// If any service returns an error, stop all services and return error
		case err := <-e:
			r.stop()
			return err
		}
	}

	func (r *Runtime) stop() error {
		for _, service := range r.services {
			err := service.Stop()
			if err != nil {
				return err
			}
		}
		return nil
	}
`

package runtime

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Runtime struct {
	services []Service

	logger Logger
}

func NewRuntime(services []Service, logger Logger) *Runtime {
	return &Runtime{
		services: services,
		logger:   logger,
	}
}

func (r *Runtime) Start() error {
	r.logger.Info("starting services")
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
		r.logger.Error("failed to start services", err)
		break
	case sig := <-osSignal:
		r.logger.Info(fmt.Sprintf("%s signal received", sig.String()))
		break
	}

	// Attempt to stop services gracefully
	r.logger.Info("attempting to stop services gracefully")
	var done = make(chan bool, 1)
	go func() {
		r.stop()
		wg.Wait()
		done <- true
	}()

	// Wait for app to stop or kill after graceful shutdown timeout
	select {
	case <-done:
		r.logger.Info("stopped all services")
		break
	case <-time.After(5 * time.Second):
		r.logger.Info("app stopped forcefully")
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

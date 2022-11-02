package runtime

import (
	"fmt"
	"os"
	"sync"
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
	r.logger.Info("starting")
	wg := sync.WaitGroup{}

	// Start app
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := r.start()
		if err != nil {
			panic(err)
		}
	}()

	// Wait for signals
	signal := ListenForSignals()
	r.logger.Info(fmt.Sprintf("%s signal received", signal.String()))

	// Attempt to stop app gracefully
	r.logger.Info("stopping app gracefully")
	var done = make(chan bool, 1)
	go func() {
		r.stop()
		wg.Wait()
		done <- true
	}()

	// Wait for app to stop or kill after graceful shutdown timeout
	select {
	case <-done:
		r.logger.Info("app stopped gracefully")
		os.Exit(0)
	case <-time.After(5 * time.Second):
		r.logger.Info("app stopped forcefully")
		os.Exit(1)
	}

	return nil
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

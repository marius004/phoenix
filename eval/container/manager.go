package container

import (
	"context"
	"log"

	"github.com/marius004/phoenix/eval"
	"github.com/marius004/phoenix/models"
	"golang.org/x/sync/semaphore"
)

// Manager implements eval.SandboxManager
type Manager struct {
	logger    *log.Logger
	semaphore *semaphore.Weighted
	config    *models.Config

	maxConcurrentSandboxes int64
	availableSandboxes     chan int
}

func (m *Manager) RunTask(ctx context.Context, task eval.Task) error {
	sandbox, err := m.getSandbox()

	if err != nil {
		m.logger.Println("Could not create sandbox", err)
		return err
	}

	// defer sandbox.Cleanup()
	defer m.ReleaseSandbox(sandbox.GetID())

	return task.Run(ctx, sandbox)
}

func (m *Manager) ReleaseSandbox(id int) {
	m.semaphore.Release(1)
	m.availableSandboxes <- id
}

func (m *Manager) Stop(ctx context.Context) error {
	if err := m.semaphore.Acquire(ctx, m.maxConcurrentSandboxes); err != nil {
		return err
	}
	return nil
}

func (m *Manager) getSandbox() (eval.Sandbox, error) {
	if err := m.semaphore.Acquire(context.Background(), 1); err != nil {
		return nil, err
	}
	return m.newSandbox(<-m.availableSandboxes)
}

func (m *Manager) newSandbox(id int) (*Container, error) {
	return newContainer(id, m.config, m.logger)
}

func NewManager(maxConcurrentSandboxes int64, config *models.Config, logger *log.Logger) eval.SandboxManager {
	manager := &Manager{
		logger: logger,
		config: config,

		semaphore:              semaphore.NewWeighted(maxConcurrentSandboxes),
		availableSandboxes:     make(chan int, maxConcurrentSandboxes),
		maxConcurrentSandboxes: maxConcurrentSandboxes,
	}

	// filling the channel with ids for sandboxes
	for i := 1; i <= int(maxConcurrentSandboxes); i++ {
		manager.availableSandboxes <- i
	}

	return manager
}

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

	defer m.ReleaseSandbox(sandbox)

	return task.Run(ctx, sandbox)
}

func (m *Manager) ReleaseSandbox(sandbox eval.Sandbox) {
	//m.semaphore.Release(1)
	sandbox.Cleanup()
	m.availableSandboxes <- sandbox.GetID()
}

func (m *Manager) Stop(ctx context.Context) error {
	// if err := m.semaphore.Acquire(ctx, m.maxConcurrentSandboxes); err != nil {
	// 	return err
	// }

	close(m.availableSandboxes)
	return nil
}

func (m *Manager) getSandbox() (eval.Sandbox, error) {
	// if err := m.semaphore.Acquire(context.Background(), 1); err != nil {
	// 	return nil, err
	// }
	return m.newSandbox(<-m.availableSandboxes)
}

func (m *Manager) newSandbox(id int) (*Container, error) {
	return newContainer(id, m.config, m.logger)
}

func NewManager(config *models.Config, logger *log.Logger) eval.SandboxManager {
	manager := &Manager{
		logger: logger,
		config: config,

		semaphore:              semaphore.NewWeighted(int64(config.MaxSandboxes)),
		availableSandboxes:     make(chan int, config.MaxSandboxes),
		maxConcurrentSandboxes: int64(config.MaxSandboxes),
	}

	// filling the channel with ids for sandboxes
	for i := 1; i <= int(config.MaxSandboxes); i++ {
		manager.availableSandboxes <- i
	}

	return manager
}

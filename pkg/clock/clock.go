package clock

import (
	"fmt"
	"os"
	"time"

	"github.com/lucasnascimento/ballastlane/pkg/config"
	"github.com/lucasnascimento/ballastlane/pkg/repository"
)

// Clock struct holds the configuration path and the actual configuration data.
type Clock struct {
	ConfigManager     *config.ConfigManager
	TimeProvider      TimeProvider
	ExitProvider      ExitProvider
	SignalsRepository repository.ISignalsRepository
}

// ExitProvider is an interface that defines the method for exiting the program.
type ExitProvider interface {
	Exit(code int)
}

// RealExitProvider is an implementation of ExitProvider that exits the program.
type RealExitProvider struct{}

// Exit exits the program with the specified code.
func (r RealExitProvider) Exit(code int) {
	os.Exit(code)
}

// NewClock initializes a new Clock instance with the given configuration path.
// It loads the configuration from the specified path and returns a Clock instance or an error.
func NewClock(cm *config.ConfigManager, timeProvider TimeProvider, exitProvider ExitProvider, signalsRepository repository.ISignalsRepository) (*Clock, error) {
	return &Clock{
		ConfigManager:     cm,
		TimeProvider:      timeProvider,
		ExitProvider:      exitProvider,
		SignalsRepository: signalsRepository,
	}, nil
}

// Run starts the clock operation, checking for configuration changes and managing time-based actions.
func (c *Clock) Run() error {
	ticker := time.NewTicker(1 * time.Second) // Checks every 1 second
	defer ticker.Stop()

	startTime := c.TimeProvider.Now()
	secondsCount := 0
	for range ticker.C {

		now := startTime.Add(time.Duration(secondsCount) * time.Second)
		secondsCount++
		signal := ""
		if now.Second() == 0 {
			if now.Minute() == 0 {
				signal = c.ConfigManager.Config.Bong // Each hour
			} else {
				signal = c.ConfigManager.Config.Tock // Each minute
			}
		} else {
			signal = c.ConfigManager.Config.Tick // Each second
		}

		fmt.Println(signal) // Each second
		err := c.SignalsRepository.SaveSignal(signal, now)
		if err != nil {
			return err
		}

		// Check if the program has been running for 3 hours
		if secondsCount >= c.ConfigManager.Config.RunSpan {
			c.ExitProvider.Exit(0) // need this in order to test the exit
			break
		}
	}

	return nil
}

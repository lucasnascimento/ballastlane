package clock

import (
	"bytes"
	"log"
	"os"
	"testing"
	"time"

	"github.com/lucasnascimento/ballastlane/pkg/config"
	"github.com/lucasnascimento/ballastlane/pkg/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewClock(t *testing.T) {
	configPath := "testdata/config_test.json"
	cm, err := config.NewConfigManager(configPath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	mockRepo := repository.MockSignalStore{}
	clock, err := NewClock(cm, RealTimeProvider{}, RealExitProvider{}, &mockRepo)

	assert.NoError(t, err)

	assert.Equal(t, "tick", clock.ConfigManager.Config.Tick, "Wrong configuration for Tick")
	assert.Equal(t, "tock", clock.ConfigManager.Config.Tock, "Wrong configuration for Tock")
	assert.Equal(t, "bong", clock.ConfigManager.Config.Bong, "Wrong configuration for Bong")
}

func TestClock_Run(t *testing.T) {
	configPath := "testdata/config_test.json"
	cm, err := config.NewConfigManager(configPath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	t.Run("OnMinuteShouldPrintTickTockTick", func(t *testing.T) {
		mockRepo := repository.MockSignalStore{}
		mockRepo.On("SaveSignal", mock.AnythingOfType("string"), mock.AnythingOfType("time.Time")).Return(nil)
		clock, err := NewClock(
			cm, MockTimeProvider{
				FixedTime: time.Date(2021, 1, 1, 23, 58, 59, 0, time.UTC),
			},
			MockedExitProvider{},
			&mockRepo,
		)
		assert.NoError(t, err)

		// Capture output
		oldStdout := os.Stdout // keep backup of the real stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Act
		clock.Run()

		// Now read the output
		w.Close()
		os.Stdout = oldStdout // restoring the real stdout
		var buf bytes.Buffer
		_, err = buf.ReadFrom(r)
		assert.NoError(t, err)

		assert.Equal(t, "tick\ntock\ntick\n", buf.String())
		mockRepo.AssertExpectations(t)
		mockRepo.AssertNumberOfCalls(t, "SaveSignal", 3)
	})

	t.Run("OnHoutShouldPrintTickBongTick", func(t *testing.T) {
		mockRepo := repository.MockSignalStore{}
		mockRepo.On("SaveSignal", mock.AnythingOfType("string"), mock.AnythingOfType("time.Time")).Return(nil)
		clock, err := NewClock(
			cm, MockTimeProvider{
				FixedTime: time.Date(2021, 1, 1, 23, 59, 59, 0, time.UTC),
			},
			MockedExitProvider{},
			&mockRepo,
		)
		assert.NoError(t, err)

		// Capture output
		oldStdout := os.Stdout // keep backup of the real stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Act
		clock.Run()

		// Now read the output
		w.Close()
		os.Stdout = oldStdout // restoring the real stdout
		var buf bytes.Buffer
		_, err = buf.ReadFrom(r)
		assert.NoError(t, err)

		assert.Equal(t, "tick\nbong\ntick\n", buf.String())
		mockRepo.AssertExpectations(t)
		mockRepo.AssertNumberOfCalls(t, "SaveSignal", 3)
	})

}

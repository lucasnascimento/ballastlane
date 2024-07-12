package config

import (
	"encoding/json"
	"net/http"
	"os"
	"sync"
)

// Config struct holds the configuration data.
type Config struct {
	RunSpan int    `json:"run_span"`
	Tick    string `json:"tick"`
	Tock    string `json:"tock"`
	Bong    string `json:"bong"`
}

// ConfigManager struct holds the configuration path and the actual configuration data.
type ConfigManager struct {
	Config *Config
	Path   string
	mu     sync.Mutex // Grant exclusive access to the Config field
}

// NewConfigManager ...
func NewConfigManager(path string) (*ConfigManager, error) {
	cm := &ConfigManager{Path: path}
	if err := cm.LoadConfig(); err != nil {
		return nil, err
	}
	return cm, nil
}

// LoadConfig loads the configuration from the file specified in the ConfigManager's Path field.
func (cm *ConfigManager) LoadConfig() error {
	var config Config
	bytes, err := os.ReadFile(cm.Path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, &config); err != nil {
		return err
	}
	cm.Config = &config
	return nil
}

// UpdateConfigHandler manages the update of the configuration.
func (cm *ConfigManager) UpdateConfigHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	var newConfig Config
	if err := json.NewDecoder(r.Body).Decode(&newConfig); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cm.mu.Lock()
	cm.Config = &newConfig
	cm.mu.Unlock()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cm.Config)
}

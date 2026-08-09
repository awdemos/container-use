package agent

import (
	"encoding/json"
	"path/filepath"

	"github.com/dagger/container-use/rules"
)

type ConfigureCursor struct {
	Name        string
	Description string
}

func NewConfigureCursor() *ConfigureCursor {
	return &ConfigureCursor{
		Name:        "Cursor",
		Description: "AI-powered code editor",
	}
}

func (a *ConfigureCursor) name() string {
	return a.Name
}

func (a *ConfigureCursor) description() string {
	return a.Description
}

func (a *ConfigureCursor) editMcpConfig() error {
	return writeMcpConfig(filepath.Join(".cursor", "mcp.json"), func(data []byte, cfg *MCPServersConfig) error {
		return json.Unmarshal(data, cfg)
	}, a.updateMcpConfig)
}

func (a *ConfigureCursor) updateMcpConfig(config MCPServersConfig) ([]byte, error) {
	// Initialize mcpServers map if nil
	if config.MCPServers == nil {
		config.MCPServers = make(map[string]MCPServer)
	}

	config.MCPServers["container-use"] = MCPServer{
		Command: ContainerUseBinary,
		Args:    []string{"stdio"},
	}

	return json.MarshalIndent(config, "", "  ")
}

func (a *ConfigureCursor) editRules() error {
	rulesFile := filepath.Join(".cursor", "rules", "container-use.mdc")
	return saveRulesFile(rulesFile, rules.CursorRules)
}

func (a *ConfigureCursor) isInstalled() bool {
	return true
}

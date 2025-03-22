package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/gofiber/fiber/v2/log"
	mcpclient "github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
)

type MCPConfig struct {
	MCPServers map[string]struct {
		Command string   `json:"command"`
		Args    []string `json:"command"`
	} `json:"mcpServers"`
}

func mcpToolsToAnthropicTools(serverName string, mcpTools []mcp.Tool) []anthropic.ToolParam {
	anthropicTools := make([]anthropic.ToolParam, len(mcpTools))
	for i, tool := range mcpTools {
		namespacedName := fmt.Sprintf("%s__%s", serverName, tool)
		schemaMap := map[string]interface{}{
			"type":       tool.InputSchema.Type,
			"properties": tool.InputSchema.Properties,
		}
		if len(tool.InputSchema.Required) > 0 {
			schemaMap["required"] = tool.InputSchema.Required
		}
		anthropicTools[i] = anthropic.ToolParam{
			Name:        anthropic.F(namespacedName),
			Description: anthropic.F(tool.Description),
			InputSchema: anthropic.Raw[interface{}](schemaMap),
		}

	}
	return anthropicTools
}

func loadMCPConfig() (*MCPConfig, error) {
	var configPath string
	if configFile != "" {
		configPath = configFile
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("error getting home directory: %w", err)
		}
		configPath = filepath.Join(homeDir, "mcp.json")
	}

	configData, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file %s: %w", configPath, err)
	}

	var config MCPConfig
	if err := json.Unmarshal(configData, &config); err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	return &config, nil
}

func creaeteMCPClients(config *MCPConfig) (map[string]*mcpclient.StdioMCPClient, error) {
	clients := make(map[string]*mcpclient.StdioMCPClient)
	for name, server := range config.MCPServers {
		client, err := mcpclient.NewStdioMCPClient(server.Command, server.Args...)
		if err != nil {
			for _, c := range clients {
				c.Close()
			}
			return nil, fmt.Errorf("failed to create MCP client for %s: %w", name, err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
        log.Info()
	}
}

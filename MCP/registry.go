package main

import (
	"github.com/input-api/mcp-server/config"
	"github.com/input-api/mcp-server/models"
	tools_api "github.com/input-api/mcp-server/tools/api"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_api.CreatePost_api_game_responseTool(cfg),
		tools_api.CreateGet_api_high_scoresTool(cfg),
		tools_api.CreatePost_api_high_scoresTool(cfg),
	}
}

package handler

import (
	"log"
	"os"
	"sandhu-sahil/bot/variables"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	variables.PREFIX = os.Getenv("PREFIX")
	variables.ServiceUrl = os.Getenv("SERVICE_URL")
	variables.OwnerId = os.Getenv("OWNER_ID")
	variables.UseSharding = os.Getenv("USE_SHARDING") == "true"

	variables.ShardId = 0
	variables.ShardCount = 1

	variables.GuildID = os.Getenv("GUILD_ID")
	variables.DefaultStatus = os.Getenv("DEFAULT_STATUS")
	variables.RemoveCommands = os.Getenv("REMOVE_COMMANDS") == "true"
}

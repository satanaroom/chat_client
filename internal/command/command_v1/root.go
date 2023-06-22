package command_v1

import (
	"log"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "chat-client",
	Short: "Клиент для многопользовательского чат-сервера",
}

func Execute() {
	if err := root.Execute(); err != nil {
		log.Fatalf("execute root: %s", err.Error())
	}
}

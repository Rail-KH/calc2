package main

import (
	"log"

	"github.com/Rail-KH/calc2/internal/agent"
)

func main() {
	agent := agent.NewAgent()
	log.Println("Starting Agent...")
	agent.Run()
}

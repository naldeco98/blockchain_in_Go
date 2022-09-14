package main

import (
	"log"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nialdeco98/blockchain_in_Go/cmd/server/handlers"
	"github.com/nialdeco98/blockchain_in_Go/internal/blockchain"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Init blockchain
	go func() {
		t := time.Now()
		genesisBlock := blockchain.Block{Timestamp: t.String()}
		spew.Dump(genesisBlock)
		blockchain.Blockchain = append(blockchain.Blockchain, genesisBlock)
	}()

	engine := gin.Default()

	engine.GET("", handlers.GetBlockchain())
	engine.POST("", handlers.WriteBlock())

	if err := engine.Run(os.Getenv("ADDR")); err != nil {
		log.Fatal(err)
	}

}

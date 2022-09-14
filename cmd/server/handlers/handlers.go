package handlers

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/nialdeco98/blockchain_in_Go/internal/blockchain"
)

type Request struct {
	BPM int
}

func GetBlockchain() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, blockchain.Blockchain)

	}
}

func WriteBlock() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req Request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			log.Printf("%+v", err)
			ctx.JSON(400, gin.H{"error": err})
			return
		}

		newBlock, err := blockchain.GenerateBlock(blockchain.Blockchain[len(blockchain.Blockchain)-1], req.BPM)
		if err != nil {
			log.Printf("%+v", err)
			ctx.JSON(500, gin.H{"error": err})
			return
		}
		if !blockchain.IsBlockValid(newBlock, blockchain.Blockchain[len(blockchain.Blockchain)-1]) {
			log.Printf("%+v", err)
			ctx.JSON(500, gin.H{"error": err})
			return
		}
		newBlockchain := append(blockchain.Blockchain, newBlock)
		blockchain.ReplaceChain(newBlockchain)
		spew.Dump(blockchain.Blockchain)

		ctx.JSON(200, gin.H{"new": newBlock})
	}
}

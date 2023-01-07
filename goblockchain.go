package main

import (
	"fmt"
	"log"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

/**
* nonce int ナンス
* previousHash string
* return block
**/

func NewBlock(nonce int, previousHash string) *Block {

	block := new(Block)
	block.timestamp = time.Now().UnixNano()
	block.nonce = nonce
	block.previousHash = previousHash
	return block

	/*
		return &Block{
			timestamp: time.Now().UnixMicro(),
		}
	*/
}

func init() {
	log.SetPrefix("BlockChain: ")
}

func main() {
	block := NewBlock(0, "init Hash")
	fmt.Println(block)
}

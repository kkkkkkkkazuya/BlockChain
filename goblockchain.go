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

	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b

	/*
		return &Block{
			timestamp: time.Now().UnixMicro(),
		}
	*/
}

func init() {
	log.SetPrefix("BlockChain: ")
}

/**
* 見やすいコメントアウト
**/
func (b *Block) Print() {
	fmt.Printf("timestamp    %d\n", b.timestamp)
	fmt.Printf("nonce        %d\n", b.nonce)
	fmt.Printf("previousHash %s\n", b.previousHash)
	fmt.Printf("transactions %s\n", b.transactions)
}

func main() {
	b := NewBlock(0, "init Hash")
	b.Print()
}

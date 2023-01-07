package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

/**
* １ブロックの構造体
**/
type Block struct {
	// ナンス
	nonce int

	// ハッシュ
	previousHash string

	// タイムスタンプ
	timestamp int64

	// トランザクション
	transactions []string
}

/**
* ブロックの初期値
* @para nonce: ナンス
* @para previousHash: ハッシュ値
* @return ブロック情報
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

/**
* ブロックのサーバーログ用
**/
func (b *Block) Print() {
	fmt.Printf("timestamp    %d\n", b.timestamp)
	fmt.Printf("nonce        %d\n", b.nonce)
	fmt.Printf("previousHash %s\n", b.previousHash)
	fmt.Printf("transactions %s\n", b.transactions)
}

/**
* ブロックチェーンの構造体
**/
type BlockChain struct {
	// トランザクションプール（トランザクション情報を詰めていく）
	transactionPool []string

	// チェイン（ブロック情報を詰めていく）
	chain []*Block
}

/**
* ブロックチェーンの初期値導入
* @return ブロックチェーン情報
**/
func NewBlockChain() *BlockChain {
	bc := new(BlockChain)
	//一番最初のブロックには、ナンスやハッシュ値がないため
	bc.CreateBlock(0, "Init hash")
	return bc
}

/**
* ブロックの作成
* @para nonce: ナンス
* @para previousHash: ハッシュ値
* @return b: 作成したブロック
**/
func (bc *BlockChain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

/**
* ブロックチェーンのサーバーログ用
**/
func (bc *BlockChain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s \n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

/**
* ログ出力の初期値
**/
func init() {
	log.SetPrefix("BlockChain: ")
}

func main() {
	// b := NewBlock(0, "init Hash")
	// b.Print()
	blockChain := NewBlockChain()
	blockChain.Print()
	blockChain.CreateBlock(5, "hash 1")
	blockChain.Print()
	blockChain.CreateBlock(4, "hash 2")
	blockChain.Print()
}

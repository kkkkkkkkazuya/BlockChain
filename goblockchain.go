package main

import (
	"crypto/sha256"
	"encoding/json"
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
	previousHash [32]byte

	// タイムスタンプ
	timestamp int64

	// トランザクション
	transactions []*Transaction
}

/**
* ブロックの初期値
* @para nonce: ナンス
* @para previousHash: ハッシュ値
* @para transactions: トランザクション情報の配列
* @return ブロック情報
**/
func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions
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
	fmt.Printf("previousHash %x\n", b.previousHash)
	for _, t := range b.transactions {
		t.Print()
	}
}

/**
* ブロックのハッシュ化
* @return jsonでハッシュ化した情報
**/
func (b *Block) Hash() [32]byte {
	//構造体をjsonでマーシャルする
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

/**
* カスタマイズでマーシャルJSONを上書き
*　@return capitalした値
**/
func (b *Block) MarshalJSON() ([]byte, error) {
	// prinvateのままだからcapitalに変更(publicでないとマーシャルできないため)
	return json.Marshal(struct {
		Timestamp    int64          `json: "timestamp"`
		Nonce        int            `json: "nonce"`
		PreviousHash [32]byte       `json: "previous_hash"`
		Transactions []*Transaction `json: "transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

/**
* ブロックチェーンの構造体
**/
type BlockChain struct {
	// トランザクションプール（トランザクション情報を詰めていく）
	transactionPool []*Transaction

	// チェイン（ブロック情報を詰めていく）
	chain []*Block
}

/**
* ブロックチェーンの初期値導入
* @return ブロックチェーン情報
**/
func NewBlockChain() *BlockChain {
	// 初期値だから空
	b := &Block{}
	bc := new(BlockChain)

	//一番最初のブロックには、ナンスやハッシュ値がないため
	bc.CreateBlock(0, b.Hash())
	return bc
}

/**
* ブロックの作成
* @para nonce: ナンス
* @para previousHash: ハッシュ値
* @return b: 作成したブロック
**/
func (bc *BlockChain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	// ブロックを追加した後に、空にしないといけないため
	bc.transactionPool = []*Transaction{}
	return b
}

/**
* １個前のブロック情報
* @return 一個前のブロック情報
**/
func (bc *BlockChain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
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
* トランザクションプールにトランザクション情報を追加
* @para sender: 送信者
* @para recipient: 受信者
* @para value: 仮想通貨の値
**/
func (bc *BlockChain) addTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

/**
* トランザクションの構造体
**/
type Transaction struct {
	// 送信側のアドレス
	senderBlockcahinAddress string
	// 受信側のアドレス
	recipientBlockchainAddress string
	// 仮想通貨の値
	value float32
}

/**
* トランザクションの初期値
* @para sender: 送信側のアドレス
* @para recipent: 受信側のアドレス
* @para value: 仮想通貨の値
* @return トランザクション情報
**/
func NewTransaction(sender string, recipent string, value float32) *Transaction {
	return &Transaction{
		sender,
		recipent,
		value,
	}
}

/**
* トランザクションのサーバーログ用
**/
func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address    %s\n", t.senderBlockcahinAddress)
	fmt.Printf(" recipent_blockchain_address  %s\n", t.recipientBlockchainAddress)
	fmt.Printf(" value  %.1f\n", t.value)
}

/**
* カスタマイズでマーシャルJSONを上書き
*　@return capitalした値
**/
func (t *Transaction) MarshalJSON() ([]byte, error) {
	// privateだからcapitalに変更(publicでないとマーシャルできないため)
	return json.Marshal(struct {
		Sender    string  `json: "sender_blockcahin_address"`
		Recipient string  `json: "recipient_blockchain_address"`
		Value     float32 `json: "value"`
	}{
		Sender:    t.senderBlockcahinAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}

/**
* ログ出力の初期値
**/
func init() {
	log.SetPrefix("BlockChain: ")
}

/**
* メイン部分
**/
func main() {
	blockChain := NewBlockChain()
	blockChain.Print()

	// トランザクション情報の追加
	blockChain.addTransaction("A", "B", 1.0)

	// 一個前のブロックのハッシュ化した情報
	previousHash := blockChain.LastBlock().Hash()

	// ブロックチェーンの作成
	blockChain.CreateBlock(5, previousHash)
	blockChain.Print()

	// 一個前のブロックのハッシュ化した情報
	previousHash = blockChain.LastBlock().Hash()

	// ブロックチェーンの作成
	blockChain.CreateBlock(4, previousHash)
	blockChain.Print()
}

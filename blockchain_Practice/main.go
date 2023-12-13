package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
)

type Block struct {
	nouce        int
	previousHash [32]byte
	transactions []string
	timestamp    int64
}

func NewBlock(nonce int, previousHash [32]byte) *Block {
	b := new(Block)
	b.nouce = nonce
	b.previousHash = previousHash

	return b
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(&b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce        int      `json:"nonce"`
		PreviousHash [32]byte `json:"previous_hash"`
		Transactions []string `json:"transactions"`
		Timestamp    int64    `json:"timestamp"`
	}{
		Nonce:        b.nouce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
		Timestamp:    b.timestamp,
	})
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)

	return b

}

func NewBlockChain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)

	bc.CreateBlock(0, b.Hash())
	return bc
}

func init() {
	log.SetPrefix("BLOCKCHAIN: ")
}

func main() {
	log.Println("test")
	b := &Block{}

	log.Printf("%x\n", b.Hash())
}

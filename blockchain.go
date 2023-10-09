package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Block represents a block in the blockchain.
type Block struct {
	Transaction   string
	Nonce         int
	PreviousHash  string
	Hash          string
}

// Blockchain represents the blockchain itself.
type Blockchain struct {
	Blocks []*Block
}

// NewBlock creates a new block and adds it to the blockchain.
func (bc *Blockchain) NewBlock(transaction string, nonce int, previousHash string) {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
		Hash:         calculateHash(transaction, nonce, previousHash),
	}

	bc.Blocks = append(bc.Blocks, block)
}

// DisplayBlocks prints all the blocks in the blockchain.
func (bc *Blockchain) DisplayBlocks() {
	for _, block := range bc.Blocks {
		fmt.Printf("Transaction: %s\nNonce: %d\nPrevious Hash: %s\nHash: %s\n\n",
			block.Transaction, block.Nonce, block.PreviousHash, block.Hash)
	}
}

// calculateHash calculates the hash of a block.
func calculateHash(transaction string, nonce int, previousHash string) string {
	data := fmt.Sprintf("%s%d%s", transaction, nonce, previousHash)
	hashBytes := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hashBytes[:])
}

// InitializeBlockchain creates and initializes a new blockchain.
func InitializeBlockchain() *Blockchain {
	genesisBlock := &Block{
		Transaction:  "Genesis Block",
		Nonce:        0,
		PreviousHash: "",
		Hash:         calculateHash("Genesis Block", 0, ""),
	}

	return &Blockchain{Blocks: []*Block{genesisBlock}}
}

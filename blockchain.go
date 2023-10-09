//Suleiman Afzal
//20i-1791
//SE
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

// ChangeBlock changes the transaction of a block.
func (b *Block) ChangeBlock(newTransaction string) {
	b.Transaction = newTransaction
	b.Hash = b.calculateHash()
}

// VerifyChain verifies the integrity of the entire blockchain.
func VerifyChain(blocks []*Block) bool {
	for i := 1; i < len(blocks); i++ {
		currentBlock := blocks[i]
		previousBlock := blocks[i-1]

		// Check if the previous hash in the current block matches the hash of the previous block.
		if currentBlock.PreviousHash != previousBlock.calculateHash() {
			return false
		}
	}

	return true
}

// calculateHash calculates the hash of a block.
func (b *Block) calculateHash() string {
	data := fmt.Sprintf("%s%d%s", b.Transaction, b.Nonce, b.PreviousHash)
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

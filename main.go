package main

import (
	"fmt"
	"github.com/suleimanafzal/assignment01bca"
)

func main() {
	// Initialize the blockchain
	blockchain := assignment01bca.InitializeBlockchain()

	// Add some blocks
	blockchain.NewBlock("Alice to Bob", 1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.NewBlock("Bob to Carol", 2, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	// Display all blocks
	blockchain.DisplayBlocks()

	// Change the transaction of a block
	blockToChange := blockchain.Blocks[1] // Change the transaction of the second block
	blockToChange.ChangeBlock("Mallory to Eve")

	// Display all blocks after the change
	blockchain.DisplayBlocks()

	// Verify the blockchain
	isValid := assignment01bca.VerifyChain(blockchain.Blocks)
	fmt.Printf("Is the blockchain valid? %v\n", isValid)
}

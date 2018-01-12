package main

// The blockchian is just a sequence of blocks

type Blockchain struct {
	blocks []*Block
}

func (blockchain *Blockchain) AddBlock(data string) {
	prevBlock := blockchain.blocks[len(blockchain.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	blockchain.blocks = append(blockchain.blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

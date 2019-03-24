package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := newBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)

	return newBlock
}

type Block struct {
	Hash          string
	PrevBlockHash string
	Data          string
}

func (b *Block) setHash() {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data))
	b.Hash = hex.EncodeToString(hash[:])
}

func newBlock(data, prevBlockHash string) *Block {
	block := &Block{
		Data:          data,
		PrevBlockHash: prevBlockHash,
	}
	block.setHash()
	return block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		Blocks: []*Block{newGenesisBlock()},
	}
}

func newGenesisBlock() *Block {
	return newBlock("This is the Genesis Block.", "")
}

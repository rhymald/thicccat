package main

type BlockBranch struct {
	blocks []*Block
}

func (chain *BlockBranch) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockBranch() *BlockBranch {
	return &BlockBranch{[]*Block{Genesis()}}
}


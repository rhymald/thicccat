package thicket

type BlockBranch struct {
	Blocks []*Block
}

func (chain *BlockBranch) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockBranch() *BlockBranch {
	return &BlockBranch{[]*Block{Genesis()}}
}


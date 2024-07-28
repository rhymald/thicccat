package main 

import (
	"fmt"
)

var (
	branch = InitBlockBranch()
)

func init() {
	branch.AddBlock("one")
	branch.AddBlock("two")
	branch.AddBlock("three")

}
func main() {
	for _, block := range branch.blocks {
		fmt.Printf("Prsh: %X\n", block.Prsh)
		fmt.Printf("Data: %X\n", block.Data)
		fmt.Printf("Hash: %X\n", block.Hash)
		fmt.Printf("------------------------------------------------------------\n\n")
	}
	select{}
}
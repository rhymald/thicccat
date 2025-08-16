package main 

import (
	"fmt"
	"github.com/rhymald/thicccat/thicket"
)

var (
	branch = thicket.InitBlockBranch()
)

func init() {
	branch.AddBlock("one")
	branch.AddBlock("two")
	branch.AddBlock("three")

}
func main() {
	for _, block := range branch.Blocks {
		fmt.Printf("Prsh: %X\n", block.Prsh)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %X\n", block.Hash)
		fmt.Printf("=======================================================================\n\n")
	}
	select{}
}
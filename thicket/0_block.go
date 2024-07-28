package thicket 

import (
	// "bytes"
	// "crypto/sha512"
)

type Block struct {
	Trgr []byte // Which blokc triggered
	Prsh []byte // Which block paren
	Hash []byte // ID of the block
	Data []byte // Data of the block
	Nonce int64
}

// func (b *Block) DeriveHash() {
// 	info := bytes.Join(
// 		[][]byte{b.Data, b.Prsh}, 
// 		[]byte{},
// 	)
// 	fasthash := sha512.Sum512(info)
// 	b.Hash = fasthash[:]
// }

func CreateBlock(data string, prsh []byte) *Block {
	block := &Block{
		Trgr: []byte{}, 
		Prsh: prsh,
		Data: []byte(data),
		Nonce: 0,
	}
	pow := InitPoW(block)
	nonce, hash := pow.Run()
	// block.DeriveHash()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}
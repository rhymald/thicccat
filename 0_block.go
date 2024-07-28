package main 

import (
	"bytes"
	// "crypto/bcrypt"
	"crypto/sha512"
)

type Block struct {
	Trgr []byte // Which blokc triggered
	Prsh []byte // Which block paren
	Hash []byte // ID of the block
	Data []byte // Data of the block
}

func (b *Block) DeriveHash() {
	info := bytes.Join(
		[][]byte{b.Data, b.Prsh}, 
		[]byte{},
	)
	fasthash := sha512.Sum512(info)
	// slowhash := bcrypt.generateFromPassword(fasthash[:], 8)
	b.Hash = fasthash[:]
}

func CreateBlock(data string, prsh []byte) *Block {
	block := &Block{
		Trgr: []byte{}, 
		Prsh: prsh,
		Data: []byte(data),
	}
	block.DeriveHash()
	return block
}
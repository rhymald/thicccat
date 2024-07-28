package thicket 

import (
	"math/big"
	"math"
	"encoding/binary"
	"crypto/sha512"
	"bytes"
	"log"
)

const Diff = 12

type PoW struct {
	Block *Block
	Target *big.Int
}

func InitPoW(b *Block) *PoW {
	target := big.NewInt(1)
	target.Lsh(target, uint(512-Diff))
	pow := &PoW{b, target}
	return pow
}

func (pow *PoW) InitData(nonce int64) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.Data, 
			pow.Block.Prsh,
			ToHex(int64(nonce)),
			ToHex(int64(Diff)),
		}, 
		[]byte{},
	)
	return data
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil { log.Panic(err) }
	return buff.Bytes()
}

func (pow *PoW) Run() (int64, []byte) {
	var intHash big.Int 
	var hash [64]byte
	nonce := int64(0)
	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha512.Sum512(data)
		log.Printf("PoW: %x", hash)
		intHash.SetBytes(hash[:])
		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	log.Println("------------------------------------------------------------------")
	return nonce, hash[:]
}

func (pow *PoW) Validate() bool {
	var intHash big.Int 
	data := pow.InitData(pow.Block.Nonce)
	hash := sha512.Sum512(data)
	intHash.SetBytes(hash[:])
	return intHash.Cmp(pow.Target) == -1
}
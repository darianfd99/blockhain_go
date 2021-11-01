package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBLockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBLockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBLockHash: []byte(prevBlockHash),
		Hash:          []byte{},
	}

	block.SetHash()
	return block
}

func main() {
	b := NewBlock("Probando nuestro bloque", []byte("65dfs4"))
	fmt.Println("Timestamp", b.Timestamp)
	fmt.Println("Data:", string(b.Data))
	fmt.Println("PrevBLockHash:", string(b.PrevBLockHash))
	fmt.Println("Hash:", string(b.Hash))
}

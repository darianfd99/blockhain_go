package blockchain_go

import "fmt"

func main() {
	bc := NewBlockchain()
	bc.AddBlock("Probando nuestro bloque")
	bc.AddBlock("Enviamos 5 BTC a Marcos")
	bc.AddBlock("Vamos")

	for _, block := range bc.blocks {
		fmt.Printf("Hash B.A: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash : %x\n\n", block.Hash)

	}

}

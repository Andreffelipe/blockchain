package main

import (
	"fmt"
	"log"
	"time"
)

/*
	---------------------------- BLOCK ----------------------------
    +--------------------------+ 	   +--------------------------+
    | prev hash     timestamp  | ----> | prev hash     timestamp  |
    | 						   | 	   | 						  |
    |  	nounce     transaction | 	   |   nounce     transaction |
    +--------------------------+ 	   +--------------------------+

	prev hash => hash gerada do bloco anterior
	timestamp => data/hora em que o bloco foi criado
	nounce =>
	transaction => os dados a serem gravados no bloco
*/

type Block struct {
	nounce       int
	previousHash string
	timestamp    int64
	transaction  []string
}

func NewBlock(nounce int, previousHash string) *Block {
	b := new(Block)
	b.nounce = nounce
	b.timestamp = time.Now().UnixNano()
	b.previousHash = previousHash
	return b
}

func (b *Block) Print() {
	fmt.Printf("Timestamp      %d\n", b.timestamp)
	fmt.Printf("Nounce         %d\n", b.nounce)
	fmt.Printf("PreviousHash   %s\n", b.previousHash)
	fmt.Printf("Transaction    %s\n", b.transaction)
}

func init() {
	log.SetPrefix("Blockchain :")
}

func main() {
	b := NewBlock(0, "init hash")
	b.Print()
}

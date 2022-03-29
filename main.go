package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
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

/*
	Definition of type the Block
*/
type Block struct {
	nounce       int
	previousHash [32]byte
	timestamp    int64
	transaction  []string
}

func NewBlock(nounce int, previousHash [32]byte) *Block {
	b := new(Block)
	b.nounce = nounce
	b.timestamp = time.Now().UnixNano()
	b.previousHash = previousHash
	return b
}

func (b Block) Print() {
	fmt.Printf("Timestamp      %d\n", b.timestamp)
	fmt.Printf("Nounce         %d\n", b.nounce)
	fmt.Printf("PreviousHash   %x\n", b.previousHash)
	fmt.Printf("Transaction    %s\n", b.transaction)
}

/*
	Cria um hash com os dados contidos no bloco
*/
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nounce       int      `json:"nounce"`
		PreviousHash [32]byte `json:"previous_hash"`
		Timestamp    int64    `json:"timestamp"`
		Transaction  []string `json:"transaction"`
	}{
		Nounce:       b.nounce,
		PreviousHash: b.previousHash,
		Timestamp:    b.timestamp,
		Transaction:  b.transaction,
	})
}

/*
	Definition of type the Blockchain
*/
type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nounce int, previousHash [32]byte) *Block {
	b := NewBlock(nounce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
}

func init() {
	log.SetPrefix("Blockchain :")
}

func main() {
	b := NewBlockchain()

	hash := b.LastBlock().Hash()
	b.CreateBlock(0, hash)

	hash = b.LastBlock().Hash()
	b.CreateBlock(0, hash)
	b.Print()
}

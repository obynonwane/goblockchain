package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*Transaction
}

// creating a new block instance
func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions

	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp           %d\n", b.timestamp)
	fmt.Printf("nonce               %d\n", b.nonce)
	fmt.Printf("previous_hash       %x\n", b.previousHash)
	fmt.Printf("transactions        %s\n", b.transactions)
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

// this function is automatically invoked when a block struct is passed
// as an argument into json.Marshal or json.MarshalIndent
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

// define the blockchain struct
type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

// create Blochain instance
func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

// create a new block and append to blockchain
func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)

	return b
}

// method to identify the last block
func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}
func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

type Transaction struct {
	senderBlockchainAddress   string
	receiverBlockchainAddress string
	value                     float32
}

// create new transaction
func NewTransaction(sender string, receiver string, value float32) *Transaction {
	return &Transaction{
		sender,
		receiver,
		value,
	}
}

// print transaction
func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("_", 40))
	fmt.Printf(" sender_blockchain_address         %s\n", t.senderBlockchainAddress)
	fmt.Printf(" receiver_blockchain_address       %s\n", t.receiverBlockchainAddress)
	fmt.Printf(" value                             %.1f\n", t.value)
}

// martialling transaction data into json
func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender   string  `json:"sender_blockchain_address"`
		Receiver string  `json:"receiver_blockchain_address"`
		Value    float32 `json:"value"`
	}{
		Sender:   t.senderBlockchainAddress,
		Receiver: t.receiverBlockchainAddress,
		Value:    t.value,
	})
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {

	blockChain := NewBlockchain()
	blockChain.Print()

	previousHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(5, previousHash)
	blockChain.Print()

	previousHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(2, previousHash)
	blockChain.Print()
}

package main

//importing dependencies
import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// create a custom type to represent the blocks that will make up our blockchain
type Block struct {
	data         map[string]interface{}
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

// create a custom Blockchain type that contains our blocks
type Blockchain struct {
	genesisBlock Block //the first block added to the blockchain
	chain        []Block
	difficulty   int //the minimum effort miners must undertake to mine and include a block in the blockchain.
}

/*
The hash of a block is its identifier generated using cryptography.
We will derive the hash for each block in our blockchain by combining and then hashing the hash of the previous block,
the data of the current block, the current block’s timestamp, and PoW using the SHA256 algorithm.
Converted the block’s data to JSON
Concatenated the previous block’s hash, and the current block’s data, timestamp, and PoW
Hashed the earlier concatenation with the SHA256 algorithm
Returned the base 16 hash as a string
*/
func (b Block) calculateHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

// Mining is the process of adding a new block to the blockchain.
// the number of zeros is called the mining difficulty
// if the mining difficulty is three, you have to generate a block hash that starts with "000" like, "0009a1bfb506…".
func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
	}
}

// add the new block to the blockchain
func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}
	return Blockchain{
		genesisBlock, //the first block on the blockchain
		[]Block{genesisBlock},
		difficulty,
	}
}

/*
Collects the details of a transaction (sender, receiver, and transfer amount)
Creates a new block with the transaction details
Mines the new block with the previous block hash, current block data, and generated PoW
Adds the newly created block to the blockchain
*/
func (b *Blockchain) addBlock(from, to string, amount float64) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}
	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
}

// Checking the validity of the blockchain
func (b Blockchain) isValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}

// Using the blockchain to make transactions
func main() {
	// create a new blockchain instance with a mining difficulty of 2
	blockchain := CreateBlockchain(2)

	// record transactions on the blockchain for Alice, Bob, and John
	blockchain.addBlock("Alice", "Bob", 5)
	blockchain.addBlock("John", "Bob", 2)

	// check if the blockchain is valid; expecting true
	fmt.Println(blockchain.isValid())
}

//refence : https://blog.logrocket.com/build-blockchain-with-go/

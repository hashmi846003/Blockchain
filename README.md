# Blockchain
Blockchain implementation using golang
# SimpleCrypto

A simple demonstration of a basic cryptocurrency implementation in Go.

## Features

- Basic blockchain functionality
- Simple proof-of-work consensus mechanism
- Transaction handling with encryption

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.16 or higher)

### Installation

1. **Clone the repository**:
    ```bash
    git clone https://github.com/yourusername/SimpleCrypto.git
    cd SimpleCrypto
    ```

2. **Run the code**:
    ```bash
    go run main.go
    ```

## How It Works

### Blockchain Structure

A blockchain is a series of blocks, where each block contains data and is cryptographically linked to the previous block, forming a chain. This ensures data integrity and immutability.

### Block Structure

Each block in the blockchain contains:
- `Index`: Position of the block in the blockchain.
- `Timestamp`: The time when the block was created.
- `Transactions`: A list of transactions included in the block.
- `PreviousHash`: Hash of the previous block in the chain.
- `Hash`: Hash of the current block, calculated using its data and the previous hash.
- `Nonce`: A number used once for the proof-of-work algorithm.

### Transaction Structure

Transactions represent the transfer of cryptocurrency between parties. Each transaction includes:
- `From`: Sender's address.
- `To`: Recipient's address.
- `Amount`: Amount of cryptocurrency transferred.

### Proof-of-Work

Proof-of-work is a consensus mechanism used to secure the blockchain. It requires miners to solve a computational puzzle to add a new block. In this project, the hash of each block must start with "0000". The `Nonce` value is incremented until this condition is met.

### Functions

- **calculateHash**: Generates a hash for a block based on its content. This hash acts as a unique identifier and ensures data integrity.
    ```go
    func calculateHash(block Block) string {
        record := strconv.Itoa(block.Index) + block.Timestamp + fmt.Sprint(block.Transactions) + block.PreviousHash + strconv.Itoa(block.Nonce)
        h := sha256.New()
        h.Write([]byte(record))
        hashed := h.Sum(nil)
        return hex.EncodeToString(hashed)
    }
    ```

- **generateBlock**: Creates a new block given the previous block and a list of transactions. It includes a simple proof-of-work mechanism that adjusts the `Nonce` until the hash meets the required condition.
    ```go
    func generateBlock(oldBlock Block, transactions []Transaction) Block {
        var newBlock Block

        newBlock.Index = oldBlock.Index + 1
        newBlock.Timestamp = time.Now().String()
        newBlock.Transactions = transactions
        newBlock.PreviousHash = oldBlock.Hash

        for i := 0; ; i++ {
            newBlock.Nonce = i
            newBlock.Hash = calculateHash(newBlock)
            if isValidHash(newBlock.Hash) {
                break
            }
        }

        return newBlock
    }

    func isValidHash(hash string) bool {
        return hash[:4] == "0000"
    }
    ```

- **createGenesisBlock**: Initializes the blockchain with the first block, known as the genesis block.
    ```go
    func createGenesisBlock() Block {
        genesisBlock := Block{
            Index:        0,
            Timestamp:    time.Now().String(),
            Transactions: []Transaction{},
            PreviousHash: "",
            Hash:         "",
        }
        genesisBlock.Hash = calculateHash(genesisBlock)
        return genesisBlock
    }
    ```

### Example Usage

This example demonstrates how to create a simple blockchain, add a couple of transactions, and append a new block to the blockchain. The blockchain's blocks are then printed to the console.

1. **Create the Genesis Block**:
    ```go
    genesisBlock := createGenesisBlock()
    blockchain := Blockchain{[]Block{genesisBlock}}
    ```

2. **Add Transactions and Generate a New Block**:
    ```go
    transactions := []Transaction{
        {From: "Alice", To: "Bob", Amount: 10.5},
        {From: "Bob", To: "Charlie", Amount: 3.0},
    }

    newBlock := generateBlock(genesisBlock, transactions)
    blockchain.Blocks = append(blockchain.Blocks, newBlock)
    ```

3. **Print the Blockchain**:
    ```go
    for _, block := range blockchain.Blocks {
        fmt.Printf("Index: %d\n", block.Index)
        fmt.Printf("Timestamp: %s\n", block.Timestamp)
        fmt.Printf("Transactions: %v\n", block.Transactions)
        fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
        fmt.Printf("Hash: %s\n", block.Hash)
        fmt.Println()
    }
    ```


## Acknowledgments

- Inspired by various blockchain tutorials and open-source projects.
- Thanks to the Go community for their excellent documentation and resources.



package main
import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "io"
    "strconv"
    "time"
	"crypto/sha256"
)
func encrypt(data string, key string) (string, error) {
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return "", err
    }

    ciphertext := make([]byte, aes.BlockSize+len(data))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(data))

    return hex.EncodeToString(ciphertext), nil
}

func decrypt(encryptedData string, key string) (string, error) {
    ciphertext, _ := hex.DecodeString(encryptedData)

    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return "", err
    }

    iv := ciphertext[:aes.BlockSize]
    ciphertext = ciphertext[aes.BlockSize:]

    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(ciphertext, ciphertext)

    return string(ciphertext), nil
}
type Block struct {
    Index        int
    Timestamp    string
    EncryptedData string
    PreviousHash string
    Hash         string
}

func calculateHash(block Block) string {
    record := strconv.Itoa(block.Index) + block.Timestamp + block.EncryptedData + block.PreviousHash
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, data string, key string) Block {
    var newBlock Block

    encryptedData, err := encrypt(data, key)
    if err != nil {
        fmt.Println(err)
    }

    newBlock.Index = oldBlock.Index + 1
    newBlock.Timestamp = time.Now().String()
    newBlock.EncryptedData = encryptedData
    newBlock.PreviousHash = oldBlock.Hash
    newBlock.Hash = calculateHash(newBlock)

    return newBlock
}

func createGenesisBlock(key string) Block {
    genesisBlock := Block{
        Index:        0,
        Timestamp:    time.Now().String(),
        EncryptedData: "",
        PreviousHash: "",
        Hash:         "",
    }

    data := "Genesis Block"
    encryptedData, err := encrypt(data, key)
    if err != nil {
        fmt.Println(err)
    }
    genesisBlock.EncryptedData = encryptedData
    genesisBlock.Hash = calculateHash(genesisBlock)
    return genesisBlock
}

func main() {
    key := "examplekey123456"  // AES-128 requires 16 bytes key

    genesisBlock := createGenesisBlock(key)
    blockchain := []Block{genesisBlock}

    newBlock := generateBlock(genesisBlock, "Second Block", key)
    blockchain = append(blockchain, newBlock)

    newBlock = generateBlock(newBlock, "Third Block", key)
    blockchain = append(blockchain, newBlock)

    for _, block := range blockchain {
        decryptedData, err := decrypt(block.EncryptedData, key)
        if err != nil {
            fmt.Println(err)
        }
        fmt.Printf("Index: %d\n", block.Index)
        fmt.Printf("Timestamp: %s\n", block.Timestamp)
        fmt.Printf("Decrypted Data: %s\n", decryptedData)
        fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
        fmt.Printf("Hash: %s\n\n", block.Hash)
    }
}

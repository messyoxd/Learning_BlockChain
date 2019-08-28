package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

type Message struct {
	BPM int
}

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, BPM int) (Block, error) {

	var newBlock Block

	newBlock.Index = oldBlock.Index + 1
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Timestamp = time.Now().String()
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}
func isBlockValid(block Block, oldBlock Block) bool {

	if oldBlock.Index+1 != block.Index {
		return false
	}

	if block.PrevHash != oldBlock.Hash {
		return false
	}

	if block.Hash != calculateHash(block) {
		return false
	}
	return true
}

func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}

}

func handleConn(conn net.Conn) {

	defer conn.Close()

	io.WriteString(conn, "Escreva um novo BPM:")

	scanner := bufio.NewScanner(conn)

	go func() {
		for scanner.Scan() {
			bpm, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Printf("%v nao e um numero: %v", scanner.Text(), err)
				continue
			}
			newBlock, err := generateBlock(Blockchain[len(Blockchain)-1], bpm)
			if err != nil {
				log.Println(err)
				continue
			}
			if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
				newBlockchain := append(Blockchain, newBlock)
				replaceChain(newBlockchain)
			}
			bcServer <- Blockchain
			io.WriteString(conn, "\nEscreva um novo BPM:")
		}
	}()

	go func() {

		time.Sleep(30 * time.Second)
		output, err := json.Marshal(Blockchain)
		if err != nil {
			log.Fatal(err)
		}
		io.WriteString(conn, string(output))

	}()
	for _ = range bcServer {
		spew.Dump(Blockchain)
	}

}

var Blockchain []Block
var bcServer chan []Block

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	bcServer = make(chan []Block)

	t := time.Now()
	genesisBLock := Block{0, t.String(), 0, "", ""}
	spew.Dump(genesisBLock)
	Blockchain = append(Blockchain, genesisBLock)

	httpPort := os.Getenv("ADDR")

	server, err := net.Listen("tcp", ":"+httpPort)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("TCP server listening on port:" + httpPort)
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}

}

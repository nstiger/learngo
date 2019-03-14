package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

/*
1.Index 是这个块在整个链中的位置
2.Timestamp 显而易见就是块生成时的时间戳
3.Hash 是这个块通过 SHA256 算法生成的散列值
4.PrevHash 代表前一个块的 SHA256 散列值
5.BPM 每分钟心跳数，也就是心率
*/
// Block represents each 'item' in the blockchain
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

//定义一个结构表示整个链，最简单的表示形式就是一个 Block 的 slice：
// Blockchain is a series of validated Blocks
var Blockchain []Block

/*为了简化，我们直接以 JSON 格式返回整个链，
可以在浏览器中访问 localhost:8080 或者 127.0.0.1:8080 来查看（这里的8080就是你在 .env 中定义的端口号 ADDR）。
POST 请求的 handler 稍微有些复杂，我们先来定义一下 POST 请求的 payload：
*/
// Message takes incoming JSON payload for writing heart rate
type Message struct {
	BPM int
}

var mutex = &sync.Mutex{}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := Block{}
		genesisBlock = Block{0, t.String(), 0, calculateHash(genesisBlock), ""}
		spew.Dump(genesisBlock)

		mutex.Lock()
		Blockchain = append(Blockchain, genesisBlock)
		mutex.Unlock()
	}()
	log.Fatal(run())

}

/*
Web 服务
其中的端口号是通过前面提到的 .env 来获得，再添加一些基本的配置参数，
这个 web 服务就已经可以 listen and serve 了！
接下来我们再来定义不同 endpoint 以及对应的 handler。
例如，对“/”的 GET 请求我们可以查看整个链，“/”的 POST 请求可以创建块。
*/

func run() error {
	mux := makeMuxRouter()
	httpPort := os.Getenv("PORT")
	log.Println("HTTP Server Listening on port :", httpPort)
	s := &http.Server{
		Addr:           ":" + httpPort,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

// create handlers
func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
	return muxRouter
}

//GET 请求的 handler：
// write blockchain when we receive an http request
func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

//再看看 handler 的实现：
// takes JSON payload as an input for heart rate (BPM)
func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var msg Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&msg); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	mutex.Lock()
	prevBlock := Blockchain[len(Blockchain)-1]
	newBlock := generateBlock(prevBlock, msg.BPM)

	if isBlockValid(newBlock, prevBlock) {
		Blockchain = append(Blockchain, newBlock)
		spew.Dump(Blockchain)
	}
	mutex.Unlock()

	respondWithJSON(w, r, http.StatusCreated, newBlock)

}

//POST 请求处理完之后，无论创建块成功与否，需要返回客户端一个响应：
func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}


/*
校验函数：
需要有函数帮我们判断一个块是否有被篡改。检查 Index 来看这个块是否正确得递增，
检查 PrevHash 与前一个块的 Hash 是否一致，再来通过 calculateHash 检查当前块的 Hash 值是否正确
*/
// make sure block is valid by checking index, and comparing the hash of the previous block
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

//写一个函数，用来计算给定的数据的 SHA256 散列值：
// SHA256 hasing
func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

//编写一个生成块的函数：
/*Index 是从给定的前一块的 Index 递增得出，
时间戳是直接通过 time.Now() 函数来获得的，
Hash 值通过前面的 calculateHash 函数计算得出，
PrevHash 则是给定的前一个块的 Hash 值。
*/
// create a new block using previous block's hash
func generateBlock(oldBlock Block, BPM int) Block {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

/*通常来说，更长的链表示它的数据（状态）是更新的，所以我们需要一个函数
能帮我们将本地的过期的链切换成最新的链：
*/

func replaceChain(newBlocks []Block) {
    if len(newBlocks) > len(Blockchain) {
        Blockchain = newBlocks
    }
}
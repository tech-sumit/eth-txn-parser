package ethereum

import (
	"bytes"
	"encoding/json"
	"eth-indexer/pkg/parser"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	ETH_RPC         = "https://cloudflare-eth.com"
	ApplicationJson = "application/json"
)

type EthereumParser struct {
	currentBlock     int
	subscribers      map[string]bool
	transactions     map[string][]parser.Transaction
	mutex            sync.RWMutex
	firstBlockNumber int
}

func NewParser(firstBlockNumber int) *EthereumParser {
	return &EthereumParser{
		subscribers:      make(map[string]bool),
		transactions:     make(map[string][]parser.Transaction),
		firstBlockNumber: firstBlockNumber,
	}
}

func (p *EthereumParser) GetCurrentBlock() int {
	return p.currentBlock
}

func (p *EthereumParser) Subscribe(address string) bool {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.subscribers[address] = true
	return true
}

func (p *EthereumParser) GetTransactions(address string) []parser.Transaction {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	return p.transactions[address]
}

func (p *EthereumParser) parseBlock(blockNumber int) error {
	payload := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x%x", true],"id":1}`, blockNumber)

	resp, err := http.Post(ETH_RPC, ApplicationJson, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	blockData := result["result"].(map[string]interface{})
	transactions := blockData["transactions"].([]interface{})
	for _, tx := range transactions {
		txData := tx.(map[string]interface{})

		var from string
		if _, exists := txData["from"]; exists {
			fmt.Printf("from: %+v\n", txData["from"])
			from = fmt.Sprintf("%s", txData["from"])
		}

		var to string
		if _, exists := txData["to"]; exists {
			fmt.Printf("to: %+v\n", txData["to"])
			to = fmt.Sprintf("%s", txData["to"])
		}

		p.mutex.Lock()
		t := parser.TransactionFromJson(txData)
		if p.subscribers[from] {
			p.transactions[from] = append(p.transactions[from], t)
		}

		if p.subscribers[to] {
			p.transactions[to] = append(p.transactions[to], t)
		}
		p.mutex.Unlock()
	}

	p.currentBlock = blockNumber
	return nil
}

func (p *EthereumParser) StartParsing() {
	var latestBlock int
	for {
		if latestBlock == 0 || latestBlock == p.GetCurrentBlock() {
			latestBlock = p.GetLatestBlock()
			fmt.Println("Scanning latest block:", latestBlock)
			continue
		}

		err := p.parseBlock(p.firstBlockNumber)
		if err != nil {
			fmt.Printf("Error parsing block %d: %v\n", p.firstBlockNumber, err)
			break
		}
		p.firstBlockNumber++
		time.Sleep(100 * time.Millisecond)
	}
}

func (p *EthereumParser) GetLatestBlock() int {
	payload := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}`)

	resp, err := http.Post(ETH_RPC, ApplicationJson, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Printf("Error getting latest block number: %v", err)
		return 0
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error getting latest block number: %v", err)
		return 0
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("Error getting latest block number: %v", err)
		return 0
	}

	blockNumberHex := result["result"].(string)
	blockNumber, err := strconv.ParseInt(blockNumberHex[2:], 16, 64)
	if err != nil {
		fmt.Printf("Error getting latest block number: %v", err)
		return 0
	}

	return int(blockNumber)
}

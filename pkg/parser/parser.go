package parser

// Parser interface
type Parser interface {
	GetCurrentBlock() int
	Subscribe(address string) bool
	GetTransactions(address string) []Transaction
}

// Transaction represents an Ethereum transaction.
// @Description Transaction object
// @type object
// Transaction represents an Ethereum transaction.
// @Description Ethereum transaction details.
// swagger:model Transaction
type Transaction struct {
	BlockHash            string        `json:"blockHash"`            // Hash of the block where this transaction was in.
	BlockNumber          string        `json:"blockNumber"`          // Number of the block where this transaction was in.
	From                 string        `json:"from"`                 // Address of the sender.
	Gas                  string        `json:"gas"`                  // Gas provided by the sender.
	GasPrice             string        `json:"gasPrice"`             // Gas price provided by the sender in Wei.
	MaxFeePerGas         string        `json:"maxFeePerGas"`         // Maximum fee per gas willing to pay in Wei.
	MaxPriorityFeePerGas string        `json:"maxPriorityFeePerGas"` // Maximum priority fee per gas willing to pay in Wei.
	Hash                 string        `json:"hash"`                 // Hash of the transaction.
	Input                string        `json:"input"`                // The data sent along with the transaction.
	Nonce                string        `json:"nonce"`                // The number of transactions made by the sender prior to this one.
	To                   string        `json:"to"`                   // Address of the receiver. null when it's a contract creation transaction.
	TransactionIndex     string        `json:"transactionIndex"`     // Integer of the transaction's index position in the block.
	Value                string        `json:"value"`                // Value transferred in Wei.
	Type                 string        `json:"type"`                 // EIP-2718 type of the transaction.
	AccessList           []interface{} `json:"accessList"`           // EIP-2930 access list.
	ChainId              string        `json:"chainId"`              // EIP-155 chain ID. Null means legacy transaction.
	V                    string        `json:"v"`                    // ECDSA recovery id.
	R                    string        `json:"r"`                    // ECDSA signature r.
	S                    string        `json:"s"`                    // ECDSA signature s.
	YParity              string        `json:"yParity"`              // EIP-1559 transaction y-parity.
}

func TransactionFromJson(data map[string]interface{}) Transaction {
	var t Transaction
	if val, ok := data["blockHash"].(string); ok {
		t.BlockHash = val
	}
	if val, ok := data["blockNumber"].(string); ok {
		t.BlockNumber = val
	}
	if val, ok := data["from"].(string); ok {
		t.From = val
	}
	if val, ok := data["gas"].(string); ok {
		t.Gas = val
	}
	if val, ok := data["gasPrice"].(string); ok {
		t.GasPrice = val
	}
	if val, ok := data["maxFeePerGas"].(string); ok {
		t.MaxFeePerGas = val
	}
	if val, ok := data["maxPriorityFeePerGas"].(string); ok {
		t.MaxPriorityFeePerGas = val
	}
	if val, ok := data["hash"].(string); ok {
		t.Hash = val
	}
	if val, ok := data["input"].(string); ok {
		t.Input = val
	}
	if val, ok := data["nonce"].(string); ok {
		t.Nonce = val
	}
	if val, ok := data["to"].(string); ok {
		t.To = val
	}
	if val, ok := data["transactionIndex"].(string); ok {
		t.TransactionIndex = val
	}
	if val, ok := data["value"].(string); ok {
		t.Value = val
	}
	if val, ok := data["type"].(string); ok {
		t.Type = val
	}
	if val, ok := data["accessList"]; ok {
		// Assuming accessList is an array of interfaces, further parsing would be needed based on structure.
		t.AccessList = val.([]interface{})
	}
	if val, ok := data["chainId"].(string); ok {
		t.ChainId = val
	}
	if val, ok := data["v"].(string); ok {
		t.V = val
	}
	if val, ok := data["r"].(string); ok {
		t.R = val
	}
	if val, ok := data["s"].(string); ok {
		t.S = val
	}
	if val, ok := data["yParity"].(string); ok {
		t.YParity = val
	}
	return t
}

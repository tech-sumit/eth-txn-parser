package controllers

import (
	"eth-indexer/pkg/parser"
	"github.com/gin-gonic/gin"
)

// SubscribeRequest represents the request to subscribe to an address.
type SubscribeRequest struct {
	Address string `json:"address" example:"0x..."`
}

// SubscribeResponse represents the response for a subscription request.
type SubscribeResponse struct {
	Subscribed bool `json:"subscribed"`
}

// CurrentBlockResponse represents the response containing the current block number.
type CurrentBlockResponse struct {
	CurrentBlock int `json:"currentBlock" example:"1234567"`
}

type Controllers interface {
	GetCurrentBlock(c *gin.Context)
	Subscribe(c *gin.Context)
	GetTransactions(c *gin.Context)
}

func Setup(parser parser.Parser) Controllers {
	return &controllers{parser: parser}
}

type controllers struct {
	parser parser.Parser
}

// GetCurrentBlock Get current block number
// @Summary Get current block number
// @Description Retrieves the current block number that the parser is aware of.
// @Produce json
// @Success 200 {object} CurrentBlockResponse "Returns the current block number"
// @Router /currentBlock [get]
func (ct *controllers) GetCurrentBlock(c *gin.Context) {
	currentBlock := ct.parser.GetCurrentBlock() // assuming ct.parser.GetCurrentBlock() returns an int
	c.JSON(200, gin.H{
		"currentBlock": currentBlock,
	})
}

// Subscribe to an address
// @Summary Subscribe to an address
// @Description Subscribes to an Ethereum address to monitor transactions.
// @Accept json
// @Produce json
// @Param body body SubscribeRequest true "Subscription request"
// @Success 200 {object} SubscribeResponse
// @Router /subscribe [post]
func (ct *controllers) Subscribe(c *gin.Context) {
	var request SubscribeRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	success := ct.parser.Subscribe(request.Address)
	c.JSON(200, SubscribeResponse{Subscribed: success})
}

// GetTransactions Get transactions for an address
// @Summary Get transactions for an address
// @Description Retrieves the list of transactions for a subscribed address.
// @Produce json
// @Param address path string true "The Ethereum address"
// @Success 200 {array} parser.Transaction
// @Failure 404 "Not Found if the address is not subscribed or does not have transactions"
// @Router /transactions/{address} [get]
func (ct *controllers) GetTransactions(c *gin.Context) {
	address := c.Param("address")
	transactions := ct.parser.GetTransactions(address)
	c.JSON(200, transactions)
}

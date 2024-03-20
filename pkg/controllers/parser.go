package controllers

import (
	"eth-indexer/pkg/parser"
	"github.com/gin-gonic/gin"
)

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

func (ct *controllers) GetCurrentBlock(c *gin.Context) {
	var currentBlock int
	currentBlock = ct.parser.GetCurrentBlock()
	c.JSON(200, gin.H{
		"currentBlock": currentBlock,
	})
}

func (ct *controllers) Subscribe(c *gin.Context) {
	var json struct {
		Address string `json:"address"`
	}
	if c.BindJSON(&json) == nil {
		success := ct.parser.Subscribe(json.Address)
		c.JSON(200, gin.H{
			"subscribed": success,
		})
	}
}

func (ct *controllers) GetTransactions(c *gin.Context) {
	address := c.Param("address")
	transactions := ct.parser.GetTransactions(address)
	c.JSON(200, transactions)
}

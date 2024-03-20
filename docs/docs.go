// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/currentBlock": {
            "get": {
                "description": "Retrieves the current block number that the parser is aware of.",
                "produces": [
                    "application/json"
                ],
                "summary": "Get current block number",
                "responses": {
                    "200": {
                        "description": "Returns the current block number",
                        "schema": {
                            "$ref": "#/definitions/controllers.CurrentBlockResponse"
                        }
                    }
                }
            }
        },
        "/subscribe": {
            "post": {
                "description": "Subscribes to an Ethereum address to monitor transactions.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Subscribe to an address",
                "parameters": [
                    {
                        "description": "Subscription request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.SubscribeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.SubscribeResponse"
                        }
                    }
                }
            }
        },
        "/transactions/{address}": {
            "get": {
                "description": "Retrieves the list of transactions for a subscribed address.",
                "produces": [
                    "application/json"
                ],
                "summary": "Get transactions for an address",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The Ethereum address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/parser.Transaction"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found if the address is not subscribed or does not have transactions"
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.CurrentBlockResponse": {
            "type": "object",
            "properties": {
                "currentBlock": {
                    "type": "integer",
                    "example": 1234567
                }
            }
        },
        "controllers.SubscribeRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "0x..."
                }
            }
        },
        "controllers.SubscribeResponse": {
            "type": "object",
            "properties": {
                "subscribed": {
                    "type": "boolean"
                }
            }
        },
        "parser.Transaction": {
            "description": "Transaction object",
            "type": "object",
            "properties": {
                "accessList": {
                    "description": "EIP-2930 access list.",
                    "type": "array",
                    "items": {}
                },
                "blockHash": {
                    "description": "Hash of the block where this transaction was in.",
                    "type": "string"
                },
                "blockNumber": {
                    "description": "Number of the block where this transaction was in.",
                    "type": "string"
                },
                "chainId": {
                    "description": "EIP-155 chain ID. Null means legacy transaction.",
                    "type": "string"
                },
                "from": {
                    "description": "Address of the sender.",
                    "type": "string"
                },
                "gas": {
                    "description": "Gas provided by the sender.",
                    "type": "string"
                },
                "gasPrice": {
                    "description": "Gas price provided by the sender in Wei.",
                    "type": "string"
                },
                "hash": {
                    "description": "Hash of the transaction.",
                    "type": "string"
                },
                "input": {
                    "description": "The data sent along with the transaction.",
                    "type": "string"
                },
                "maxFeePerGas": {
                    "description": "Maximum fee per gas willing to pay in Wei.",
                    "type": "string"
                },
                "maxPriorityFeePerGas": {
                    "description": "Maximum priority fee per gas willing to pay in Wei.",
                    "type": "string"
                },
                "nonce": {
                    "description": "The number of transactions made by the sender prior to this one.",
                    "type": "string"
                },
                "r": {
                    "description": "ECDSA signature r.",
                    "type": "string"
                },
                "s": {
                    "description": "ECDSA signature s.",
                    "type": "string"
                },
                "to": {
                    "description": "Address of the receiver. null when it's a contract creation transaction.",
                    "type": "string"
                },
                "transactionIndex": {
                    "description": "Integer of the transaction's index position in the block.",
                    "type": "string"
                },
                "type": {
                    "description": "EIP-2718 type of the transaction.",
                    "type": "string"
                },
                "v": {
                    "description": "ECDSA recovery id.",
                    "type": "string"
                },
                "value": {
                    "description": "Value transferred in Wei.",
                    "type": "string"
                },
                "yParity": {
                    "description": "EIP-1559 transaction y-parity.",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{"http"},
	Title:            "Ethereum Parser API",
	Description:      "API for parsing Ethereum blockchain data",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

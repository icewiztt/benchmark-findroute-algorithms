package entity

import "gorm.io/gorm"

type Swap struct {
	Pool              string      `json:"pool"`
	TokenIn           string      `json:"tokenIn"`
	TokenOut          string      `json:"tokenOut"`
	SwapAmount        string      `json:"swapAmount"`
	AmountOut         string      `json:"amountOut"`
	LimitReturnAmount string      `json:"limitReturnAmount"`
	MaxPrice          string      `json:"maxPrice"`
	Exchange          string      `json:"exchange"`
	PoolLength        int         `json:"poolLength"`
	PoolType          string      `json:"poolType"`
	Extra             interface{} `json:"extra,omitempty"`
}

type TokenInfo struct {
	Address  string  `json:"address"`
	Symbol   string  `json:"symbol"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Decimals uint8   `json:"decimals"`
}

type FindRouteResponse struct {
	InputAmount     string                 `json:"inputAmount"`
	OutputAmount    string                 `json:"outputAmount"`
	TotalGas        int64                  `json:"totalGas"`
	GasPriceGwei    string                 `json:"gasPriceGwei"`
	GasUsd          float64                `json:"gasUsd"`
	AmountInUsd     float64                `json:"amountInUsd"`
	AmountOutUsd    float64                `json:"amountOutUsd"`
	ReceivedUsd     float64                `json:"receivedUsd"`
	Swaps           [][]Swap               `json:"swaps"`
	Tokens          map[string]TokenInfo   `json:"tokens"`
	EncodedSwapData string                 `json:"encodedSwapData,omitempty"`
	RouterAddress   string                 `json:"routerAddress,omitempty"`
	Debug           map[string]interface{} `json:"debug,omitempty"`
}

type TestResult struct {
	gorm.Model
	RunningTime float64
	MaxHops     uint8
	MaxPaths    uint8
	MinPartUsd  uint32

	OldNumPaths uint8
	OldNumHops  uint8
	NewNumPaths uint8
	NewNumHops  uint8

	InputAmount     string `json:"inputAmount"`
	OldOutputAmount string `json:"oldOutputAmount"`
	NewOutputAmount string `json:"newOutputAmount"`

	AmountInUsd     float64 `json:"amountInUsd"`
	OldAmountOutUsd float64 `json:"oldAmountOutUsd"`
	NewAmountOutUsd float64 `json:"newAmountOutUsd"`

	Diff float64
}

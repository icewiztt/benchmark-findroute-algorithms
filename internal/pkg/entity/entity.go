package entity

import (
	"time"

	"gorm.io/gorm"
)

type FindRouteRequest struct {
	TokenIn    string `form:"tokenIn" binding:"required"`
	TokenOut   string `form:"tokenOut" binding:"required"`
	AmountIn   string `form:"amountIn" binding:"required"`
	SaveGas    string `form:"saveGas"`
	Dexes      string `form:"dexes"`
	GasInclude string `form:"gasInclude"`
	GasPrice   string `form:"gasPrice"`
	Debug      string `form:"debug"`
}

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

type RouteResponse struct {
	InputAmount     string                 `json:"inputAmount"`
	OutputAmount    string                 `json:"outputAmount"`
	MinOutputAmount string                 `json:"minOutputAmount"`
	AmountInUsd     float64                `json:"amountInUsd"`
	AmountOutUsd    float64                `json:"amountOutUsd"`
	Swaps           [][]Swap               `json:"swaps"`
	Tokens          map[string]TokenInfo   `json:"tokens"`
	EncodedSwapData string                 `json:"encodedSwapData,omitempty"`
	RouterAddress   string                 `json:"routerAddress,omitempty"`
	Debug           map[string]interface{} `json:"debug,omitempty"`
}

type TestResult struct {
	gorm.Model
	TokenIn, TokenOut string
	AmountIn          string

	V1TotalNumberOfSwaps, V2TotalNumberOfSwaps, V3TotalNumberOfSwaps int
	V1AmountOut, V2AmountOut, V3AmountOut                            string
	V1ResponseTime, V2ResponseTime, V3ResponseTime                   time.Duration
	V1andV2Same                                                      bool
	V3BetterThanV1                                                   bool
}

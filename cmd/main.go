package main

import (
	"fmt"
	"math"
	"math/big"

	"benchmark-find-route/internal/pkg/entity"
	"benchmark-find-route/internal/pkg/fetcher"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Run(db *gorm.DB, originalTest, newWayTest fetcher.InputRequestParamFindRoute) error {
	originalResult, responseTimeOfOriginalFetch, err := fetcher.NewFindRouteFetcher(&originalTest).Get()
	newWayResult, responseTimeOfNewWayFetch, err := fetcher.NewFindRouteFetcher(&newWayTest).Get()

	oldResultAmountOut, _ := new(big.Float).SetString(originalResult.OutputAmount)
	newResultAmountOut, _ := new(big.Float).SetString(newWayResult.OutputAmount)

	if err != nil {
		return err
	}

	diff, _ :=
		new(big.Float).Mul(
			new(big.Float).Quo(
				new(big.Float).Sub(newResultAmountOut, oldResultAmountOut),
				oldResultAmountOut,
			), big.NewFloat(100.0)).Float64()

	fmt.Println(originalResult.OutputAmount, newWayResult.OutputAmount, diff)

	newNumHops := 0
	for _, arr := range newWayResult.Swaps {
		for range arr {
			newNumHops++
		}
	}

	oldNumHops := 0
	for _, arr := range originalResult.Swaps {
		for range arr {
			oldNumHops++
		}
	}

	result := entity.TestResult{
		RunningTime: math.Max(responseTimeOfOriginalFetch.Seconds(), responseTimeOfNewWayFetch.Seconds()),
		MaxHops:     6,
		MaxPaths:    2,
		MinPartUsd:  500,

		OldNumPaths: uint8(len(originalResult.Swaps)),
		OldNumHops:  uint8(oldNumHops),

		NewNumPaths: uint8(len(newWayResult.Swaps)),
		NewNumHops:  uint8(newNumHops),

		InputAmount: newWayResult.InputAmount,
		AmountInUsd: newWayResult.AmountInUsd,

		OldOutputAmount: originalResult.OutputAmount,
		NewOutputAmount: newWayResult.OutputAmount,

		OldAmountOutUsd: originalResult.AmountOutUsd,
		NewAmountOutUsd: newWayResult.AmountOutUsd,

		Diff: diff,
	}
	db.Create(&result)
	return nil
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&entity.TestResult{})

	var (
		test = fetcher.InputRequestParamFindRoute{
			Url:      fetcher.BaseUrl,
			TokenIn:  "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
			TokenOut: "0xdac17f958d2ee523a2206206994597c13d831ec7",
			//AmountIn:          "1000000000000000000000",
			AmountIn:          "1000000000000000000",
			SaveGas:           0,
			GasInclude:        0,
			SlippageTolerance: 50,
			Deadline:          2664879454,
			To:                "0x7446c5C01b8E627cBD55702C81779671b3b00124",
			ChargeFeeBy:       "",
			FeeReceiver:       "",
			IsInBps:           "",
			FeeAmount:         "",
			ClientData:        "ks",
		}
	)
	originalTest := test
	newWayTest := test
	newWayTest.Url = fetcher.BaseUrlV2

	base18, _ := new(big.Int).SetString("1000000000000000000", 10)
	for i := 900; i <= 1000; i++ {
		amountIn := new(big.Int).Mul(big.NewInt(int64(i)), base18).String()
		originalTest.AmountIn = amountIn
		newWayTest.AmountIn = amountIn
		Run(db, originalTest, newWayTest)
	}
}

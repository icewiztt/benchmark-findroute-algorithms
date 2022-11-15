package main

import (
	"fmt"
	"math/big"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"benchmark-find-route/internal/pkg/entity"
	"benchmark-find-route/internal/pkg/fetcher"
)

func Run(db *gorm.DB, urlV1, urlV2, urlV3, tokenIn, tokenOut, amountIn, dexes string) error {
	v1Result, v1responseTime, err := fetcher.NewFindRouteFetcher(urlV1).Get(tokenIn, tokenOut, amountIn, dexes)
	v2Result, v2responseTime, err := fetcher.NewFindRouteFetcher(urlV2).Get(tokenIn, tokenOut, amountIn, dexes)
	v3Result, v3responseTime, err := fetcher.NewFindRouteFetcher(urlV3).Get(tokenIn, tokenOut, amountIn, dexes)

	v1ResultAmountOut, _ := new(big.Float).SetString(v1Result.OutputAmount)
	v2ResultAmountOut, _ := new(big.Float).SetString(v2Result.OutputAmount)
	v3ResultAmountOut, _ := new(big.Float).SetString(v3Result.OutputAmount)

	if err != nil {
		return err
	}

	v1NumHops := 0
	for _, arr := range v1Result.Swaps {
		for range arr {
			v1NumHops++
		}
	}

	v2NumHops := 0
	for _, arr := range v2Result.Swaps {
		for range arr {
			v2NumHops++
		}
	}

	v3NumHops := 0
	for _, arr := range v3Result.Swaps {
		for range arr {
			v3NumHops++
		}
	}

	fmt.Println(v1ResultAmountOut, v2ResultAmountOut, v3ResultAmountOut)
	fmt.Println(v1NumHops, v2NumHops, v3NumHops)
	fmt.Println(v1responseTime, v2responseTime, v3responseTime)

	result := entity.TestResult{
		TokenIn:              tokenIn,
		TokenOut:             tokenOut,
		AmountIn:             amountIn,
		V1TotalNumberOfSwaps: v1NumHops,
		V2TotalNumberOfSwaps: v2NumHops,
		V3TotalNumberOfSwaps: v3NumHops,
		V1AmountOut:          v1ResultAmountOut.String(),
		V2AmountOut:          v2ResultAmountOut.String(),
		V3AmountOut:          v3ResultAmountOut.String(),
		V1ResponseTime:       v1responseTime,
		V2ResponseTime:       v2responseTime,
		V3ResponseTime:       v3responseTime,
		V1andV2Same:          v1ResultAmountOut.Cmp(v2ResultAmountOut) == 0,
		V3BetterThanV1:       v3ResultAmountOut.Cmp(v1ResultAmountOut) > 0,
	}
	db.Save(&result)
	return nil
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&entity.TestResult{})
	if err != nil {
		panic("failed to migrate schema")
	}
	amountInList := []string{
		"1000000",
		"1000000000",
		"1000000000000",
		"1000000000000000",
	}
	for _, tokenIn := range entity.Tokens {
		for _, tokenOut := range entity.Tokens {
			if tokenIn != tokenOut {
				for _, amountIn := range amountInList {
					Run(db, fetcher.BaseUrlV1, fetcher.BaseUrlV2, fetcher.BaseUrlV3,
						tokenIn, tokenOut, amountIn, entity.DefaultDexes)
				}
			}
		}
	}
}

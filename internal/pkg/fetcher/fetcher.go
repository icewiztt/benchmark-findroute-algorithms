package fetcher

import (
	"encoding/json"
	"time"

	"benchmark-find-route/internal/pkg/entity"
	"github.com/go-resty/resty/v2"
)

const (
	BaseUrl   = "http://localhost:8081/ethereum/route/encode"
	BaseUrlV2 = "http://localhost:8080/ethereum/route/encode"
)

type InputRequestParamFindRoute struct {
	Url               string
	TokenIn           string
	TokenOut          string
	AmountIn          string
	SaveGas           uint8
	GasInclude        uint8
	SlippageTolerance uint8
	Deadline          uint64
	To                string
	ChargeFeeBy       string
	FeeReceiver       string
	IsInBps           string
	FeeAmount         string
	ClientData        string
}
type FindRouteFetcher struct {
	Client *resty.Client
	*InputRequestParamFindRoute
}

func NewFindRouteFetcher(
	inputParams *InputRequestParamFindRoute,
) *FindRouteFetcher {
	client := resty.New()
	return &FindRouteFetcher{
		Client:                     client,
		InputRequestParamFindRoute: inputParams,
	}
}

func (f *FindRouteFetcher) Get() (*entity.FindRouteResponse, time.Duration, error) {
	resp, err := f.Client.R().EnableTrace().
		SetQueryParams(map[string]string{
			"url":               f.Url,
			"tokenIn":           f.TokenIn,
			"tokenOut":          f.TokenOut,
			"amountIn":          f.AmountIn,
			"saveGas":           string(f.SaveGas),
			"gasInclude":        string(f.GasInclude),
			"slippageTolerance": string(f.SlippageTolerance),
			"deadline":          string(f.Deadline),
			"to":                f.To,
			"chargeFeeBy":       f.ChargeFeeBy,
			"feeReceiver":       f.FeeReceiver,
			"isInBps":           f.IsInBps,
			"feeAmount":         string(f.FeeAmount),
			"clientData":        f.ClientData,
		}).SetHeader("Accept", "application/json").Get(f.Url)

	if err != nil {
		panic(err)
		return nil, 0, err
	}

	raw := resp.Body()
	findRouteResponse := entity.FindRouteResponse{}
	err = json.Unmarshal(raw, &findRouteResponse)
	if err != nil {
		return nil, 0, err
	}

	return &findRouteResponse, resp.Time(), nil
}

package fetcher

import (
	"encoding/json"
	"time"

	"github.com/go-resty/resty/v2"

	"benchmark-find-route/internal/pkg/entity"
)

const (
	BaseUrlV1 = "http://localhost:8080/solana/route/"
	BaseUrlV2 = "http://localhost:8081/solana/route/"
	BaseUrlV3 = "http://localhost:8082/solana/route/"
)

type FindRouteFetcher struct {
	Client *resty.Client
	Url    string
}

func NewFindRouteFetcher(url string) *FindRouteFetcher {
	client := resty.New()
	return &FindRouteFetcher{
		Client: client,
		Url:    url,
	}
}

func (f *FindRouteFetcher) Get(tokenIn, tokenOut, amountIn, dexes string) (*entity.RouteResponse, time.Duration, error) {
	resp, err := f.Client.R().EnableTrace().
		SetQueryParams(map[string]string{
			"tokenIn":  tokenIn,
			"tokenOut": tokenOut,
			"amountIn": amountIn,
			// "dexes":    dexes,
		}).SetHeader("Accept", "application/json").Get(f.Url)

	if err != nil {
		panic(err)
		return nil, 0, err
	}

	raw := resp.Body()
	findRouteResponse := entity.RouteResponse{}
	err = json.Unmarshal(raw, &findRouteResponse)
	if err != nil {
		return nil, 0, err
	}

	return &findRouteResponse, resp.Time(), nil
}

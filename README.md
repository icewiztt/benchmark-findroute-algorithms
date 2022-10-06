# Benchmark find route algorithms


## Prerequisites

* Go 1.17+
* Redis


## Quick Starts guide

From the root folder of the project:

### Start Redis
- Ensure that redis server which run in localhost has existing find route data crawled from ethereum-network scanners.
- Note that this quickstart guide is used only for Ethereum mainnet network
### Prepare old kyberswap algorithm

1. Clone project
```bash
git clone https://github.com/KyberNetwork/kyberswap-aggregator
go mod download
```
2. Edit `bindAddress` in `internal/pkg/config/ethereum.yaml` to `:8081`
3. Run project
```bash
go run ./cmd/app -c internal/pkg/config/ethereum.yaml api
```

### Prepare new brute-force kyberswap algorithm

4. Clone project
```bash
git clone -b fix/find-route-algorithm https://github.com/KyberNetwork/kyberswap-aggregator
go mod download
```
5. Edit `bindAddress` in `internal/pkg/config/ethereum.yaml` to `:8080`
6. Run project
```bash
go run ./cmd/app -c internal/pkg/config/ethereum.yaml api
```

### Run benchmark script
7. Clone project
```bash
git clone https://github.com/datluongductuan/benchmark-findroute-algorithms
go mod download
```

8. Edit parameters in `internal/pkg/entity/constants.go` as you want, note that increasing in MaxPaths and MaxHops will result in larger complexity.
Recommend to use `MaxHops = 5` and `MaxPaths = 2`, this config will give good result when run with `amountIn <= 100ETH`
9. Edit parameters in new-algorithm-project at step 4 (project _kyberswap-aggreagator, branch fix/find-route-algorithm)_
- In [service/route.go](https://github.com/KyberNetwork/kyberswap-aggregator/blob/c1d4446a1b2748aa00b5b07f62a29e9d1222c9b8/internal/pkg/service/route.go#L628), line 628 and 630, edit value of `MathPaths` and `MaxHops` as same as the edited values at Step 8.
```bash
go run ./cmd/main.go 
```
_Note that every time you changed config at Step 8 and Step 9, you need to restart new-algortihm-project._

10. Result will save in `test.db` file at root project directory. You should install SQLite Extension to view/query.

[See my benchmark result](https://docs.google.com/spreadsheets/d/1kuzs04SFMnulamJ1xNAzJJIh_VRdexF6slzqIaPIkGQ/edit?usp=sharing)

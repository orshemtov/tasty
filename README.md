# Tasty

Unofficial Go client for the TastyWorks API.

## Authentication

The following environment variables are required:

1. `TASTYWORKS_PASSWORD`
2. `TASTYWORKS_USERNAME`

You'll need a token to authenticate with the API, and an account ID.

There's a `LoadConfig` function that will take care of saving both to `~/.tastyworks/config.json`, the function will run on init() when loading the library.

## Usage

```go
// Get account balance
balances, err := Balances()
 if err != nil {
  panic(err)
 }
    
// Get market data
symbols := []string{"AAPL", "GOOG"}
 data, err := MarketMetrics(symbols)
 if err != nil {
  panic(err)
 }

 for _, item := range data.Data.Items {
  fmt.Println(item)
 }
```

## TODO

- Save auth token and account ID at `~/.tastyworks/config`
- Cache token, invalidate the cache if the token has expired
- `PositionsResponse` struct from JSON response
- Improve typing on structs that parse JSON responses
- Add documentation comments

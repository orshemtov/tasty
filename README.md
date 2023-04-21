# Tasty

Unofficial Go client for the TastyWorks API.

## Authentication

The following environment variables are required:

1. `TASTYWORKS_PASSWORD`
2. `TASTYWORKS_USERNAME`

You'll need a token to authenticate with the API, and an account ID.

### Token

```go
r, err := Auth()
if errr != nil {
    panic(err)
}
fmt.Printf("token: %s\n", r.Data.SessionToken)
```

Save the response from the API in `.cache/token.json`:

```json
{
    "data": {
        "user": {
            "email": "me@email.com",
            "username": "myusername",
            "external-id": "..."
        },
        "session-token": "session-token-here"
    },
    "context": "/sessions"
}
```

### Account ID

```go
r, err := Accounts()
 if err != nil {
  panic(err)
 }
 for _, a := range r.Data.Items {
  fmt.Printf("account ID: %s\n", a.Account.AccountNumber)
 }
```

Save the account ID in `.cache/account.json` like so:

```json
{
    "number": "5AB25D45"
}
```

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


package main
// go get github.com/machinebox/graphql
import (
    "context"
    "fmt"
    "github.com/machinebox/graphql"
)

func main() {
    graphqlClient := graphql.NewClient("https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v3")
    graphqlRequest := graphql.NewRequest(`
      {
        bundles(first: 1) {
        id
        ethPriceUSD
      }
    }
    `)
    var graphqlResponse interface{}
    if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
        panic(err)
    }
    fmt.Println(graphqlResponse)
}

// go run uniswap-graphql-feed.go

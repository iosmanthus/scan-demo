package main

import (
	"context"
	"flag"
	"log"

	"github.com/tikv/client-go/v2/config"
	"github.com/tikv/client-go/v2/rawkv"
)

var (
	pdAddr = flag.String("pd", "", "address of PD")
	limit  = flag.Int("limit", 0, "limit of keys")
)

func main() {
	flag.Parse()
	ctx := context.Background()
	client, err := rawkv.NewClient(ctx, []string{*pdAddr}, config.Security{})
	if err != nil {
		log.Fatal(err)
	}

	keys, _, err := client.Scan(ctx, nil, nil, *limit)
	log.Printf("scaned %v keys\n", len(keys))
}

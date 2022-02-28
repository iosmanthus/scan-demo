package main

import (
	"context"
	"flag"
	"fmt"
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
	defer client.Close()

	keys, values, err := client.Scan(ctx, nil, nil, *limit)
	fmt.Printf("scaned %v keys\n", len(keys))
	for i := 0; i < len(keys); i++ {
		fmt.Printf("[%s:%s]\n", string(keys[i]), string(values[i]))
	}
}

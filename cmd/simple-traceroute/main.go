package main

import (
	"fmt"
	"time"

	tracelib "github.com/andylibrian/tracelib/pkg"
)

func main() {
	cache := tracelib.NewLookupCache()

	hops, err := tracelib.RunTrace("www.detik.com", "0.0.0.0", "::", time.Second, 64, cache)

	if nil != err {
		fmt.Println("Traceroute error:", err)
		return
	}

	for i, hop := range hops {
		fmt.Printf("%d. %v(%s)/AS%d %v (final:%v timeout:%v error:%v down:%v)\n", i+1, hop.Host, hop.Addr, hop.AS, hop.RTT, hop.Final, hop.Timeout, hop.Error, hop.Down)
	}
}

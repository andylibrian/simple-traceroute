# Simple Trace route

This is a simplified version of https://github.com/kanocz/tracelib for learning purpose.

Traceroute implementation in go including multi-round trace (returns min/max/avg/lost) and AS number detection for IPv4.
Usage example of regular traceroute (only IPs without hostnames and AS numbers):

```go
hops, err := tracelib.RunTrace("google.com", "0.0.0.0", time.Second, 64)
for i, hop := range hops {
	fmt.Printf("%d. %v(%s)/AS%d %v (final:%v timeout:%v error:%v)\n",
      i+1, hop.Host, hop.Addr, hop.AS, hop.RTT, hop.Final, hop.Timeout, hop.Error)
}
```

## Build

```
go build ./cmd/...
```

### Run

```
sudo ./simple-traceroute
```

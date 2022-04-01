# This project is for practice in Go.


## How to use

You can save **any type data** by `Set` method. It will save data for 30 seconds.

```go
cache = localcache.NewCacheV3()
cache.Set("key", 1)
cache.Get("key") // Got 1
// After 30 second
cache.Get("key") // Got nil
```

## How to test

Just run `go test`.
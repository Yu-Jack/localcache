# This project is for practice in Go.


## How to use

You can save **any type data** by `Set` method. It will save data for 30 seconds.

```go
localcache.Set("key", 1)
localcache.Get("key") // Got 1
// After 30 second
localcache.Get("key") // Got nil
```

## How to test

Just run `go test`.
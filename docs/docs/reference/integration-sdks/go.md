# Go SDK Reference

## Install

```sh
go get github.com/outagelab/go-sdk
```

## Initialize

```go
import "github.com/outagelab/go-sdk/outagelab"

func initOutageLab() {
    outagelab.Start(outagelab.Options{
		Application: "my-service",
		Environment: "local",
		ApiKey:      os.Getenv("OUTAGELAB_API_KEY"),
		Host:        "https://app.outagelab.com",
	})
}
```

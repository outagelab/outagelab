### Install

<!-- #region install -->

```sh
go get github.com/outagelab/go-sdk
```

<!-- #endregion install -->

### Setup

<!-- #region setup -->

```go
import "github.com/outagelab/go-sdk/outagelab"

func main() {
  outagelab.Start(outagelab.Options{
		Application: "my-service",
		Environment: "local",
		ApiKey:      os.Getenv("OUTAGELAB_API_KEY"),
		Host:        "https://app.outagelab.com",
	})
}
```

<!-- #endregion setup -->

### Install

<!-- #region install -->

```sh
pip install outagelab
```

<!-- #endregion install -->

### Setup

<!-- #region setup -->

```py
import outagelab

outagelab.start(
    application="my-service",
    environment="local",
    api_key=os.environ.get("OUTAGELAB_API_KEY"),
    host="https://app.outagelab.com"
)
```

<!-- #endregion setup -->

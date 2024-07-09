# Python SDK Reference

## Install

```sh
pip install outagelab
```

## Initialize

```python
import outagelab

outagelab.init(
    application="my-service",
    environment="local",
    api_key=os.environ.get("OUTAGELAB_API_KEY"),
    host="https://app.outagelab.com"
)
```

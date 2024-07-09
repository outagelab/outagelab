# Node.js SDK Reference

## Install

```sh
npm install -s outagelab
```

## Initialize

```js
const OutageLabClient = require("outagelab");

let client = new OutageLabClient({
  application: "my-service",
  environment: "local",
  apiKey: process.env.OUTAGELAB_API_KEY,
  host: "https://app.outagelab.com",
});
```

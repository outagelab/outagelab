### Install

<!-- #region install -->

```sh
npm install -s outagelab
```

OR

```sh
yarn add outagelab
```

<!-- #endregion install -->

### Setup

<!-- #region setup -->

```js
const outagelab = require("outagelab");

outagelab.start({
  application: "my-service",
  environment: "local",
  apiKey: process.env.OUTAGELAB_API_KEY,
  host: "https://app.outagelab.com",
});
```

<!-- #endregion setup -->

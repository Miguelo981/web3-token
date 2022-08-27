# Web3 Token

Web3 Token is a new way to authenticate users inspired in [web3-token](https://github.com/bytesbay/web3-token) for Golang and use it as a middleware in your backends.

---
## Install

```bash
$ go get https://github.com/Miguelo981/web3-token
```

---

## Example usage (Server side)
```go

import (
	web3Token "github.com/Miguelo981/web3-token"
)

decoded, err := web3Token.Verify(token, "")

if err != nil {
	//Token is invalid
	return
}

// Now you can find that user by his address or use the signed attributes in the token
req.user = await User.findOne({ decoded.Address });
```

---

## API

### Descrypt(signer, options)
Name | Description | Required | Example
--- | --- | --- | ---
`token` | A function that returns a the decrypted body of the signed token | `required` | `web3Token.Descrypt(`eyJzaWduYXR1cmUiOi`)`


### Verify(token, options)
Name | Description | Required | Example
--- | --- | --- | ---
`token` | A token string that is generated from `sign()` | `required` | `...`
`options` | An options object | `required` | `{ domain: 'example.com' }`
`options.domain` | The domain you want to accept | `required` | `'example.com'`

---

## License
Web3 Token is released under the MIT license. © 2022 Miguelo
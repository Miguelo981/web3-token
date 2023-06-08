# Web3 Token

Web3 Token is an open-source Golang library that empowers you to decode and verify web3 signed tokens within the Ethereum network. It offers a seamless solution for utilizing these tokens as authentication tokens in your Golang backends, facilitating the creation of middlewares and custom services with utmost convenience. Inspired in [web3-token](https://github.com/bytesbay/web3-token).

---
## Install

```bash
$ go get github.com/Miguelo981/web3-token
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

## Contributing

We welcome contributions to enhance the Web3-token and make it even better. To contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your modifications and commit the changes.
4. Push your changes to your forked repository.
5. Submit a pull request to the main repository.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgements

We would like to express our gratitude to the open-source community for their invaluable contributions and support.

## Contact

If you have any questions, suggestions, or feedback, please feel free to reach out to us at [email protected]

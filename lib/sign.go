package lib

import (
	"errors"
	"github.com/Miguelo981/web3-token/utils"
)

/*type Signature struct {
	Raw   []byte
	Hash  [32]byte
	R     [32]byte
	S     [32]byte
	V     uint8
}*/

/*
raw256, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
ecdsaKey, _ := ssh.NewSignerFromKey(raw256)

var key, _ = crypto.GenerateKey()
var p = hex.EncodeToString(key.D.Bytes())

func Sign2(message string) Signature {
	hashRaw := crypto.Keccak256([]byte(message))
	signature, err := crypto.Sign(hashRaw, p.ecdsa)
	p.errorHandler.Handle(err, "Signature error")

	return Signature{
		signature,
		p.bytes32(hashRaw),
		p.bytes32(signature[:32]),
		p.bytes32(signature[32:64]),
		uint8(int(signature[65])) + 27, // Yes add 27, weird Ethereum quirk
	}
}*/

type Token struct {
	signature string
	body interface{}
}

type TokenParams struct {
	expires_in string
}

func Sign(message string, expiration string) (string, error) {
	if expiration == "" {
		expiration = "1d"
	}

	/*params := TokenParams{
		expires_in: expiration,
	}*/

	//validateParams(params)

	//processParams(params)

	//msg := buildMessage(params)

	//crypto.Sign()
	//signature := signer(msg)
	signature := ""
	msg := ""
	if signature == "" { //nil
		return "", errors.New("'signer' argument should be a function that returns a signature string (Promise<string>)")
	}

	token, err := utils.EncodeToBase64(Token{signature: signature, body: msg})
	if err == nil {
		return "", err
	}

	return token.String(), nil
}
package lib

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/Miguelo981/web3-token/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"regexp"
	"strconv"
	"strings"
)

type Signature struct {
	r []byte
	s []byte
	v uint8
}

type DecryptedToken struct {
	Version int
	Address string
	StringBody string
	Body map[string]string
	Signature string
	PublicKey ed25519.PublicKey
}

func GetVersion(body string) (int, error) {
	r, _ := regexp.Compile("Web3[\\s-]+Token[\\s-]+Version: \\d")
	str := r.FindString(body)
	parsedStr := strings.Split(strings.Replace(str, " ", "", -1), ":")[1]

	version, err := strconv.Atoi(parsedStr)
	if err != nil {
		return 0, err
	}

	return version, nil
}

func HashPersonalMessage(message []byte) []byte {
	prefix := []byte("\u0019Ethereum Signed Message:\n" + strconv.Itoa(len(message)))

	return crypto.Keccak256(append(prefix, message...))
}

func FromRpcSig(sig []byte) (*Signature, error) {
	if len(sig) < 65 {
		return nil, errors.New("Invalid signature length")
	}

	v, _ := strconv.ParseInt(hex.EncodeToString(sig[64:]), 16, 32) //uint8(int(sig[64])) + 27 //uint8(binary.LittleEndian.Uint32(signature[:65]))+27
	if v < 27 {
		v += 27
	}

	return &Signature{r: sig[:32], s: sig[32:64], v: uint8(v)}, nil
}

func PublicKeyBytesToAddress(publicKey []byte) common.Address {
	var buf []byte

	hash := crypto.NewKeccakState()
	hash.Write(publicKey[1:]) // remove EC prefix 04
	buf = hash.Sum(nil)
	address := buf[12:]

	return common.HexToAddress(hex.EncodeToString(address))
}

func Decrypt(token string) (*DecryptedToken, error) {
	if len(token) < 1 || token == "" {
		return nil, errors.New("Token required.")
	}

	base64Decoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err //return errors.New("Token malformed (must be base64 encoded)")
	}

	var privateKeyJSON map[string]string//interface{}
	err = json.Unmarshal(base64Decoded, &privateKeyJSON)
	if err != nil {
		return nil, errors.New("Token malformed (unparsable JSON)")
	}

	body, signature := privateKeyJSON["body"],  privateKeyJSON["signature"]

	if body == "" {
		return nil, errors.New("Token malformed (empty message)")
	}

	if signature == "" {
		return nil, errors.New("Token malformed (empty signature)")
	}

	msgBuffer, err := utils.ToBuffer("0x" + hex.EncodeToString([]byte(body)))
	if err != nil {
		return nil, err
	}

	msgHash := HashPersonalMessage(msgBuffer)
	signatureBuffer, err := utils.ToBuffer(signature) //"0" + strings.Split(signature, "0x")[1]
	if err != nil {
		return nil, err
	}
	signatureParams, err := FromRpcSig(signatureBuffer)
	print(signatureParams)
	if err != nil {
		return nil, err
	}

	rsv := append(append(signatureParams.r, signatureParams.s...), signatureParams.v - 27)

	publicKey, err := crypto.Ecrecover(msgHash, rsv) //signatureParams
	if err != nil {
		return nil, err
	}

	address := PublicKeyBytesToAddress(publicKey)

	version, err := GetVersion(body)
	if err != nil {
		return nil, err
	}

	decryptedToken := DecryptedToken{ Version: version, Address: address.String(), StringBody: body, Signature: signature, PublicKey: publicKey }

	return &decryptedToken, nil
}
package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"strings"
)

func ToBuffer(value string) ([]byte, error) {
	if value[:2] == "0x" {
		value = value[2:]
	}

	if len(value) % 2 == 1 {
		value = "0"+value
	}

	buffer, err := hex.DecodeString(value)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func EncodeToBase64(v interface{}) (bytes.Buffer, error) {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	err := json.NewEncoder(encoder).Encode(v)
	if err != nil {
		return bytes.Buffer{}, err
	}
	encoder.Close()
	return buf, nil
}

func DecodeFromBase64(v interface{}, enc string) error {
	return json.NewDecoder(base64.NewDecoder(base64.StdEncoding, strings.NewReader(enc))).Decode(v)
}

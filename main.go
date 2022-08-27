package main

import (
	"github.com/Miguelo981/web3-token/lib"
)

func main() {
	token := "eyJzaWduYXR1cmUiOiIweDFiYmY3N2FjNzhmYzgwZDY1ZTA0MmNkYjgxOTRiYzQyZmE0OGRmMjRhNTE4MDZkNDU2ZjA5MTc2NzUyYWMyNzY2YWQ1MmJmOWQxZmIzMDgyMmU4ZmI1ZjQwZTFmNDQzYjg4N2U2NTViMDg4MWFhNzJlMzU3MDJmOTVkOGUyMGM4MWMiLCJib2R5IjoiVVJJOiBodHRwOi8vbG9jYWxob3N0OjMwMDAvcG9zdC8xXG5XZWIzIFRva2VuIFZlcnNpb246IDJcbk5vbmNlOiAxOTM1NzEwMVxuSXNzdWVkIEF0OiAyMDIyLTA1LTE3VDE3OjE3OjE4LjU5MlpcbkV4cGlyYXRpb24gVGltZTogMjAyMi0wNS0xOFQxNzoxNzoxOC4wMDBaIn0="
	//decoded, err := lib.Decrypt(token)
	decoded, err := lib.Verify(token, "")
	if err != nil {
		print(err)
		return
	}
	print(decoded.Body)
	println(decoded)
}
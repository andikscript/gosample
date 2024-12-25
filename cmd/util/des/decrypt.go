package des

import (
	"crypto/des"
	"encoding/hex"
	"fmt"
	"samplecode/cmd/util"
)

func DesDecryption(ct string) string {
	keys := "\u0001\u0002\u0003\u0004\u0005\u0006\u0007\u0008"
	text, _ := hex.DecodeString(ct)

	block, err := des.NewCipher([]byte(keys))
	if err != nil {
		util.LogError(fmt.Sprintf("Error when make new cipher or '%v'", err))
		return ""
	}

	mode := NewECBDecrypter(block)
	pt := make([]byte, len(text))
	mode.CryptBlocks(pt, text)
	padder := NewPkcs7Padding(block.BlockSize())

	pt, err = padder.Unpad(pt)
	if err != nil {
		util.LogError(fmt.Sprintf("Error when unpad value or '%v'", err))
		return ""
	}

	return string(pt)
}

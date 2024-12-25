package des

import (
	"bytes"
	"errors"
)

type Padding interface {
	Pad(p []byte) ([]byte, error)
	Unpad(p []byte) ([]byte, error)
}

type Padder struct {
	blockSize int
}

func NewPkcs5Padding() Padding {
	return &Padder{blockSize: 8}
}

func NewPkcs7Padding(blockSize int) Padding {
	return &Padder{blockSize: blockSize}
}

func (p *Padder) Pad(buf []byte) ([]byte, error) {
	bufLen := len(buf)
	padLen := p.blockSize - (bufLen % p.blockSize)
	padText := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(buf, padText...), nil
}

func (p *Padder) Unpad(buf []byte) ([]byte, error) {
	bufLen := len(buf)
	if bufLen == 0 {
		return nil, errors.New("crypto/padding: invalid padding size")
	}

	pad := buf[bufLen-1]
	if pad == 0 {
		return nil, errors.New("crypto/padding: invalid last byte of padding")
	}

	padLen := int(pad)
	if padLen > bufLen || padLen > p.blockSize {
		return nil, errors.New("crypto/padding: invalid padding size")
	}

	for _, v := range buf[bufLen-padLen : bufLen-1] {
		if v != pad {
			return nil, errors.New("crypto/padding: invalid padding")
		}
	}

	return buf[:bufLen-padLen], nil
}

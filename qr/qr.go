package qr

import (
	"errors"

	"github.com/skip2/go-qrcode"
)

func GenerateQR(uri string) ([]byte, error) {
	code, err := qrcode.Encode(uri, qrcode.Medium, 256)
	if err != nil {
		return nil, errors.New("Could not generate a QR code")
	}
	return code, nil
}

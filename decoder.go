package ngssp

import (
	"encoding/xml"
	"io"
)

func decode[T any](rd io.Reader) (T, error) {
	var payload T
	err := xml.NewDecoder(rd).Decode(&payload)
	return payload, err
}

type NgsspPackageDefDecoder struct {
	rd io.Reader
}

func (pd *NgsspPackageDefDecoder) Decode() (*NgsspPackage, error) {
	payload, err := decode[NgsspPackageEnvelope](pd.rd)
	return &payload.NgsspPackage, err
}

func NewNgsspPackageDefDecoder(rd io.Reader) *NgsspPackageDefDecoder {
	return &NgsspPackageDefDecoder{rd: rd}
}

type NgsspPackageSelectorDecoder struct {
	rd io.Reader
}

func (pd *NgsspPackageSelectorDecoder) Decode() (*NgsspVariantPackageSelector, error) {
	payload, err := decode[NgsspSelectorEnvelope](pd.rd)
	return &payload.NgsspVariantPackageSelector, err
}

func NewNgsspPackageSelectorDecoder(rd io.Reader) *NgsspPackageSelectorDecoder {
	return &NgsspPackageSelectorDecoder{rd: rd}
}

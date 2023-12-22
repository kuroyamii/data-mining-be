package model

import (
	"encoding/json"
	"io"
)

type ClassifyResponseFromFlask struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Errors  interface{}    `json:"errors"`
	Data    ClassifyResult `json:"data"`
}

type ClassifyResult struct {
	ResNet string `json:"resnet"`
	VGG16  string `json:"vgg"`
}

func (fr *ClassifyResponseFromFlask) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(fr)
}

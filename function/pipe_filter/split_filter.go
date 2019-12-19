package main

import (
	"errors"
	"strings"
)

var SplitFilterFormatError = errors.New("split filter error")

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response, error) {
	str, ok := data.(string)
	if !ok {
		return nil, SplitFilterFormatError
	}
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}

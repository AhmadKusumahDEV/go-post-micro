package types

import (
	"errors"
)

var (
	ErrNotFound = errors.New("not found query")
)

type Header string

const (
	HeaderKey Header = "header"
)

package discogs

import (
	"fmt"
	"strings"
)

type Error struct {
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("discogs error: %s", strings.ToLower(e.Message))
}

var (
	ErrCurrencyNotSupported = &Error{"currency does not supported"}
)

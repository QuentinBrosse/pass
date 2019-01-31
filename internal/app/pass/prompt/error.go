package prompt

import "errors"

var ErrAbort = errors.New("")
var ErrEOF = errors.New("^D")
var ErrInterrupt = errors.New("^C")

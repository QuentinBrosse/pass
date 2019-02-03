package prompt

import "errors"

// ErrAborted is an error who signify that the user has intentionally aborted a prompt operation.
var ErrAborted = errors.New("aborted")

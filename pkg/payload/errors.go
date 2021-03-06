package payload

import (
	"net/http"

	"github.com/ppacher/webthings-mqtt-gateway/pkg/errors"
)

// Common errors
var (
	ErrNoType            = errors.NewWithStatus(http.StatusBadRequest, "no handler type")
	ErrInvalidType       = errors.NewWithStatus(http.StatusBadRequest, "invalid handler type")
	ErrAlreadyRegistered = errors.NewWithStatus(http.StatusInternalServerError, "handler type already registered")
)

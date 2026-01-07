package errors

import (
	"github.com/valyala/fasthttp"

	"github.com/mbatimel/mobile_phone_only_menu/pkg/models/consts"
)

var (
	ForbiddenError        = func() *Error { return New("forbidden", fasthttp.StatusForbidden, consts.ErrForbidden) }
	MethodNotAllowedError = func() *Error { return New("method not allowed", fasthttp.StatusBadRequest, consts.ErrMethodNotAllowed) }
	InternalServerError   = func() *Error {
		return New("internal server error", fasthttp.StatusInternalServerError, consts.ErrInternal)
	}
	NotFound       = func() *Error { return New("not found", fasthttp.StatusBadRequest, consts.ErrNotFound) }
	AlreadyExists  = func() *Error { return New("key already exists", fasthttp.StatusConflict, consts.ErrBadRequest) }
	BadMetaValue   = func() *Error { return New("bad meta", fasthttp.StatusBadRequest, consts.ErrBadRequest) }
	InvalidRequest = func() *Error { return New("invalid request", fasthttp.StatusBadRequest, consts.ErrBadRequest) }
)

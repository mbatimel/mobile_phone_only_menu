package publicapi

import "go/types"

type Resp200 struct {
	// @tg desc=`Флаг показывающий, что ответ пришел с ошибкой`
	// @tg example=false
	Error bool `json:"error"`
	// @tg example=nil
	ErrorText string `json:"errorText"`
	// @tg example=nil
	Data types.Nil `json:"data"`
	// @tg example=nil
	AdditionalErrors types.Nil `json:"additionalErrors"`
}

type Err400 struct {
	// @tg example=`nil`
	Data types.Nil `json:"data,omitempty"`
	// @tg desc=`Флаг показывающий, что ответ пришел с ошибкой`
	Error bool `json:"error"`
	// @tg desc=`Заголовок ошибки`
	// @tg example=`content.api.errors.retention.badRequest`
	ErrorText string `json:"errorText"`
	// @tg desc=`Текст ошибки, при ответе`
	AdditionalErrors struct {
		Errors []struct {
			TrKey string `json:"trKey"`
			// @tg example=`{"userID": "2354532", "supplierID": "dc76595d-88bf-4cf4-ac49-4698649358a0"}`
			Params map[string]string `json:"params"`
		} `json:"errors"`
	} `json:"additionalErrors"`
}

type Err403 struct {
	// @tg example=`nil`
	Data types.Nil `json:"data,omitempty"`
	// @tg desc=`Флаг показывающий, что ответ пришел с ошибкой`
	Error bool `json:"error"`
	// @tg desc=`Заголовок ошибки`
	// @tg example=`content.api.errors.retention.accessDenied`
	ErrorText string `json:"errorText"`
	// @tg desc=`Текст ошибки, при ответе, со статус кодом 403, не указывается`
	AdditionalErrors types.Nil `json:"additionalErrors,omitempty"`
}

type Err405 struct {
	// @tg example=`nil`
	Data types.Nil `json:"data,omitempty"`
	// @tg desc=`Флаг показывающий, что ответ пришел с ошибкой`
	Error bool `json:"error"`
	// @tg desc=`Заголовок ошибки`
	// @tg example=`content.api.errors.retention.methodNotAllowed`
	ErrorText string `json:"errorText"`
	// @tg desc=`Текст ошибки, при ответе, со статус кодом 403, не указывается`
	AdditionalErrors types.Nil `json:"additionalErrors,omitempty"`
}

type Err409 struct {
	// @tg example=`nil`
	Data types.Nil `json:"data,omitempty"`
	// @tg desc=`Флаг показывающий, что ответ пришел с ошибкой`
	Error bool `json:"error"`
	// @tg desc=`Заголовок ошибки`
	// @tg example=`content.api.errors.retention.conflict`
	ErrorText string `json:"errorText"`
	// @tg desc=`Текст ошибки, при ответе, со статус кодом 409, не указывается`
	AdditionalErrors types.Nil `json:"additionalErrors,omitempty"`
}

type Err500 struct {
	// @tg example=`nil`
	Data types.Nil `json:"data,omitempty"`
	// @tg desc=`Флаг показывающий, что ответ пришел с ошибкой`
	Error bool `json:"error"`
	// @tg desc=`Заголовок ошибки`
	// @tg example=`content.api.errors.retention.internalError`
	ErrorText string `json:"errorText"`
	// @tg desc=`Текст ошибки, при ответе, со статус кодом 403, не указывается`
	AdditionalErrors types.Nil `json:"additionalErrors,omitempty"`
}

package consts

const (
	ErrInternal                 = "errors.menu.internalError"            // Внутренняя ошибка
	ErrBadRequest               = "errors.menu.badRequest"               // Плохой запрос
	ErrMethodNotAllowed         = "errors.menu.methodNotAllowed"         // Метод не поддерживается
	ErrForbidden                = "errors.menu.forbidden"                // Доступ запрещен
	ErrInvalidRequest           = "errors.menu.invalidRequest"           // Неправильный запрос
	ErrAccessDenied             = "errors.menu.accessDenied"             // Отказано в доступе
	ErrFailedGetOldSupplierID   = "errors.menu.failedGetOldSupplierId"   //"failed get old supplierID"
	ErrNotEnoughMoney           = "errors.menu.notEnoughMoney"           // На балансе недостаточно средств 402
	ErrUnknownCountryCode       = "errors.menu.unknownCountryCode"       // Неизвестный код страны
	ErrUnavailableInDemo        = "errors.menu.unavailableInDemo"        // Данный сервис недоступен в демо периоде
	ErrPromoUnavailable         = "errors.menu.promoUnavailable"         // Промокод недоступен
	ErrPremiumOptionUnavailable = "errors.menu.premiumOptionUnavailable" // Промокод недоступен
	ErrNotFound                 = "errors.menu.notFound"                 // Данные по запросу не найдены
	ErrTariffReadOnly           = "errors.menu.tariffReadOnly"           // Данные по запросу не найдены

)

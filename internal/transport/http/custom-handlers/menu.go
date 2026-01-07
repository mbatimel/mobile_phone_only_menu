package customhandlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mbatimel/mobile_phone_only_menu/pkg/errors"
	"github.com/mbatimel/mobile_phone_only_menu/pkg/interfaces/publicapi"
	"github.com/rs/zerolog/log"
)

const MenuDishServiceName = "Menu-Dish"

func CreateDish(ctx *fiber.Ctx, svc publicapi.PublicApi, secretId uuid.UUID, dish string, categoty string) error {
	var (
		methodName = "CreateDish"
		err        error
	)

	defer func(begin time.Time) {
		fields := map[string]interface{}{
			"method":      "post",
			"path":        "/menu/api/create/dish",
			"handlerName": methodName,
			"dish":        dish,
			"categoty":    categoty,
			"service":     MenuDishServiceName,
			"took":        time.Since(begin).String(),
		}
		l := log.Info()
		if err != nil {
			if errors.Is(err, errors.ForbiddenError()) {
				l = log.Warn().Err(err)
			} else {
				l = log.Error().Err(err)
			}
		}
		l.Fields(fields).Msg("call")

	}(time.Now())

	err = svc.CreateDish(ctx.UserContext(), secretId, dish, categoty)
	if err != nil {
		sendResponse(ctx, log.Logger, nil, err)
		return nil
	}

	sendResponse(ctx, log.Logger, nil, nil)
	return err
}

func MarkFavoriteDish(ctx *fiber.Ctx, svc publicapi.PublicApi, secretId uuid.UUID, ids []uint64) error {
	var (
		methodName = "MarkFavoriteDish"
		err        error
	)

	defer func(begin time.Time) {
		fields := map[string]interface{}{
			"method":      "post",
			"path":        "/menu/api/mark",
			"handlerName": methodName,
			"ids":         ids,

			"service": MenuDishServiceName,
			"took":    time.Since(begin).String(),
		}
		l := log.Info()
		if err != nil {
			if errors.Is(err, errors.ForbiddenError()) {
				l = log.Warn().Err(err)
			} else {
				l = log.Error().Err(err)
			}
		}
		l.Fields(fields).Msg("call")

	}(time.Now())

	err = svc.MarkFavoriteDish(ctx.UserContext(), secretId, ids)
	if err != nil {
		sendResponse(ctx, log.Logger, nil, err)
		return nil
	}

	sendResponse(ctx, log.Logger, nil, nil)
	return err
}

func DeleteDish(ctx *fiber.Ctx, svc publicapi.PublicApi, secretId uuid.UUID, id uint64) error {
	var (
		methodName = "DeleteDish"
		err        error
	)

	defer func(begin time.Time) {
		fields := map[string]interface{}{
			"method":      "delete",
			"path":        "/menu/api/delete",
			"handlerName": methodName,
			"id":          id,

			"service": MenuDishServiceName,
			"took":    time.Since(begin).String(),
		}
		l := log.Info()
		if err != nil {
			if errors.Is(err, errors.ForbiddenError()) {
				l = log.Warn().Err(err)
			} else {
				l = log.Error().Err(err)
			}
		}
		l.Fields(fields).Msg("call")

	}(time.Now())

	err = svc.DeleteDish(ctx.UserContext(), secretId, id)
	if err != nil {
		sendResponse(ctx, log.Logger, nil, err)
		return nil
	}

	sendResponse(ctx, log.Logger, nil, nil)
	return err
}

func CreateChef(ctx *fiber.Ctx, svc publicapi.PublicApi, secretId uuid.UUID, name string) error {
	var (
		methodName = "CreateChef"
		err        error
	)

	defer func(begin time.Time) {
		fields := map[string]interface{}{
			"method":      "post",
			"path":        "/menu/api/create/chef",
			"handlerName": methodName,
			"name":        name,
			"service":     MenuDishServiceName,
			"took":        time.Since(begin).String(),
		}
		l := log.Info()
		if err != nil {
			if errors.Is(err, errors.ForbiddenError()) {
				l = log.Warn().Err(err)
			} else {
				l = log.Error().Err(err)
			}
		}
		l.Fields(fields).Msg("call")

	}(time.Now())

	err = svc.CreateChef(ctx.UserContext(), secretId, name)
	if err != nil {
		sendResponse(ctx, log.Logger, nil, err)
		return nil
	}

	sendResponse(ctx, log.Logger, nil, nil)
	return err
}

func UpdateDish(ctx *fiber.Ctx, svc publicapi.PublicApi, secretId uuid.UUID, id uint64, text string) error {
	var (
		methodName = "UpdateDish"
		err        error
	)

	defer func(begin time.Time) {
		fields := map[string]interface{}{
			"method":      "update",
			"path":        "/menu/api/update",
			"handlerName": methodName,
			"id":          id,
			"text":        text,
			"service":     MenuDishServiceName,
			"took":        time.Since(begin).String(),
		}
		l := log.Info()
		if err != nil {
			if errors.Is(err, errors.ForbiddenError()) {
				l = log.Warn().Err(err)
			} else {
				l = log.Error().Err(err)
			}
		}
		l.Fields(fields).Msg("call")

	}(time.Now())

	err = svc.UpdateDish(ctx.UserContext(), secretId, id, text)
	if err != nil {
		sendResponse(ctx, log.Logger, nil, err)
		return nil
	}

	sendResponse(ctx, log.Logger, nil, nil)
	return err
}

func GetAllDish(ctx *fiber.Ctx, svc publicapi.PublicApi, secretId uuid.UUID) error {
	var (
		methodName = "GetAllDish"
		err        error
	)

	defer func(begin time.Time) {
		fields := map[string]interface{}{
			"method":      "get",
			"path":        "/menu/api/all",
			"handlerName": methodName,
			"service":     MenuDishServiceName,
			"took":        time.Since(begin).String(),
		}
		l := log.Info()
		if err != nil {
			if errors.Is(err, errors.ForbiddenError()) {
				l = log.Warn().Err(err)
			} else {
				l = log.Error().Err(err)
			}
		}
		l.Fields(fields).Msg("call")

	}(time.Now())

	response, err := svc.GetAllDish(ctx.UserContext(), secretId)
	if err != nil {
		sendResponse(ctx, log.Logger, response, err)
		return nil
	}

	sendResponse(ctx, log.Logger, response, nil)
	return err
}
func GetFavoriteDish(ctx *fiber.Ctx, svc publicapi.PublicApi, secretId uuid.UUID) error {
	var (
		methodName = "GetFavoriteDish"
		err        error
	)

	defer func(begin time.Time) {
		fields := map[string]interface{}{
			"method":      "get",
			"path":        "/menu/api/favorite",
			"handlerName": methodName,
			"service":     MenuDishServiceName,
			"took":        time.Since(begin).String(),
		}
		l := log.Info()
		if err != nil {
			if errors.Is(err, errors.ForbiddenError()) {
				l = log.Warn().Err(err)
			} else {
				l = log.Error().Err(err)
			}
		}
		l.Fields(fields).Msg("call")

	}(time.Now())

	response, err := svc.GetFavoriteDish(ctx.UserContext(), secretId)
	if err != nil {
		sendResponse(ctx, log.Logger, response, err)
		return nil
	}

	sendResponse(ctx, log.Logger, response, nil)
	return err
}

func DeleteAllMenu(ctx *fiber.Ctx, svc publicapi.PublicApi, secretId uuid.UUID) error {
	var (
		methodName = "DeleteAllMenu"
		err        error
	)

	defer func(begin time.Time) {
		fields := map[string]interface{}{
			"method":      "delete",
			"path":        "/menu/api/all",
			"handlerName": methodName,
			"service":     MenuDishServiceName,
			"took":        time.Since(begin).String(),
		}
		l := log.Info()
		if err != nil {
			if errors.Is(err, errors.ForbiddenError()) {
				l = log.Warn().Err(err)
			} else {
				l = log.Error().Err(err)
			}
		}
		l.Fields(fields).Msg("call")

	}(time.Now())

	err = svc.DeleteAllMenu(ctx.UserContext(), secretId)
	if err != nil {
		sendResponse(ctx, log.Logger, nil, err)
		return nil
	}

	sendResponse(ctx, log.Logger, nil, nil)
	return err
}

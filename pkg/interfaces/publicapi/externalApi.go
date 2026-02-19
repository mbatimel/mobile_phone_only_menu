// Package publicapi
// @tg version=0.0.1
// @tg backend=retention
// @tg title=`Retention`
// @tg servers=http://45.129.128.131:80|http://45.129.128.131:80
//
//go:generate tg transport --services . --out ../../../internal/transport/jsonRPC/externalapi --outSwagger ../../../swaggers/publicapi/swagger.yaml

package publicapi

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/mbatimel/mobile_phone_only_menu/internal/consts"
)

// PublicApi
// @tg http-server metrics log
// @tg http-prefix=/menu/api
type PublicApi interface {
	// CreateDish
	// @tg http-method=POST
	// @tg http-path=/create/dish
	// @tg http-cookies=secretId|x-secret-id
	// @tg http-response=github.com/mbatimel/mobile_phone_only_menu/internal/transport/http/custom-handlers:CreateDish
	// @tg summary=`Создание позиции в меню`
	// @tg desc=`создание позиции в меню`
	// @tg 200=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Resp200
	// @tg 500=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err500
	// @tg 400=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err400
	// @tg 403=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err403
	// @tg 405=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err405
	CreateDish(ctx context.Context, secretId uuid.UUID, dish string, category string) (err error)

	// MarkFavoriteDish
	// @tg http-method=POST
	// @tg http-path=mark
	// @tg http-cookies=secretId|x-secret-id
	// @tg http-response=github.com/mbatimel/mobile_phone_only_menu/internal/transport/http/custom-handlers:MarkFavoriteDish
	// @tg summary=`пометить позицию как нравится `
	// @tg desc=`пометить позицию как нравится `
	// @tg 200=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Resp200
	// @tg 500=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err500
	// @tg 400=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err400
	// @tg 403=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err403
	// @tg 405=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err405
	MarkFavoriteDish(ctx context.Context, secretId uuid.UUID, ids []uint64) (err error)

	// MarkUnFavoriteDish
	// @tg http-method=POST
	// @tg http-path=unmark
	// @tg http-cookies=secretId|x-secret-id
	// @tg http-response=github.com/mbatimel/mobile_phone_only_menu/internal/transport/http/custom-handlers:MarkUnFavoriteDish
	// @tg summary=`пометить позицию как не нравится `
	// @tg desc=`пометить позицию как не нравится `
	// @tg 200=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Resp200
	// @tg 500=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err500
	// @tg 400=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err400
	// @tg 403=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err403
	// @tg 405=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err405
	MarkUnFavoriteDish(ctx context.Context, secretId uuid.UUID, ids []uint64) (err error)

	// DeleteDish
	// @tg http-method=DELETE
	// @tg http-path=/delete
	// @tg http-cookies=secretId|x-secret-id
	// @tg http-response=github.com/mbatimel/mobile_phone_only_menu/internal/transport/http/custom-handlers:DeleteDish
	// @tg summary=`удалить позицию в меню`
	// @tg desc=`удалить позицию в меню`
	// @tg 200=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Resp200
	// @tg 500=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err500
	// @tg 400=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err400
	// @tg 403=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err403
	// @tg 405=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err405
	DeleteDish(ctx context.Context, secretId uuid.UUID, id uint64) (err error)

	// CreateChef
	// @tg http-method=POST
	// @tg http-path=/create/chef
	// @tg http-cookies=secretId|x-secret-id
	// @tg http-response=github.com/mbatimel/mobile_phone_only_menu/internal/transport/http/custom-handlers:CreateChef
	// @tg summary=`добавить шефа`
	// @tg desc=`добавить шефа`
	// @tg 200=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Resp200
	// @tg 500=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err500
	// @tg 400=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err400
	// @tg 403=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err403
	// @tg 405=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err405
	CreateChef(ctx context.Context, secretId uuid.UUID, name string) (err error)

	// DeleteChef
	// @tg http-method=DELETE
	// @tg http-path=/chef
	// @tg http-cookies=secretId|x-secret-id
	// @tg http-response=github.com/mbatimel/mobile_phone_only_menu/internal/transport/http/custom-handlers:DeleteChef
	// @tg summary=`удалсть шефа`
	// @tg desc=`удалить шефа`
	// @tg 200=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Resp200
	// @tg 500=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err500
	// @tg 400=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err400
	// @tg 403=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err403
	// @tg 405=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err405
	DeleteChef(ctx context.Context, secretId uuid.UUID) (err error)
	// GetChef
	// @tg http-method=GET
	// @tg http-path=/chef
	// @tg http-cookies=secretId|x-secret-id
	// @tg http-response=github.com/mbatimel/mobile_phone_only_menu/internal/transport/http/custom-handlers:GetChef
	// @tg summary=`удалсть шефа`
	// @tg desc=`удалить шефа`
	// @tg 200=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Resp200
	// @tg 500=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err500
	// @tg 400=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err400
	// @tg 403=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err403
	// @tg 405=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err405
	GetChef(ctx context.Context, secretId uuid.UUID) (name string, err error)

	// UpdateDish
	// @tg http-method=PUT
	// @tg http-path=/update
	// @tg http-cookies=secretId|x-secret-id
	// @tg http-response=github.com/mbatimel/mobile_phone_only_menu/internal/transport/http/custom-handlers:UpdateDish
	// @tg summary=`обновить поззицию в меню`
	// @tg desc=`обновить поззицию в меню`
	// @tg 200=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Resp200
	// @tg 500=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err500
	// @tg 400=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err400
	// @tg 403=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err403
	// @tg 405=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err405
	UpdateDish(ctx context.Context, secretId uuid.UUID, id uint64, text string, category string) (err error)

	// GetAllDish
	// @tg http-method=GET
	// @tg http-path=/all
	// @tg http-cookies=secretId|x-secret-id
	// @tg http-args=date|date
	// @tg http-response=github.com/mbatimel/mobile_phone_only_menu/internal/transport/http/custom-handlers:GetAllDish
	// @tg summary=`получить полный список позиций меню
	// @tg desc=`получить полный список позиций меню`
	// @tg 200=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Resp200
	// @tg 500=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err500
	// @tg 400=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err400
	// @tg 403=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err403
	// @tg 405=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err405
	GetAllDish(ctx context.Context, secretId uuid.UUID, date time.Time) (resp []consts.MenuDish, err error)

	// GetFavoriteDish
	// @tg http-method=GET
	// @tg http-path=/favorite
	// @tg http-cookies=secretId|x-secret-id
	// @tg http-args=date|date
	// @tg http-response=github.com/mbatimel/mobile_phone_only_menu/internal/transport/http/custom-handlers:GetFavoriteDish
	// @tg summary=`Получить сприсок желаймого`
	// @tg desc=`Получить сприсок желаймого`
	// @tg 200=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Resp200
	// @tg 500=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err500
	// @tg 400=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err400
	// @tg 403=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err403
	// @tg 405=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err405
	GetFavoriteDish(ctx context.Context, secretId uuid.UUID, date time.Time) (resp []consts.MenuDish, err error)

	// DeleteAllMenu
	// @tg http-method=DELETE
	// @tg http-path=/all
	// @tg http-cookies=secretId|x-secret-id
	// @tg http-response=github.com/mbatimel/mobile_phone_only_menu/internal/transport/http/custom-handlers:DeleteAllMenu
	// @tg summary=`опустошить меню`
	// @tg desc=`опустошить меню`
	// @tg 200=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Resp200
	// @tg 409=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err409
	// @tg 500=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err500
	// @tg 400=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err400
	// @tg 403=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err403
	// @tg 405=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err405
	DeleteAllMenu(ctx context.Context, secretId uuid.UUID) (err error)
}

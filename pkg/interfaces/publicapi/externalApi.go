// Package publicapi
// @tg version=0.0.1
// @tg backend=retention
// @tg title=`Retention`
// @tg servers=localhost:9000|localhost:9001
//
//go:generate tg transport --services . --out ../../../../internal/transport/http/  --outSwagger ../../../../swaggers/publicapi/swagger.yaml

package publicapi

import "context"

// PublicApi
// @tg http-server metrics log
// @tg http-prefix=/menu/api
type PublicApi interface {
	// CreateDish
	// @tg http-method=POST
	// @tg http-path=/create/dish
	// @tg http-response=github.com/mbatimel/mobile_phone_only_menu/internal/transport/http/custom-handlers:CreateDish
	// @tg summary=`Создание позиции в меню`
	// @tg desc=`создание позиции в меню`
	// @tg 200=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Resp200
	// @tg 500=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err500
	// @tg 400=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err400
	// @tg 403=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err403
	// @tg 405=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err405
	CreateDish(ctx context.Context, dish string, categoty string) (err error)

	// MarkFavoriteDish
	// @tg http-method=POST
	// @tg http-path=mark
	// @tg http-response=github.com/mbatimel/mobile_phone_only_menu/internal/transport/http/custom-handlers:MarkFavoriteDish
	// @tg summary=`пометить позицию как нравится `
	// @tg desc=`пометить позицию как нравится `
	// @tg 200=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Resp200
	// @tg 500=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err500
	// @tg 400=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err400
	// @tg 403=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err403
	// @tg 405=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err405
	MarkFavoriteDish(ctx context.Context, ids []uint64) (err error)

	// DeleteDish
	// @tg http-method=DELETE
	// @tg http-path=/delete
	// @tg http-response=github.com/mbatimel/mobile_phone_only_menu/internal/transport/http/custom-handlers:DeleteDish
	// @tg summary=`удалить позицию в меню`
	// @tg desc=`удалить позицию в меню`
	// @tg 200=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Resp200
	// @tg 500=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err500
	// @tg 400=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err400
	// @tg 403=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err403
	// @tg 405=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err405
	DeleteDish(ctx context.Context, id uint64) (err error)

	// CreateChef
	// @tg http-method=POST
	// @tg http-path=/create/chef
	// @tg http-response=github.com/mbatimel/mobile_phone_only_menu/internal/transport/http/custom-handlers:CreateChef
	// @tg summary=`добавить шефа`
	// @tg desc=`добавить шефа`
	// @tg 200=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Resp200
	// @tg 500=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err500
	// @tg 400=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err400
	// @tg 403=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err403
	// @tg 405=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err405
	CreateChef(ctx context.Context, name string) (err error)

	// UpdateDish
	// @tg http-method=UPDATE
	// @tg http-path=/update
	// @tg http-response=github.com/mbatimel/mobile_phone_only_menu/internal/transport/http/custom-handlers:UpdateDish
	// @tg summary=`обновить поззицию в меню`
	// @tg desc=`обновить поззицию в меню`
	// @tg 200=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Resp200
	// @tg 500=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err500
	// @tg 400=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err400
	// @tg 403=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err403
	// @tg 405=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err405
	UpdateDish(ctx context.Context, id uint64, text string) (err error)

	// GetAllDish
	// @tg http-method=GET
	// @tg http-path=/all
	// @tg http-response=github.com/mbatimel/mobile_phone_only_menu/internal/transport/http/custom-handlers:GetAllDish
	// @tg summary=`получить полный список позиций меню
	// @tg desc=`получить полный список позиций меню`
	// @tg 200=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Resp200
	// @tg 500=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err500
	// @tg 400=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err400
	// @tg 403=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err403
	// @tg 405=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err405
	GetAllDish(ctx context.Context) (resp *models.MenuDish, err error)

	// GetFavoriteDish
	// @tg http-method=GET
	// @tg http-path=/favorite
	// @tg http-response=github.com/mbatimel/mobile_phone_only_menu/internal/transport/http/custom-handlers:GetFavoriteDish
	// @tg summary=`Получить сприсок желаймого`
	// @tg desc=`Получить сприсок желаймого`
	// @tg 200=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Resp200
	// @tg 500=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err500
	// @tg 400=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err400
	// @tg 403=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err403
	// @tg 405=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err405
	GetFavoriteDish(ctx context.Context) (resp *models.MenuDish, err error)

	// DeleteAllMenu
	// @tg http-method=DELETE
	// @tg http-path=/all
	// @tg http-response=github.com/mbatimel/mobile_phone_only_menu/internal/transport/http/custom-handlers:DeleteAllMenu
	// @tg summary=`опустошить меню`
	// @tg desc=`опустошить меню`
	// @tg 200=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Resp200
	// @tg 409=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err409
	// @tg 500=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err500
	// @tg 400=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err400
	// @tg 403=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err403
	// @tg 405=github.com/mbatimel/mobile_phone_only_menu/swaggers/publicapi/models:Err405
	DeleteAllMenu(ctx context.Context) (err error)
}

package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/mbatimel/mobile_phone_only_menu/internal/consts"
	"github.com/mbatimel/mobile_phone_only_menu/internal/storage/postgres"
	"github.com/mbatimel/mobile_phone_only_menu/pkg/interfaces/publicapi"
	"github.com/rs/zerolog"
)

type menuDishService struct {
	logger  zerolog.Logger
	storage postgres.Storage

	serviceID uuid.UUID
}

func (m *menuDishService) CreateDish(ctx context.Context, secretId uuid.UUID, dish string, categoty string) (err error) {
	return m.storage.CreateDish(ctx, dish, categoty)
}

func (m *menuDishService) MarkFavoriteDish(ctx context.Context, secretId uuid.UUID, ids []uint64) (err error) {
	return m.storage.MarkFavoriteDish(ctx, ids)
}

func (m *menuDishService) DeleteDish(ctx context.Context, secretId uuid.UUID, id uint64) (err error) {
	return m.storage.DeleteDish(ctx, id)
}

func (m *menuDishService) CreateChef(ctx context.Context, secretId uuid.UUID, name string) (err error) {
	return m.storage.CreateChef(ctx, name)
}

func (m *menuDishService) UpdateDish(ctx context.Context, secretId uuid.UUID, id uint64, text string) (err error) {
	return m.storage.UpdateDish(ctx, id, text)
}

func (m *menuDishService) GetAllDish(ctx context.Context, secretId uuid.UUID) (resp []consts.MenuDish, err error) {
	return m.storage.GetAllDish(ctx)
}

func (m *menuDishService) GetFavoriteDish(ctx context.Context, secretId uuid.UUID) (resp []consts.MenuDish, err error) {
	return m.storage.GetFavoriteDish(ctx)
}

func (m *menuDishService) DeleteAllMenu(ctx context.Context, secretId uuid.UUID) (err error) {
	return m.storage.DeleteAllMenu(ctx)
}
func NewMenuService(
	logger zerolog.Logger,
	storage postgres.Storage,
	serviceID uuid.UUID,

) publicapi.PublicApi {

	return &menuDishService{
		logger:  logger,
		storage: storage,

		serviceID: serviceID,
	}
}

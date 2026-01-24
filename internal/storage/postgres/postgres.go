package postgres

import (
	"context"
	"fmt"

	"github.com/mbatimel/mobile_phone_only_menu/internal/config"
	models "github.com/mbatimel/mobile_phone_only_menu/internal/consts"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

//go:generate mockgen -source=interface.go -destination storage_mock.go -package postgres
type Storage interface {
	CreateDish(ctx context.Context, dish string, category string) (err error)

	MarkFavoriteDish(ctx context.Context, ids []uint64) (err error)
	MarkUnFavoriteDish(ctx context.Context, ids []uint64) (err error)

	DeleteDish(ctx context.Context, id uint64) (err error)

	CreateChef(ctx context.Context, name string) (err error)

	UpdateDish(ctx context.Context, id uint64, text string, category string) (err error)

	GetAllDish(ctx context.Context) (resp []models.MenuDish, err error)
	GetChef(ctx context.Context) (resp string, err error)

	GetFavoriteDish(ctx context.Context) (resp []models.MenuDish, err error)

	DeleteAllMenu(ctx context.Context) (err error)
	DeleteChef(ctx context.Context) (err error)
}
type storage struct {
	connectManager ConnectManager
}

func New(cfg config.Postgres, logger zerolog.Logger) (Storage, error) {
	connectManager, err := NewManager(logger, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "newManager error")
	}

	return &storage{
		connectManager: connectManager,
	}, nil
}
func (s *storage) CreateDish(ctx context.Context, dish string, category string) (err error) {
	masterConn := s.connectManager.GetConnection(MASTER)

	tx, err := masterConn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("tx.Begin error: %w", err)
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	return insertDish(ctx, tx, dish, category)
}

func (s *storage) MarkFavoriteDish(ctx context.Context, ids []uint64) (err error) {
	masterConn := s.connectManager.GetConnection(MASTER)

	tx, err := masterConn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("tx.Begin error: %w", err)
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()
	return updateFavoriteDish(ctx, tx, ids)
}
func (s *storage) MarkUnFavoriteDish(ctx context.Context, ids []uint64) (err error) {
	masterConn := s.connectManager.GetConnection(MASTER)

	tx, err := masterConn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("tx.Begin error: %w", err)
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()
	return updateUnFavoriteDish(ctx, tx, ids)
}

func (s *storage) DeleteDish(ctx context.Context, id uint64) (err error) {
	masterConn := s.connectManager.GetConnection(MASTER)

	tx, err := masterConn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("tx.Begin error: %w", err)
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()
	return deleteDish(ctx, tx, id)
}

func (s *storage) CreateChef(ctx context.Context, name string) (err error) {
	masterConn := s.connectManager.GetConnection(MASTER)

	tx, err := masterConn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("tx.Begin error: %w", err)
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()
	return insertChef(ctx, tx, name)
}

func (s *storage) UpdateDish(ctx context.Context, id uint64, text string, category string) (err error) {
	masterConn := s.connectManager.GetConnection(MASTER)

	tx, err := masterConn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("tx.Begin error: %w", err)
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()
	return updateDish(ctx, tx, id, text, category)
}

func (s *storage) GetAllDish(ctx context.Context) (resp []models.MenuDish, err error) {
	masterConn := s.connectManager.GetConnection(REPLICA)

	tx, err := masterConn.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("tx.Begin error: %w", err)
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()
	return selectAllDish(ctx, tx)
}
func (s *storage) GetChef(ctx context.Context) (resp string, err error) {
	masterConn := s.connectManager.GetConnection(REPLICA)

	tx, err := masterConn.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("tx.Begin error: %w", err)
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()
	return selectChef(ctx, tx)
}

func (s *storage) GetFavoriteDish(ctx context.Context) (resp []models.MenuDish, err error) {
	masterConn := s.connectManager.GetConnection(REPLICA)

	tx, err := masterConn.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("tx.Begin error: %w", err)
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()
	return selectFavoriteDish(ctx, tx)
}

func (s *storage) DeleteAllMenu(ctx context.Context) (err error) {
	masterConn := s.connectManager.GetConnection(MASTER)

	tx, err := masterConn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("tx.Begin error: %w", err)
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()
	return deleteAll(ctx, tx)
}
func (s *storage) DeleteChef(ctx context.Context) (err error) {
	masterConn := s.connectManager.GetConnection(MASTER)

	tx, err := masterConn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("tx.Begin error: %w", err)
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()
	return deleteChef(ctx, tx)
}

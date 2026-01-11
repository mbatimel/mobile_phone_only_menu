package postgres

import (
	"context"
	_ "embed"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/mbatimel/mobile_phone_only_menu/internal/consts"
)

const defaultTimeout = 3000 * time.Millisecond

//go:embed sql/insert_menu.sql
var sqlInsertMenu string

//go:embed sql/update_favorite_dish.sql
var sqlUpdateFavoriteDish string

//go:embed sql/update_unfavorite_dish.sql
var sqlUpdateUnFavoriteDish string

//go:embed sql/delete_dish.sql
var sqlDeleteDish string

//go:embed sql/insert_chef.sql
var sqlInsertChef string

//go:embed sql/update_dish.sql
var sqlUpdateDish string

//go:embed sql/select_all_dish.sql
var sqlSelectAllDish string

//go:embed sql/select_favorite_dish.sql
var sqlSelectFavoriteDish string

//go:embed sql/delete_all.sql
var sqlDeleteAll string

//go:embed sql/delete_chef.sql
var sqlDeleteChef string

func insertDish(ctx context.Context, conn pgx.Tx, dish string, categoty string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	_, err = conn.Exec(
		ctx,
		sqlInsertMenu,
		dish,
		categoty,
	)
	if err != nil {
		return fmt.Errorf("postgresql: %w", err)
	}
	err = conn.Commit(ctx)
	if err != nil {
		return fmt.Errorf("tx.Commit error: %w", err)
	}
	return nil
}

func updateFavoriteDish(ctx context.Context, conn pgx.Tx, ids []uint64) (err error) {
	nCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	_, err = conn.Exec(nCtx, sqlUpdateFavoriteDish, ids)
	if err != nil {
		return fmt.Errorf("postgresql: %w", err)
	}
	err = conn.Commit(ctx)
	if err != nil {
		return fmt.Errorf("tx.Commit error: %w", err)
	}
	return nil
}
func updateUnFavoriteDish(ctx context.Context, conn pgx.Tx, ids []uint64) (err error) {
	nCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	_, err = conn.Exec(nCtx, sqlUpdateUnFavoriteDish, ids)
	if err != nil {
		return fmt.Errorf("postgresql: %w", err)
	}
	err = conn.Commit(ctx)
	if err != nil {
		return fmt.Errorf("tx.Commit error: %w", err)
	}
	return nil
}

func deleteDish(ctx context.Context, conn pgx.Tx, id uint64) (err error) {
	nCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	_, err = conn.Exec(nCtx, sqlDeleteDish, id)
	if err != nil {
		return fmt.Errorf("postgresql: %w", err)
	}
	err = conn.Commit(ctx)
	if err != nil {
		return fmt.Errorf("tx.Commit error: %w", err)
	}
	return nil
}

func insertChef(ctx context.Context, conn pgx.Tx, name string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	_, err = conn.Exec(
		ctx,
		sqlInsertChef,
		name,
	)
	if err != nil {
		return fmt.Errorf("postgresql: %w", err)
	}
	err = conn.Commit(ctx)
	if err != nil {
		return fmt.Errorf("tx.Commit error: %w", err)
	}

	return nil
}

func updateDish(ctx context.Context, conn pgx.Tx, id uint64, text string) (err error) {
	nCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	_, err = conn.Exec(nCtx, sqlUpdateDish, text, id)
	if err != nil {
		return fmt.Errorf("postgresql: %w", err)
	}
	err = conn.Commit(ctx)
	if err != nil {
		return fmt.Errorf("tx.Commit error: %w", err)
	}
	return nil
}

func selectAllDish(ctx context.Context, conn pgx.Tx) ([]consts.MenuDish, error) {
	internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	rows, err := conn.Query(internalCtx, sqlSelectAllDish)
	if err != nil {
		return nil, fmt.Errorf("could not query: %w", err)
	}
	var result []consts.MenuDish
	for rows.Next() {
		res := consts.MenuDish{}
		err := rows.Scan(
			&res.Id,
			&res.Name,
			&res.Category,
			&res.Choice,
		)
		if err != nil {
			return nil, fmt.Errorf("could not scan: %w", err)
		}
		result = append(result, res)
	}
	return result, nil
}

func selectFavoriteDish(ctx context.Context, conn pgx.Tx) ([]consts.MenuDish, error) {
	internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	rows, err := conn.Query(internalCtx, sqlSelectFavoriteDish)
	if err != nil {
		return nil, fmt.Errorf("could not query: %w", err)
	}
	var result []consts.MenuDish
	for rows.Next() {
		res := consts.MenuDish{}
		err := rows.Scan(
			&res.Id,
			&res.Name,
			&res.Category,
			&res.Choice,
		)
		if err != nil {
			return nil, fmt.Errorf("could not scan: %w", err)
		}
		result = append(result, res)
	}
	return result, nil
}

func deleteAll(ctx context.Context, conn pgx.Tx) (err error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	_, err = conn.Exec(ctx, sqlDeleteAll)
	if err != nil {
		return fmt.Errorf("postgresql: %w", err)
	}
	err = conn.Commit(ctx)
	if err != nil {
		return fmt.Errorf("tx.Commit error: %w", err)
	}
	return nil
}
func deleteChef(ctx context.Context, conn pgx.Tx) (err error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	_, err = conn.Exec(ctx, sqlDeleteChef)
	if err != nil {
		return fmt.Errorf("postgresql: %w", err)
	}
	err = conn.Commit(ctx)
	if err != nil {
		return fmt.Errorf("tx.Commit error: %w", err)
	}
	return nil
}

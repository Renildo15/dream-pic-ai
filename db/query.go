package db

import (
	"context"
	"dreampicai/types"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func CreateImage(tx bun.Tx, image *types.Image) error {
	_, err := tx.NewInsert().Model(image).Exec(context.Background())
	return err
}

func UpdateImage(tx bun.Tx, image *types.Image) error {
	_, err := tx.NewUpdate().
		Model(image).
		WherePK().
		Exec(context.Background())
	return err
}
func GetImageByID(ID int) (types.Image, error) {
	var image types.Image
	err := Bun.NewSelect().Model(&image).Where("id = ?", ID).Scan(context.Background())
	return image, err
}

func GetImagesByBatchID(batchID uuid.UUID) ([]types.Image, error) {
	var images []types.Image
	err := Bun.NewSelect().
		Model(&images).
		Where("batch_id = ?", batchID).
		Scan(context.Background())
	return images, err
}

func GetImagesByUserID(userID uuid.UUID) ([]types.Image, error) {
	var images []types.Image
	err := Bun.NewSelect().
		Model(&images).
		Where("deleted = ?", false).
		Where("user_id = ?", userID).
		Order("created_at desc").
		Scan(context.Background())
	return images, err
}

func UpdateAccount(account *types.Account) error {
	_, err := Bun.NewUpdate().
		Model(account).
		WherePK().
		Exec(context.Background())
	return err
}

func GetAccountByUserID(userID uuid.UUID) (types.Account, error) {
	var account types.Account
	err := Bun.NewSelect().Model(&account).Where("user_id = ?", userID).Scan(context.Background())
	return account, err
}

func CreateAccount(account *types.Account) error {
	_, err := Bun.NewInsert().Model(account).Exec(context.Background())
	return err
}

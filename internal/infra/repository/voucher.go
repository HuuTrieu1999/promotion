package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"promotion/internal/core/entity"
	"promotion/internal/core/port/repository"
)

type voucherRepo struct {
	col *mongo.Collection
}

func (u voucherRepo) Insert(voucherEntity entity.VoucherEntity) error {
	_, err := u.col.InsertOne(context.Background(), voucherEntity)
	return err
}

func NewVoucherRepo(col *mongo.Collection) repository.VoucherRepository {
	return &voucherRepo{
		col: col,
	}
}

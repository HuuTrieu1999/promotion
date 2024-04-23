package repository

import (
	"promotion/internal/core/entity"
)

type VoucherRepository interface {
	Insert(voucherEntity entity.VoucherEntity) error
}

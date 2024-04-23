package voucher

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"promotion/internal/core/entity"
	"promotion/internal/core/port/repository"
	"promotion/internal/core/port/service"
	"time"
)

type voucherService struct {
	voucherRepo repository.VoucherRepository
}

func (v voucherService) CreateDiscountVoucherForTopupFee(userID string, discount int) error {
	return v.voucherRepo.Insert(entity.VoucherEntity{
		VoucherID:   primitive.NewObjectID(),
		UserID:      userID,
		Discount:    discount,
		Description: "",
		ExpireDate:  time.Now().Add(time.Hour * 24 * 30),
		VoucherType: "TOPUP_MOBILE_PHONE_FEE",
	})
}

func NewVoucherService(voucherRepo repository.VoucherRepository) service.VoucherService {
	return &voucherService{
		voucherRepo: voucherRepo,
	}
}

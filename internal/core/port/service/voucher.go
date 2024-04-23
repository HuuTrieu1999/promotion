package service

type VoucherService interface {
	CreateDiscountVoucherForTopupFee(userID string, discount int) error
}

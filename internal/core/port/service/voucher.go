package service

type VoucherService interface {
	CreateDiscountVoucher(userID string, discount int) error
}

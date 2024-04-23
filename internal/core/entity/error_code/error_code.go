package error_code

type ErrorCode string

// error code
const (
	Success                 ErrorCode = "SUCCESS"
	InvalidRequest          ErrorCode = "INVALID_REQUEST"
	InternalError           ErrorCode = "INTERNAL_ERROR"
	WrongPasswordError                = "WRONG_PASSWORD"
	InvalidEmailError                 = "INVALID_EMAIL"
	InvalidPhoneNumberError           = "INVALID_PHONE_NUMBER"
	InvalidUserNameError              = "INVALID_USER_NAME"
	AccountNotExistError              = "ACCOUNT_NOT_EXIST"
	InvalidAccountError               = "INVALID_ACCOUNT"
	InvalidPasswordError              = "INVALID_PASSWORD"
	InvalidFullNameError              = "INVALID_FULL_NAME"
	InvalidBirthdayError              = "INVALID_BIRTHDAY"
	AccountExistError                 = "ACCOUNT_EXIST"
)

// error message
const (
	SuccessErrMsg         = "success"
	InternalErrMsg        = "internal error"
	InvalidRequestErrMsg  = "invalid request"
	WrongPasswordMsg      = "wrong password"
	InvalidEmailMsg       = "invalid email"
	InvalidPhoneNumberMsg = "invalid phone number"
	InvalidUserNameMsg    = "invalid campaign name"
	AccountNotExistMsg    = "account not exist"
	InvalidAccountMsg     = "invalid account"
	InvalidPasswordMsg    = "invalid password"
	InvalidFullNameMsg    = "invalid full name"
	InvalidBirthdayMsg    = "invalid birthday"
	AccountExistMsg       = "account exist"
)

package serviceValidator

type PaymentURLPaystack struct {
	Amount      string `json:"amount"       binding:"required"`
	Currency    string `json:"currency"     binding:"required"`
	CallbackUrl string `json:"callback_url" binding:"required, default="`
	Email       string `json:"email"        binding:"required, email"`
	Reference   string `json:"reference"`
}

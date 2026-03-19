package micha

type Invoice struct {
	Title          string
	Description    string
	StartParameter string
	Currency       string
	TotalAmount    int
}

type OrderInfo struct {
	Name            string           `json:"name,omitempty"`
	PhoneNumber     string           `json:"phone_number,omitempty"`
	Email           string           `json:"email,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

type SuccessfulPayment struct {
	// TODO
}

type ShippingAddress struct {
	// TODO
}

type PassportData struct {
	// TODO
}

type ShippingQuery struct {
	// TODO
}

type PreCheckoutQuery struct {
	ID               string     `json:"id"`
	From             User       `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int64      `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionID string     `json:"shipping_option_id,omitempty"`
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`
}

package args

// ReceiptArgs is the default args for receipt
type ReceiptArgs struct {
	SharedArgs
	OrderNumber   string   `json:"order_number"`
	OrderCreated  string   `json:"order_created"`
	PaymentOption string   `json:"payment_option"`
	Items         []Item   `json:"items"`
	Total         Total    `json:"total"`
	Options       []Option `json:"options"`
}

// Item ...
type Item struct {
	ProductName  string `json:"product_nickname"`
	ProductQty   int    `json:"product_qty"`
	ProductPrice string `json:"product_price"`
}

// Total ...
type Total struct {
	SubTotal            string  `json:"order_subtotal"`
	TaxRate             *string `json:"order_tax_rate,omitempty"`
	Tax                 *string `json:"order_tax,omitempty"`
	Discount            *string `json:"order_discount,omitempty"`
	Total               string  `json:"order_total"`
	OrderShippingOption *string `json:"order_shipping_option,omitempty"`
	OrderShipping       *string `json:"order_shipping,omitempty"`
}

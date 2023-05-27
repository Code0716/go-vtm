package domain

import "time"

// Invoice is model
type Invoice struct {
	ID            string     `json:"id"`
	UserID        string     `json:"userId"`
	AuthorizerID  string     `json:"authorizerId"`
	BillingDate   string     `json:"billingDate"`
	BillingAmount *int       `json:"billingAmount,omitempty"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     *time.Time `json:"deletedAt,omitempty"`
}

// TableName GORMにテーブル名を教える
func (i Invoice) TableName() string {
	return "invoice"
}

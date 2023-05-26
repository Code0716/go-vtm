package domain

import "time"

type Address struct {
	ID        string     `json:"id"`
	PostCode  string     `json:"postCode"`
	Address   string     `json:"address"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

// TableName GORMにテーブル名を教える
func (i Address) TableName() string {
	return "address"
}

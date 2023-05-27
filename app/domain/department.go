package domain

import "time"

// Department struct
type Department struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	AddressID string     `json:"addressId"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

// TableName GORMにテーブル名を教える
func (i Department) TableName() string {
	return "department"
}

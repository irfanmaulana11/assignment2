package models

type Items struct {
	ID          int    `json:"id,omitempty" gorm:"type:int;primary_key;auto_increment;not_null"`
	ItemCode    string `json:"item_code,omitempty" gorm:"type:varchar(10)"`
	Description string `json:"description,omitempty" gorm:"type:varchar(100)"`
	Quantity    int    `json:"quantity,omitempty" gorm:"type:int"`
	OrderID     int    `json:"order_id,omitempty" gorm:"type:int"`
}

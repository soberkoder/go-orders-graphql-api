package models

// Order model
type Order struct {
    ID           int     `json:"id" gorm:"primary_key"`
    CustomerName string  `json:"customerName"`
    OrderAmount  float64 `json:"orderAmount"`
    Items        []Item  `json:"items" gorm:"foreignkey:OrderID"`
}

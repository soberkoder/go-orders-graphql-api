package models;

// Item model
type Item struct {
    ID          int `json:"id" gorm:"primary_key"`
    ProductCode string `json:"productCode"`
    ProductName string `json:"productName"`
    Quantity    int    `json:"quantity"`
    OrderID     int    `json:"-"`
}

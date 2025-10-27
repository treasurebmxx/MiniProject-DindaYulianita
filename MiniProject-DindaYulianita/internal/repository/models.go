package repository

import "time"

type User struct {
    ID        uint   `gorm:"primaryKey" json:"id"`
    Name      string `json:"name"`
    Email     string `gorm:"uniqueIndex;size:100" json:"email"`
    Phone     string `gorm:"uniqueIndex;size:20" json:"phone"`
    Password  string `json:"-"`
    IsAdmin   bool   `json:"is_admin" gorm:"default:false"`
    CreatedAt time.Time
    UpdatedAt time.Time
    Store     Store  `gorm:"foreignKey:OwnerID"`
    StoreID   *uint
}

type Store struct {
    ID        uint   `gorm:"primaryKey" json:"id"`
    Name      string `json:"name"`
    OwnerID   uint   `json:"owner_id"`
    CreatedAt time.Time
    UpdatedAt time.Time
    Products  []Product `gorm:"foreignKey:StoreID"`
}

type Address struct {
    ID        uint   `gorm:"primaryKey" json:"id"`
    UserID    uint   `json:"user_id"`
    Street    string `json:"street"`
    City      string `json:"city"`
    Province  string `json:"province"`
    Postal    string `json:"postal"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

type Category struct {
    ID        uint   `gorm:"primaryKey" json:"id"`
    Name      string `json:"name" gorm:"uniqueIndex"`
    CreatedAt time.Time
    UpdatedAt time.Time
    Products  []Product `gorm:"foreignKey:CategoryID"`
}

type Product struct {
    ID          uint    `gorm:"primaryKey" json:"id"`
    StoreID     uint    `json:"store_id"`
    CategoryID  *uint   `json:"category_id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    ImageURL    string  `json:"image_url"`
    Stock       int     `json:"stock"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

type Transaction struct {
    ID         uint    `gorm:"primaryKey" json:"id"`
    UserID     uint    `json:"user_id"`
    StoreID    uint    `json:"store_id"`
    TotalPrice float64 `json:"total_price"`
    CreatedAt  time.Time
    UpdatedAt  time.Time
    Items      []TransactionItem `gorm:"foreignKey:TransactionID"`
}

type TransactionItem struct {
    ID            uint    `gorm:"primaryKey" json:"id"`
    TransactionID uint    `json:"transaction_id"`
    ProductID     uint    `json:"product_id"`
    ProductName   string  `json:"product_name"`
    Price         float64 `json:"price"`
    Quantity      int     `json:"quantity"`
    CreatedAt     time.Time
}

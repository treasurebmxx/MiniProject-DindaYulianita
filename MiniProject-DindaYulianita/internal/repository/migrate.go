package repository

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) error {
    return db.AutoMigrate(
        &User{},
        &Store{},
        &Address{},
        &Category{},
        &Product{},
        &Transaction{},
        &TransactionItem{},
    )
}

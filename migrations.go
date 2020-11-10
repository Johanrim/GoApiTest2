package main

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var testMigrations = []*gormigrate.Migration{
	{
		ID: "202011101700",
		Migrate: func(tx *gorm.DB) (err error) {
			book := new(Book)
			category := new(Category)

			tx.LogMode(true)
			err = tx.AutoMigrate(book, category).Error
			if err != nil {
				return
			}

			err = tx.Model(book).AddForeignKey("category_id", "category(id)", "CASCADE", "CASCADE").Error
			if err != nil {
				return
			}

			tx.LogMode(false)
			return
		},
		Rollback: func(tx *gorm.DB) (err error) {
			book := new(Book)
			category := new(Category)

			tx.LogMode(true)
			err = tx.Model(book).RemoveForeignKey("category_id", "category(id)").Error
			if err != nil {
				return
			}
			err = tx.DropTableIfExists(book, category).Error
			if err != nil {
				return
			}
			tx.LogMode(false)
			return
		},
	},
}

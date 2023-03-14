package database

import "gorm.io/gorm"

type Product struct {
	ID    string
	Name  string
	Price float64
}

type Person struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Phone string

	/*
		HasOne
	*/
	Document Document

	/*
		BelongsTo
	*/
	AddressID int
	Address   Address

	/*
		CreatedAt, UpdatedAt, DeletedAt
	*/
	gorm.Model
}

type Address struct {
	ID       int `gorm:"primaryKey"`
	Location string
}

type Document struct {
	ID             int `gorm:"primaryKey"`
	Number         string
	DocumentTypeID int
	DocumentType   DocumentType
	PersonID       int
}

type DocumentType struct {
	ID        int `gorm:"primaryKey"`
	Value     string
	Documents []Document
}

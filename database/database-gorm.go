package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func RunDatabaseGorm() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Person{}, &Address{}, &Document{}, &DocumentType{})

	//crud(db)
	//belogsTo(db)
	//hasOne(db)
	//hasMany(db)
	//manyToMany(db) ???
	pessimisticLocking(db)
}

func pessimisticLocking(db *gorm.DB) {
	tx := db.Begin()

	var d DocumentType
	err := tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&d, 1).Error
	if err != nil {
		panic(err)
	}
	// TODO realizar as alterações
	tx.Debug().Commit()
}

func hasMany(db *gorm.DB) {
	tx := db.Begin()

	address := Address{
		Location: "Rua Washington Luiz, 210E, bairro São Cristóvão, Chapecó/SC",
	}

	if err := tx.Debug().Create(&address).Error; err != nil {
		tx.Debug().Rollback()
		panic(err)
	}

	person := Person{
		Name:      "Edson",
		Phone:     "46991034819",
		AddressID: address.ID,
	}

	if err := tx.Debug().Create(&person).Error; err != nil {
		tx.Debug().Rollback()
		panic(err)
	}

	documentType := DocumentType{
		Value: "CPF",
	}

	if err := tx.Debug().Create(&documentType).Error; err != nil {
		tx.Debug().Rollback()
		panic(err)
	}

	document := Document{
		Number:         "06714214928",
		DocumentTypeID: documentType.ID,
		PersonID:       person.ID,
	}

	if err := tx.Debug().Create(&document).Error; err != nil {
		tx.Debug().Rollback()
		panic(err)
	}

	address = Address{
		Location: "Rua Bolivia, 1172, bairro Luther King, Francisco Beltrao/PR",
	}

	if err := tx.Debug().Create(&address).Error; err != nil {
		tx.Debug().Rollback()
		panic(err)
	}

	person = Person{
		Name:      "Teresinha",
		Phone:     "46991176565",
		AddressID: address.ID,
	}

	if err := tx.Debug().Create(&person).Error; err != nil {
		tx.Debug().Rollback()
		panic(err)
	}

	document = Document{
		Number:         "12345678909",
		DocumentTypeID: documentType.ID,
		PersonID:       person.ID,
	}

	if err := tx.Debug().Create(&document).Error; err != nil {
		tx.Debug().Rollback()
		panic(err)
	}

	if err := tx.Debug().Commit().Error; err != nil {
		tx.Debug().Rollback()
		panic(err)
	}

	tx = db.Begin(&sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  true,
	})

	fmt.Println("Searching all")

	var people1 []Person
	tx.Debug().Preload("Address").Preload("Document").Find(&people1)
	for _, person := range people1 {
		fmt.Println(person)
	}

	fmt.Println("Searching all documents")

	var documentTypes []DocumentType
	tx.Debug().Model(&DocumentType{}).Preload("Documents").Find(&documentTypes)
	for _, documentType := range documentTypes {
		fmt.Printf("Type %v has documents:\n", documentType.Value)
		for _, document := range documentType.Documents {
			fmt.Printf("Document: %v\n", document)
		}
	}

	if err := tx.Debug().Commit().Error; err != nil {
		tx.Debug().Rollback()
		panic(err)
	}
}

func hasOne(db *gorm.DB) {
	tx := db.Begin()

	address := Address{
		Location: "Rua Washington Luiz, 210E, bairro São Cristóvão, Chapecó/SC",
	}

	if err := tx.Create(&address).Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	person := Person{
		Name:      "Edson",
		Phone:     "46991034819",
		AddressID: address.ID,
	}

	if err := tx.Create(&person).Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	documentType := DocumentType{
		Value: "CPF",
	}

	if err := tx.Create(&documentType).Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	document := Document{
		Number:         "06714214928",
		DocumentTypeID: documentType.ID,
		PersonID:       person.ID,
	}

	if err := tx.Create(&document).Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	tx = db.Begin(&sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  true,
	})

	fmt.Println("Searching all")

	var people1 []Person
	tx.Preload("Address").Preload("Document").Find(&people1)
	for _, person := range people1 {
		fmt.Println(person)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

func belogsTo(db *gorm.DB) {
	tx := db.Begin()

	address := Address{
		Location: "Rua Washington Luiz, 210E, bairro São Cristóvão, Chapecó/SC",
	}

	if err := tx.Create(&address).Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	person := Person{
		Name:      "Edson",
		Phone:     "46991034819",
		AddressID: address.ID,
	}

	if err := tx.Create(&person).Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	tx = db.Begin(&sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  true,
	})

	fmt.Println("Searching all")

	var people1 []Person
	tx.Preload("Address").Find(&people1)
	for _, person := range people1 {
		fmt.Println(person)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

func crud(db *gorm.DB) {
	tx := db.Begin()

	fmt.Println("Deleting all people")

	if err := tx.Debug().Delete(&Person{}, "id > 0").Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	fmt.Println("Creating person Edson")

	err := tx.Debug().Create(&Person{
		Name:  "Edson",
		Phone: "46991034819",
	}).Error

	if err != nil {
		tx.Rollback()
		panic(err)
	}

	fmt.Println("Creating people Irla, Luiza and Lorena")

	err = tx.Debug().Create(&[]Person{
		{Name: "Irla", Phone: "49991115566"},
		{Name: "Luiza", Phone: "49985235566"},
		{Name: "Lorena", Phone: "49123456789"},
	}).Error
	if err != nil {
		tx.Debug().Rollback()
		panic(err)
	}

	if err = tx.Debug().Commit().Error; err != nil {
		panic(err)
	}

	tx = db.Begin(&sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  true,
	})

	//err = tx.Create(&Person{
	//	Name:  "ERR",
	//	Phone: "ERR",
	//}).Error
	//if err != nil {
	//	tx.Rollback()
	//	panic(err)
	//}

	fmt.Println("Searching by ID")

	var person Person
	tx.Debug().First(&person, 14)
	fmt.Println(person)

	fmt.Println("Searching by name")

	tx.Debug().First(&person, "name = ?", "Irla")
	fmt.Println(person)

	fmt.Println("Searching all")

	var people1 []Person
	tx.Debug().Find(&people1)
	for _, person := range people1 {
		fmt.Println(person)
	}

	fmt.Println("Searching by limit and offset")

	var people2 []Person
	tx.Debug().Limit(2).Offset(2).Find(&people2)
	for _, person := range people2 {
		fmt.Println(person)
	}

	fmt.Println("Searching by Phone")

	var people3 []Person
	tx.Debug().Where("phone = ?", "49123456789").Find(&people3)
	for _, person := range people3 {
		fmt.Println(person)
	}

	fmt.Println("Searching by Name like")

	var people4 []Person
	tx.Debug().Where("name LIKE ?", "%ore%").Find(&people4)
	for _, person := range people4 {
		fmt.Println(person)
	}

	if err = tx.Debug().Commit().Error; err != nil {
		tx.Debug().Rollback()
		panic(err)
	}

	tx = db.Begin()

	person = people1[0]

	fmt.Printf("Changing name from %v to Lorelai\n", person.Name)

	person.Name = "Lorelai"

	tx.Debug().Save(person)

	person = people1[2]

	fmt.Printf("Deleting person %v", person.Name)

	tx.Debug().Delete(&person)

	if err = tx.Debug().Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

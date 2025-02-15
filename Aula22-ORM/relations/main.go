package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Name    string
	TutorID uint `gorm:"index"`
	Tutor   Tutor
}

type Tutor struct {
	gorm.Model
	Name  string
	Phone string
	Pets  []Pet `gorm:"foreignKey:TutorID"`
}

/*
* Nesse ponto acho importante explicar alguns tipos de relacionamentos que o GORM suporta:
* 1. One-To-One
	Esse tipo de relacionamento é quando uma entidade está relacionada a apenas uma outra entidade.
	Exemplo: Um tutor tem um endereço.
* 2. One-To-Many
*   Esse tipo de relacionamento é quando uma entidade está relacionada a várias outras entidades.
*   Exemplo: Um tutor tem vários pets.
* 3. Many-To-Many
    Esse tipo de relacionamento é quando várias entidades estão relacionadas a várias outras entidades.
	Exemplo: Um tutor tem vários pets e um pet tem vários tutores.
* 4. Belongs-To
	Esse tipo de relacionamento é quando uma entidade pertence a outra entidade.
	Exemplo: Um pet pertence a um tutor.
*/

func main() {
	databaseConnectionString := "aula22:aula22@tcp(localhost:3306)/aula22?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(databaseConnectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Migrator().DropTable(&Pet{}, &Tutor{})
	db.AutoMigrate(&Pet{}, &Tutor{})

	// Create a Tutor
	matheus := Tutor{Name: "Matheus", Phone: "123456789"}
	db.Create(&matheus)

	// Create a Pet
	thor := Pet{Name: "Thor", TutorID: matheus.ID}
	db.Create(&thor)

	mel := Pet{Name: "Mel", TutorID: matheus.ID}
	db.Create(&mel)

	var tutor Tutor
	db.Preload("Pets").First(&tutor, "id = ?", matheus.ID)

	tutor.about()
}

func (tutor Tutor) about() {
	fmt.Println("Tutor:", tutor.Name)
	fmt.Println("Phone:", tutor.Phone)
	fmt.Println("Pets:")
	for _, pet := range tutor.Pets {
		fmt.Println("  ->", pet.Name)
	}
	fmt.Println("---------------------------------")
}

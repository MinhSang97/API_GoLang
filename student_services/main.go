package main

import (
	"app/dbutil"
	"app/payload"
	"app/repo/mysql"
	"context"
	"fmt"
	"os"
)

func main() {

	eric := payload.Student{}

	JSONString := `{"FirstName":"Eric","ho":"Nguyen","tuoi":0,"Grade":0,"ClassName":""}`
	eric.FromJson(JSONString)

	db := dbutil.ConnectDB()

	repo := mysql.NewStudentRepository(db)
	student := eric.ToModel()
	err := repo.InsertOne(context.Background(), &student)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(student)

	students, err := repo.GetAll(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(students)

}

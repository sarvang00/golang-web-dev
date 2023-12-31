package main

import (
	"fmt"
	"html/template"
	"os"
)

type UserMeta struct {
	Visits int
}

type User struct {
	Name string
	Age  int
	Meta UserMeta
	Bio1 string
	Bio2 template.HTML
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "John Doe",
		Age:  12,
		Meta: UserMeta{
			Visits: 3,
		},
		Bio1: `<script>alert("Haha, you have been h4x0r3d!");</script>`,
		Bio2: `<script>alert("Haha, you have been h4x0r3d!");</script>`,
	}

	user2 := struct {
		Name string
		Age  int
	}{
		Name: "Jane Doe",
	}

	fmt.Println("User meta visits: ", user.Meta.Visits)

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, user2)
	if err != nil {
		panic(err)
	}
}

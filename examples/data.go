package examples

import (
	"encoding/json"
	"fmt"
)

func TestData1() {
	type Foo struct {
		Name     string
		Age      int
		NickName string
	}
	realFoo := Foo{"Pasha", 35, "Goshnik"}

	fakeFoo := struct {
		Name     string
		Age      int
		NickName string
	}{"Pasha", 35, "Goshnik"}

	compareStruct(realFoo, fakeFoo)
}

// false
func compareStruct(a, b interface{}) {
	fmt.Println(a == b)
}

type MyData struct {
	One int    `json:"one"`
	Two string `json:"two"`
}

// examples.MyData{One:1, Two:"two"}
// {"one":1,"two":"two"}
// examples.MyData{One:1, Two:"two"}
func TestData2() {
	in := MyData{1, "two"}

	fmt.Printf("%#v\n", in) // что отобразится после вызова?
	encoded, _ := json.Marshal(in)
	fmt.Println(string(encoded)) // что отобразится после вызова?
	var out MyData
	json.Unmarshal(encoded, &out)
	fmt.Printf("%#v\n", out) // что отобразится после вызова?
}

type data struct {
	name string
	age  int
}

func (d data) displayName() {
	fmt.Println("My name is", d.name)
}

func (d *data) setAge(age int) {
	d.age = age
	fmt.Println(d.name, "is age", d.age)
}

func TestData4() {
	d := data{
		name: "Paul",
	}
	f1 := d.displayName
	f1()
	// нет указетеля на display name, два раза вывыдетеся paul
	d.name = "Max"
	f1()
	// нет указетеля на display name, два раза вывыдетеся paul

	f2 := d.setAge
	f2(35)
	d.name = "Ian"
	f2(40)
}

type user struct {
	id   int
	name string
}

// указатели не равны, false
func TestData5() {
	u1 := getUser()
	u2 := getUser()
	fmt.Println(u1 == u2)
}

func getUser() *user {
	return &user{1000, "Misha"}
}

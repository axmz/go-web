package main

import (
	"os"
	"text/template"
)

type person struct {
	Fname string
	Lname string
}

type car struct {
	Model string
	Doors int
}

type data struct {
	People []person
	Cars   []car
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p1 := person{"alex", "moroz"}
	p2 := person{"alexa", "m"}
	people := []person{p1, p2}

	c1 := car{"vw", 5}
	c2 := car{"bmw", 3}
	cars := []car{c1, c2}

	d := data{people, cars}

	tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", d)
}

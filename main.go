package main

import (
	"fmt"
)

type person struct {
	name string
	family string
	age int
}

func main()  {
	fmt.Println("LOVE GOLANG")
 example :=  person{
  	name:  "AmirAbbas",
  	family : "KhoshBayan",
  	age : 21,
  }
   fmt.Println(example)
   fmt.Printf("i am %s and Love golang language,my family is %s,The best age to learn %d",example.name,example.family,example.age)
}

func (p *person)  test(){
	ptest := new(person)
	ptest.name = "Hossain"
	ptest.family = "KhoshBayan"
	ptest.age = 25
}

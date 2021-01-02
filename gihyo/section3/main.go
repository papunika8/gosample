package main

import (
	"fmt"
)

type Person struct {
   firstName string 
   age int
}

func (p Person) intro(greetings string) string{
    return greetings + " I am " + p.firstName
}

func main(){
  bob := Person{"Bob", 30}
  fmt.Println(bob.intro("Hello")) //=> Hello I am Bob
}
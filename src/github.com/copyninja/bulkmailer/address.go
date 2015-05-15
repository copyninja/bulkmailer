package main

import (
	"fmt"
	"strings"
)

type Address struct {
	FirstName, LastName string
	Email               string
}

func (a Address) String() string {
	return fmt.Sprintf("%s %s <%s>",
		a.FirstName, a.LastName, a.Email)
}

func NewAddress(value string) Address {
	tmp := strings.Split(value, " ")
	tmpAddress := tmp[len(tmp)-1]
	fmt.Println(tmpAddress)
	fmt.Println(tmpAddress[1:len(tmpAddress)-1])
	return Address{tmp[0], strings.Join(tmp[1:len(tmp)-1], " "),
		tmpAddress[1 : len(tmpAddress)-1]}
}

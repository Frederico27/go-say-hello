package gosayhello

import "strconv"

func SayHello(name string, age int) string {
	return "Hello " + name + " Your age is " + strconv.Itoa(age)
}

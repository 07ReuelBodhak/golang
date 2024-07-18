package main

import (
	"errors"
	"fmt"
	"time"
)

// Basic Error
type error interface {
	Error() string
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a/b, nil
}

// Custom Error

type MyError struct {
	when time.Time
	what string
}

func (e *MyError) Error() string{
	return fmt.Sprintf("%v, %s",e.when,e.what)
}

func checkPassword(pass int)(string,error){
	if pass != 123{
		return "none", &MyError{
			when: time.Now(),
			what: "invalid password",
		}
	}
	return "welcome back", nil
}

func main() {
	a,b := 32,4
	res,err := divide(a,b)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(res)
	}

	if msg,err := checkPassword(125); err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(msg)
	}
}
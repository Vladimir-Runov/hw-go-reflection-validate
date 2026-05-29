package main

import (
    "errors"
    "fmt" )

type User struct {
    Name  string `validate:"min=3"`
    Age   int    `validate:"min=18;max=65"`
    Email string `validate:"regexp=^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"`
}

// Validate(v interface{}) error, которая проверяет структуру на соответствие правилам, заданным через теги. 
func Validate(v interface{}) error {
    return nil
}



func main() {
}
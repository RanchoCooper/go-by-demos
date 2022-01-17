package main

import (
    "fmt"

    "github.com/go-playground/validator/v10"
)

/**
 * @author Rancho
 * @date 2022/1/17
 */

type User struct {
    Username string `validate:"min=6,max=10"`
    Age      uint8  `validate:"gte=1,lte=10"`
    Sex      string `validate:"oneof=female male"`
}

func main() {
    validate := validator.New()

    user1 := User{Username: "asong", Age: 11, Sex: "null"}
    err := validate.Struct(user1)
    if err != nil {
        fmt.Println(err)
    }
}

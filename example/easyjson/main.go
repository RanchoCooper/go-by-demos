package main

import (
    "fmt"

    "github.com/mailru/easyjson"
    "github.com/mailru/easyjson/jwriter"
)

/**
 * @author Rancho
 * @date 2021/12/16
 */

type T struct {
    Name string
    Age  int
}

func (t T) MarshalEasyJSON(w *jwriter.Writer) {
}

func main() {
    t := T{
        Name: "rancho",
        Age:  20,
    }

    fmt.Println(easyjson.Marshal(&t))
}

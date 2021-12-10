package main

import (
    "fmt"
    "net/http"

    "github.com/fatih/structs"
)

/**
 * @author Rancho
 * @date 2021/12/11
 */

type Server struct {
    Name        string `json:"name,omitempty"`
    ID          int
    Enabled     bool
    users       []string // not exported
    http.Server          // embedded
}

func main() {
    server := &Server{
        Name:    "gopher",
        ID:      123456,
        Enabled: true,
    }
    // Convert a struct to a map[string]interface{}
    // => {"Name":"gopher", "ID":123456, "Enabled":true}
    m := structs.Map(server)
    fmt.Println(m)

    // Convert the values of a struct to a []interface{}
    // => ["gopher", 123456, true]
    values := structs.Values(server)
    fmt.Println(values)

    // Convert the names of a struct to a []string
    // (see "Names methods" for more info about fields)
    names := structs.Names(server)
    fmt.Println(names)

    // Convert the values of a struct to a []*Field
    // (see "Field methods" for more info about fields)
    fields := structs.Fields(server)
    fmt.Println(fields)

    // Return the struct name => "Server"
    name := structs.Name(server)
    fmt.Println(name)

    // Check if any field of a struct is initialized or not.
    fmt.Println(structs.HasZero(server))

    // Check if all fields of a struct is initialized or not.
    fmt.Println(structs.IsZero(server))

    // Check if server is a struct or a pointer to struct
    fmt.Println(structs.IsStruct(server))
}

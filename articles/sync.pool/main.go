package main

import (
    "sync"
)

/**
 * @author Rancho
 * @date 2021/12/6
 */

type Person struct {
    Name string
    Age  int
}

var personPool = sync.Pool{
    New: func() interface{} { return new(Person) },
}

// https://medium.com/swlh/go-the-idea-behind-sync-pool-32da5089df72
func main() {
    // Get hold of an instance
    newPerson := personPool.Get().(*Person)
    defer personPool.Put(newPerson)

    // Using the instance
    newPerson.Name = "Jack"
}

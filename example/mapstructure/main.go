package main

import (
    "encoding/json"
    "fmt"

    "github.com/mitchellh/mapstructure"
)

type Outer struct {
    Type      int         `json:"type"`
    RawConfig interface{} `json:"config"`
    Config1   ConfigType1 `json:"-"`
    // Config2   ConfigType2 `json:"-"`
    // Config3   ConfigType3 `json:"-"`
}

type ConfigType1 struct {
    A int
    B int
}

func main() {
    // emulate
    input1 := "{\"type\":1,\"config\":{\"A\":1,\"B\":2}}"
    outer := &Outer{}
    _ = json.Unmarshal([]byte(input1), outer)

    // parse rawConfig to Config1 field
    if outer.Type == 1 {
        configType1 := ConfigType1{}
        mapstructure.Decode(outer.RawConfig, &configType1)
        outer.Config1 = configType1
        fmt.Println(outer.Config1.A)
    }

    if outer.Type == 2 {
        // handle other type
    }

}

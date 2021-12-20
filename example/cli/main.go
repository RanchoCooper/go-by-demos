package main

import (
    "os"

    "github.com/urfave/cli/v2"
)

/**
 * @author Rancho
 * @date 2021/12/20
 */

func main() {
    (&cli.App{}).Run(os.Args)
}

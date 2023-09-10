# Discord Emoji Map

[![Build status](https://github.com/Bios-Marcel/discordemojimap/actions/workflows/go.yml/badge.svg)](https://github.com/Bios-Marcel/discordemojimap/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/Bios-Marcel/discordemojimap?status.svg)](https://pkg.go.dev/github.com/Bios-Marcel/discordemojimap/v2)
[![codecov](https://codecov.io/gh/Bios-Marcel/discordemojimap/branch/master/graph/badge.svg)](https://codecov.io/gh/Bios-Marcel/discordemojimap)

This library allows you to substitute Discord emoji codes with their respective
emoji.

## Example

```go
package main

import (
    "fmt"
    "github.com/Bios-Marcel/discordemojimap"
)

func main() {
    fmt.Println(discordemojimap.Replace("What a wonderful day :sun_with_face:, am I right?"))
}
```

## Update Mapping

To regenerate `mapping.go`, run these commands:

```sh
wget http://discord.com/assets/5c193e4366261ef233e1.js
go run ./cmd/extractmap -path ./5c193e4366261ef233e1.js -out ./mapping.go
```

This was last updated on September 10th, 2023.

Note that the name of the asset containing the mapping may change in the future.

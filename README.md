# Discord Emoji Map

[![Build status](https://github.com/Bios-Marcel/discordemojimap/actions/workflows/go.yml/badge.svg)](https://github.com/Bios-Marcel/discordemojimap/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/Bios-Marcel/discordemojimap?status.svg)](https://pkg.go.dev/github.com/Bios-Marcel/discordemojimap/v2)
[![codecov](https://codecov.io/gh/Bios-Marcel/discordemojimap/branch/master/graph/badge.svg)](https://codecov.io/gh/Bios-Marcel/discordemojimap)

This is the map of emojis that Discord uses. This includes skin tones.

## Usage

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

## Generating

To regenerate `mapping.go`, run these commands:

```sh
wget https://discord.com/assets/b38205c8085075585265.js
go run ./cmd/extractmap -path ./b38205c8085075585265.js
```

This was last updated on the 4th of July, 2020.

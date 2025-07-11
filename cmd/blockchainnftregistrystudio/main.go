// cmd/blockchainnftregistrystudio/main.go
package main

import (
"flag"
"log"
"os"

"blockchainnftregistrystudio/internal/blockchainnftregistrystudio"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainnftregistrystudio.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}

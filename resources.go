//go:build !wasm

package main

import "embed"

//go:embed resources media
var resources embed.FS

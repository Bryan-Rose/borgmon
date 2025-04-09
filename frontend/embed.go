// Package site handles the Beszel frontend embedding.
package frontend

import (
	"embed"
	"io/fs"
)

//go:embed build/*
var distDir embed.FS

// DistDirFS contains the embedded dist directory files (without the "dist" prefix)
var DistDirFS, _ = fs.Sub(distDir, "build")

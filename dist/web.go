package dist

import "embed"

//go:embed index.html
var IndexHtml embed.FS

//go:embed static/* img/*
var Assets embed.FS

//go:embed index.html
var IndexByte []byte

//go:embed favicon.ico
var Favicon embed.FS

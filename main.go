package main

//go:generate go run gen.go salt
//go:generate go run gen.go
//go:generate go fmt .
//go:generate go fmt ./internal/...

// Perintah-perintah `go:generate` di atas akan dieksekusi saat perintah
// `go generate` diberikan.

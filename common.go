package main

import (
	"io"
)

func OpenAssets(filePath string) []byte {
	file, err := assets.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bytes, _ := io.ReadAll(file)
	return bytes
}

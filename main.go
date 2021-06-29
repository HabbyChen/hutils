package main

import (
	"kitutils/utils"
	"os"
	"path"
)

func main() {
	str, _ := os.Getwd()
	utils.Init(path.Join(str, "yml"))
}

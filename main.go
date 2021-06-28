package main

import (
	"kitutils/utils"
	"os"
	"path"
)

func main() {
	str, _ := os.Getwd()
	utils.InitNeed(path.Join(str, "yml"))
}

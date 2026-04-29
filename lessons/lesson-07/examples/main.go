package main

import (
	"example.com/lesson07/internal/auth"
	"example.com/lesson07/pkg/utils"
)

func main() {
	auth.Login("JavaDev")
	utils.PrintHello()
}

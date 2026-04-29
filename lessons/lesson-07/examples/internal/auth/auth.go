package auth

import "fmt"

// Login 是导出的函数（首字母大写）
func Login(username string) {
	fmt.Printf("User %s logged in\n", username)
}

// secret 是私有的（首字母小写），外部包无法访问
func secret() {
	fmt.Println("This is a secret")
}

package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// 测试多个密码
	passwords := []string{"admin", "admin123", "123456"}

	for _, password := range passwords {
		fmt.Printf("\n=== 测试密码: %s ===\n", password)

		// 生成哈希
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Printf("Error generating hash: %v\n", err)
			continue
		}
		fmt.Printf("Hash: %s\n", string(hash))

		// 验证正确密码
		err = bcrypt.CompareHashAndPassword(hash, []byte(password))
		if err == nil {
			fmt.Println("✅ 正确密码验证: SUCCESS")
		} else {
			fmt.Printf("❌ 正确密码验证: FAILED - %v\n", err)
		}

		// 验证错误密码
		err = bcrypt.CompareHashAndPassword(hash, []byte("wrong_password"))
		if err != nil {
			fmt.Println("✅ 错误密码验证: FAILED (正确行为)")
		} else {
			fmt.Println("❌ 错误密码验证: SUCCESS (不应该成功)")
		}
	}
}

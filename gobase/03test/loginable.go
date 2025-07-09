package main

import "fmt"

// 1. 定义一个接口 Loginable，要求实现 Login() 方法，并实现一个结构体 User 来实现它。
type Loginable interface {
	Login()
}

type User struct {
	name   string
	online bool
}

func (user *User) Login() {
	if !user.online {
		user.online = true
	}
}

// 2. 写一个函数 LoadConfig(path string) (Config, error)，模拟文件加载失败时返回错误。
type Config struct {
	Host string
	Port int
}

func LoadConfig(path string) (Config, error) {
	if path == "" {
		return Config{}, fmt.Errorf("配置文件路径为空")
	}
	// 模拟加载成功
	return Config{Host: "localhost", Port: 8080}, nil
}

// 3. 写一个闭包函数 AddN(n int)，返回一个函数 func(x int) int，返回 x + n。
func AddN(n int) func(x int) int {
	return func(x int) int {
		return x + n
	}
}

func main() {
	// 1.
	var u Loginable = &User{"Alice", false}
	u.Login()

	// 2.
	config, err := LoadConfig("")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(config.Host, " ", config.Port)
	}

	// 3.
	sum1 := AddN(2)
	fmt.Println(sum1(1)) // 2 + 1 = 3
	fmt.Println(sum1(3)) // 2 + 3 = 5

	sum2 := AddN(100)
	fmt.Println(sum2(1)) // 100 + 1 = 101
	fmt.Println(sum2(3)) // 100 + 3 = 103
}

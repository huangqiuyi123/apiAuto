package util

import (
	"math/rand"
	"time"
)

/*
   随机生成指定数量和指定长度的字符串
*/

func RandomString(count, length int) []string {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 用于生成随机数的随机数生成器
	// 使用当前时间戳
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 生成指定数量的随机字符串
	result := make([]string, count)
	for i := 0; i < count; i++ {
		// 生成随机字符串
		bytes := make([]byte, length)
		for j := 0; j < length; j++ {
			bytes[j] = alphabet[r.Intn(len(alphabet))] // r.Intn(len(alphabet)) 生成一个 [0, len(alphabet)-1] 范围内的随机整数
		}

		// 保存随机字符串到结果列表中
		result[i] = string(bytes)
	}

	return result
}

//func main() {
//	a := RandomString(1, 10)
//	fmt.Println(a)
//}

package tools

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func GetMd5(data interface{}) string {
	bytes, _ := json.Marshal(data)
	res := md5.Sum(bytes)
	md5 := string(res[:])
	return md5
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func GenerateUniqueRandomNumbers(count, max int) ([]int, error) {
	if count > max {
		return nil, fmt.Errorf("count cannot be greater than max")
	}
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子
	uniqueNumbers := make(map[int]bool)
	numbers := make([]int, 0, count)

	for len(numbers) < count {
		num := rand.Intn(max) + 1 // 生成 1 到 max 之间的随机数
		if _, exists := uniqueNumbers[num]; !exists {
			uniqueNumbers[num] = true
			numbers = append(numbers, num)
		}
	}
	return numbers, nil
}

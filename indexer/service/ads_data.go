package service

import (
	"ads-bool-indexer/indexer/ad_model"
	"ads-bool-indexer/indexer/tools/convert"
	"fmt"
	"math/rand"
	"time"
)

func GetADSFromDB() []ad_model.Ad {
	var ads []ad_model.Ad
	// 广告1
	var predicates []ad_model.Predicate
	predicates = append(predicates, ad_model.Predicate{
		Contains: true,
		Attr:     "gender", // 属性名称（唯一）
		ValueIDs: []int{2},
	})
	a1 := ad_model.Ad{
		ID:         1,
		Name:       "开屏广告", // 广告名称
		Predicates: predicates,
	}
	ads = append(ads, a1)

	// 其他广告
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子
	for a := 1; a < 1000; a++ {
		var predicates []ad_model.Predicate
		// 广告定向数量
		predicateNum := rand.Intn(1000) + 1
		predicateIndex, _ := generateUniqueRandomNumbers(predicateNum, 1000)
		for _, index := range predicateIndex {
			predicates = append(predicates, ad_model.Predicate{
				Contains: true,
				Attr:     convert.ToString(index), // 属性名称（唯一）
				ValueIDs: []int{rand.Intn(100)},
			})
		}
		var ad = ad_model.Ad{
			ID:         10000 + a,
			Name:       "广告" + convert.ToString(a), // 广告名称
			Predicates: predicates,
		}
		ads = append(ads, ad)
	}
	return ads
}

func generateUniqueRandomNumbers(count, max int) ([]int, error) {
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

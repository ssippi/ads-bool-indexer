package ad_model

import (
	"ads-bool-indexer/indexer/tools"
	"ads-bool-indexer/indexer/tools/convert"
	"encoding/json"
	"math/rand"
	"time"
)

// Ad 广告类
type Ad struct {
	ID         int         //主键id
	Name       string      // 广告名称
	Predicates []Predicate // 广告定向
	Creatives  Creative    // 广告素材
}

// Creative 广告素材
type Creative struct {
}

// Predicate 广告定向
type Predicate struct {
	Contains bool   // true : Attr在你ValueIDs范围内
	Attr     string // 属性名称
	ValueIDs []int  // PredicateEnumID
}

func GetADSFromDB() []Ad {
	// 广告1定向要求 性别女
	var str = `[{"ID":1,"Name":"开屏广告","Predicates":[{"Contains":true,"Attr":"gender","ValueIDs":[2]}]}]`
	var ads []Ad
	json.Unmarshal([]byte(str), &ads)

	// 其他广告
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子
	for a := 1; a < 1000; a++ {
		var predicates []Predicate
		// 广告定向数量
		predicateNum := rand.Intn(1000) + 1
		predicateIndex, _ := tools.GenerateUniqueRandomNumbers(predicateNum, 1000)
		for _, index := range predicateIndex {
			predicates = append(predicates, Predicate{
				Contains: true,
				Attr:     convert.ToString(index), // 属性名称（唯一）
				ValueIDs: []int{rand.Intn(100)},
			})
		}
		var ad = Ad{
			ID:         10000 + a,
			Name:       "广告_" + convert.ToString(10000+a), // 广告名称
			Predicates: predicates,
		}
		ads = append(ads, ad)
	}
	return ads
}

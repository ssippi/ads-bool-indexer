package indexer

import (
	"ads-bool-indexer/indexer/ad_model"
	"ads-bool-indexer/indexer/service"
	"fmt"
	"testing"
)

func TestSliceCapLen(t *testing.T) {

	// 匹配广告定向。转成内部枚举id
	predicateValueService := service.PredicateValueService{}
	predicateValueService.InitPredicateMap()
	var userMap = make(map[string]string)
	userMap[ad_model.AttrAge] = "20"
	userMap[ad_model.AttrGender] = "female"
	userMap[ad_model.AttrProvince] = "上海"
	predicateValueIds := predicateValueService.GetPredicateValueIds(userMap)

	// 创建全量索引
	indexService := service.IndexService{}
	indexService.BuildIndex()
	adIDs := indexService.Match(predicateValueIds)
	fmt.Println(adIDs)
}

var Ary = []int{1, 2, 3, 4, 5}

func BenchmarkSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sort()
	}
}

func BenchmarkSort2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sort2()
	}
}

func sort2() {
	var l = len(Ary)
	var min = 0
	for i := 0; i < l; i++ {
		if min > Ary[i] {
			min = Ary[i]
		}
	}
	var lm = 0
	for i := 0; i < l; i++ {
		if min == Ary[i] {
			lm++
		}
	}
}

func sort() {
	// 从小到大排序
	var n = len(Ary)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			a := Ary[j-1]
			b := Ary[j]

			if a > b || (a == b) {
				Ary[j-1], Ary[j] = Ary[j], Ary[j-1]
			}
			j = j - 1
		}
	}
}

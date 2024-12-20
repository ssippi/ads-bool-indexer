package indexer

import (
	"ads-bool-indexer/indexer/service"
	"fmt"
	"testing"
)

func TestSliceCapLen(t *testing.T) {

	// 匹配广告定向。转成内部枚举id
	PredicateEnumService := service.PredicateEnumService{}
	PredicateEnumService.InitPredicateMap()
	var userMap = make(map[string]string)
	userMap["age"] = "20"
	userMap["gender"] = "female"
	userMap["province"] = "上海"
	PredicateEnumIds := PredicateEnumService.GetPredicateEnumIds(userMap)

	// 创建全量索引
	indexService := service.IndexService{}
	indexService.BuildIndex()
	adIDs := indexService.Match(PredicateEnumIds)
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

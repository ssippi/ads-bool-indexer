package indexer

import (
	"ads-bool-indexer/indexer/ad_model"
	"ads-bool-indexer/indexer/service"
	"fmt"
	"testing"
)

func TestSliceCapLen(t *testing.T) {

	predicateValueService := service.PredicateValueService{}
	indexService := service.IndexService{}

	// 创建全量索引
	predicateValueService.InitPredicateMap()
	indexService.BuildIndex()

	// 匹配广告
	user := ad_model.User{}
	user.Age = 20
	user.Gender = "male"
	user.Province = "shanghai"
	predicateValueIds := predicateValueService.GetPredicateValueIds(user)

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

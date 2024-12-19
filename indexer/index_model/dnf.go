package index_model

import "ads-bool-indexer/indexer/ad_model"

type DNF struct {
	Conjunctions []Conjunction
}

type Conjunction struct {
	ConjId     int // 数据index
	Predicates []ad_model.Predicate
	AdId       []int
}

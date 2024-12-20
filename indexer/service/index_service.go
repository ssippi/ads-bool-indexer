package service

import (
	"ads-bool-indexer/indexer/ad_model"
	"ads-bool-indexer/indexer/index_model"
	"ads-bool-indexer/indexer/tools"
)

var gIndexer index_model.Indexer

type IndexService struct {
}

func (a IndexService) BuildIndex() {
	// 获取广告数据
	adsData := ad_model.GetADSFromDB()
	// 构建索引
	conjunctions := a.getConjunctions(adsData)

	// 构建conjID和广告id的对应关系
	conjAdIdMap := make(map[int]index_model.AdIdAry) // map[conjID]AdId
	for i, conjunction := range conjunctions {
		conjAdIdMap[i] = conjunction.AdId
	}

	// 构建predicateConjMap核心索引
	conjunctionLevelMap := a.convertToMap(conjunctions)
	predicateConjMap := a.buildPredicateConjMap(conjunctionLevelMap)

	var maxConjSize int
	for level, _ := range conjunctionLevelMap {
		if level >= maxConjSize {
			maxConjSize = level
		}
	}

	var data = index_model.Indexer{}
	data.Conjunctions = conjunctions
	data.ConjAdIdMap = conjAdIdMap
	data.PredicateConjMap = predicateConjMap
	data.MaxConjSize = maxConjSize
	gIndexer = data
}

func (a IndexService) getConjunctions(ads []ad_model.Ad) []index_model.Conjunction {
	// 1 解析广告camps成 conj
	conjMap := make(map[string]*index_model.Conjunction) // map[md5]广告定向
	for _, ad := range ads {
		// 相同的广告定向，用md5去重
		md5 := tools.GetMd5(ad.Predicates)
		if conj, ok := conjMap[md5]; ok {
			conj.AdId = append(conj.AdId, ad.ID)
		} else {
			var newConj index_model.Conjunction
			newConj.Predicates = ad.Predicates
			newConj.AdId = []int{ad.ID}
			conjMap[md5] = &newConj
		}
	}
	// 2 生成 conj_id = conjunctions数组index
	var conjunctions []index_model.Conjunction
	var index int
	for _, conjunction := range conjMap {
		conjunction.ConjId = index
		conjunctions = append(conjunctions, *conjunction)
		index++
	}
	return conjunctions
}

// Conjunction 按定向纬度数量做分层
func (a IndexService) convertToMap(conjunctions []index_model.Conjunction) map[int][]index_model.Conjunction {
	conjunctionLevelMap := make(map[int][]index_model.Conjunction)
	for _, conjunction := range conjunctions {
		// 定向纬度数量
		var conjLevel = 0
		for _, predicate := range conjunction.Predicates {
			if predicate.Contains {
				conjLevel++
			}
		}
		conjunctionTemps, ok := conjunctionLevelMap[conjLevel]
		if !ok {
			conjunctionTemps = []index_model.Conjunction{}
			conjunctionLevelMap[conjLevel] = conjunctionTemps
		}
		conjunctionTemps = append(conjunctionTemps, conjunction)
		conjunctionLevelMap[conjLevel] = conjunctionTemps
	}
	return conjunctionLevelMap
}

func (a IndexService) buildPredicateConjMap(conjunctionLevelMap map[int][]index_model.Conjunction) map[int]index_model.PredicateConjMap {
	conjLevelMap := make(map[int]index_model.PredicateConjMap)
	for level, conjS := range conjunctionLevelMap {
		conjAry := make(index_model.PredicateConjMap) // // map[PredId][]conjId
		for _, conjunction := range conjS {
			for _, predicate := range conjunction.Predicates {
				entry := index_model.Entry{
					ConjId:   conjunction.ConjId,
					Contains: predicate.Contains,
				}
				for _, PredicateEnumID := range predicate.ValueIDs {
					pIds, _ := conjAry[PredicateEnumID]
					pIds = append(pIds, entry)
					conjAry[PredicateEnumID] = pIds
				}
			}
		}
		conjLevelMap[level] = conjAry
	}
	return conjLevelMap
}

func (a IndexService) Match(PredicateEnumIds []int) []int {
	conjIDs := a.matchConjID(PredicateEnumIds)
	adIds := a.getAdIdsByConjID(conjIDs)
	return adIds
}

func (a IndexService) matchConjID(PredicateEnumIds []int) []int {
	var conjunctionIds []int
	minLevel := tools.Min(len(PredicateEnumIds), gIndexer.MaxConjSize)
	for level := minLevel; level >= 0; level-- {
		// 匹配conj
		pLists := a.getIds(level, PredicateEnumIds)

		// k 是满足conj的条件数量
		K := level
		if K == 0 {
			K = 1
		}
		// 如果匹配出来的属性数据，小于k，可以判断这些conjId都不满足
		if len(pLists) < K {
			continue
		}
		pListService := PListService{pLists}
		resultTemp := pListService.GetMatch(K)
		conjunctionIds = append(conjunctionIds, resultTemp...)
	}
	return conjunctionIds
}

func (a IndexService) getIds(level int, predicateIds []int) []index_model.PList {
	predicateConjMap := gIndexer.PredicateConjMap[level]
	var pLists []index_model.PList
	for _, id := range predicateIds {
		// 条件匹配的conj
		entryS := predicateConjMap[id]
		l := len(entryS)
		if l > 0 {
			pLists = append(pLists, index_model.PList{EntryS: entryS, PredicateId: id, Curr: 0, Len: l})
		}
	}
	return pLists
}

func (a IndexService) getAdIdsByConjID(IDs []int) []int {
	var adIds []int
	for _, ID := range IDs {
		datas, _ := gIndexer.ConjAdIdMap[ID]
		adIds = append(adIds, datas...)
	}
	return adIds
}

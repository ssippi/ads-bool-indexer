package service

import (
	"ads-bool-indexer/indexer/ad_model"
	"strconv"
)

var (
	// map[age_18]id
	PredicateEnumMap = map[string]int{}

	// map[age][]PredicateEnum-范围匹配
	PredicateRangeMap = map[string][]ad_model.PredicateEnum{}

	// map[PredicateEnumId]id
	PredicateParentMap = map[int]int{}
)

type PredicateEnumService struct {
}

func (a PredicateEnumService) GetPredicateEnumIds(userData map[string]string) []int {
	// 根据熟悉值获取ValueId
	var predicateIds []int
	for key, value := range userData {
		if key == "province" {
			// 父子数据类型
			var ids = a.getIdsByValue("province", value)
			predicateIds = append(predicateIds, ids...)
		} else {
			id := a.getIdByValue(key, value)
			predicateIds = append(predicateIds, id)
		}
	}
	// 排出非法用户定向
	predicateIds = getIdWithOutZero(predicateIds)
	return predicateIds
}

func (a PredicateEnumService) InitPredicateMap() {
	var dataS = ad_model.GetPredicateEnumFromDB()
	var PredicateEnumMap = make(map[string]int)
	var predicateRangeMap = make(map[string][]ad_model.PredicateEnum)
	var predicateParentMap = make(map[int]int)
	for _, v := range dataS {
		if v.ValueType == ad_model.Value {
			var key = v.Attr + "_" + v.Value
			PredicateEnumMap[key] = v.ID
		} else if v.ValueType == ad_model.Range {
			var key = v.Attr
			if predicates, ok := predicateRangeMap[key]; ok {
				predicates = append(predicates, v)
			} else {
				// todo
				predicates = []ad_model.PredicateEnum{}
				predicates = append(predicates, v)
				predicateRangeMap[key] = predicates
			}
		}
		if v.ParentID != 0 {
			predicateParentMap[v.ID] = v.ParentID
		}
	}
	PredicateEnumMap = PredicateEnumMap
	PredicateRangeMap = predicateRangeMap
	PredicateParentMap = predicateParentMap
}

func getIdWithOutZero(ids []int) []int {
	var predicateIds []int
	for _, id := range ids {
		if id > 0 {
			predicateIds = append(predicateIds, id)
		}
	}
	return predicateIds
}

func (a PredicateEnumService) getIdByValue(attr, value string) int {
	var key = attr + "_" + value
	if id, ok := PredicateEnumMap[key]; ok {
		return id
	}
	num, _ := strconv.Atoi(value)
	if ps, ok := PredicateRangeMap[attr]; ok {
		for _, p := range ps {
			if p.BeginValue <= num && num <= p.EndValue {
				return p.ID
			}
		}
	}
	return 0
}

func (a PredicateEnumService) getIdsByValue(attr, value string) []int {
	id := a.getIdByValue(attr, value)
	if id == 0 {
		return []int{}
	}
	ids := a.getIdsByPredicateId(id)
	return ids
}

func (a PredicateEnumService) getIdsByPredicateId(id int) []int {
	var ids = []int{id}
	for {
		if parentID, ok := PredicateParentMap[id]; ok {
			id = parentID
			ids = append(ids, parentID)
		} else {
			break
		}
	}
	return ids
}

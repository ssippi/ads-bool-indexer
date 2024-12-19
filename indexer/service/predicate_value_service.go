package service

import (
	"ads-bool-indexer/indexer/ad_model"
	"strconv"
)

var (
	PredicateValueMap  = map[string]int{}                       // map[age_2]PredicateValueId-id匹配
	PredicateRangeMap  = map[string][]ad_model.PredicateValue{} // map[age][]PredicateValue-范围匹配
	PredicateParentMap = map[int]int{}                          // map[PredicateValueId]ParentId
)

type PredicateValueService struct {
}

func (a PredicateValueService) GetPredicateValueIds(userData map[string]string) []int {
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

func (a PredicateValueService) InitPredicateMap() {
	var dataS = GetPredicateValueFromDB()
	var predicateValueMap = make(map[string]int)
	var predicateRangeMap = make(map[string][]ad_model.PredicateValue)
	var predicateParentMap = make(map[int]int)
	for _, v := range dataS {
		if v.ValueType == ad_model.Value {
			var key = v.Attr + "_" + v.Value
			predicateValueMap[key] = v.ID
		} else if v.ValueType == ad_model.Range {
			var key = v.Attr
			if predicates, ok := predicateRangeMap[key]; ok {
				predicates = append(predicates, v)
			} else {
				// todo
				predicates = []ad_model.PredicateValue{}
				predicates = append(predicates, v)
				predicateRangeMap[key] = predicates
			}
		}
		if v.ParentID != 0 {
			predicateParentMap[v.ID] = v.ParentID
		}
	}
	PredicateValueMap = predicateValueMap
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

func (a PredicateValueService) getIdByValue(attr, value string) int {
	var key = attr + "_" + value
	if id, ok := PredicateValueMap[key]; ok {
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

func (a PredicateValueService) getIdsByValue(attr, value string) []int {
	id := a.getIdByValue(attr, value)
	if id == 0 {
		return []int{}
	}
	ids := a.getIdsByPredicateId(id)
	return ids
}

func (a PredicateValueService) getIdsByPredicateId(id int) []int {
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

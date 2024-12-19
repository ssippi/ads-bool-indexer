package service

import "ads-bool-indexer/indexer/index_model"

type PListService struct {
	pListAry []index_model.PList
}

func (p PListService) SortByCurrent() {
	// 从小到大排序
	var n = len(p.pListAry)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			a := p.pListAry[j-1].Current()
			b := p.pListAry[j].Current()

			if a.ConjId > b.ConjId || (a.ConjId == b.ConjId) && a.Contains && !b.Contains {
				p.pListAry[j-1], p.pListAry[j] = p.pListAry[j], p.pListAry[j-1]
			}
			j = j - 1
		}
	}
}

// GetMatch k是 匹配conj的条件数量
func (p PListService) GetMatch(K int) []int {
	var result []int
	p.SortByCurrent()
	for p.pListAry[K-1].Curr != index_model.EOL {
		var nextID int
		if p.pListAry[0].Current().ConjId == p.pListAry[K-1].Current().ConjId {
			if p.pListAry[0].Current().Contains == false {
				rejectId := p.pListAry[0].Current().ConjId
				// 补充逻辑（排除前k行的头部数据）+（排除剩余的头部数据）
				for L := K; L <= len(p.pListAry)-1; L++ {
					if p.pListAry[L].Current().ConjId == rejectId {
						p.pListAry[L].SkipTo(nextID)
					} else {
						break
					}
				}
			} else {
				// ConjId条件都满足
				result = append(result, p.pListAry[0].Current().ConjId)
			}
			nextID = p.pListAry[K-1].Current().ConjId + 1
		} else {
			nextID = p.pListAry[K-1].Current().ConjId
		}

		// 排除前k行的头部数据
		for L := 0; L <= K-1; L++ {
			p.pListAry[L].SkipTo(nextID)
		}
		// 排序性能不高，可优化
		p.SortByCurrent()
	}
	return result
}

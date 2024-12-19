package index_model

const (
	EOL = -1
)

type Indexer struct {
	PredicateConjMap map[int]PredicateConjMap // map[level]map[PredId][]conjId
	Conjunctions     []Conjunction            //生成conj_id
	ConjAdIdMap      map[int]AdIdAry          // map[ConjId]AdIdAry
	MaxConjSize      int
}

type PredicateConjMap = map[int][]Entry // map[PredId][]conjId
type AdIdAry = []int

type PList struct {
	EntryS      []Entry
	PredicateId int
	Curr        int
	Len         int
}

type Entry struct {
	ConjId   int
	Contains bool // ∈：true, ∉: false
}

func (p *PList) Current() Entry {
	return p.EntryS[p.Curr]
}
func (p *PList) SkipTo(id int) {
	if p.Curr == p.Len-1 {
		p.Curr = EOL
		return
	}
	for p.EntryS[p.Curr].ConjId < id {
		p.Curr++
	}
}

package ad_model

const (
	// AttrAge 广告纬度
	AttrGender   = "gender"   // 性别
	AttrCountry  = "country"  // 国家
	AttrProvince = "province" // 城市
	Attr1        = "1"        // 其他纬度
	Attr2        = "2"
	Attr3        = "3"
	Attr_1000    = "1000"
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
	ValueIDs []int  // PredicateValueID
}

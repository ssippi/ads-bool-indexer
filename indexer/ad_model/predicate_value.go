package ad_model

import (
	"ads-bool-indexer/indexer/tools/convert"
	"encoding/json"
)

var (
	Value ValueType = 1 // 准确
	Range ValueType = 2 // 范围
)

type ValueType int

// PredicateEnum 广告定向
type PredicateEnum struct {
	ID         int       //主键id
	ParentID   int       //所属父类属性id
	Attr       string    // 广告纬度
	ValueType  ValueType // 值类型
	Value      string    // 属性值
	BeginValue int       // 范围属性值（开始）左闭右闭
	EndValue   int       // 范围属性值（结束）左闭右闭
}

// 获取定向枚举枚举
func GetPredicateEnumFromDB() []PredicateEnum {
	var str = `[
	{"ID":1,"ParentID":0,"Attr":"gender","ValueType":1,"Value":"male","BeginValue":0,"EndValue":0},
	{"ID":2,"ParentID":0,"Attr":"gender","ValueType":1,"Value":"female","BeginValue":0,"EndValue":0},

	{"ID":3,"ParentID":0,"Attr":"age","ValueType":2,"Value":"","BeginValue":18,"EndValue":200},
	{"ID":4,"ParentID":0,"Attr":"age","ValueType":1,"Value":"2","BeginValue":0,"EndValue":0},

	{"ID":14,"ParentID":0,"Attr":"country","ValueType":1,"Value":"CN","BeginValue":0,"EndValue":0},
	{"ID":15,"ParentID":14,"Attr":"province","ValueType":1,"Value":"shanghai","BeginValue":0,"EndValue":0},
	{"ID":16,"ParentID":0,"Attr":"country","ValueType":1,"Value":"CA","BeginValue":0,"EndValue":0}
	]`
	var vs []PredicateEnum
	json.Unmarshal([]byte(str), &vs)

	// 其他自定义定向
	for a := 1; a <= 1000; a++ {
		for b := 1; b <= 100; b++ {
			vs = append(vs, PredicateEnum{
				ID:        10000 + a,
				Attr:      convert.ToString(a),
				ValueType: Value,
				Value:     convert.ToString(b),
			})
		}
	}

	return vs
}

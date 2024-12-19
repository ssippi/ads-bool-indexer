package ad_model

var (
	Value ValueType = 1 // 准确值
	Range ValueType = 2 // 范围只
)

// PredicateValue 枚举值
type PredicateValue struct {
	ID         int       //主键id
	ParentID   int       //所属父类属性id
	Attr       string    // 属性名称（唯一）
	ValueType  ValueType // 值类型
	Value      string    // 属性值
	BeginValue int       // 范围属性值（开始）左闭右闭
	EndValue   int       // 范围属性值（结束）左闭右闭
}

type ValueType int

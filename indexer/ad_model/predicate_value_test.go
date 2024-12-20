package ad_model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetPredicateEnumFromDB(t *testing.T) {
	var str = `[{"ID":1,"ParentID":0,"Attr":"gender","ValueType":1,"Value":"male","BeginValue":0,"EndValue":0},{"ID":2,"ParentID":0,"Attr":"gender","ValueType":1,"Value":"female","BeginValue":0,"EndValue":0},{"ID":3,"ParentID":0,"Attr":"age","ValueType":2,"Value":"","BeginValue":18,"EndValue":200},{"ID":4,"ParentID":0,"Attr":"country","ValueType":1,"Value":"CHN","BeginValue":0,"EndValue":0},{"ID":5,"ParentID":4,"Attr":"province","ValueType":1,"Value":"shanghai","BeginValue":0,"EndValue":0}]`
	var vs []PredicateEnum
	json.Unmarshal([]byte(str), &vs)
	fmt.Println(vs)
}

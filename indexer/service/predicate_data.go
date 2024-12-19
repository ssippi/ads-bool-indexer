package service

import (
	"ads-bool-indexer/indexer/ad_model"
	"ads-bool-indexer/indexer/tools/convert"
)

func GetPredicateValueFromDB() []ad_model.PredicateValue {
	// 读mysql
	var vs []ad_model.PredicateValue
	// 性别定向
	vs = append(vs, ad_model.PredicateValue{
		ID:        1,
		Attr:      ad_model.AttrGender,
		ValueType: ad_model.Value,
		Value:     "male",
	})
	vs = append(vs, ad_model.PredicateValue{
		ID:        2,
		Attr:      ad_model.AttrGender,
		ValueType: ad_model.Value,
		Value:     "female",
	})

	// 年龄定向
	vs = append(vs, ad_model.PredicateValue{
		ID:         3,
		Attr:       ad_model.AttrAge,
		ValueType:  ad_model.Range,
		BeginValue: 18,
		EndValue:   200,
	})

	// 国家定向
	vs = append(vs, ad_model.PredicateValue{
		ID:        4,
		Attr:      ad_model.AttrCountry,
		ValueType: ad_model.Value,
		Value:     "CHN",
	})
	vs = append(vs, ad_model.PredicateValue{
		ID:        5,
		ParentID:  4,
		ValueType: ad_model.Value,
		Attr:      ad_model.AttrProvince,
		Value:     "shanghai",
	})

	// other 定向
	for a := 1; a <= 1000; a++ {
		for b := 1; b <= 100; b++ {
			vs = append(vs, ad_model.PredicateValue{
				ID:        10000 + a,
				Attr:      convert.ToString(a),
				ValueType: ad_model.Value,
				Value:     convert.ToString(b),
			})
		}
	}

	return vs
}

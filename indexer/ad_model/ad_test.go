package ad_model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetADSFromDB(t *testing.T) {
	//ads := GetADSFromDB()
	//adsBytes, _ := json.Marshal(ads)
	//fmt.Println(string(adsBytes))
	var str = `[{"ID":1,"Name":"开屏广告","Predicates":[{"Contains":true,"Attr":"gender","ValueIDs":[2]}]}]`
	var adsNew []Ad
	json.Unmarshal([]byte(str), &adsNew)
	fmt.Println(adsNew)
}

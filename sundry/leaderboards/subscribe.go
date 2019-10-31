// Time        : 2019/10/22
// Description :

package leaderboards

import (
	"fmt"
	"reflect"
	"strings"
)

const Separator = "::"

// subscriber
func subscribe(event interface{}) {
	info, ok := extractEventInfo(event)
	if !ok {
		return
	}
	typ := info["type"].(uint16)
	_ = typ
	identity := strings.Join(append([]string{info["identity"].(string)}, info["group"].([]string)...), Separator)
	_ = identity
	value := info["value"].(int64)
	extend := info["extend"]

	var rtr RunningTotalRule
	//rtr = rr.GetRunningTotalRule(typ)

	var runTotal RunningTotal
	//runTotal, ok = repository.GetByIdentify(identity)
	//if !ok {
	//	runTotal = repository.Create(identity, value, extend)
	//}

	GatherMap[rtr.GatherType](&runTotal, value)
	if extend != nil {
		runTotal.Extend = extend
	}
	//repository.Save(runTotal)
}

func extractEventInfo(event interface{}) (info map[string]interface{}, ok bool) {
	info = make(map[string]interface{})
	rv := reflect.ValueOf(event)
	rt := rv.Type()
	group := make([]string, 0)
	info["group"] = group
	for i := 0; i < rv.NumField(); i++ {
		tag := rt.Field(i).Tag.Get("rt")
		if tag == "group" {
			gp := fmt.Sprintf("%v", rv.Field(i).Interface())
			if gp == "" {
				continue
			}
			group = append(group, gp)
			info["group"] = group
		} else if tag != "" {
			info[tag] = rv.Field(i).Interface()
		}
	}
	if _, ok = info["type"]; !ok {
		return
	}
	if _, ok = info["identity"]; !ok {
		panic("must given identity")
	}
	return
}

var GatherMap = map[GatherType]GatherFunc{
	Sum: GatherSum,
}

type GatherFunc func(rt *RunningTotal, value int64)

func GatherSum(rt *RunningTotal, value int64) {
	rt.Value += value
}

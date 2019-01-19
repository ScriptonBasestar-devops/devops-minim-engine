package util

import (
	"fmt"
	"strings"
)

func ArgDataArrToMap(arr []string) map[string]*string {
	//result := map[string]*string{}
	result := make(map[string]*string)
	for _, item := range arr {
		slist := strings.Split(item, "=")
		if len(slist) != 2 {
			panic(fmt.Sprintf("%: parsing error string pattern should be 'key=val'", item))
		}
		result[slist[0]] = &slist[1]
	}
	return result
}

func VolumeDataArrToMap(arr []string) map[string]struct{} {
	result := make(map[string]struct{})
	for _, item := range arr {
		//slist := regexp.MustCompile("[:=]").Split(item, -1)
		//if len(slist) != 2 {
		//	panic(fmt.Sprintf("%: parsing error string pattern should be 'key=val'", item))
		//}
		//result[slist[0]] = struct{}{}
		result[item] = struct{}{}
	}
	return result
}

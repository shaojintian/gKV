package src

import (
	"gKV/utils"
	"log"
	"strconv"
)

var GlobalMap map[string]string

const MAP_INIT_SIZE int = 100

func zMapSize() int {
	return len(GlobalMap)
}

func printZMap() {
	for k, v := range GlobalMap {
		log.Printf("k=%s,v=%s\n", k, v)
	}
}

func doSet(opts []string) string {
	GlobalMap[opts[0]] = opts[1]
	return utils.OK
}

func doGet(opts []string) string {
	k := opts[0]
	v, ok := GlobalMap[k]
	if ok {
		return v
	}
	return utils.ERR_NIL
}

//del k
func doDel(opts []string) string {
	k := opts[0]
	if _, ok := GlobalMap[k]; ok {
		//del k
		delete(GlobalMap, k)
		return utils.INTEGER + "1"
	}
	return utils.INTEGER + "0"
}

//append [k v]
func doAppend(opts []string) string {
	k := opts[0]
	v := opts[1]
	if value, ok := GlobalMap[k]; ok {
		//do append
		GlobalMap[k] = value + v
		return utils.INTEGER + strconv.Itoa(len(value+v))
	}
	// no this k,do set k ,v
	GlobalMap[k] = v
	return utils.INTEGER + strconv.Itoa(len(v))
}

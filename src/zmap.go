package src

import (
	"gKV/utils"
	"log"
)

var GlobalMap map[string]string

const MAP_INIT_SIZE int = 100

func zMapSize() int {
	return len(GlobalMap)
}

func printZMap(){
	for k,v:=range GlobalMap{
		log.Printf("k=%s,v=%s\n",k,v)
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

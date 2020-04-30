package src

import "log"

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
package src

import (
	"container/list"
	"gKV/utils"
	"log"
	"strconv"
)
var ZlistCounter map[string]*Zlist

//const ZLIST_CAPACITY uint32 = (1 << 32) - 1 // 2^32 -1


type Zlist struct{
	name string
	l *list.List

}

func newZlist(name string)*Zlist{
	return &Zlist{
		name:name,
		l:list.New(),
	}
}

//[name value(v)]
func lpush(opts []string) string{
	name := opts[0]
	v := opts[1]
	var size int
	if zlistObj,ok:= ZlistCounter[name];ok{
		zlistObj.l.PushFront(v)
		size = zlistObj.l.Len()
		//maybe bug of old zlistObj
	}else{
		zl := newZlist(name)
		zl.l.PushFront(v)
		size = zl.l.Len()
		log.Printf("current zlist:%s size is %d",name,size)
		ZlistCounter[name]=zl
	}

	return utils.INTEGER+strconv.Itoa(size)
}

//lrange name start end  [start,end]
func lrange(opts []string) []interface{}{

	name := opts[0]
	if zlistObj,ok:= ZlistCounter[name];ok{
		res := make([]interface{}, zlistObj.l.Len())
		for i:=zlistObj.l.Front();i!=nil;i=i.Next(){
			res = append(res,i.Value)
		}
		return res
	}else{
		return []interface{}{utils.EMPTY_LIST_OR_SET}
	}
}

func llen(opts []string) int {
	name := opts[0]
	if zlistObj,ok:= ZlistCounter[name];ok{
		return zlistObj.l.Len()
	}else{
		return 0
	}

}



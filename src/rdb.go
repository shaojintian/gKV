package src

import (
	"gKV/utils"
	"log"
)

type redisObject struct {
	Type     uint
	Encoding uint
	Lru      uint /* lru time (relative to server.lruclock) */
	Refcount int
	Ptr      interface{}
}
type robj redisObject

func newRobj() *robj {
	return &robj{
		//....
	}
}

func rdbSaveType(t *rune) int {
	return 0
}
func rdbLoadType() {

}
func rdbSaveTime(time *time_t) {

}
func rdbLoadTime() time_t {
	return 0
}
func rdbSaveLen(len uint32) int {
	return 0
}
func rdbLoadLen(isencoded *int) uint32 {
	return 0
}
func rdbSaveObjectType(o *robj) int {
	return 0
}
func rdbLoadObjectType() int {
	return 0
}
func rdbLoad(filename *rune) {

}
func rdbSaveToSlavesSockets() {

}
func rdbRemoveTempFile(childpid pid_t) {

}

func rdbSaveObject(o *robj) int64 {
	return 0
}
func rdbSavedObjectLen(o *robj) uint64 {
	return 0
}
func rdbLoadObject(typ int) *robj {
	return newRobj()
}
func backgroundSaveDoneHandler(exitcode int, bysignal int) {

}
func rdbSaveKeyValuePair(key, val *robj, expiretime, now int64) {
}
func rdbLoadStringObject() *robj {

	return newRobj()
}

func saveCommand() string {
	filename := "default.rdb"
	if rdbSave(filename) == utils.OK {
		return utils.OK
	} else {
		return utils.ERR_SAVE
	}
}

func rdbSave(filename string) string {
	var (
		tmpfile [256]rune
		cwd     [utils.MAX_PATH_LEN]rune
	)
}

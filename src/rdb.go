package src

import ()

type redisObject struct {
	Type     uint
	Encoding uint
	Lru      uint /* lru time (relative to server.lruclock) */
	Refcount int
	Ptr      interface{}
}

type robj redisObject

func rdbSaveType(t *rune) int {

}
func rdbLoadType() {

}
func rdbSaveTime(time *time_t) {

}
func rdbLoadTime() time_t {

}
func rdbSaveLen(len uint32) int {

}
func rdbLoadLen(isencoded *int) uint32 {

}
func rdbSaveObjectType(o *robj) int {

}
func rdbLoadObjectType() int {

}
func rdbLoad(filename *rune) {

}
func rdbSaveToSlavesSockets() {

}
func rdbRemoveTempFile(childpid pid_t) {

}
func rdbSave(filename *rune) {

}
func rdbSaveObject(o *robj) int64 {

}
func rdbSavedObjectLen(o *robj) uint64 {

}
func rdbLoadObject(typ int) *robj {

}
func backgroundSaveDoneHandler(exitcode int, bysignal int) {

}
func rdbSaveKeyValuePair(key, val *robj, expiretime, now int64) {
}
func rdbLoadStringObject() *robj {

}

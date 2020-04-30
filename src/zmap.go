package src

var GlobalMap map[string]string

const MAP_INIT_SIZE int = 100

func zMapSize() int {
	return len(GlobalMap)
}

package utils

import "log"

const (
	//err
	ERR_SET           string = "(error) ERR wrong number of arguments for 'set' command"
	ERR_GET           string = "(error) ERR wrong number of arguments for 'get' command"
	ERR_NIL           string = "(nil)"
	OK                string = "OK"
	INTEGER           string = "(integer) "
	EMPTY_LIST_OR_SET string = "(empty list or set)"
	ERR_SAVE          string = "(err rdb save)"
	ERR_FILE          string = "(err file operation)"

	//other
	MAX_PATH_LEN int = 1024
)

func CheckErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

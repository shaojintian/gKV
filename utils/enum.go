package utils

import "log"

const (
	ERR_SET string = "(error) ERR wrong number of arguments for 'set' command"
	ERR_GET string = "(error) ERR wrong number of arguments for 'get' command"
	ERR_NIL string = "(nil)"
	OK      string = "OK"
)

func CheckErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

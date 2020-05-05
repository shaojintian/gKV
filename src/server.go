package src

import (
	"gKV/utils"
	"log"
	"net"
	"strconv"
	"strings"
	//"reflect"
)

func ReceiveFromClient(conn net.Conn) []byte {
	log.Println("start read data from client ...")
	operation := make([]byte, 1024)
	n, err := conn.Read(operation)
	utils.CheckErr(err)
	log.Println("read operation:" + string(operation) + " FROM client successfully!")
	return operation[:n]
}

func Send2Client(res string, conn net.Conn) {
	log.Println("start send " + res + " to client...!")
	_, err := conn.Write([]byte(res))
	utils.CheckErr(err)
	log.Println("return " + res + " to client successfully!\n")
}

func Handle(conn net.Conn) {

	//Receive operation from client
	operation := ReceiveFromClient(conn)
	//handle xxx k v
	res := doHandle(operation)
	//Send2Client
	Send2Client(res, conn)
}

func doHandle(operation []byte) string {
	var res string
	//handle set/get
	//delete space in operation begin or end
	//log.Println(operation)
	opStr := strings.TrimSpace(string(operation))
	//split opStr set k v
	opts := strings.Split(opStr, " ")
	//print operation like "set key value"
	log.Println(opStr)
	if len(opts) == 3 {
		//set k v
		if opts[0] == "set" {
			res = doSet(opts[1:])

			//lpush listName string
		} else if opts[0] == "lpush" || opts[0] == "LPUSH" {
			res = lpush(opts[1:])
			//append k v
		} else if checkOpt("append", opts) {
			res = doAppend(opts[1:])
		}
	} else if len(opts) == 2 {
		// get k
		if opts[0] == "get" {
			res = doGet(opts[1:])
			//llen [name]
		} else if opts[0] == "llen" || opts[0] == "LLEN" {
			res = strconv.Itoa(llen(opts[1:]))
			//del k
		} else if opts[0] == "del" || opts[0] == "DEL" {
			res = doDel(opts[1:])
		}
	} else if len(opts) == 4 {
		//lrange name start end  [s,e] len== e-s+1
		if opts[0] == "lrange" || opts[0] == "LRANGE" {
			res = lrange(opts[1:])
		}
	} else {
		res = "(error) ERR wrong number of arguments for  command"
		log.Printf("zMap size is:%d\n", zMapSize())
		printZMap()
	}

	return res
}

func checkOpt(opt string, opts []string) bool {
	if opts[0] == opt || opts[0] == strings.ToUpper(opt) {
		return true
	}
	return false
}

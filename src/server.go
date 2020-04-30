package src

import (
	"gKV/utils"
	"log"
	"net"
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
		}else if opts[0]=="lpush" || opts[0]=="LPUSH" {
			res = lpush(opts[1:])
		}
	} else if len(opts) == 2 {
		if opts[0] == "get" {
			res = doGet(opts[1:])
			//llen [name]
		}else if opts[0] == "llen"||opts[0]=="LLEN"{
			res = string(llen(opts[1:]))
		}
	} else if len(opts)==4{
		//lrange name start end
		if opts[0]=="lrange"||opts[0]=="LRANGE"{
			//....
		}
	} else {
		res = "(error) ERR wrong number of arguments for  command"
		log.Printf("zMap size is:%d\n",zMapSize())
		printZMap()
	}

	return res
}




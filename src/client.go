package src

import (
	"container/list"
	"gKV/utils"
	"net"
)

//-------------Global variable-------------------
var Client = newClient()

type gkvClient struct {
	// 套接字描述符
	fd int

	// 当前正在使用的数据库
	db *gkvDB

	// 当前正在使用的数据库的 id （号码）
	dictid int

	// 客户端的名字
	name *robj /* As set by CLIENT SETNAME */

	// 查询缓冲区
	querybuf string /* As set by CLIENT SETNAME */

	// 查询缓冲区长度峰值
	querybufPeak size_t /* Recent (100ms or more) peak of querybuf size */

	// 参数数量
	argc int /* Recent (100ms or more) peak of querybuf size */

	// 参数对象数组
	argv **robj /* Recent (100ms or more) peak of querybuf size */

	// 记录被客户端执行的命令
	cmd, lastcmd *gkvCommand /* Recent (100ms or more) peak of querybuf size */

	// 请求的类型：内联命令还是多条命令
	reqtype int /* Recent (100ms or more) peak of querybuf size */

	// 剩余未读取的命令内容数量
	multibulklen int /* number of multi bulk arguments left to read */

	// 命令内容的长度
	bulklen int32 /* length of bulk argument in multi bulk request */

	// 回复链表
	reply *list.List /* length of bulk argument in multi bulk request */

	// 回复链表中对象的总大小
	replyBytes uint32 /* Tot bytes of objects in reply list */

	// 已发送字节，处理 short write 用
	sentlen int /* Amount of bytes already sent in the current
	buffer or object being sent. */

	// 创建客户端的时间
	ctime time_t /* Client creation time */

	// 客户端最后一次和服务器互动的时间
	lastInteraction time_t /* time of the last interaction, used for timeout */

	// 客户端的输出缓冲区超过软性限制的时间
	obufSoftLimitReachedTime time_t /* time of the last interaction, used for timeout */

	// 客户端状态标志
	flags int /* gkv_SLAVE | gkv_MONITOR | gkv_MULTI ... */

	// 当 server.requirepass 不为 NULL 时
	// 代表认证的状态
	// 0 代表未认证， 1 代表已认证
	authenticated int /* when requirepass is non-NULL */

	// 复制状态
	replState int /* replication state if this is a slave */
	// 用于保存主服务器传来的 RDB 文件的文件描述符
	replDbFd int /* replication DB file descriptor */

	// 读取主服务器传来的 RDB 文件的偏移量
	repldbOff off_t /* replication DB file offset */
	// 主服务器传来的 RDB 文件的大小
	repldbSize off_t /* replication DB file size */

	replpreamble string /* replication DB preamble. */

	// 主服务器的复制偏移量
	reploff int64 /* replication offset if this is our master */
	// 从服务器最后一次发送 REPLCONF ACK 时的偏移量
	replAckOff int64 /* replication ack offset, if this is a slave */
	// 从服务器最后一次发送 REPLCONF ACK 的时间
	replAckTime int64 /* replication ack time, if this is a slave */
	// 主服务器的 master run ID
	// 保存在客户端，用于执行部分重同步
	replrunid []rune /* master run id if this is a master */
	// 从服务器的监听端口号
	slaveListeningPort int /* As configured with: SLAVECONF listening-port */

	// 事务状态
	mstate multiState /* MULTI/EXEC state */

	// 阻塞类型
	btype int /* Type of blocking op if gkv_BLOCKED. */
	// 阻塞状态
	bpop blockingState /* blocking state */

	// 最后被写入的全局复制偏移量
	woff int64 /* Last write global replication offset. */

	// 被监视的键
	watchedKeys *list.List /* Keys WATCHED for MULTI/EXEC CAS */

	// 这个字典记录了客户端所有订阅的频道
	// 键为频道名字，值为 NULL
	// 也即是，一个频道的集合
	pubsubChannels *dict /* channels a client is interested in (SUBSCRIBE) */

	// 链表，包含多个 pubsubPattern 结构
	// 记录了所有订阅频道的客户端的信息
	// 新 pubsubPattern 结构总是被添加到表尾
	pubsubPatterns *list.List /* patterns a client is interested in (SUBSCRIBE) */
	peerid         string     /* Cached peer ID. */

	/* Response buffer */
	// 回复偏移量
	bufpos int /* Cached peer ID. */
	// 回复缓冲区
	buf []rune /* Cached peer ID. */
}

func newClient() *gkvServer {
	return &gkvServer{
		id:         1024,
		operations: make([]byte, 1024),
	}
}

/*
bio send
*/
func Send2Server(operation string, conn *net.TCPConn) {
	//log.Println("start send operation:"+string(operation)+ " to server... ")
	_, err := conn.Write([]byte(operation))
	utils.CheckErr(err)
	//log.Println("send operation:"+string(operation)+ " to server successfully!")
}

/*
bio receive
*/
func ReceiveFromServer(conn *net.TCPConn) string {
	//log.Println("start receive res FROM server....")
	//read data from TCPConn
	res := make([]byte, 1024)
	n, err := conn.Read(res)
	utils.CheckErr(err)
	//log.Printf(string(res[:n]))
	utils.CheckErr(err)
	//log.Println("receive res:"+string(res[:n])+ " FROM server successfully!")
	return string(res[:n])
}

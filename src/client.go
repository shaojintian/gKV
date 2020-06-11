package src

import (
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
	dictId int

	// 客户端的名字
	name           *gkvObject  /* As set by CLIENT SETNAME */

	// 查询缓冲区
	queryBuf	   *string

	// 查询缓冲区长度峰值
	querybufPeak   size_t/* Recent (100ms or more) peak of querybuf size */

	// 参数数量
	argc			int

	// 参数对象数组
	argv			**gkvObject

	// 记录被客户端执行的命令
	struct redisCommand *cmd, *lastcmd

		// 请求的类型：内联命令还是多条命令
		int reqtype

		// 剩余未读取的命令内容数量
		int multibulklen       /* number of multi bulk arguments left to read */

		// 命令内容的长度
		long bulklen           /* length of bulk argument in multi bulk request */

		// 回复链表
		list *reply

		// 回复链表中对象的总大小
		unsigned long reply_bytes /* Tot bytes of objects in reply list */

		// 已发送字节，处理 short write 用
		int sentlen            /* Amount of bytes already sent in the current
	   buffer or object being sent. */

		// 创建客户端的时间
		time_t ctime           /* Client creation time */

		// 客户端最后一次和服务器互动的时间
		time_t lastinteraction /* time of the last interaction, used for timeout */

		// 客户端的输出缓冲区超过软性限制的时间
		time_t obuf_soft_limit_reached_time

		// 客户端状态标志
		int flags              /* REDIS_SLAVE | REDIS_MONITOR | REDIS_MULTI ... */

		// 当 server.requirepass 不为 NULL 时
		// 代表认证的状态
		// 0 代表未认证， 1 代表已认证
		int authenticated      /* when requirepass is non-NULL */

		// 复制状态
		int replstate          /* replication state if this is a slave */
		// 用于保存主服务器传来的 RDB 文件的文件描述符
		int repldbfd           /* replication DB file descriptor */

		// 读取主服务器传来的 RDB 文件的偏移量
		off_t repldboff        /* replication DB file offset */
		// 主服务器传来的 RDB 文件的大小
		off_t repldbsize       /* replication DB file size */

		sds replpreamble       /* replication DB preamble. */

		// 主服务器的复制偏移量
		long long reploff      /* replication offset if this is our master */
		// 从服务器最后一次发送 REPLCONF ACK 时的偏移量
		long long repl_ack_off /* replication ack offset, if this is a slave */
		// 从服务器最后一次发送 REPLCONF ACK 的时间
		long long repl_ack_time/* replication ack time, if this is a slave */
		// 主服务器的 master run ID
		// 保存在客户端，用于执行部分重同步
		char replrunid[REDIS_RUN_ID_SIZE+1] /* master run id if this is a master */
		// 从服务器的监听端口号
		int slave_listening_port /* As configured with: SLAVECONF listening-port */

		// 事务状态
		multiState mstate      /* MULTI/EXEC state */

		// 阻塞类型
		int btype              /* Type of blocking op if REDIS_BLOCKED. */
		// 阻塞状态
		blockingState bpop     /* blocking state */

		// 最后被写入的全局复制偏移量
		long long woff         /* Last write global replication offset. */

		// 被监视的键
		list *watched_keys     /* Keys WATCHED for MULTI/EXEC CAS */

		// 这个字典记录了客户端所有订阅的频道
		// 键为频道名字，值为 NULL
		// 也即是，一个频道的集合
		dict *pubsub_channels  /* channels a client is interested in (SUBSCRIBE) */

		// 链表，包含多个 pubsubPattern 结构
		// 记录了所有订阅频道的客户端的信息
		// 新 pubsubPattern 结构总是被添加到表尾
		list *pubsub_patterns  /* patterns a client is interested in (SUBSCRIBE) */
		sds peerid             /* Cached peer ID. */

		/* Response buffer */
		// 回复偏移量
		int bufpos
		// 回复缓冲区
		char buf[REDIS_REPLY_CHUNK_BYTES]

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

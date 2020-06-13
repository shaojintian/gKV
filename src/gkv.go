package src

/*
* 跳跃表
 */
type zskiplist struct {

	// 表头节点和表尾节点
	header,	tail	*zskiplistNode

	// 表中节点的数量
	length	uint32

	// 表中层数最大的节点的层数
	level	int

}

type zskiplistNode struct {
	// 成员对象
	obj *gkvObject

	// 分值
	score double

	// 后退指针
	backward *zskiplistNode

	// 层
	zskiplistLevel struct {
		// 前进指针
		forward *zskiplistNode

		// 跨度
		span uint
	}
}
/*
* 事务命令
 */
type multiCmd struct {
	// 参数
	argv **robj

	// 参数数量
	argc int

	// 命令指针
	cmd *gkvCommand
}

type multiState struct {
	// 事务队列，FIFO 顺序
	commands *multiCmd /* Array of MULTI commands */

	// 已入队命令计数
	count              int    /* Total number of MULTI commands */
	minreplicas        int    /* MINREPLICAS for synchronous replication */
	minreplicasTimeout time_t /* MINREPLICAS timeout as unixtime. */
}

// 阻塞状态
type blockingState struct {
	/* Generic fields. */
	// 阻塞时限
	timeout mstime_t /* Blocking operation timeout. If UNIX current time
	* is > timeout then the operation timed out. */

	/* REDIS_BLOCK_LIST */
	// 造成阻塞的键
	keys *dict /* The keys we are waiting to terminate a blocking
	* operation such as BLPOP. Otherwise NULL. */
	// 在被阻塞的键有新元素进入时，需要将这些新元素添加到哪里的目标键
	// 用于 BRPOPLPUSH 命令
	target *robj /* The key that should receive the element,
	* for BRPOPLPUSH. */

	/* REDIS_BLOCK_WAIT */
	// 等待 ACK 的复制节点数量
	numreplicas int /* Number of replicas we are waiting for ACK. */
	// 复制偏移量
	reploffset int64 /* Replication offset to reach. */
}

// 客户端缓冲区限制
type clientBufferLimitsConfig struct {
	// 硬限制
	hardLimitBytes	uint64
	// 软限制
	softLimitBytes	uint64
	// 软限制时限
	softLimitSeconds	time_t
}
/* The redisOp structure defines a Redis Operation, that is an instance of
 * a command with an argument vector, database ID, propagation target
 * (REDIS_PROPAGATE_*), and command pointer.
 *
 * redisOp 结构定义了一个 Redis 操作，
 * 它包含指向被执行命令的指针、命令的参数、数据库 ID 、传播目标（REDIS_PROPAGATE_*）。
 *
 * Currently only used to additionally propagate more commands to AOF/Replication
 * after the propagation of the executed command.
 *
 * 目前只用于在传播被执行命令之后，传播附加的其他命令到 AOF 或 Replication 中。
 */
type redisOp struct {

	// 参数
	argv	**robj

	// 参数数量、数据库 ID 、传播目标
	argc,	dbid,	target	int

	// 被执行命令的指针
	cmd	*gkvCommand

}


type gkvOpArray struct {
	ops	*redisOp
	numops	int
}

package src

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


package src

import "container/list"

// 集群状态，每个节点都保存着一个这样的状态，记录了它们眼中的集群的样子。
// 另外，虽然这个结构主要用于记录集群的属性，但是为了节约资源，
// 有些与节点有关的属性，比如 slots_to_keys 、 failover_auth_count
// 也被放到了这个结构里面。
type clusterState struct {
	// 指向当前节点的指针
	myself *clusterNode /* This node */

	// 集群当前的配置纪元，用于实现故障转移
	currentEpoch uint64_t /* This node */

	// 集群当前的状态：是在线还是下线
	state int /* REDIS_CLUSTER_OK, REDIS_CLUSTER_FAIL, ... */

	// 集群中至少处理着一个槽的节点的数量。
	size int /* Num of master nodes with at least one slot */

	// 集群节点名单（包括 myself 节点）
	// 字典的键为节点的名字，字典的值为 clusterNode 结构
	nodes *dict /* Hash table of name -> clusterNode structures */

	// 节点黑名单，用于 CLUSTER FORGET 命令
	// 防止被 FORGET 的命令重新被添加到集群里面
	// （不过现在似乎没有在使用的样子，已废弃？还是尚未实现？）
	nodesBlackList *dict /* Nodes we don't re-add for a few seconds. */

	// 记录要从当前节点迁移到目标节点的槽，以及迁移的目标节点
	// migrating_slots_to[i] = NULL 表示槽 i 未被迁移
	// migrating_slots_to[i] = clusterNode_A 表示槽 i 要从本节点迁移至节点 A
	migratingSlotsTo []*clusterNode /* Nodes we don't re-add for a few seconds. */

	// 记录要从源节点迁移到本节点的槽，以及进行迁移的源节点
	// importing_slots_from[i] = NULL 表示槽 i 未进行导入
	// importing_slots_from[i] = clusterNode_A 表示正从节点 A 中导入槽 i
	importingSlotsFrom []*clusterNode /* Nodes we don't re-add for a few seconds. */

	// 负责处理各个槽的节点
	// 例如 slots[i] = clusterNode_A 表示槽 i 由节点 A 处理
	slots []*clusterNode /* Nodes we don't re-add for a few seconds. */

	// 跳跃表，表中以槽作为分值，键作为成员，对槽进行有序排序
	// 当需要对某些槽进行区间（range）操作时，这个跳跃表可以提供方便
	// 具体操作定义在 db.c 里面
	slotsToKeys *zskiplist /* Nodes we don't re-add for a few seconds. */

	/* The following fields are used to take the slave state on elections. */
	// 以下这些域被用于进行故障转移选举

	// 上次执行选举或者下次执行选举的时间
	failoverAuthTime mstime_t /* Time of previous or next election. */

	// 节点获得的投票数量
	failoverAuthCount int /* Number of votes received so far. */

	// 如果值为 1 ，表示本节点已经向其他节点发送了投票请求
	failoverAuthSent int /* True if we already asked for votes. */

	failoverAuthRank int /* This slave rank for current auth request. */

	failoverAuthEpoch uint64_t /* Epoch of the current election. */

	/* Manual failover state in common. */
	/* 共用的手动故障转移状态 */

	// 手动故障转移执行的时间限制
	mfEnd mstime_t /* Manual failover time limit (ms unixtime).
	It is zero if there is no MF in progress. */
	/* Manual failover state of master. */
	/* 主服务器的手动故障转移状态 */
	mfSlave *clusterNode /* Slave performing the manual failover. */
	/* Manual failover state of slave. */
	/* 从服务器的手动故障转移状态 */
	mfMasterOffset int64 /* Master offset the slave needs to start MF
	or zero if stil not received. */
	// 指示手动故障转移是否可以开始的标志值
	// 值为非 0 时表示各个主服务器可以开始投票
	mfCanStart int /* If non-zero signal that the manual failover
	can start requesting masters vote. */

	/* The followign fields are uesd by masters to take state on elections. */
	/* 以下这些域由主服务器使用，用于记录选举时的状态 */

	// 集群最后一次进行投票的纪元
	lastVoteEpoch uint64_t /* Epoch of the last vote granted. */

	// 在进入下个事件循环之前要做的事情，以各个 flag 来记录
	todoBeforeSleep int /* Things to do in clusterBeforeSleep(). */

	// 通过 cluster 连接发送的消息数量
	statsBusMessagesSent int64 /* Num of msg sent via cluster bus. */

	// 通过 cluster 接收到的消息数量
	statsBusMessagesReceived int64 /* Num of msg rcvd via cluster bus.*/
}

// 节点状态
type clusterNode struct {

	// 创建节点的时间
	ctime	mstime_t /* Node object creation time. */

	// 节点的名字，由 40 个十六进制字符组成
	// 例如 68eef66df23420a5862208ef5b1a7005b806f2ff
	name	[]rune /* Node name, hex string, sha1-size */

	// 节点标识
	// 使用各种不同的标识值记录节点的角色（比如主节点或者从节点），
	// 以及节点目前所处的状态（比如在线或者下线）。
	flags	int      /* REDIS_NODE_... */

	// 节点当前的配置纪元，用于实现故障转移
	configEpoch	uint64_t /* Last configEpoch observed for this node */

	// 由这个节点负责处理的槽
	// 一共有 REDIS_CLUSTER_SLOTS / 8 个字节长
	// 每个字节的每个位记录了一个槽的保存状态
	// 位的值为 1 表示槽正由本节点处理，值为 0 则表示槽并非本节点处理
	// 比如 slots[0] 的第一个位保存了槽 0 的保存情况
	// slots[0] 的第二个位保存了槽 1 的保存情况，以此类推
	slots[REDIS_CLUSTER_SLOTS/8]	urune /* slots handled by this node */

	// 该节点负责处理的槽数量
	numslots	int   /* Number of slots handled by this node */

	// 如果本节点是主节点，那么用这个属性记录从节点的数量
	numslaves	int  /* Number of slave nodes, if this is a master */

	// 指针数组，指向各个从节点
	slaves	**clusterNode /* pointers to slave nodes */

	// 如果这是一个从节点，那么指向主节点
	slaveof	*clusterNode /* pointer to the master node */

	// 最后一次发送 PING 命令的时间
	pingSent	mstime_t      /* Unix time we sent latest ping */

	// 最后一次接收 PONG 回复的时间戳
	pongReceived	mstime_t  /* Unix time we received the pong */

	// 最后一次被设置为 FAIL 状态的时间
	failTime	mstime_t      /* Unix time when FAIL flag was set */

	// 最后一次给某个从节点投票的时间
	votedTime	mstime_t     /* Last time we voted for a slave of this master */

	// 最后一次从这个节点接收到复制偏移量的时间
	replOffsetTime	mstime_t  /* Unix time we received offset for this node */

	// 这个节点的复制偏移量
	replOffset	int64      /* Last known repl offset for this node. */

	// 节点的 IP 地址
	ip	[]rune  /* Latest known IP address of this node */

	// 节点的端口号
	port	int                   /* Latest known port of this node */

	// 保存连接节点所需的有关信息
	link	*clusterLink          /* TCP/IP link with this node */

	// 一个链表，记录了所有其他节点对该节点的下线报告
	failReports	*list.List         /* List of nodes signaling this as failing */

}

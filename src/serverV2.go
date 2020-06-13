package src

import (
	"container/list"
	"time"
)

type gkvServerV2 struct {
	/* General */

	// 配置文件的绝对路径
	configfile string /* Absolute config file path, or NULL */

	// serverCron() 每秒调用的次数
	hz int /* serverCron() calls frequency in hertz */

	// 数据库
	db *gkvDB /* serverCron() calls frequency in hertz */

	// 命令表（受到 rename 配置选项的作用）
	commands *dict /* Command table */
	// 命令表（无 rename 配置选项的作用）
	origCommands *dict /* Command table before command renaming. */

	// 事件状态
	el *aeEventLoop /* Command table before command renaming. */

	// 最近一次使用时钟
	lruclock time.Time /* Clock for LRU eviction */

	// 关闭服务器的标识
	shutdownAsap int /* SHUTDOWN needed ASAP */

	// 在执行 serverCron() 时进行渐进式 rehash
	activerehashing int /* Incremental rehash in serverCron() */

	// 是否设置了密码
	requirepass string /* Pass for AUTH command, or NULL */

	// PID 文件
	pidfile string /* PID file path */

	// 架构类型
	archBits int /* 32 or 64 depending on sizeof(long) */

	// serverCron() 函数的运行次数计数器
	cronloops int /* Number of times the cron function run */

	// 本服务器的 RUN ID
	runid []rune /* ID always different at every exec. */

	// 服务器是否运行在 SENTINEL 模式
	sentinelMode int /* True if this instance is a Sentinel. */

	/* Networking */

	// TCP 监听端口
	port int /* TCP listening port */

	tcpBacklog int /* TCP listen() backlog */

	// 地址
	bindaddr []*rune /* Addresses we should bind to */
	// 地址数量
	bindaddrCount int /* Number of addresses in server.bindaddr[] */

	// UNIX 套接字
	unixsocket     string /* UNIX socket path */
	unixsocketperm mode_t /* UNIX socket permission */

	// 描述符
	ipfd []int /* TCP socket file descriptors */
	// 描述符数量
	ipfdCount int /* Used slots in ipfd[] */

	// UNIX 套接字文件描述符
	sofd int /* Unix socket file descriptor */

	cfd      []int /* Cluster bus listening socket */
	cfdCount int   /* Used slots in cfd[] */

	// 一个链表，保存了所有客户端状态结构
	clients *list.List /* List of active clients */
	// 链表，保存了所有待关闭的客户端
	clientsToClose *list.List /* Clients to close asynchronously */

	// 链表，保存了所有从服务器，以及所有监视器
	slaves, monitors *list.List /* List of slaves and MONITORs */

	// 服务器的当前客户端，仅用于崩溃报告
	currentClient *gkvClient /* Current client, only used on crash report */

	clientsPaused       int      /* True if clients are currently paused */
	clientsPauseEndTime mstime_t /* Time when we undo clients_paused */

	// 网络错误
	neterr []rune /* Error buffer for anet.c */

	// MIGRATE 缓存
	migrateCachedSockets *dict /* MIGRATE cached sockets */

	/* RDB / AOF loading information */

	// 这个值为真时，表示服务器正在进行载入
	loading int /* We are loading data from disk if true */

	// 正在载入的数据的大小
	loadingTotalBytes off_t /* We are loading data from disk if true */

	// 已载入数据的大小
	loadingLoadedBytes off_t /* We are loading data from disk if true */

	// 开始进行载入的时间
	loadingStartTime                  time_t /* We are loading data from disk if true */
	loadingProcessEventsIntervalBytes off_t  /* We are loading data from disk if true */

	/* Fast pointers to often looked up command */
	// 常用命令的快捷连接
	delCommand, multiCommand, lpushCommand, lpopCommand, rpopCommand *gkvCommand /* We are loading data from disk if true */

	/* Fields used only for stats */

	// 服务器启动时间
	statStarttime time_t /* Server start time */

	// 已处理命令的数量
	statNumcommands int64 /* Number of processed commands */

	// 服务器接到的连接请求数量
	statNumconnections int64 /* Number of connections received */

	// 已过期的键数量
	statExpiredkeys int64 /* Number of expired keys */

	// 因为回收内存而被释放的过期键的数量
	statEvictedkeys int64 /* Number of evicted keys (maxmemory) */

	// 成功查找键的次数
	statKeyspaceHits int64 /* Number of successful lookups of keys */

	// 查找键失败的次数
	statKeyspaceMisses int64 /* Number of failed lookups of keys */

	// 已使用内存峰值
	statPeakMemory size_t /* Max used memory record */

	// 最后一次执行 fork() 时消耗的时间
	statForkTime int64 /* Time needed to perform latest fork() */

	// 服务器因为客户端数量过多而拒绝客户端连接的次数
	statRejectedConn int64 /* Clients rejected because of maxclients */

	// 执行 full sync 的次数
	statSyncFull int64 /* Number of full resyncs with slaves. */

	// PSYNC 成功执行的次数
	statSyncPartialOk int64 /* Number of accepted PSYNC requests. */

	// PSYNC 执行失败的次数
	statSyncPartialErr int64 /* Number of unaccepted PSYNC requests. */

	/* slowlog */

	// 保存了所有慢查询日志的链表
	slowlog *list.List /* SLOWLOG list of commands */

	// 下一条慢查询日志的 ID
	slowlogEntryId int64 /* SLOWLOG current entry ID */

	// 服务器配置 slowlog-log-slower-than 选项的值
	slowlogLogSlowerThan int64 /* SLOWLOG time limit (to get logged) */

	// 服务器配置 slowlog-max-len 选项的值
	slowlogMaxLen   uint32 /* SLOWLOG max number of items logged */
	residentSetSize size_t /* RSS sampled in serverCron(). */
	/* The following two are used to track instantaneous "load" in terms
	 * of operations per second. */
	// 最后一次进行抽样的时间
	opsSecLastSampleTime int64 /* Timestamp of last sample (in ms) */
	// 最后一次抽样时，服务器已执行命令的数量
	opsSecLastSampleOps int64 /* numcommands in last sample */
	// 抽样结果
	opsSecSamples     []int64 /* numcommands in last sample */
	// 数组索引，用于保存抽样结果，并在需要时回绕到 0
	opsSecIdx int /* numcommands in last sample */

	/* Configuration */

	// 日志可见性
	verbosity int /* Loglevel in gkv.conf */

	// 客户端最大空转时间
	maxidletime int /* Client timeout in seconds */

	// 是否开启 SO_KEEPALIVE 选项
	tcpkeepalive         int    /* Set SO_KEEPALIVE if non-zero. */
	activeExpireEnabled  int    /* Can be disabled for testing purposes. */
	clientMaxQuerybufLen size_t /* Limit for client query buffer length */
	dbnum                int    /* Total number of configured DBs */
	daemonize            int    /* True if running as a daemon */
	// 客户端输出缓冲区大小限制
	// 数组的元素有 gkv_CLIENT_LIMIT_NUM_CLASSES 个
	// 每个代表一类客户端：普通、从服务器、pubsub，诸如此类
	clientObufLimits []clientBufferLimitsConfig /* True if running as a daemon */

	/* AOF persistence */

	// AOF 状态（开启/关闭/可写）
	aofState int /* gkv_AOF_(ON|OFF|WAIT_REWRITE) */

	// 所使用的 fsync 策略（每个写入/每秒/从不）
	aofFsync            int    /* Kind of fsync() policy */
	aofFilename         string /* Name of the AOF file */
	aofNoFsyncOnRewrite int    /* Don't fsync if a rewrite is in prog. */
	aofRewritePerc      int    /* Rewrite AOF if % growth is > M and... */
	aofRewriteMinSize   off_t  /* the AOF file is at least N bytes. */

	// 最后一次执行 BGREWRITEAOF 时， AOF 文件的大小
	aofRewriteBaseSize off_t /* AOF size on latest startup or rewrite. */

	// AOF 文件的当前字节大小
	aofCurrentSize      off_t /* AOF current size. */
	aofRewriteScheduled int   /* Rewrite once BGSAVE terminates. */

	// 负责进行 AOF 重写的子进程 ID
	aofChildPid pid_t /* PID if rewriting process */

	// AOF 重写缓存链表，链接着多个缓存块
	aofRewriteBufBlocks *list.List /* Hold changes during an AOF rewrite. */

	// AOF 缓冲区
	aofBuf string /* AOF buffer, written before entering the event loop */

	// AOF 文件的描述符
	aofFd int /* File descriptor of currently selected AOF file */

	// AOF 的当前目标数据库
	aofSelectedDb int /* Currently selected DB in AOF */

	// 推迟 write 操作的时间
	aofFlushPostponedStart time_t /* UNIX time of postponed AOF flush */

	// 最后一直执行 fsync 的时间
	aofLastFsync       time_t /* UNIX time of last fsync() */
	aofRewriteTimeLast time_t /* Time used by last AOF rewrite run. */

	// AOF 重写的开始时间
	aofRewriteTimeStart time_t /* Current AOF rewrite start time. */

	// 最后一次执行 BGREWRITEAOF 的结果
	aofLastbgrewriteStatus int /* gkv_OK or gkv_ERR */

	// 记录 AOF 的 write 操作被推迟了多少次
	aofDelayedFsync uint32 /* delayed AOF fsync() counter */

	// 指示是否需要每写入一定量的数据，就主动执行一次 fsync()
	aofRewriteIncrementalFsync int /* fsync incrementally while rewriting? */
	aofLastWriteStatus         int /* gkv_OK or gkv_ERR */
	aofLastWriteErrno          int /* Valid if aof_last_write_status is ERR */
	/* RDB persistence */

	// 自从上次 SAVE 执行以来，数据库被修改的次数
	dirty int64 /* Changes to DB from the last save */

	// BGSAVE 执行前的数据库被修改次数
	dirtyBeforeBgsave int64 /* Used to restore dirty on failed BGSAVE */

	// 负责执行 BGSAVE 的子进程的 ID
	// 没在执行 BGSAVE 时，设为 -1
	rdbChildPid    pid_t      /* PID of RDB saving child */
	saveparams     *saveparam /* Save points array for RDB */
	saveparamslen  int        /* Number of saving points */
	rdbFilename    string     /* Name of RDB file */
	rdbCompression int        /* Use compression in RDB? */
	rdbChecksum    int        /* Use RDB checksum? */

	// 最后一次完成 SAVE 的时间
	lastsave time_t /* Unix time of last successful save */

	// 最后一次尝试执行 BGSAVE 的时间
	lastbgsaveTry time_t /* Unix time of last attempted bgsave */

	// 最近一次 BGSAVE 执行耗费的时间
	rdbSaveTimeLast time_t /* Time used by last RDB save run. */

	// 数据库最近一次开始执行 BGSAVE 的时间
	rdbSaveTimeStart time_t /* Current RDB save start time. */

	// 最后一次执行 SAVE 的状态
	lastbgsaveStatus      int /* gkv_OK or gkv_ERR */
	stopWritesOnBgsaveErr int /* Don't allow writes if can't BGSAVE */

	/* Propagation of commands in AOF / replication */
	alsoPropagate gkvOpArray /* Additional command to propagate. */

	/* Logging */
	logfile        string /* Path of log file */
	syslogEnabled  int    /* Is syslog enabled? */
	syslogIdent    string /* Syslog ident */
	syslogFacility int    /* Syslog facility */

	/* Replication (master) */
	slaveseldb int /* Last SELECTed DB in replication output */
	// 全局复制偏移量（一个累计值）
	masterReplOffset int64 /* Global replication offset */
	// 主服务器发送 PING 的频率
	replPingSlavePeriod int /* Master pings the slave every N seconds */

	// backlog 本身
	replBacklog string /* Replication backlog for partial syncs */
	// backlog 的长度
	replBacklogSize int64 /* Backlog circular buffer size */
	// backlog 中数据的长度
	replBacklogHistlen int64 /* Backlog actual data length */
	// backlog 的当前索引
	replBacklogIdx int64 /* Backlog circular buffer current offset */
	// backlog 中可以被还原的第一个字节的偏移量
	replBacklogOff int64 /* Replication offset of first byte in the
	backlog buffer. */
	// backlog 的过期时间
	replBacklogTimeLimit time_t /* Time without slaves after the backlog
	gets released. */

	// 距离上一次有从服务器的时间
	replNoSlavesSince time_t /* We have no slaves since that time.
	Only valid if server.slaves len is 0. */

	// 是否开启最小数量从服务器写入功能
	replMinSlavesToWrite int /* Min number of slaves to write. */
	// 定义最小数量从服务器的最大延迟值
	replMinSlavesMaxLag int /* Max lag of <count> slaves to write. */
	// 延迟良好的从服务器的数量
	replGoodSlavesCount int /* Number of slaves with lag <= max_lag. */

	/* Replication (slave) */
	// 主服务器的验证密码
	masterauth string /* AUTH with this password with master */
	// 主服务器的地址
	masterhost string /* Hostname of master */
	// 主服务器的端口
	masterport int /* Port of master */
	// 超时时间
	replTimeout int /* Timeout after N seconds of master idle */
	// 主服务器所对应的客户端
	master *gkvClient /* Client that is master for this slave */
	// 被缓存的主服务器，PSYNC 时使用
	cachedMaster      *gkvClient /* Cached master to be reused for PSYNC. */
	replSyncioTimeout int        /* Timeout for synchronous I/O calls */
	// 复制的状态（服务器是从服务器时使用）
	replState int /* Replication status if the instance is a slave */
	// RDB 文件的大小
	replTransferSize off_t /* Size of RDB to read from master during sync. */
	// 已读 RDB 文件内容的字节数
	replTransferRead off_t /* Amount of RDB read from master during sync. */
	// 最近一次执行 fsync 时的偏移量
	// 用于 sync_file_range 函数
	replTransferLastFsyncOff off_t /* Offset when we fsync-ed last time. */
	// 主服务器的套接字
	replTransferS int /* Slave -> Master SYNC socket */
	// 保存 RDB 文件的临时文件的描述符
	replTransferFd int /* Slave -> Master SYNC temp file descriptor */
	// 保存 RDB 文件的临时文件名字
	replTransferTmpfile string /* Slave-> master SYNC temp file name */
	// 最近一次读入 RDB 内容的时间
	replTransferLastio time_t /* Unix time of the latest read, for timeout */
	replServeStaleData int    /* Serve stale data when link is down? */
	// 是否只读从服务器？
	replSlaveRo int /* Slave is read only? */
	// 连接断开的时长
	replDownSince time_t /* Unix time at which link with master went down */
	// 是否要在 SYNC 之后关闭 NODELAY ？
	replDisableTcpNodelay int /* Disable TCP_NODELAY after SYNC? */
	// 从服务器优先级
	slavePriority int /* Reported in INFO and used by Sentinel. */
	// 本服务器（从服务器）当前主服务器的 RUN ID
	replMasterRunid []rune /* Master run id for PSYNC. */
	// 初始化偏移量
	replMasterInitialOffset int64 /* Master PSYNC offset. */

	/* Replication script cache. */
	// 复制脚本缓存
	// 字典
	replScriptcacheDict *dict /* SHA1 all slaves are aware of. */
	// FIFO 队列
	replScriptcacheFifo *list.List /* First in, first out LRU eviction. */
	// 缓存的大小
	replScriptcacheSize int /* Max number of elements. */

	/* Synchronous replication. */
	clientsWaitingAcks *list.List /* Clients waiting in WAIT command. */
	getAckFromSlaves   int        /* If true we send REPLCONF GETACK. */
	/* Limits */
	maxclients       int    /* Max number of simultaneous clients */
	maxmemory        uint64 /* Max number of memory bytes to use */
	maxmemoryPolicy  int    /* Policy for key eviction */
	maxmemorySamples int    /* Pricision of random sampling */

	/* Blocked clients */
	bpopBlockedClients uint       /* Number of clients blocked by lists */
	unblockedClients   *list.List /* list of clients to unblock before next loop */
	readyKeys          *list.List /* List of readyList structures for BLPOP & co */

	/* Sort parameters - qsort_r() is only available under BSD so we
	 * have to take this state global, in order to pass it to sortCompare() */
	sortDesc      int /* List of readyList structures for BLPOP & co */
	sortAlpha     int /* List of readyList structures for BLPOP & co */
	sortBypattern int /* List of readyList structures for BLPOP & co */
	sortStore     int /* List of readyList structures for BLPOP & co */

	/* Zip structure config, see gkv.conf for more information  */
	hashMaxZiplistEntries size_t /* List of readyList structures for BLPOP & co */
	hashMaxZiplistValue   size_t /* List of readyList structures for BLPOP & co */
	listMaxZiplistEntries size_t /* List of readyList structures for BLPOP & co */
	listMaxZiplistValue   size_t /* List of readyList structures for BLPOP & co */
	setMaxIntsetEntries   size_t /* List of readyList structures for BLPOP & co */
	zsetMaxZiplistEntries size_t /* List of readyList structures for BLPOP & co */
	zsetMaxZiplistValue   size_t /* List of readyList structures for BLPOP & co */
	hllSparseMaxBytes     size_t /* List of readyList structures for BLPOP & co */
	unixtime              time_t /* Unix time sampled every cron cycle. */
	mstime                int64  /* Like 'unixtime' but with milliseconds resolution. */

	/* Pubsub */
	// 字典，键为频道，值为链表
	// 链表中保存了所有订阅某个频道的客户端
	// 新客户端总是被添加到链表的表尾
	pubsubChannels *dict /* Map channels to list of subscribed clients */

	// 这个链表记录了客户端订阅的所有模式的名字
	pubsubPatterns *list.List /* A list of pubsub_patterns */

	notifyKeyspaceEvents int /* Events to propagate via Pub/Sub. This is an
	xor of gkv_NOTIFY... flags. */

	/* Cluster */

	clusterEnabled     int           /* Is cluster enabled? */
	clusterNodeTimeout mstime_t      /* Cluster node timeout. */
	clusterConfigfile  string        /* Cluster auto-generated config file name. */
	cluster            *clusterState /* State of the cluster */

	clusterMigrationBarrier int /* Cluster replicas migration barrier. */
	/* Scripting */

	// Lua 环境
	lua *lua_State /* The Lua interpreter. We use just one for all clients */

	// 复制执行 Lua 脚本中的 gkv 命令的伪客户端
	luaClient *gkvClient /* The "fake client" to query gkv from Lua */

	// 当前正在执行 EVAL 命令的客户端，如果没有就是 NULL
	luaCaller *gkvClient /* The client running EVAL right now, or NULL */

	// 一个字典，值为 Lua 脚本，键为脚本的 SHA1 校验和
	luaScripts *dict /* A dictionary of SHA1 -> Lua scripts */
	// Lua 脚本的执行时限
	luaTimeLimit mstime_t /* Script timeout in milliseconds */
	// 脚本开始执行的时间
	luaTimeStart mstime_t /* Start time of script, milliseconds time */

	// 脚本是否执行过写命令
	luaWriteDirty int /* True if a write command was called during the
	execution of the current script. */

	// 脚本是否执行过带有随机性质的命令
	luaRandomDirty int /* True if a random command was called during the
	execution of the current script. */

	// 脚本是否超时
	luaTimedout int /* True if we reached the time limit for script
	execution. */

	// 是否要杀死脚本
	luaKill int /* Kill the script if true. */

	/* Assert & bug reporting */

	assertFailed   string /* Kill the script if true. */
	assertFile     string /* Kill the script if true. */
	assertLine     int    /* Kill the script if true. */
	bugReportStart int    /* True if bug report header was already logged. */
	watchdogPeriod int    /* Software watchdog period in ms. 0 = off */
}

// 服务器的保存条件（BGSAVE 自动执行的条件）
type saveparam struct {
	// 多少秒之内
	seconds time_t

	// 发生多少次修改
	changes int
}

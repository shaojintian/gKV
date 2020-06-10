package src

import "container/list"

type gkvServerV2 struct {
	// 配置文件的绝对路径
	configFile      *string     /* Absolute config file path, or NULL */

	// serverCron() 每秒调用的次数
	hz          	int           /* serverCron() calls frequency in hertz */

	// 数据库
	db 				*gkvDB

	// 命令表（受到 rename 配置选项的作用）
	commands  		*dict          /* Command table */
	// 命令表（无 rename 配置选项的作用）
	origCommands   	*dict    /* Command table before command renaming. */

	// 事件状态
	el				*aeEventLoop

	// 最近一次使用时钟
	//unsigned lruclock:REDIS_LRU_BITS /* Clock for LRU eviction */

	// 关闭服务器的标识
	shutdownAsap   int       /* SHUTDOWN needed ASAP */

	// 在执行 serverCron() 时进行渐进式 rehash
	activeRehashing      int  /* Incremental rehash in serverCron() */

	// 是否设置了密码
	requirePass       	*string   /* Pass for AUTH command, or NULL */

	// PID 文件
	pidFile             *string /* PID file path */

	// 架构类型
	archBits            int  /* 32 or 64 depending on sizeof(long) */

	// serverCron() 函数的运行次数计数器
	cronLoops         	int     /* Number of times the cron function run */

	// 本服务器的 RUN ID
	runId				[GKV_RUN_ID_SIZE+1]rune  /* ID always different at every exec. */

	// 服务器是否运行在 SENTINEL 模式
	sentinelMode          int/* True if this instance is a Sentinel. */


	/* Networking */

	// TCP 监听端口
	port       			int            /* TCP listening port */

	tcpBacklog          int  /* TCP listen() backlog */

	// 地址
	bindAddr			*[GKV_BINDADDR_MAX]rune /* Addresses we should bind to */
	// 地址数量
	bindAddrCount       int /* Number of addresses in server.bindaddr[] */

	// UNIX 套接字
	unixSocket          *string /* UNIX socket path */
	unixSocketPerm      modeT/* UNIX socket permission */

	// 描述符
	ipFd				[GKV_BINDADDR_MAX]int /* TCP socket file descriptors */
	// 描述符数量
	ipFdCount           int  /* Used slots in ipfd[] */

	// UNIX 套接字文件描述符
	soFd                int   /* Unix socket file descriptor */

	cfd					[GKV_BINDADDR_MAX]int/* Cluster bus listening socket */
	cfdCount            	int   /* Used slots in cfd[] */

	// 一个链表，保存了所有客户端状态结构
	clients              *list.List/* List of active clients */
	// 链表，保存了所有待关闭的客户端
	clientsToClose     	*list.List/* Clients to close asynchronously */

	// 链表，保存了所有从服务器，以及所有监视器
	slaves, monitors    *list.List/* List of slaves and Monitors */

	// 服务器的当前客户端，仅用于崩溃报告
	redisClient *current_client /* Current client, only used on crash report */

	int clients_paused         /* True if clients are currently paused */
	mstime_t clients_pause_end_time /* Time when we undo clients_paused */

	// 网络错误
	char neterr[ANET_ERR_LEN]   /* Error buffer for anet.c */

	// MIGRATE 缓存
	dict *migrate_cached_sockets/* MIGRATE cached sockets */


	/* RDB / AOF loading information */

	// 这个值为真时，表示服务器正在进行载入
	int loading                /* We are loading data from disk if true */

	// 正在载入的数据的大小
	off_t loading_total_bytes

	// 已载入数据的大小
	off_t loading_loaded_bytes

	// 开始进行载入的时间
	time_t loading_start_time
	off_t loading_process_events_interval_bytes

	/* Fast pointers to often looked up command */
	// 常用命令的快捷连接
struct redisCommand *delCommand, *multiCommand, *lpushCommand, *lpopCommand,
	*rpopCommand


	/* Fields used only for stats */

	// 服务器启动时间
	time_t stat_starttime          /* Server start time */

	// 已处理命令的数量
	long long stat_numcommands     /* Number of processed commands */

	// 服务器接到的连接请求数量
	long long stat_numconnections  /* Number of connections received */

	// 已过期的键数量
	long long stat_expiredkeys     /* Number of expired keys */

	// 因为回收内存而被释放的过期键的数量
	long long stat_evictedkeys     /* Number of evicted keys (maxmemory) */

	// 成功查找键的次数
	long long stat_keyspace_hits   /* Number of successful lookups of keys */

	// 查找键失败的次数
	long long stat_keyspace_misses /* Number of failed lookups of keys */

	// 已使用内存峰值
	size_t stat_peak_memory        /* Max used memory record */

	// 最后一次执行 fork() 时消耗的时间
	long long stat_fork_time       /* Time needed to perform latest fork() */

	// 服务器因为客户端数量过多而拒绝客户端连接的次数
	long long stat_rejected_conn   /* Clients rejected because of maxclients */

	// 执行 full sync 的次数
	long long stat_sync_full       /* Number of full resyncs with slaves. */

	// PSYNC 成功执行的次数
	long long stat_sync_partial_ok /* Number of accepted PSYNC requests. */

	// PSYNC 执行失败的次数
	long long stat_sync_partial_err/* Number of unaccepted PSYNC requests. */


	/* slowlog */

	// 保存了所有慢查询日志的链表
	list *slowlog                  /* SLOWLOG list of commands */

	// 下一条慢查询日志的 ID
	long long slowlog_entry_id     /* SLOWLOG current entry ID */

	// 服务器配置 slowlog-log-slower-than 选项的值
	long long slowlog_log_slower_than /* SLOWLOG time limit (to get logged) */

	// 服务器配置 slowlog-max-len 选项的值
	unsigned long slowlog_max_len     /* SLOWLOG max number of items logged */
	size_t resident_set_size       /* RSS sampled in serverCron(). */
	/* The following two are used to track instantaneous "load" in terms
	 * of operations per second. */
	// 最后一次进行抽样的时间
	long long ops_sec_last_sample_time /* Timestamp of last sample (in ms) */
	// 最后一次抽样时，服务器已执行命令的数量
	long long ops_sec_last_sample_ops  /* numcommands in last sample */
	// 抽样结果
	long long ops_sec_samples[REDIS_OPS_SEC_SAMPLES]
	// 数组索引，用于保存抽样结果，并在需要时回绕到 0
	int ops_sec_idx


	/* Configuration */

	// 日志可见性
	int verbosity                  /* Loglevel in redis.conf */

	// 客户端最大空转时间
	int maxidletime                /* Client timeout in seconds */

	// 是否开启 SO_KEEPALIVE 选项
	int tcpkeepalive               /* Set SO_KEEPALIVE if non-zero. */
	int active_expire_enabled      /* Can be disabled for testing purposes. */
	size_t client_max_querybuf_len /* Limit for client query buffer length */
	int dbnum                      /* Total number of configured DBs */
	int daemonize                  /* True if running as a daemon */
	// 客户端输出缓冲区大小限制
	// 数组的元素有 REDIS_CLIENT_LIMIT_NUM_CLASSES 个
	// 每个代表一类客户端：普通、从服务器、pubsub，诸如此类
	clientBufferLimitsConfig client_obuf_limits[REDIS_CLIENT_LIMIT_NUM_CLASSES]


	/* AOF persistence */

	// AOF 状态（开启/关闭/可写）
	int aof_state                  /* REDIS_AOF_(ON|OFF|WAIT_REWRITE) */

	// 所使用的 fsync 策略（每个写入/每秒/从不）
	int aof_fsync                  /* Kind of fsync() policy */
	char *aof_filename             /* Name of the AOF file */
	int aof_no_fsync_on_rewrite    /* Don't fsync if a rewrite is in prog. */
	int aof_rewrite_perc           /* Rewrite AOF if % growth is > M and... */
	off_t aof_rewrite_min_size     /* the AOF file is at least N bytes. */

	// 最后一次执行 BGREWRITEAOF 时， AOF 文件的大小
	off_t aof_rewrite_base_size    /* AOF size on latest startup or rewrite. */

	// AOF 文件的当前字节大小
	off_t aof_current_size         /* AOF current size. */
	int aof_rewrite_scheduled      /* Rewrite once BGSAVE terminates. */

	// 负责进行 AOF 重写的子进程 ID
	pid_t aof_child_pid            /* PID if rewriting process */

	// AOF 重写缓存链表，链接着多个缓存块
	list *aof_rewrite_buf_blocks   /* Hold changes during an AOF rewrite. */

	// AOF 缓冲区
	sds aof_buf      /* AOF buffer, written before entering the event loop */

	// AOF 文件的描述符
	int aof_fd       /* File descriptor of currently selected AOF file */

	// AOF 的当前目标数据库
	int aof_selected_db /* Currently selected DB in AOF */

	// 推迟 write 操作的时间
	time_t aof_flush_postponed_start /* UNIX time of postponed AOF flush */

	// 最后一直执行 fsync 的时间
	time_t aof_last_fsync            /* UNIX time of last fsync() */
	time_t aof_rewrite_time_last   /* Time used by last AOF rewrite run. */

	// AOF 重写的开始时间
	time_t aof_rewrite_time_start  /* Current AOF rewrite start time. */

	// 最后一次执行 BGREWRITEAOF 的结果
	int aof_lastbgrewrite_status   /* REDIS_OK or REDIS_ERR */

	// 记录 AOF 的 write 操作被推迟了多少次
	unsigned long aof_delayed_fsync  /* delayed AOF fsync() counter */

	// 指示是否需要每写入一定量的数据，就主动执行一次 fsync()
	int aof_rewrite_incremental_fsync/* fsync incrementally while rewriting? */
	int aof_last_write_status      /* REDIS_OK or REDIS_ERR */
	int aof_last_write_errno       /* Valid if aof_last_write_status is ERR */
	/* RDB persistence */

	// 自从上次 SAVE 执行以来，数据库被修改的次数
	long long dirty                /* Changes to DB from the last save */

	// BGSAVE 执行前的数据库被修改次数
	long long dirty_before_bgsave  /* Used to restore dirty on failed BGSAVE */

	// 负责执行 BGSAVE 的子进程的 ID
	// 没在执行 BGSAVE 时，设为 -1
	pid_t rdb_child_pid            /* PID of RDB saving child */
struct saveparam *saveparams   /* Save points array for RDB */
	int saveparamslen              /* Number of saving points */
	char *rdb_filename             /* Name of RDB file */
	int rdb_compression            /* Use compression in RDB? */
	int rdb_checksum               /* Use RDB checksum? */

	// 最后一次完成 SAVE 的时间
	time_t lastsave                /* Unix time of last successful save */

	// 最后一次尝试执行 BGSAVE 的时间
	time_t lastbgsave_try          /* Unix time of last attempted bgsave */

	// 最近一次 BGSAVE 执行耗费的时间
	time_t rdb_save_time_last      /* Time used by last RDB save run. */

	// 数据库最近一次开始执行 BGSAVE 的时间
	time_t rdb_save_time_start     /* Current RDB save start time. */

	// 最后一次执行 SAVE 的状态
	int lastbgsave_status          /* REDIS_OK or REDIS_ERR */
	int stop_writes_on_bgsave_err  /* Don't allow writes if can't BGSAVE */


	/* Propagation of commands in AOF / replication */
	redisOpArray also_propagate    /* Additional command to propagate. */


	/* Logging */
	char *logfile                  /* Path of log file */
	int syslog_enabled             /* Is syslog enabled? */
	char *syslog_ident             /* Syslog ident */
	int syslog_facility            /* Syslog facility */


	/* Replication (master) */
	int slaveseldb                 /* Last SELECTed DB in replication output */
	// 全局复制偏移量（一个累计值）
	long long master_repl_offset   /* Global replication offset */
	// 主服务器发送 PING 的频率
	int repl_ping_slave_period     /* Master pings the slave every N seconds */

	// backlog 本身
	char *repl_backlog             /* Replication backlog for partial syncs */
	// backlog 的长度
	long long repl_backlog_size    /* Backlog circular buffer size */
	// backlog 中数据的长度
	long long repl_backlog_histlen /* Backlog actual data length */
	// backlog 的当前索引
	long long repl_backlog_idx     /* Backlog circular buffer current offset */
	// backlog 中可以被还原的第一个字节的偏移量
	long long repl_backlog_off     /* Replication offset of first byte in the
	   backlog buffer. */
	// backlog 的过期时间
	time_t repl_backlog_time_limit /* Time without slaves after the backlog
	   gets released. */

	// 距离上一次有从服务器的时间
	time_t repl_no_slaves_since    /* We have no slaves since that time.
	   Only valid if server.slaves len is 0. */

	// 是否开启最小数量从服务器写入功能
	int repl_min_slaves_to_write   /* Min number of slaves to write. */
	// 定义最小数量从服务器的最大延迟值
	int repl_min_slaves_max_lag    /* Max lag of <count> slaves to write. */
	// 延迟良好的从服务器的数量
	int repl_good_slaves_count     /* Number of slaves with lag <= max_lag. */


	/* Replication (slave) */
	// 主服务器的验证密码
	char *masterauth               /* AUTH with this password with master */
	// 主服务器的地址
	char *masterhost               /* Hostname of master */
	// 主服务器的端口
	int masterport                 /* Port of master */
	// 超时时间
	int repl_timeout               /* Timeout after N seconds of master idle */
	// 主服务器所对应的客户端
	redisClient *master     /* Client that is master for this slave */
	// 被缓存的主服务器，PSYNC 时使用
	redisClient *cached_master /* Cached master to be reused for PSYNC. */
	int repl_syncio_timeout /* Timeout for synchronous I/O calls */
	// 复制的状态（服务器是从服务器时使用）
	int repl_state          /* Replication status if the instance is a slave */
	// RDB 文件的大小
	off_t repl_transfer_size /* Size of RDB to read from master during sync. */
	// 已读 RDB 文件内容的字节数
	off_t repl_transfer_read /* Amount of RDB read from master during sync. */
	// 最近一次执行 fsync 时的偏移量
	// 用于 sync_file_range 函数
	off_t repl_transfer_last_fsync_off /* Offset when we fsync-ed last time. */
	// 主服务器的套接字
	int repl_transfer_s     /* Slave -> Master SYNC socket */
	// 保存 RDB 文件的临时文件的描述符
	int repl_transfer_fd    /* Slave -> Master SYNC temp file descriptor */
	// 保存 RDB 文件的临时文件名字
	char *repl_transfer_tmpfile /* Slave-> master SYNC temp file name */
	// 最近一次读入 RDB 内容的时间
	time_t repl_transfer_lastio /* Unix time of the latest read, for timeout */
	int repl_serve_stale_data /* Serve stale data when link is down? */
	// 是否只读从服务器？
	int repl_slave_ro          /* Slave is read only? */
	// 连接断开的时长
	time_t repl_down_since /* Unix time at which link with master went down */
	// 是否要在 SYNC 之后关闭 NODELAY ？
	int repl_disable_tcp_nodelay   /* Disable TCP_NODELAY after SYNC? */
	// 从服务器优先级
	int slave_priority             /* Reported in INFO and used by Sentinel. */
	// 本服务器（从服务器）当前主服务器的 RUN ID
	char repl_master_runid[REDIS_RUN_ID_SIZE+1]  /* Master run id for PSYNC. */
	// 初始化偏移量
	long long repl_master_initial_offset         /* Master PSYNC offset. */

}

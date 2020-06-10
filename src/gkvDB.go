package src


type gkvDB struct {
	// 数据库键空间，保存着数据库中的所有键值对
	dict			map[string]interface{}              	/* The keyspace for this DB */

	// 键的过期时间，字典的键为键，字典的值为过期事件 UNIX 时间戳
	expires     	map[string]time_t      					/* Timeout of keys with a timeout set */

	// 正处于阻塞状态的键
	blockingKeys    map[string]interface{} 					/* Keys with clients waiting for data (BLPOP) */

	// 可以解除阻塞的键
	readyKeys      	map[string]interface{}  					/* Blocked keys that received a PUSH */

	// 正在被 WATCH 命令监视的键
	watchedKeys    	map[string]interface{}     				/* WATCHED keys for MULTI/EXEC CAS */

	evictionPool   	evictionPoolEntry  			/* Eviction pool of keys */

	// 数据库号码
	id             	int     /* Database ID */

	// 数据库的键的平均 TTL ，统计信息
	avgTTL 			int64          /* Average TTL, just for stats */
}

type evictionPoolEntry struct {
	idle uint64
	key  *rune
}

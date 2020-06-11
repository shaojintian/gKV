package src

type gkvCommand struct {
	// 命令名字
	name string

	// 实现函数
	proc *gkvCommandProc

	// 参数个数
	arity int

	// 字符串表示的 FLAG
	sFlags string /* Flags as string representation, one char per flag. */

	// 实际 FLAG
	flags int /* The actual flags, obtained from the 'sflags' field. */

	/* Use a function to determine keys arguments in a command line.
	 * Used for Redis Cluster redirect. */
	// 从命令中判断命令的键参数。在 Redis 集群转向时使用。
	getKeysProc *gkvGetKeysProc

	/* What keys should be loaded in background when calling this command? */
	// 指定哪些参数是 key
	firstKey int/* The first argument that's a key (0 = no keys) */
	lastKey  int /* The last argument that's a key */
	keyStep  int/* The step between first and last key */

	// 统计信息
	// microseconds 记录了命令执行耗费的总毫微秒数
	// calls 是命令被执行的总次数
	microseconds, calls  int64
}

type gkvCommandProc func(c *gkvClient)
type gkvGetKeysProc func(cmd *gkvCommand, argv **gkvObject, argc int, numKeys int)

package src

/*
A event-driven library.  (ae means A Event-driven lib)

TimeEvent:

FileEvent:
 */

type aeEventLoop struct {

	// 目前已注册的最大描述符
	maxFd    			int				/* highest file descriptor currently registered */

	// 目前已追踪的最大描述符
	setSize				int			 	/* max number of file descriptors tracked */

	// 用于生成时间事件 id
	timeEventNextId 	int64

	// 最后一次执行时间事件的时间
	lastTime     		time_t					/* Used to detect system clock skew */

	// 已注册的文件事件
	events 				*aeFileEvent				/* Registered events */

	// 已就绪的文件事件
	fired 				*aeFiredEvent			/* Fired events */

	// 时间事件
	timeEventHead		*aeTimeEvent

	// 事件处理器的开关
	stop				int

	// 多路复用库的私有数据
	apiData 			*interface{}/* This is used for polling API specific data */

	// 在处理事件前要执行的函数
	beforeSleep			*aeBeforeSleepProc
}

type aeFileEvent struct {
	// 监听事件类型掩码，
	// 值可以是 AE_READABLE 或 AE_WRITABLE ，
	// 或者 AE_READABLE | AE_WRITABLE
	mask 	int/* one of AE_(READABLE|WRITABLE) */

	// 读事件处理器
	rfileProc 	aeFileProc

	// 写事件处理器
	wfileProc	aeFileProc

	// 多路复用库的私有数据
	clientData *interface{}

}

type aeTimeEvent struct {
	// 时间事件的唯一标识符
	id  		int64/* time event identifier. */

	// 事件的到达时间
	whenSec	 	int32/* seconds */
	whenMs  	int32/* milliseconds */

	// 事件处理函数
	timeProc	*aeTimeProc

	// 事件释放函数
	finalizerProc *aeEventFinalizerProc

	// 多路复用库的私有数据
	clientData   *interface{}

	// 指向下个时间事件结构，形成链表
	next		*aeTimeEvent

}

type aeFiredEvent struct {
	// 已就绪文件描述符
	fd  int

	// 事件类型掩码，
	// 值可以是 AE_READABLE 或 AE_WRITABLE
	// 或者是两者的或
	mask int


}


/* Types and data structures
 *
 * 事件接口
 */
type aeBeforeSleepProc func(eventLoop *aeEventLoop)
type aeFileProc func(eventLoop *aeEventLoop,fd int,clientData *interface{},mask int)
type aeTimeProc func(eventLoop *aeEventLoop,id int64,clientData *interface{})
type aeEventFinalizerProc func(eventLoop *aeEventLoop,clientData *interface{})



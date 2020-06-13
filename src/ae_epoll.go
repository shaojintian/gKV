package src

type aeApiState struct {

	// epoll_event 实例描述符
	epfd	int

	// 事件槽
	events	*epollEvent

}

type epollEvent struct {

}
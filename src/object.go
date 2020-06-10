package src

type gkvObject struct {
	// 类型
	typ  uint8

	// 编码
	encoding uint8

	// 对象最后一次被访问的时间
	lru  uint16

	// 引用计数
	refCount	int

	// 指向实际值的指针
	ptr *interface{}

}


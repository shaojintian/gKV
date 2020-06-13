package src

type aofrwblock struct {

	// 缓存块已使用字节数和可用字节数
	used,	free	uint32

	// 缓存块
	buf	[]rune

}


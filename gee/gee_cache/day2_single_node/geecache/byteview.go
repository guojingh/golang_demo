package geecache

type ByteView struct {
	// 存储真正的缓存值，例如字符串，图片等
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b)
}

// b 是只读的，防止缓存值被外部程序修改
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

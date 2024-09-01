package geecache

// 不可变的字节视图
type ByteView struct {
	b []byte
}

// 实现 Value 接口，使得lru可以存储ByteView，并且可以拥有 Len() 方法
func (v ByteView) Len() int {
	return len(v.b)
}

// c.b 的复制品，防止被修改
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// v.b 的字符串形式
func (v ByteView) String() string {
	return string(v.b)
}

// 克隆一个字节数组
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

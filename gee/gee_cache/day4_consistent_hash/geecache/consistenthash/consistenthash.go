package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// Hash maps
type Hash func(data []byte) uint32

// Map 包含所有的 hash keys
type Map struct {
	hash     Hash           // hash 函数
	replicas int            // 虚拟节点倍数，虚拟节点主要用来解决当节点较少时产生的数据倾斜问题
	keys     []int          // 哈希环
	hashMap  map[int]string // 虚拟节点与真实节点的映射表，键是虚拟节点的哈希值，值是真实节点的名称
}

// 创建 Map 实例，允许自定义虚拟节点倍数和Hash函数
func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}

	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}

	return m
}

// 向 hash 中添加真实节点
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			// 将虚拟节点添加到环上
			m.keys = append(m.keys, hash)
			// 添加虚拟节点和真实节点的映射关系
			m.hashMap[hash] = key
		}
	}
	// 对环上的哈希值进行排序
	sort.Ints(m.keys)
}

func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	// 计算 key 的hash值
	hash := int(m.hash([]byte(key)))

	// 顺时针找到第一个匹配的虚拟节点的下标
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	// 通过 hashMap 找到真实节点
	return m.hashMap[m.keys[idx%len(m.keys)]]
}

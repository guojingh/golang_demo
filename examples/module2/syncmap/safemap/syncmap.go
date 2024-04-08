package safemap

/*
1.分布加锁的思路将大块的数据进行切分成小块的数据，从而减少冲突导致锁阻塞的可能性。
如果在一些特殊场景下，将读写数据分来，是不是能再进一步提高性能呢？
2.在 Go 1.9+ 内置的 sync 包中，有线程安全的 sync.map
3.这种方式一般不经常用，但是可能在某些场景下会有一些效率提升。官方给出的建议：
a) when the entry for a given key is only ever written once but read many times, as in caches that only grow.
b) when multiple goroutines read, write, and overwrite entries for disjoint sets of keys.
要么是一些多读，要么是多个协程操作的 key 集合没有交集。所以在使用前，应该做一下性能测试。
*/

# golang 实现并发安全的三种方式

1. `rwmutex.go` 通过 `RWMutex` 实现 `map `的高并发安全
2. 通过 分区加锁 实现 `map `的并发安全
   go社区常用 map 加锁算法：https://github.com/orcaman/concurrent-map?tab=readme-ov-file#concurrent-map-
3. 通过 `sync.map` 实现 `map `的并发安全

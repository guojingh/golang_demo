package singleton

// Hungry 单例模式 饿汉方式
type Hungry struct {
}

var hungry *Hungry = &Hungry{}

func GetInsOr() *Hungry {
	return hungry
}

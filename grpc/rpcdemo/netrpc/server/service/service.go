package service

type Args struct {
	X, Y int
}

// ServiceA 自定义的一个结构体
type ServiceA struct {
}

// Add 为ServiceA类型增加一个可导出的Add方法
func (s *ServiceA) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

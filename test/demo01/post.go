package demo01

type Post struct {
	Name    string
	Address string
}

type Service interface {
	ListPosts() ([]*Post, error)
}

func ListPosts(svc Service) ([]*Post, error) {
	return svc.ListPosts()
}

func main() {
	v := []string{"1", "2", "3"}

	for _ = range v {
	}
}

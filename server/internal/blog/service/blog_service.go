package service

// logic
type blogService struct{}

// new blog service create new blog service
func NewBlogService() *blogService {
	return &blogService{}
}

// list blogs return all blog service

func (s *blogService) ListBlogs() ([]any, error) {
	return []any{}, nil
}

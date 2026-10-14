package service

import "testing"

func TestBlogService_ListBlogs(t *testing.T) {

	t.Run("should return an empty blog list", func(t *testing.T) {

		blogService := NewBlogService()

		records, err := blogService.ListBlogs()

		if err != nil {
			t.Fatalf("expected no error , got %v", err)
		}

		if records == nil {
			t.Fatalf("experect an empty slice , got nil")
		}

		if len(records) != 0 {
			t.Fatalf("expected 0 blogs records , %v", len(records))
		}

	})
}

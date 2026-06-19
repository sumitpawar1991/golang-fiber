package model

import "testing"

func TestBlogModel(t *testing.T) {

	blog := Blog{
		Title: "First Blog",
		Post:  "This is first blog post",
	}

	if blog.Title == "" {
		t.Fatal("title should not be empty")
	}

	if blog.Post == "" {
		t.Fatal("Post should not be blank")
	}
}

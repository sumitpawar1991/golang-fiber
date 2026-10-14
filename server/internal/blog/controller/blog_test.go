package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestBlogList(t *testing.T) {

	t.Run("should return blog list", func(t *testing.T) {
		app := fiber.New()

		app.Get("/blog", BlogList)

		req := httptest.NewRequest(http.MethodGet, "/blog", nil)

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if resp.StatusCode != fiber.StatusOK {

			t.Fatalf(
				"expected status code %d, got %d",
				fiber.StatusOK,
				resp.StatusCode,
			)
		}

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			t.Fatalf("expected to read response body , got %v", err)
		}

		// expected := `{status:ok,message:blog list,blog_records:[]}`

		// if string(body) != expected {
		// 	t.Fatalf("expected body %s, got %s", expected, string(body))
		// }

		var result map[string]any

		if err := json.Unmarshal(body, &result); err != nil {
			t.Fatalf("expected valid json , got error: %v", err)
		}

		if result["status"] != "ok" {

			t.Errorf("expected status ok , got %v", result["status"])
		}

		if result["message"] != "blog list" {
			t.Errorf("expected message blog list , got %v", result["message"])
		}

		records, ok := result["blog_records"].([]any)

		if !ok {
			t.Fatalf("expected blog_results to an array")
		}

		if len(records) != 0 {
			t.Errorf("expected empty blog_results, got %v", records)
		}

	})
}

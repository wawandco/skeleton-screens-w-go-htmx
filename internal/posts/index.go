package posts

import (
	"net/http"

	"github.com/leapkit/leapkit/core/render"
	"github.com/leapkit/leapkit/core/server"
)

// Renders the skeleton posts page.
func Index(w http.ResponseWriter, r *http.Request) {
	rw := render.FromCtx(r.Context())

	err := rw.Render("posts/index.html")
	if err != nil {
		server.Errorf(w, http.StatusInternalServerError, "error rendering template: %s", err.Error())
	}
}

// Renders the posts list.
func All(w http.ResponseWriter, r *http.Request) {
	rw := render.FromCtx(r.Context())

	rw.Set("posts", AllPosts)

	err := rw.RenderClean("posts/posts.html")
	if err != nil {
		server.Errorf(w, http.StatusInternalServerError, "error rendering posts template: %s", err.Error())
	}
}

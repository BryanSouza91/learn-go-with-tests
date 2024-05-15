package blogrenderer

import (
	"embed"
	"html/template"
	"io"

	"github.com/BryanSouza91/learn-go-with-tests/blogposts"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.html")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p blogposts.Post) error {

	if err := r.templ.ExecuteTemplate(w, "blog.html", p); err != nil {
		return err
	}

	return nil
}

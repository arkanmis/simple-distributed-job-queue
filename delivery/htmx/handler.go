package htmx

import (
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

// HelloHandler handles HTMX view requests for the hello world page.
type HelloHandler struct{}

// NewHelloHandler creates a HelloHandler. No dependencies are required for a
// pure view that does not go through the service layer.
func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

// Page serves the full HTML shell for the hello world page.
func (h *HelloHandler) Page(c echo.Context) error {
	tmpl, err := template.ParseFiles("./web/htmx/hello.html")
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return tmpl.Execute(c.Response().Writer, nil)
}

// Message serves the HTMX HTML fragment swapped into the page on demand.
func (h *HelloHandler) Message(c echo.Context) error {
	return c.HTML(http.StatusOK, `<p id="message">Hello, World!</p>`)
}

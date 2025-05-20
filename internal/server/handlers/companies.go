package handlers

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/valyala/fasthttp"

	"github.com/Semerokozlyat/autoreview/internal/database"
)

var (
	//go:embed all:templates/*
	templateFS embed.FS

	//go:embed css/output.css
	cssFS embed.FS

	htmlTpl *template.Template
)

func init() {
	var err error
	htmlTpl, err = parseTemplates(templateFS, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func NewCSSFSHandler() fasthttp.RequestHandler {
	fsForCSS := &fasthttp.FS{FS: cssFS}
	return fsForCSS.NewRequestHandler()
}

func IndexHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html")
	htmlTpl.ExecuteTemplate(ctx, "index.html", database.Data.Companies)
}

func CompanyAddHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html")
	htmlTpl.ExecuteTemplate(ctx, "company-add.html", database.Data.Companies)
}

func CompanyCreateHandler(ctx *fasthttp.RequestCtx) {
	name := ctx.FormValue("company")
	contact := ctx.FormValue("contact")
	country := ctx.FormValue("country")

	database.Data.Add(database.Company{
		Company: string(name),
		Contact: string(contact),
		Country: string(country),
	})
	ctx.SetContentType("text/html")
	htmlTpl.ExecuteTemplate(ctx, "companies.html", database.Data.Companies)
}

func CompanyDeleteHandler(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id")
	database.Data.Delete(id.(string))

	ctx.SetContentType("text/html")
	htmlTpl.ExecuteTemplate(ctx, "companies.html", database.Data.Companies)
}

type CustomHandler struct {
	msg string
}

func NewCustomHandler(msg string) *CustomHandler {
	return &CustomHandler{msg: msg}
}

func (h *CustomHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("X-Custom-Response", "True")
	fmt.Fprintf(ctx, "Request path is %s, message is %s", ctx.Path(), h.msg)
}

func parseTemplates(templates fs.FS, funcMap template.FuncMap) (*template.Template, error) {
	root := template.New("")
	err := fs.WalkDir(templates, "templates", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		b, err := fs.ReadFile(templates, path)
		if err != nil {
			return fmt.Errorf("read file %s: %w", path, err)
		}
		tplName := strings.Split(path, string(os.PathSeparator))[1]
		t := root.New(tplName).Funcs(funcMap)
		_, err = t.Parse(string(b))
		if err != nil {
			return fmt.Errorf("parse template %s: %w", path, err)
		}
		return nil
	})
	return root, err
}

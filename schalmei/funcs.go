package schalmei

import (
	"html/template"
)

func routerUrl(name string, pairs ...string) string {
	u, _ := router.Get(name).URL(pairs...)
	return u.String()
}

var funcs = template.FuncMap{
	"url": routerUrl,
}

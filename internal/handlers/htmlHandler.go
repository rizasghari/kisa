package handlers

import (
	"github.com/a-h/templ"
	web "kisa-url-shortner/web/templ"
)

type HtmlHandler struct {
}

func NewHtmlHandler() *HtmlHandler {
	return &HtmlHandler{}
}

func (hh *HtmlHandler) GetIndexPage() templ.Component {
	return web.Index()
}

package handler

import (
	"net/http"
	"text/template"
)

const tplPath = "views/index.tpl"

func Index(w http.ResponseWriter, req *http.Request) {
	// テンプレートをパース
	tpl := template.Must(template.ParseFiles(tplPath))
	m := map[string]string{}

	// テンプレートを描画
	tpl.Execute(w, m)
}

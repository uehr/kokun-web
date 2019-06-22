package handler

import (
	"net/http"
	"text/template"
	"time"
)

const tplPath = "views/index.tpl"

func Index(w http.ResponseWriter, req *http.Request) {
	// テンプレートをパース
	tpl := template.Must(template.ParseFiles(tplPath))

	m := map[string]string{
		"Date": time.Now().Format("2006-01-02"),
		"Time": time.Now().Format("15:04:05"),
	}

	// テンプレートを描画
	tpl.Execute(w, m)
}

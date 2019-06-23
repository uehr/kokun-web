package handler

import (
	"net/http"
	"net/url"
	"text/template"

	"github.com/uehr/kokun/imageProcess"
	"github.com/uehr/kokun/senryu"
)

const tplPath = "views/index.tpl"

const FirstSentenceQueryKey = "first"
const SecondSentenceQueryKey = "second"
const ThirdSentenceQueryKey = "third"
const AuthorNameQueryKey = "author"

func hasRequiredQuery(query url.Values) bool {
	_, isExists := query[FirstSentenceQueryKey]
	if !isExists {
		return false
	}

	_, isExists = query[SecondSentenceQueryKey]
	if !isExists {
		return false
	}

	_, isExists = query[ThirdSentenceQueryKey]
	if !isExists {
		return false
	}

	_, isExists = query[AuthorNameQueryKey]
	if !isExists {
		return false
	}

	return true
}

func Index(w http.ResponseWriter, req *http.Request) {
	// テンプレートをパース
	tpl := template.Must(template.ParseFiles(tplPath))

	// クエリを読み込み川柳を作成
	query := req.URL.Query()
	m := map[string]string{}

	if hasRequiredQuery(query) {
		sen := senryu.Senryu{
			FirstSentence:  query.Get(FirstSentenceQueryKey),
			SecondSentence: query.Get(SecondSentenceQueryKey),
			ThirdSentence:  query.Get(ThirdSentenceQueryKey),
			AuthorName:     query.Get(AuthorNameQueryKey),
		}

		option := senryu.SenryuImageOption{}
		senryuImage, err := senryu.CreateImage(&sen, &option)
		if err != nil {
			return
		}

		base64Image := imageProcess.ImageBase64Encode(senryuImage)
		senryuImageHTML := "<img src='data:image/png;base64," + base64Image + "' />"

		m = map[string]string{"senryuImage": senryuImageHTML}
	}

	// テンプレートを描画
	tpl.Execute(w, m)
}

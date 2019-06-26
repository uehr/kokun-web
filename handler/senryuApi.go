package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"unicode/utf8"

	"github.com/uehr/kokun/imageProcess"
	"github.com/uehr/kokun/senryu"
)

type SenryuRequest struct {
	first  string
	second string
	third  string
	author string
}

type SenryuResponse struct {
	base64Img string
}

const FirstSentenceMaxLength = 10
const SecondSentenceMaxLength = 10
const ThirdSentenceMaxLength = 10
const AuthorNameMaxLength = 15

func returnRequestEntityTooLarge(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusRequestEntityTooLarge)
	w.Write([]byte("☄ HTTP status code returned!"))
}

func isValidLengthSenryu(sen *senryu.Senryu) bool {
	firstSentenceLength := utf8.RuneCountInString(sen.FirstSentence)
	secondSentenceLength := utf8.RuneCountInString(sen.SecondSentence)
	thirdSentenceLength := utf8.RuneCountInString(sen.ThirdSentence)
	authorNameLength := utf8.RuneCountInString(sen.AuthorName)

	return firstSentenceLength <= FirstSentenceMaxLength && secondSentenceLength <= SecondSentenceMaxLength && thirdSentenceLength <= ThirdSentenceMaxLength && authorNameLength <= AuthorNameMaxLength
}

func SenryuApi(w http.ResponseWriter, req *http.Request) {
	method := req.Method

	if method == "POST" {
		defer req.Body.Close()
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var sen senryu.Senryu
		json.Unmarshal([]byte(string(body)), &sen)

		if !isValidLengthSenryu(&sen) {
			returnRequestEntityTooLarge(w, req)
			return
		}

		option := senryu.SenryuImageOption{}
		senryuImage, err := senryu.CreateImage(&sen, &option)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// base64形式へ変換
		base64Img := imageProcess.ImageBase64Encode(senryuImage)

		// jsonRes, err := json.Marshal(SenryuResponse{base64Img: base64Img})
		mapRes := map[string]interface{}{"base64Img": base64Img}

		jsonRes, err := json.Marshal(mapRes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// json形式で川柳画像を送信
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonRes)
	}
}

package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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

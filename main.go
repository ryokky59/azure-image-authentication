package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	FaceAPIKey string = os.Getenv("FACE_API_KEY")
	EndPoint   string = os.Getenv("FACE_END_POINT")
)

func main() {
	filename := "bassy.png"
	img, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("画像の読み込みに失敗しました: ", err)
	}

	param := "/detect?returnFaceAttributes=age,gender"
	url := EndPoint + param

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(img))
	if err != nil {
		log.Println("リクエストの作成に失敗しました: ", err)
		return
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Ocp-Apim-Subscription-Key", FaceAPIKey)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println("リクエストに失敗しました: ", err)
		return
	}
	defer resp.Body.Close()

	var result []struct {
		FaceAttributes struct {
			Age    float64 `json:"age"`
			Gender string  `json:"gender"`
		} `json:"FaceAttributes"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("フォーマットに失敗しました: ", err)
		return
	}
	log.Println(result)
}

package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type resValue struct {
	ImgPath string `json:"imgPath"`
	Message string `json:"message"`
	Err     bool   `json:"error"`
}

func UploadImg(res http.ResponseWriter, req *http.Request) {
	//폼 파일 포맷의 image를 가져옴
	//header에는 img의 정보 img는 그냥 img만
	img, imgHeader, err := req.FormFile("image")

	if err != nil {
		fmt.Println("err in uploadImg.go 24")
		fmt.Println(err)
		res.WriteHeader(400)
		fmt.Fprint(res, "error during get image data")
		return
	}

	//이미지 인코딩
	encodedImg, _ := ioutil.ReadAll(img)

	//이미지가 들어갈 위치와 이름 설정
	imgName := time.Now().String() + imgHeader.Filename
	imgName = strings.ReplaceAll(imgName, " ", "") //img이름 설정시 생기는 공백 제거
	imgName = strings.ReplaceAll(imgName, ":", "") //img이름 설정시 생기는 공백 제거

	//파일 업로드
	err = ioutil.WriteFile("public/"+imgName, encodedImg, 0644)

	if err != nil {
		fmt.Println(err)
		res.WriteHeader(400)
		fmt.Fprint(res, "cannot uplaod the img")
		return
	}

	resValue := resValue{
		ImgPath: imgName,
		Message: "success",
		Err:     false,
	}

	resJson, _ := json.Marshal(resValue)

	fmt.Fprint(res, string(resJson))
}

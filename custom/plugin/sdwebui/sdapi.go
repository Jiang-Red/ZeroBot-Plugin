package sdwebui

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/FloatTech/floatbox/web"
)

var (
	baseurl = "http://127.0.0.1:7860"
)

func txt2imgapi() string {
	return baseurl + "/sdapi/v1/txt2img"
}

func sdmodelsapi() string {
	return baseurl + "/sdapi/v1/sd-models"
}

func sdoptions() string {
	return baseurl + "/sdapi/v1/options"
}

func posttxt2img(c txt2img) (imginfo imginfo, err error) {
	buffer := bytes.Buffer{}
	err = json.NewEncoder(&buffer).Encode(&c)
	if err != nil {
		return
	}
	data, err := web.PostData(txt2imgapi(), "application/json", &buffer)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &imginfo)
	if err != nil {
		return
	}
	return
}

func getmodels() (models models, err error) {
	data, err := web.GetData(sdmodelsapi())
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &models)
	if err != nil {
		return
	}
	return
}

func changemodel(modelname string) (err error) {
	models, err := getmodels()
	if err != nil {
		return
	}
	ok := false
	for _, v := range models {
		if v.ModelName == modelname {
			ok = true
			break
		}
	}
	if !ok {
		err = errors.New("unkown model")
		return
	}
	buffer := bytes.Buffer{}
	options := options{SDModelCheckpoint: modelname,
		CLIPStopAtLastLayers: 2}
	err = json.NewEncoder(&buffer).Encode(&options)
	if err != nil {
		return
	}
	_, err = web.PostData(sdoptions(), "application/json", &buffer)
	return
}

//@author D-S
//Created on 2019/9/18 5:52 下午
package net

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Post(url string, body interface{}) (error, map[string]interface{}) {
	bodyByte, _ := json.Marshal(body)
	requestBody := bytes.NewBuffer(bodyByte)
	rsp, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		return err, nil
	}
	rb, _ := ioutil.ReadAll(rsp.Body)
	var resBody map[string]interface{}
	if err := json.Unmarshal(rb, &resBody); err != nil {
		return err, nil
	}
	return nil, resBody
}

func PostForm(u string, v url.Values) (error, map[string]interface{}) {
	resp, err := http.PostForm(u, v)

	if err != nil {
		// handle error
		return err, nil
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var resBody map[string]interface{}
	if err := json.Unmarshal(body, &resBody); err != nil {
		return err, nil
	}
	return nil, resBody
}

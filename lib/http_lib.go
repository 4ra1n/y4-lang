package lib

import (
	"bytes"
	"io"
	"net/http"

	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/log"
	"github.com/4ra1n/y4-lang/native"
)

const (
	HttpLibName = "http"
)

var (
	HttpLib = base.NewList[*native.NativeFunction]()
)

func init() {
	doGetFun, err := native.NewNativeFunction(HttpLibName+SEP+"doGet", doGet)
	if err != nil {
		return
	}
	doPostFun, err := native.NewNativeFunction(HttpLibName+SEP+"doPost", doPost)
	if err != nil {
		return
	}

	HttpLib.Add(doGetFun)
	HttpLib.Add(doPostFun)
}

func doGet(url string) *base.Map[string, interface{}] {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("http lib new request error: %s", err.Error())
		return nil
	}
	req.Header.Set("User-Agent", "Y4-Lang")
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("http lib do request error: %s", err.Error())
		return nil
	}

	result := base.NewMap[string, interface{}]()
	result.Set("code", resp.StatusCode)
	headers := base.NewMap[string, interface{}]()
	for k, v := range resp.Header {
		var val string
		for i, vi := range v {
			if i != len(v)-1 {
				val += vi
				val += ", "
			}
		}
		headers.Set(k, val)
	}
	result.Set("headers", headers)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("http lib read body error: %s", err.Error())
		return nil
	}
	result.Set("body", string(body))
	return result
}

func doPost(url string, body string, ct string) *base.Map[string, interface{}] {
	client := http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(body))
	if err != nil {
		log.Errorf("http lib new request error: %s", err.Error())
		return nil
	}
	req.Header.Set("User-Agent", "Y4-Lang")
	if ct == "form" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else if ct == "json" {
		req.Header.Set("Content-Type", "application/json")
	} else {
		log.Errorf("only support json and form content-type")
		return nil
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("http lib do request error: %s", err.Error())
		return nil
	}

	result := base.NewMap[string, interface{}]()
	result.Set("code", resp.StatusCode)
	headers := base.NewMap[string, interface{}]()
	for k, v := range resp.Header {
		var val string
		for i, vi := range v {
			if i != len(v)-1 {
				val += vi
				val += ", "
			}
		}
		headers.Set(k, val)
	}
	result.Set("headers", headers)
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("http lib read body error: %s", err.Error())
		return nil
	}
	result.Set("body", string(respBody))
	return result
}

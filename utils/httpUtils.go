package utils

import (
	"bytes"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"time"
)

// SendRequest ...
func SendRequest(method string, uri string, body *bytes.Buffer) (int, string, error) {
	var (
		responseBody string
	)

	client := &http.Client{}
	client.Timeout = time.Minute * 3

	glog.V(6).Info(method)
	glog.V(6).Info(uri)
	glog.V(6).Info(body.String())

	if body == nil {
		body = bytes.NewBuffer([]byte("{}"))
	}
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		glog.Fatalf("http.NewRequest() failed with '%s'\n", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		glog.Warningf("client.Do() failed with '%s'\n", err)
		return http.StatusBadRequest, "", err
	}
	glog.V(6).Info(resp.StatusCode)
	glog.V(6).Infof("RESP: %+v", resp)

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	responseBody = string(bodyBytes)

	return resp.StatusCode, responseBody, nil
}

// SendRequestCustom ...
func SendRequestCustom(method string, uri string, body *bytes.Buffer, customRequestFunc func(req *http.Request)) (int, string, error) {
	var (
		responseBody string
	)

	client := &http.Client{}
	client.Timeout = time.Minute * 3

	glog.V(6).Info(method)
	glog.V(6).Info(uri)
	glog.V(6).Info(body.String())

	if body == nil {
		body = bytes.NewBuffer([]byte("{}"))
	}
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		glog.Fatalf("http.NewRequest() failed with '%s'\n", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	customRequestFunc(req)
	resp, err := client.Do(req)
	if err != nil {
		glog.Warningf("client.Do() failed with '%s'\n", err)
		return http.StatusBadRequest, "", err
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	responseBody = string(bodyBytes)

	glog.V(6).Info(resp.StatusCode)
	glog.V(6).Info(responseBody)

	return resp.StatusCode, responseBody, nil
}

package slack

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (e *Engine) do(method string, path string, param []byte) ([]byte, error) {
	var (
		req     *http.Request
		res     *http.Response
		reqURL  *url.URL
		reqBody io.Reader
		resBody []byte
		err     error
	)

	reqURL, err = url.Parse(path)
	if err != nil {
		return resBody, err
	}

	reqBody = bytes.NewBuffer(param)

	req, err = http.NewRequest(method, reqURL.String(), reqBody)
	if err != nil {
		return resBody, err
	}

	res, err = e.client.Do(req)
	if err != nil {
		return resBody, err
	}
	defer res.Body.Close()

	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return resBody, err
	}

	return resBody, nil
}

func (e *Engine) doJSON(method string, path string, param, response interface{}) error {

	var (
		jsonByte []byte
		resBody  []byte
		err      error
	)

	jsonByte, err = json.Marshal(param)
	if err != nil {
		return err
	}

	resBody, err = e.do(method, path, jsonByte)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resBody, response)
	return err
}

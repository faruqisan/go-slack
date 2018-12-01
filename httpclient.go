package slack

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (e *Engine) doJSON(method string, path string, param, response interface{}) error {

	var (
		req      *http.Request
		res      *http.Response
		reqURL   *url.URL
		reqBody  io.Reader
		resBody  []byte
		jsonByte []byte
		err      error
	)

	reqURL, err = url.Parse(path)
	if err != nil {
		return err
	}

	jsonByte, err = json.Marshal(param)
	if err != nil {
		return err
	}

	reqBody = bytes.NewBuffer(jsonByte)

	req, err = http.NewRequest(method, reqURL.String(), reqBody)
	if err != nil {
		return err
	}

	res, err = e.client.Do(req)
	if err != nil {
		return err
	}

	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = json.Unmarshal(resBody, response)

	return err
}

package dlsite_api

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"

func getJsonDataFromURL(method, url, requestBody string, client *http.Client) ([]byte, error) {
	var req *http.Request
	var err error
	if len(requestBody) == 0 {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bytes.NewBuffer([]byte(requestBody)))
	}
	if err != nil {
		errMsg := errors.New("getJsonDataFromURL：不正确的请求")
		return nil, errMsg
	}

	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Authorization", "Bearer "+Token)
	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		errMsg := errors.New("getJsonDataFromURL：连接失败或超时: " + err.Error())
		return nil, errMsg
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg := errors.New("getJsonDataFromURL：错误的返回码:" + strconv.Itoa(resp.StatusCode))
		return nil, errMsg
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errMsg := errors.New("getJsonDataFromURL：获取信息失败")
		return nil, errMsg
	}

	return body, nil
}

func doKeyword(keyword string) string {
	return strings.Replace(keyword, " ", "+", -1)
}

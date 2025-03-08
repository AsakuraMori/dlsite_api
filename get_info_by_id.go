package dlsite_api

import (
	"errors"
	"fmt"
	"net/http"
)

func GetInfoById(Id, typeName string, client *http.Client) ([]byte, error) {
	sType := ""
	switch typeName {
	case "游戏":
		sType = "pro"
	case "同人":
		sType = "maniax"
	case "漫画":
		sType = "book"
	case "手机游戏":
		sType = "appx"
	default:
		errMsg := errors.New("未匹配到类型")
		return nil, errMsg
	}

	apiURL := fmt.Sprintf("https://www.dlsite.com/%s/api/=/product.json?workno=%s",
		sType,
		Id)
	return getJsonDataFromURL("GET", apiURL, "", client)
}

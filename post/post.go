package post

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

type Item struct {
	title string `json:"title"`
	body  string `json:"body"`
}

func getQiita(c echo.Context) error {
	resp, err := http.Get("http://qiita.com/api/v2/items?page=1")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var data []Item
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)

}

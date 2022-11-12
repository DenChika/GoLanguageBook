package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func getComicsInfo(URL string) (*ComicsInfo, error) {
	jsonUrl := URL + "info.0.json"
	resp, err := http.Get(jsonUrl)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Сбой запроса: %s", resp.Status)
	}
	var result ComicsInfo
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	result.URL = URL
	resp.Body.Close()
	return &result, nil
}

func GetComics(args []string) ([]*ComicsInfo, error) {
	var comics []*ComicsInfo
	maxNum := getMaxPictureNum()
	for i := 1; i <= maxNum; i++ {
		if i == 404 {
			continue
		}
		url := MainURL + "/" + strconv.Itoa(i) + "/"
		result, err := getComicsInfo(url)
		if err != nil {
			return nil, err
		}
		alt := strings.ToLower(result.Alt)
		flag := true
		for _, arg := range args {
			normalArg := strings.ToLower(arg)
			if !strings.Contains(alt, normalArg) {
				flag = false
				break
			}
		}
		if flag {
			comics = append(comics, result)
		}
	}
	return comics, nil
}

func getMaxPictureNum() int {
	resp, err := http.Get(MainURL + "/info.0.json")
	if err != nil {
		return 0
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0
	}
	var nums struct{ Num int }
	if err := json.Unmarshal(body, &nums); err != nil {
		log.Fatalf("Сбой демаршалинга JSON: %s", err)
		return 0
	}
	return nums.Num
}

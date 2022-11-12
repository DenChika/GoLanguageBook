package poster

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetMovie(title string) (*Movie, error) {
	resp, err := http.Get(URL + "?t=" + title + "&apikey=" + APIKey)
	if err != nil {
		log.Fatalf("Сбой запроса: %v\n", err)
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("%v\n", err)
		return nil, err
	}
	resp.Body.Close()
	var movie Movie
	if err := json.Unmarshal(data, &movie); err != nil {
		log.Fatalf("Сбой демаршалинга: %v\n", err)
		return nil, err
	}
	return &movie, nil
}

func UploadPoster(movie *Movie) {
	posterUrl := movie.Poster
	resp, err := http.Get(posterUrl)
	if err != nil {
		log.Fatalf("Сбой запроса: %v\n", err)
		return
	}
	defer resp.Body.Close()
	file, err := os.Create(fmt.Sprintf("poster\\%s.jpg", movie.Title))
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %v\n", err)
		return
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatalf("Ошибка при копировании: %v\n", err)
		return
	}
}

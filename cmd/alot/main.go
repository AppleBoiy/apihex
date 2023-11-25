package main

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"sync"
)

const MAX = 20

var wg sync.WaitGroup

func main() {
	wg.Add(MAX)
	for id := 0; id < MAX; id++ {
		go request(id)
	}
	wg.Wait()
}

func request(id int) {
	defer wg.Done()
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		os.Getenv("TRANSFER_URL")+"/"+strconv.Itoa(id),
		nil,
	)
	if err != nil {
		slog.Error(err.Error(), slog.Int("id", id))
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error(err.Error(), slog.Int("id", id))
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error(err.Error(), slog.Int("id", id))
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			slog.Error(err.Error(), slog.Int("id", id))
			return
		}
		slog.Error(string(b), slog.Int("id", id))
		return
	}

	type responseMessage struct {
		Message string `json:"message"`
	}

	var response responseMessage
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		slog.Error(err.Error(), slog.Int("id", id))
		return
	}

	slog.Info(response.Message, slog.Int("id", id))
}

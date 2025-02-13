package main



import (
	"context"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	RANDOM_DURATION := time.Duration(rand.Intn(10)+1) * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), RANDOM_DURATION)

	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)

	if err != nil {
		panic(err)
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	io.Copy(os.Stdout, response.Body)
}

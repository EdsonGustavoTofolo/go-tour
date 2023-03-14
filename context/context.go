package context

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func RunContextWithTimeout() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)

	defer cancel()

	bookHotel(ctx)
}

func RunContextWithValue() {
	ctx := context.WithValue(context.Background(), "token", "senha")

	value := ctx.Value("token")

	fmt.Println(value)
}

func RunContextForWebServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		log.Println("Request iniciado")
		defer log.Println("Request finalizada")

		select {
		case <-time.After(time.Second * 5):
			log.Println("Request processada com sucesso")
			w.Write([]byte("Request processada com sucesso"))
		case <-ctx.Done():
			log.Fatal("Request cancelada pelo cliente")
		}
	})
	http.ListenAndServe(":8585", nil)
}

func RunContextClientServer() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8585", nil)

	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking canceled. Timeout reached.")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Hotel booked.")
	}
}

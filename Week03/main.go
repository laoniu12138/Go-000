package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

func main() {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan)
	stop := make(chan string)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		defer fmt.Println("httpserver is stoping")
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello"))
		})

		server := http.Server{
			Addr:    ":8080",
			Handler: mux,
		}

		errChan := make(chan string)
		go func() {
			defer fmt.Println("server is stop....")

			err := server.ListenAndServe()
			if err != nil {
				errChan <- err.Error()
			}
		}()

		errInfo := ""
		select {
		case <-stop:
			server.Close()
		case errInfo = <-errChan:
			close(stop)
		}

		return errors.New(errInfo)
	})

	g.Go(func() error {
		select {
		case s := <-signalChan:
			close(stop)
			return errors.New(s.String())
		case <-stop:
			return nil
		}

	})

	fmt.Println("httpserver waiting...")
	err := g.Wait()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ctx.Err())
}

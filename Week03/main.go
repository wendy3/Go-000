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
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "geekbang go go go!!!")
		})
		go func() {
			<-ctx.Done()
			fmt.Println("server context done")
			s.Shutdown(ctx)
		}()
		return http.ListenAndServe(":8080", nil)
	})

	g.Go(func() error {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt)
		select {
		case <-sig:
			_, cancel := context.WithCancel(ctx)
			cancel()
			return errors.New("interrupt signal")
		case <-ctx.Done():
			fmt.Println("signal context done")
			return ctx.Err()
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("error is %v\n", err)
	}
}

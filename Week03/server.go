package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

//https://pkg.go.dev/context
//https://pkg.go.dev/golang.org/x/sync/errgroup

//视频1时:45分左右
func main() {

	ctx, cancel := context.WithCancel(context.Background())

	g, ctx0 := errgroup.WithContext(ctx)

	// 监听系统信号
	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-sigs:
			log.Println("catch quit signal.")
			cancel()
			return nil
		case <-ctx0.Done():
			return nil
		}
	})

	g.Go(func() error {
		return newServer(ctx0, ":1234")
	})

	g.Go(func() error {
		return newServer(ctx0, ":1235")
	})

	g.Go(func() error {
		return newServer(ctx0, ":1236")
	})

	if err := g.Wait(); errors.Is(err, context.Canceled) {
		log.Printf("%v", err)
	} else if err != nil {
		log.Printf("errorgroup error: %s\n", err)
	}

	log.Println("ctx0 error: ", ctx0.Err())
}

func newServer(ctx context.Context, addr string) error {
	s := &http.Server{Addr: addr}

	go func() {
		select {
		case <-ctx.Done():
			ctx1, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			log.Println("shutdown server with Ader: ", addr)
			s.Shutdown(ctx1)
		}
	}()

	log.Println(addr + " server is starting")
	err := s.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			err = nil
		} else {
			log.Println(addr+" server started failed", err)
		}
	}
	return err
}

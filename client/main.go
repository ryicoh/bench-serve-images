package main

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	count := 1000

	dir := "images"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	targets := make(map[string][]byte, len(files))
	for _, f := range files {
		if f.Name() == ".keep" {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, f.Name()))
		if err != nil {
			panic(err)
		}
		targets["http://localhost:8080/"+f.Name()] = data
	}

	cli := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			MaxConnsPerHost:       100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	ctx := context.TODO()
	s := semaphore.NewWeighted(1)
	var wg sync.WaitGroup


	for i := 0; i < count; i++ {
		for t := range targets {
			wg.Add(1)
			t := t
			go func() {
				defer wg.Done()
				if err := s.Acquire(ctx, 1); err != nil {
					os.Stderr.WriteString(err.Error())
					return
				}
				defer s.Release(1)

				req, err := http.NewRequest(http.MethodGet, t, bytes.NewBufferString(""))
				if err != nil {
					os.Stderr.WriteString(err.Error())
					return
				}
				defer req.Body.Close()

				resp, err := cli.Do(req)
				if err != nil {
					os.Stderr.WriteString(err.Error())
					return
				}
				defer resp.Body.Close()

				io.Copy(io.Discard, resp.Body)
			}()
		}
	}

	wg.Wait()
}

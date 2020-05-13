package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/acim/go-csp/pkg/seen"
)

type message struct {
	id        int
	timestamp int64
}

func sender(id int, wg *sync.WaitGroup, ch chan<- message, quit <-chan struct{}) {
	defer wg.Done()
	fmt.Printf("Sender %d started\n", id)
	for {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		select {
		case ch <- message{id, time.Now().Unix()}:
		case <-quit:
			fmt.Printf("Sender %d done\n", id)
			return
		}
	}
}

func receiver(ns int, ch <-chan message, quit chan<- struct{}) {
	fmt.Printf("Receiver started\n")
	s := seen.CreateMap(ns)
	for {
		m := <-ch
		fmt.Printf("Received message from %d with timestamp %d\n", m.id, m.timestamp)
		s.SetSeen(m.id)
		if s.AllSeen() {
			close(quit)
			fmt.Printf("Receiver done\n")
			return
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: go-csp number-of-senders\nnumber-of-senders is required.\n")
		os.Exit(1)
	}

	ns, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Usage: go-csp number-of-senders\nnumber-of-senders should be an integer.\n")
		os.Exit(1)
	}

	ch := make(chan message)
	quit := make(chan struct{})
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	for i := 1; i <= ns; i++ {
		wg.Add(1)
		go sender(i, &wg, ch, quit)
	}

	go receiver(ns, ch, quit)

	wg.Wait()
	close(ch)
}

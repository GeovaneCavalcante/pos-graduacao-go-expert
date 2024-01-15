package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var number uint64 = 0

func main() {
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		atomic.AddUint64(&number, 1)
		number++
		// m.Unlock()
		w.Write([]byte(fmt.Sprintf("Você é o visitante numero %d", number)))
	})
	http.ListenAndServe(":8080", nil)
}

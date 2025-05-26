package main

import (
	"ClientServerApplication/lib"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	maxReportsProcessingAmount = 5
	minProcessingTime          = 1
	maxProcessingTime          = 180
)

func main() {
	reportQueue, reports := lib.NewReportQueue(), make(chan lib.Report)
	for i := 0; i < maxReportsProcessingAmount; i++ {
		go process(i, reports)
	}

	go func() {
		for {
			for reportQueue.Size() > 0 {
				reports <- *reportQueue.Remove()
			}
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var report lib.Report
		err := json.NewDecoder(r.Body).Decode(&report)
		if err != nil {
			http.Error(w, fmt.Sprintf("decoding error: `%v`", err), http.StatusBadRequest)
			return
		}

		if reportQueue.Exists(report) {
			http.Error(w, fmt.Sprintf("report already exists: `%v`", report), http.StatusConflict)
			return
		}
		reportQueue.Add(&report)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func process(id int, reports chan lib.Report) {
	for r := range reports {
		log.Printf("process %d started report `%v`\n", id, r)
		time.Sleep(10 * time.Second)
		log.Printf("process %d finished report `%v`\n", id, r)
	}
}

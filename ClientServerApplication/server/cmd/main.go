package main

import (
	"ClientServerApplication/lib/doubleset"
	"ClientServerApplication/lib/report"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// consts
const (
	maxReportsProcessingAmount = 5
	minProcessingTime          = 1
	maxProcessingTime          = 180
)

var (
	doubleSet   = doubleset.NewDoubleSet[int, report.ReportType]()
	reportQueue = report.NewReportQueue()
	reports     = make(chan *report.Report)
)

func main() {
	// amount of processes handling requests
	for id := 0; id < maxReportsProcessingAmount; id++ {
		go runner(id, reports)
	}

	// reports collects data from the reportQueue and sends it to processes
	go func() {
		for {
			for reportQueue.Size() > 0 {
				reports <- reportQueue.Remove()
			}
		}
	}()

	// handler
	http.HandleFunc("/", handler)

	// run server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// handler processes requests on the path `/`.
func handler(w http.ResponseWriter, r *http.Request) {
	// json decoding
	var report report.Report
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, fmt.Sprintf("decoding error: `%v`", err), http.StatusBadRequest)
		return
	}

	// check, if the report type already exists in the doubleSet
	if doubleSet.Exists(report.Id, report.Type) {
		log.Printf("report already exists: `{%v, %v}`\n", report.Id, report.Type)
		http.Error(w, fmt.Sprintf("report already exists: `{%v, %v}`", report.Id, report.Type), http.StatusConflict)
		return
	}

	// adding report info into doubleSet and reportQueue
	doubleSet.Add(report.Id, report.Type)
	reportQueue.Add(&report)
	log.Printf("amount of requests in queue: %d\n", reportQueue.Size())
	log.Printf("report added: `{%v, %v}`\n", report.Id, report.Type)
}

// process is a function for the ability to process data from different goroutines.
func runner(id int, reports chan *report.Report) {
	for r := range reports {
		delay := time.Duration(minProcessingTime+rand.Intn(maxProcessingTime-minProcessingTime+1)) * time.Second
		log.Printf("runner %d started report `{%v, %v}`. Processing time: %s\n", id, r.Id, r.Type, delay.String())
		time.Sleep(delay)

		doubleSet.Remove(r.Id, r.Type)

		log.Printf("runner %d finished report `{%v, %v}`\n", id, r.Id, r.Type)
		log.Printf("amount of requests in queue: %d\n", reportQueue.Size())
	}
}

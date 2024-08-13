package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Println("1st log")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("2nd log")
	log.Println("3rd log")

	prefixLog := log.New(os.Stdout, "ERROR: ", log.LstdFlags)
	prefixLog.Println("4th log")
	prefixLog.SetPrefix("WARNING: ")
	prefixLog.Println("5th log")

	var buf bytes.Buffer
	bufLog := log.New(&buf, "WARNING 1: ", log.LstdFlags)
	bufLog.Print("6th log", buf.String())
	fmt.Println(buf.String())

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")
	myslog.Info("hello again", "key1", "val1", "key2", "val2", "age", 25)
}

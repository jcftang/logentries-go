package main

import (
	"encoding/csv"
	"fmt"
	"github.com/jcftang/logentries-go"
	flag "github.com/spf13/pflag"
	"log"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var config = flag.String("config", "logs.csv", "CSV file containing a list of 'tokens,/file/to/follow'")
	flag.Parse()

	if _, err := os.Stat(*config); os.IsNotExist(err) {
		fmt.Printf("no such file or directory: %s\n", *config)
		return
	}

	f, err := os.Open(*config)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 2
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, each := range rawCSVdata {
		// first column is the token, the second is the path to the log to follow
		fmt.Println("Following : ", each[1], each[0])
		t, err := multi.NewTailer(each[1], each[0])
		if err != nil {
			log.Fatal(err)
		}
		wg.Add(1)
		go func(t *multi.Tailer) {
			t.Do()
			wg.Done()
		}(t)
	}
	wg.Wait()
}

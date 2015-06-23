package main

import (
	"fmt"
	"github.com/activestate/tail"
	"github.com/jcftang/le_go"
	flag "github.com/spf13/pflag"
	"os"
)

func main() {
	var token = flag.String("token", "nil", "log token")
	var logfile = flag.String("logfile", "/tmp/foo.txt", "log file to follow")
	var seekInfoOnStart = &tail.SeekInfo{Offset: 0, Whence: os.SEEK_END}

	flag.Parse()

	fmt.Println("using token: ", *token)

	if _, err := os.Stat(*logfile); os.IsNotExist(err) {
		fmt.Printf("no such file or directory: %s\n", *logfile)
		return
	}

	le, err := le_go.Connect(*token) // replace with token
	if err != nil {
		panic(err)
	}

	defer le.Close()
	t, err := tail.TailFile(*logfile, tail.Config{Follow: true, ReOpen: true, Location: seekInfoOnStart, Logger: tail.DiscardingLogger})
	if err == nil {
		for line := range t.Lines {
			le.Println(line.Text)
		}
	}
}

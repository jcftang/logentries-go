package multi

import (
	"fmt"
	"os"

	"github.com/ActiveState/tail"
	"github.com/jcftang/le_go"
)

var (
	seekInfoOnStart = &tail.SeekInfo{Offset: 0, Whence: os.SEEK_END}
)

type Tailer struct {
	*le_go.Logger
	*tail.Tail
	path  string
	token string
}

func NewTailer(filename string, token string) (*Tailer, error) {
	t, err := tail.TailFile(filename, tail.Config{
		Follow:   true,
		Location: seekInfoOnStart,
		Logger:   tail.DiscardingLogger,
	})

	if err != nil {
		return nil, err
	}

	le, err := le_go.Connect(token)
		if err != nil {
			panic(err)
		}
	defer le.Close()

	return &Tailer{
		Logger: le,
		Tail: t,
		path: filename,
		token: token,
	}, nil
}

func (t Tailer) Do() {
	for line := range t.Lines {
		fmt.Println(line.Text, t.token)
		t.Println(line.Text)
	}
}

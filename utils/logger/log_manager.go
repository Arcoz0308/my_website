package logger

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/arcoz0308/arcoz0308.tech/utils"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	maxLines        = 10000
	requestLogLines int
	requestLogFile  *os.File
	awaitRequestLog []string
)

func Init() {
	if err := os.MkdirAll("request_logs", os.ModePerm); err != nil {
		panic(err)
	}
	if err := os.MkdirAll("rate_limit_logs", os.ModePerm); err != nil {
		panic(err)
	}
	loadRequestLogger()
	_, err := utils.Cron.AddFunc("0 * * * * *", func() {
		if len(awaitRequestLog) != 0 {
			requestLogFile.WriteString(strings.Join(awaitRequestLog, "\n") + "\n")
			requestLogLines += len(awaitRequestLog)
			awaitRequestLog = []string{}
			if requestLogLines >= maxLines {
				loadRequestLogger()
			}
		}
	})
	if err != nil {
		panic(err)
	}
	_, err = utils.Cron.AddFunc("@daily", func() {
		loadRequestLogger()
	})
	if err != nil {
		panic(err)
	}
}

func loadRequestLogger() {
	if requestLogFile != nil {
		requestLogFile.Close()
	}
	var f *os.File
	// format YYYY-MM-DD-n.log
	files, err := os.ReadDir("request_logs")
	if err != nil {
		panic(err)
	}
	if len(files) == 0 {
		t := time.Now()
		f, err = os.Create(fmt.Sprintf("request_logs/%s-0.log", t.Format("2006-01-02")))
		if err != nil {
			panic(err)
		}
	} else {
		i := 0
		found := false
		for _, file := range files {
			m, err := regexp.Match(time.Now().Format("2006-01-02")+"-[0-9].log", []byte(file.Name()))
			if err != nil {
				panic(err)
			}
			if m {
				f2, err := os.Open("request_logs/" + file.Name())
				if err != nil {
					panic(err)
				}
				l, err := LineCounter(f2)
				if err != nil {
					panic(err)
				}
				if l < maxLines {
					f = f2
					found = true
					break
				} else {
					i++
				}
			}
		}
		if !found {
			f, err = os.Create(fmt.Sprintf("request_logs/%s-%d.log", time.Now().Format("2006-01-02"), i))
			if err != nil {
				panic(err)
			}
		}
	}
	name := f.Name()
	f.Close()

	file, err := os.OpenFile(name, os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	l, err := LineCounter(file)
	if err != nil {
		panic(err)
	}
	requestLogLines = l
	requestLogFile = file
}
func RequestLog(msg string) {
	awaitRequestLog = append(awaitRequestLog, msg)
}

func LineCounter(r io.Reader) (int, error) {

	var count int
	const lineBreak = '\n'

	buf := make([]byte, bufio.MaxScanTokenSize)

	for {
		bufferSize, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], lineBreak)
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}

	return count, nil
}

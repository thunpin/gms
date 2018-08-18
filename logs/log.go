package logs

import "log"

type Header struct {
	Tag       string
	RequestID string
	UUID      string
	Action    string
}

type Info struct {
	Header
	Content interface{}
}

type Error struct {
	Header
	Error error
}

type Log interface {
	Info(Info)
	Warning(Info)
	Error(Error)
	Close()
}

type Logger struct{}

func (logger Logger) Info(info Info) {
	log.Printf("[INFO] %+v", info)
}

func (logger Logger) Warning(info Info) {
	log.Printf("[WARN] %+v", info)
}

func (logger Logger) Error(err Error) {
	log.Printf("[ERROR] %+v", err)
}

func (logger Logger) Close() {}

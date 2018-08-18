package logs

var instance Log

func DefineInstance(log Log) {
	instance = log
}

func Instance() Log {
	if instance == nil {
		instance = &Logger{}
	}

	return instance
}

func Close() {
	if instance != nil {
		instance.Close()
		instance = nil
	}
}

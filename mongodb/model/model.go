package model

type ExecTime struct {
	StartTime int64 `bson:"start_time"`
	EndTime   int64 `bson:"end_time"`
}

type LogRecord struct {
	JobName string `bson:"jobName"`
	Command string `bson:"command"`
	Err     string `bson:"err"`
	Content string `bson:"content"`
	Tp      ExecTime
}

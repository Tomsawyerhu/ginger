package model

import "time"

type Account struct {
	Uuid uint64
	Username string
	CreateTime time.Time
	Auth uint //权限(上传、下载、上传下载)
}

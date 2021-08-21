package model

import "time"

type File struct {
	Uuid uint64
	CreateTime time.Time
	ModifiedTime time.Time
	Size int64
	Type uint  //文件类型(图片、音频等)
	Decoding string //解码格式
	NeedVerified uint //是否需要校验
	VerifiedMethod string //校验方式(md5,sha256等)
	VerifiedCode string //校验码
}

type FileSegment struct {
	Fid uint64 //文件uuid
	Id uint64
	Size int64
	IsLast int
}

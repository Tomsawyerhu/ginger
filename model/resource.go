package model

type ResourceControl struct {
	Aid uint64 //账号
	Fid uint64 //文件
	Mode int   //权限(只读、只写、读写)
}
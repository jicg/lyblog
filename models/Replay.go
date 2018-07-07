package models

import "time"

type Replay struct {
	Id         int
	Uid        int
	NoteId     int
	Puid       int
	Content    string
	Status     int //帖子状态，0未解决，1已解决，2精华，3删除
	Top        int //置顶，0否，1是
	ThumbCount int
	CTime      time.Time
	UTime      time.Time
}

package models

import (
	"time"
)

type Note struct {
	Id         int
	Uid        int
	U          *User `orm:"-"`
	Title      string
	Content    string
	ReplyCount int
	VisitCount int
	Status     int //帖子状态，0未解决，1已解决，2精华，3删除
	Top        int //置顶，0否，1是
	ThumbCount int
	Experience int
	CTime      time.Time
}

func QueryTopNotes() ([]*Note, error) {
	var notes []*Note
	if _, err := o.QueryTable(&Note{Top: 1}).Limit(5).All(&notes); err != nil {
		return nil, err
	}
	for _, n := range notes {
		var user = &User{}
		o.QueryTable(&User{Id: n.Id}).One(user)
		n.U = user
	}
	return notes, nil
}

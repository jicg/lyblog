package models

import (
	"time"
)

type Replay struct {
	Id         int
	Uid        int
	U          *User     `orm:"-"`
	NoteId     int
	Puid       int
	Content    string
	Status     int //帖子状态，0未解决，1已解决，2精华，3删除
	Top        int //置顶，0否，1是
	ThumbCount int
	CTime      time.Time `orm:"null"`
	UTime      time.Time `orm:"null"`
}

func QueryRespsCountByNoteId(noteid, page, pagesize int) (int64, error) {
	//if page<=0{
	//	return nil,errors.New("")
	//}
	cnt := int64(0)
	if cnt, err := o.QueryTable(&Replay{NoteId: noteid}).Count(); err != nil {
		return cnt, err
	}
	return cnt, nil
}

func QueryRespsByNoteIdAndPage(noteid, page, pagesize int) ([]*Replay, error) {
	//if page <= 0 {
	//	return nil, errors.New("页数非法")
	//}
	var replays []*Replay
	if _, err := o.QueryTable(&Replay{NoteId: noteid}).Offset((page - 1) * pagesize).Limit(pagesize).All(&replays); err != nil {
		return nil, err
	}
	for _, rep := range replays {
		var user User
		if err := o.QueryTable(&User{Id: noteid}).Limit(1).One(&user); err != nil {
			return nil, err
		}
		rep.U = &user
	}

	//if len(replays) <= 0 {
	//	return nil, errors.New("没有数据")
	//}
	return replays, nil
}

package comet

import (
	"errors"
	"fmt"
)

// セッションのcookieのkey, UUID, データ送信用のチャンネル
type SessionList struct {
	Map map[string]chan interface{}
}

func NewSessionList() *SessionList {
	return &SessionList{map[string]chan interface{}{}}
}

// リストに新しいチャンネルを自動追加
func (sl *SessionList) Set(id string) error {
	_, ok := sl.Map[id]
	if ok { // あれば拒否
		return errors.New(id + ":idは既にセットされています。")
	}
	fmt.Println("set", id)
	sl.Map[id] = make(chan interface{}, 10)
	return nil
}

// リストを取得
func (sl *SessionList) GetList() map[string]chan interface{} {
	return sl.Map
}

package comet

import (
	"errors"
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
	sl.Map[id] = make(chan interface{}, 10)
	return nil
}

func (sl *SessionList) Delete(id string) {
	ch, ok := sl.Map[id]
	if !ok { // 無ければ拒否
		return
	}
	close(ch)
	delete(sl.Map, id)
}

// リストを取得
func (sl *SessionList) GetList() map[string]chan interface{} {
	return sl.Map
}

package types

// セッションのcookieのkey, UUID, データ送信用のチャンネル
type SessionList map[string]map[string]chan interface{}

func (sl *SessionList) Init() {
	*sl = map[string]map[string]chan interface{}{}
}
func (sl *SessionList) Set() {}
func (sl *SessionList) Get() {}

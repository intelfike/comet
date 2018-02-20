package comet

import (
	"errors"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type Comet struct {
	sessionList *SessionList
	key         string
}

// cookieに使うkeyを指定
func NewComet(key string) *Comet {
	cmt := &Comet{
		sessionList: NewSessionList(),
		key:         key,
	}
	return cmt
}

// 正常にcookieをセットした場合、既にcookieがセットされている場合にnilを返します。
func (c *Comet) Start(w http.ResponseWriter, r *http.Request) error {
	cdata, err := r.Cookie(c.key)
	if err == nil { // 既にセットされているならやめ
		c.sessionList.Set(cdata.Value)
		return nil
	}
	// セッションIDを生成してセットする
	u1 := uuid.Must(uuid.NewV4()).String()
	http.SetCookie(w, &http.Cookie{Name: c.key, Value: u1, Path: "/"})
	c.sessionList.Set(u1)
	return nil
}

// done me
func (c *Comet) Done(r *http.Request, i interface{}) error {
	cdata, err := r.Cookie(c.key)
	if err != nil {
		return errors.New("セッションがセットされていません Start()を実行済みか確認してください。")
	}
	slist := c.sessionList.GetList()
	ch, ok := slist[cdata.Value]
	if !ok {
		return errors.New("セッションのリストから見つかりませんでした")
	}
	ch <- i
	return nil
}

// done all
func (c *Comet) DoneAll(i interface{}) {
	slist := c.sessionList.GetList()
	for _, ch := range slist {
		ch <- i
	}
}

// done other
// 自分以外を終了
func (c *Comet) DoneOther(r *http.Request, i interface{}) error {
	cdata, err := r.Cookie(c.key)
	if err != nil {
		return errors.New("セッションがセットされていません Start()を実行済みか確認してください。")
	}
	slist := c.sessionList.GetList()
	for key, ch := range slist {
		if cdata.Value == key {
			continue
		}
		ch <- i
	}
	return nil
}

func (c *Comet) Wait(r *http.Request) (interface{}, error) {
	cdata, err := r.Cookie(c.key)
	if err != nil {
		return nil, errors.New("セッションがセットされていません Start()を実行済みか確認してください。")
	}

	slist := c.sessionList.GetList()
	ch, ok := slist[cdata.Value]
	if !ok {
		return nil, errors.New("セッションのリストから見つかりませんでした")
	}
	i := <-ch

	return i, nil
}

func (c *Comet) End(r *http.Request) error {
	cdata, err := r.Cookie(c.key)
	if err != nil {
		return errors.New("セッションがセットされていません Start()を実行済みか確認してください。")
	}

	c.sessionList.Delete(cdata.Value)
	return nil
}

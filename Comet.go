package comet

import (
	"fmt"
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
func (c *Comet) Done(i interface{}) {
	slist := c.sessionList.GetList()
	for _, ch := range slist {
		fmt.Println("send", ch, i)
		ch <- i
	}
}
func (c *Comet) Wait(r *http.Request) interface{} {
	cdata, err := r.Cookie(c.key)
	if err != nil {
		return nil
	}

	slist := c.sessionList.GetList()
	ch, ok := slist[cdata.Value]
	if !ok {
		return nil
	}
	i := <-ch
	fmt.Println("done -", cdata.Value, i)

	return i
}

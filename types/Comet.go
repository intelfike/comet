package types

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type Comet struct{}

func New() *Comet {
	cmt := &Comet{}
	return cmt
}

// 正常にcookieをセットした場合、既にcookieがセットされている場合にnilを返します。
func (c *Comet) Start(r *http.Request, w http.ResponseWriter, key string) error {
	cdata, err := r.Cookie(key)
	if cdata != nil || cdata.String() != "" || err == nil {
		return nil
	}
	u1 := uuid.Must(uuid.NewV4()).String()
	http.SetCookie(w, &http.Cookie{Name: key, Value: u1})

}
func (c *Comet) Send(i interface{}) {

}
func (c *Comet) Wait() interface{} {

}

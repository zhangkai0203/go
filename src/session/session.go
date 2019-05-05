package session

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
)


type session struct {
	Name string
	Value string
	Path string
	Max int
	Only bool
}

var sess *session

func SessionID() (uuid.UUID)  {
	u1 := uuid.Must(uuid.NewV4())
	return u1
}

func Set(w http.ResponseWriter)  {
	sess = &session{"sessionID",SessionID().String(),"/",3600,true}
	EditCookie(w)
}

func Get(w http.ResponseWriter,r *http.Request)  {
      r.ParseForm()
      cookie,_ := r.Cookie(sess.Name)
      fmt.Println(cookie.Value)
}

func Delete(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	cookie,_ := r.Cookie(sess.Name)
	sess = &session{"sessionID",cookie.Value,"/",-1,true}
	EditCookie(w)
}

func EditCookie(w http.ResponseWriter)  {
	cookie := http.Cookie{
		Name:sess.Name,
		Value:sess.Value,
		Path:sess.Path,
		MaxAge:sess.Max,
		HttpOnly:sess.Only,
	}
	http.SetCookie(w,&cookie)
}

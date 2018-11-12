package raven

import (
	"github.com/cocopc/gcommons/log"
	"github.com/getsentry/raven-go"
)



func init(){

	l.Define("[gfilter-api]",l.Blue,"info")

}

var (
	RavenClient *raven.Client
)

func InitRavenClient(dsn string,username string, email string) {

	l.Logf("init raven client dsn: %s username: %s  email: %s  ",dsn,username,email)
	client, err := raven.New(dsn)
	if err != nil {
		l.Error("init raven client fail", err)
	}
	user := &raven.User{Username: username, Email: email}
	client.SetUserContext(user)
	RavenClient = client
}


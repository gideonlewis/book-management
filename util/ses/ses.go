package ses

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"

	mySES "git.teqnological.asia/teq-go/teq-echo/client/ses"
)

type SES struct {
	session *session.Session
	svc     *ses.SES
}

func NewSES() ISES {
	return SES{
		session: mySES.GetSession(),
		svc:     mySES.GetService(),
	}
}

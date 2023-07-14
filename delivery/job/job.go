package job

import "git.teqnological.asia/teq-go/teq-echo/config"

type IJob interface {
	Run()
}

type Jobs []IJob

func (js Jobs) Run() {
	for _, j := range js {
		go j.Run()
	}
}

func New() Jobs {
	return Jobs{
		NewHealthChecks(config.GetConfig().HealthCheck.HealthCheckEndPoint),
	}
}

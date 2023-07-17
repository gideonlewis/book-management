package job

import (
	"fmt"

	"git.teqnological.asia/teq-go/teq-echo/config"
)

type IJob interface {
	Run()
}

type Jobs []IJob

func (js Jobs) Run() {
	fmt.Println("Running jobs...")
	for _, j := range js {
		go j.Run()
	}
}

func New() Jobs {
	return Jobs{
		NewBorrowingOverdueCheck(),
		NewHealthChecks(config.GetConfig().HealthCheck.HealthCheckEndPoint),
	}
}

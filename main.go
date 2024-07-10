package outagelab

import "github.com/outagelab/outagelab/internal/app"

type OutageLabApp app.OutageLabApp

func NewApp() *OutageLabApp {
	return (*OutageLabApp)(app.New())
}

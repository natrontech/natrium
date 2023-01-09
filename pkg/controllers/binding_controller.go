package controllers

import (
	"github.com/natrongmbh/natrium/pkg/bindings"
	"github.com/natrongmbh/natrium/pkg/util"
)

var (
	err error
)

type BindingController struct {
	BaseController
}

func (c *BindingController) Setup() error {
	err = bindings.BindDBHooks()
	if err != nil {
		util.Logger.Error(err)
	}

	err = bindings.BindAppHooks()
	if err != nil {
		util.Logger.Error(err)
	}
	return nil
}

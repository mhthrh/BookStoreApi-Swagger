package Controller

import (
	"fmt"
	"github.com/mhthrh/ApiStore/Utility/ConfigUtil"
	"github.com/mhthrh/ApiStore/Utility/ExceptionUtil"
	"github.com/mhthrh/ApiStore/Utility/ValidationUtil"
	"github.com/sirupsen/logrus"
)

type KeyBook struct{}

type Controller struct {
	l *logrus.Entry
	v *ValidationUtil.Validation
	e *ExceptionUtil.Exception
	c *ConfigUtil.Config
}

var InvalidPath = fmt.Errorf("invalid Path, path must be /Controller/[id]")

type GenericError struct {
	Message string `json:"message"`
}

type ValidationError struct {
	Messages []string `json:"messages"`
}

func NewBooks(l *logrus.Entry, v *ValidationUtil.Validation, e *ExceptionUtil.Exception, c *ConfigUtil.Config) *Controller {
	return &Controller{l, v, e, c}
}

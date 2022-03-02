package Controller

import (
	"fmt"
	"github.com/mhthrh/ApiStore/Utility/ConfigUtil"
	"github.com/mhthrh/ApiStore/Utility/DbUtil/DbPool"
	"github.com/mhthrh/ApiStore/Utility/ExceptionUtil"
	"github.com/mhthrh/ApiStore/Utility/ValidationUtil"
	"github.com/sirupsen/logrus"
)

type Key struct{}

type Controller struct {
	l  *logrus.Entry
	v  *ValidationUtil.Validation
	e  *ExceptionUtil.Exception
	c  *ConfigUtil.Config
	db *DbPool.DBs
}

var InvalidPath = fmt.Errorf("invalid Path, path must be /Controller/[id]")

type GenericError struct {
	Message string `json:"message"`
}

type ValidationError struct {
	Messages []string `json:"messages"`
}

func NewBook(l *logrus.Entry, v *ValidationUtil.Validation, e *ExceptionUtil.Exception, c *ConfigUtil.Config, db *DbPool.DBs) *Controller {
	vvvvv := &Controller{l, v, e, c, db}
	return vvvvv
}

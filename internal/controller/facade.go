package controller

import (
	"github.com/alexpts/edu-go/internal/controller/home"
	"github.com/alexpts/edu-go/internal/controller/notfound"
)

// aliases in one namespace package, re-export

type Home = home.ControllerHome
type NotFound = notfound.ControllerNotFound

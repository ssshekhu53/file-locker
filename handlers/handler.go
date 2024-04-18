package handlers

import (
	"errors"
	"gofr.dev/pkg/gofr"
	"ssshekhu53/file-locker/services"
)

type Handler interface {
	Init(ctx *gofr.Context) (interface{}, error)
	Lock(ctx *gofr.Context) (interface{}, error)
	Unlock(ctx *gofr.Context) (interface{}, error)
	Help(_ *gofr.Context) (interface{}, error)
}

type handler struct {
	service services.FileLocker
}

func New(service services.FileLocker) Handler {
	return &handler{service: service}
}

func (h *handler) Init(ctx *gofr.Context) (interface{}, error) {
	password := ctx.Param("password")
	if password == "" {
		return nil, errors.New("password is required")
	}

	err := h.service.Init(password)
	if err != nil {
		return nil, err
	}

	return "file locker initialized", nil
}

func (h *handler) Lock(_ *gofr.Context) (interface{}, error) {
	err := h.service.Lock()
	if err != nil {
		return nil, err
	}

	return "file locked", nil
}

func (h *handler) Unlock(ctx *gofr.Context) (interface{}, error) {
	password := ctx.Param("password")
	if password == "" {
		return nil, errors.New("password is required")
	}

	err := h.service.Unlock(password)
	if err != nil {
		return nil, err
	}

	return "file unlocked", nil
}

func (h *handler) Help(_ *gofr.Context) (interface{}, error) {
	return `File Locker CLI Tool

Usage:
  file-locker [command]

Available Commands:
  init      Create a directory named private and initialize the file locker
  lock      Hide the private directory
  unlock    Unhides the private directory`, nil
}

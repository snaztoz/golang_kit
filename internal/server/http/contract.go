package http

import (
	"context"
	"template/internal/router"
	mail "template/utils/mailer"

	"gorm.io/gorm"
)

type Server interface {
	Run(ctx context.Context, port int)
	Done()
}

type HttpServerCfg struct {
	DB        *gorm.DB
	SMTP      mail.Mailer
	Secret    string
	AesSecret string
}

func NewServer(h *HttpServerCfg) Server {
	return &httpServer{
		router: router.NewRouter(&router.RouterCfg{
			DB:        h.DB,
			SMTP:      h.SMTP,
			Secret:    h.Secret,
			AesSecret: h.AesSecret,
		}).Route(),
	}
}

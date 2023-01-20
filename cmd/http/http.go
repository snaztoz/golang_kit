package http

import (
	"context"
	h "template/internal/server/http"
)

func StartServer(ctx context.Context) {

	ht := h.NewServer()
	defer ht.Done()
	ht.Run(ctx)

	// return
	// http.ListenAndServe(":3000", r)
}

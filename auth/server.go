//
// Copyright (c) 2017 Mainflux
//
// SPDX-License-Identifier: Apache-2.0
//

package auth

import (
	"fmt"
	"net/http"

	"github.com/go-zoo/bone"
	"go.uber.org/zap"
)

// HTTPServer function
func httpServer() http.Handler {
	mux := bone.New()

	// Status
	mux.Get("/status", http.HandlerFunc(getStatus))

	mux.Get("/users", http.HandlerFunc(getAllUsers))
	mux.Get("/users/:id", http.HandlerFunc(getUserById))
	mux.Post("/users", http.HandlerFunc(createUser))
	mux.Delete("/users/:id", http.HandlerFunc(deleteUser))
	mux.Post("/login", http.HandlerFunc(login))
	return mux
}

func StartHTTPServer(config Config, errChan chan error) {
	cfg = config
	go func() {
		p := fmt.Sprintf(":%d", cfg.Port)
		logger.Info("Starting Export Client", zap.String("url", p))
		errChan <- http.ListenAndServe(p, httpServer())
	}()
}

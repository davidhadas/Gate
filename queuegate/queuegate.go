package queuegate

import (
	"net/http"

	"go.uber.org/zap"
)

func GateHandler(logger *zap.SugaredLogger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer h.ServeHTTP(w, r)

		logger.Info("DH> list headers ")
		for name, values := range r.Header {
			// Loop over all values for the name.
			for _, value := range values {
				logger.Info("%s: %s", name, value)
			}
		}
		/*
			userAgent := r.Header.Get("User-Agent")
			contentEncoding := r.Header.Get("content-encoding")
			transferEncoding := r.Header.Get("transfer-encoding")
			keepAlive := r.Header.Get("keep-alive")
			connection := r.Header.Get("Connection")
			xForwardedFor := r.Header.Get("x-forwarded-for")
			cacheControl := r.Header.Get("cache-control")
			via := r.Header.Get("via")

			logger.Info("DH> userAgent ", userAgent)
			logger.Info("DH> contentEncoding ", contentEncoding)
			logger.Info("DH> transferEncoding ", transferEncoding)
			logger.Info("DH> keepAlive ", keepAlive)
			logger.Info("DH> connection ", connection)
			logger.Info("DH> xForwardedFor ", xForwardedFor)
			logger.Info("DH> cacheControl ", cacheControl)
			logger.Info("DH> via ", via)
		*/
	})
}

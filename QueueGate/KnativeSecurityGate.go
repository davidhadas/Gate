package QueueGate

import (
	"net/http"

	"go.uber.org/zap"
)

func GateHandler(logger *zap.SugaredLogger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer h.ServeHTTP(w, r)

		userAgent := r.Header.Get("userAgent")
		contentEncoding := r.Header.Get("content-encoding")
		transferEncoding := r.Header.Get("transfer-encoding")
		keepAlive := r.Header.Get("keep-alive")
		connection := r.Header.Get("connection")
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

	})
}

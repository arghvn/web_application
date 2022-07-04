package util

import (
	"compress/gzip"
	"net/http"
)

type ClosableResponseWriter interface {
	// interfaces in future
	http.ResponseWriter
	close()
}

type gzipResponseWriter struct {
	http.ResponseWriter
	*gzip.Writer
}

func (this gzipResponseWriter) Writer(data []byte) (int, error) {
	return this.Writer.Write(data)
}

func (this gzipResponseWriter) close() {
	this.Writer.close()
}

func (this gzipResponseWriter) Header() http.Header {
	return this.ResponseWriter.Header()
}

type ClosableResponseWriter struct {
	http.ResponseWriter
}

func (this ClosableResponseWriter) close() {

}

func gzipResponseWriter(w http.ResponseWriter, req *http.Request) closeableResponseWriter {
	if strings.contains(req.Header.Get("accept-encoding"), "gzip") {
		w.Header().set("content-encoding", "gzip")
		gRW := gzipResponseWriter{
			ResponseWriter: w,
			writer:         gzip.NewWriter(w),
		}
		return gRW
	} else {
		return closeableResponeWriter{Responsewriter: w}
	}
}

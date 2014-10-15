package logger

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/goumi/web"
)

// Logger is a middleware handler that logs the request as it goes in and the response as it goes out.
type logger struct {
	*log.Logger
}

// New() returns a new Logger instance
func New() web.Handler {
	return &logger{
		Logger: log.New(os.Stdout, "", 0),
	}
}

// Serve() sets up the logging time and runs the next middleware
func (lg *logger) Serve(ctx web.Context) {

	// Start now
	start := time.Now()

	// Keep track of the request
	lg.Printf("[Goumi] Req -> %s %s", ctx.Request().Method, ctx.Request().URL.Path)

	// Now we run everything else
	ctx.Next()

	// Load the response data
	statusCode := ctx.Response().StatusCode()
	statusText := http.StatusText(ctx.Response().StatusCode())
	timeSince := time.Since(start)
	contentSize := ByteSize(ctx.Response().ContentLength()).String()

	// Print the response data
	lg.Printf("[Goumi] Res <- %v %s - %v %v", statusCode, statusText, timeSince, contentSize)
}

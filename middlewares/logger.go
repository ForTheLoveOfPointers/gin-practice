package middlewares

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupLogOutput() {
	file, err := os.Create("gin.log")

	if err != nil {
		errFmt := fmt.Errorf("ERROR || Unable to create or access gin.log file --- %w", err)
		fmt.Println(errFmt)
		return
	}

	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

}

func Logger() gin.HandlerFunc {

	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC822),
			params.Method,
			params.Path,
			params.StatusCode,
			params.Latency,
		)
	})
}

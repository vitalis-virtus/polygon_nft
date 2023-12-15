package logger

import (
	"log"
	"os"
)

func New() *log.Logger {
	return log.New(
		os.Stderr,
		"POLYGON_NFT",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC|log.Lshortfile,
	)
}

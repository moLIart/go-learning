package main

import (
	"time"

	"github.com/moLIart/go-course/internal/service"
)

func main() {
	service.StartProcessing(7 * time.Second)
}

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// blank context, only started
	ctx := context.Background()
	ctxTime := time.Second * 3
	ctx, cancel := context.WithTimeout(ctx, ctxTime)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached")
		return
	case <-time.After(1 * time.Second):
		fmt.Println("Hotel booked")
	}
}

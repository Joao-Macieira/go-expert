package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "password")
	bookHotel(ctx)
}

// In standard ctx is the first parameter in functions
func bookHotel(ctx context.Context) {
	token := ctx.Value("token")

	fmt.Println(token)
}

package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "user", "gopher")
	name := getUserName(ctx)

	fmt.Println(name)
}

func getUserName(ctx context.Context) string {
	return ctx.Value("user").(string)
}

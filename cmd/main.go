package main

import (
	"context"

	"ad-api/internal"
)

func main() {
	ctx := context.Background()
	internal.New(ctx).Run(ctx)
}

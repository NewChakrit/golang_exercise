package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("-------------")

	// result
	//context:         context.Background
	//context err:     <nil>
	//context type:   context.backgroundCtx
	//-------------

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("-------------")

	//	context:         context.Background.WithCancel
	//	context err:     <nil>
	//	context type:   *context.cancelCtx
	//	-------------

	cancel()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("context err:\t\t", cancel)
	fmt.Printf("context type:\t%T\n", cancel)
	fmt.Println("-------------")

	//	context:         context.Background.WithCancel
	//	context err:     context canceled
	//	context type:   *context.cancelCtx
	//	context err:             0x1047084d0
	//	context type:   context.CancelFunc
	//	-------------

}

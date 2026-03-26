package pkgdemo

import "context"

func ContextDemo() {
	bgCtx := context.Background()
	ctx, cancel := context.WithCancel(bgCtx)
	defer cancel()
	_ = ctx
}

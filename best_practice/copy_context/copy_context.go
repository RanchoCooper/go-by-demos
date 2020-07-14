package copy_context

import (
	"context"

	"git.llsapp.com/cc/pkg/context/tag"

	"golang.org/x/net/trace"
)

func CopyContextValue(ctx context.Context) context.Context {
	newCtx := context.Background()
	if ctx == nil {
		return newCtx
	}

	if span, _ := trace.FromContext(ctx); span != nil {
		newCtx = trace.NewContext(newCtx, span)
	}

	if tag.Has(ctx) {
		var newTags *tag.Tags
		_, tags := tag.Extract(ctx)
		newCtx, newTags = tag.Extract(newCtx)
		for k, v := range tags.Values() {
			newTags.Set(k, v)
		}
	}

	return newCtx
}

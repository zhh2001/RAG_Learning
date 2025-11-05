package main

import (
	"context"

	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/markdown"
	"github.com/cloudwego/eino/components/document"
)

func NewTrans(ctx context.Context) document.Transformer {
	splitter, err := markdown.NewHeaderSplitter(ctx, &markdown.HeaderConfig{
		Headers: map[string]string{
			"#":   "h1",
			"##":  "h2",
			"###": "h3",
		},
		TrimHeaders: false, // 是否在输出的内容中移除标题行
	})
	if err != nil {
		panic(err)
	}
	return splitter
}

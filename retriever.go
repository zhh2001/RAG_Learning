package main

import (
	"context"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	rr "github.com/cloudwego/eino-ext/components/retriever/redis"
)

func NewRetriever(ctx context.Context, embedder *ark.Embedder) *rr.Retriever {
	retriever, err := rr.NewRetriever(ctx, &rr.RetrieverConfig{
		Client:       redisClient,
		Index:        "my_index",
		VectorField:  "vector_content",
		ReturnFields: []string{"content", "vector_content"},
		TopK:         2,
		Embedding:    embedder,
	})
	if err != nil {
		panic(err)
	}
	return retriever
}

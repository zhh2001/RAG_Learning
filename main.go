package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	InitRedis()
	StoreData()

	ctx := context.Background()
	embedder := NewArkEmbedder(ctx)
	retriever := NewRetriever(ctx, embedder)

	results, err := retriever.Retrieve(ctx, "刘氏家族")
	if err != nil {
		panic(err)
	}

	for i, doc := range results {
		fmt.Println(i, doc.Content)
	}
}

func StoreData() {
	ctx := context.Background()
	embedder := NewArkEmbedder(ctx)
	indexer := NewArkIndexer(ctx, embedder)
	splitter := NewTrans(ctx)

	bs, err := os.ReadFile("./doc.md")
	if err != nil {
		panic(err)
	}
	docs := []*schema.Document{
		{
			ID:      "doc1",
			Content: string(bs),
		},
	}
	results, err := splitter.Transform(ctx, docs)
	if err != nil {
		panic(err)
	}
	for i, doc := range results {
		doc.ID = docs[0].ID + "_" + strconv.Itoa(i+1)
	}
	ids, err := indexer.Store(ctx, results)
	if err != nil {
		panic(err)
	}
	fmt.Println("ids:", ids)
}

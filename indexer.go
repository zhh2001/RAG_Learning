package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	ri "github.com/cloudwego/eino-ext/components/indexer/redis"
	goredis "github.com/redis/go-redis/v9"
)

func createIndex() {
	ctx := context.Background()

	keyPrefix := "eino:"    // keyPrefix should be the prefix of keys you write to redis and want to retrieve.
	indexName := "my_index" // indexName should be used in redis retriever.

	// schemas should match DocumentToHashes configured in IndexerConfig.
	schemas := []*goredis.FieldSchema{
		{
			FieldName: "content",
			FieldType: goredis.SearchFieldTypeText,
			Weight:    1,
		},
		{
			FieldName: "vector_content",
			FieldType: goredis.SearchFieldTypeVector,
			VectorArgs: &goredis.FTVectorArgs{
				FlatOptions: &goredis.FTFlatOptions{
					Type:           "FLOAT32", // BFLOAT16 / FLOAT16 / FLOAT32 / FLOAT64. BFLOAT16 and FLOAT16
					Dim:            2560,      // keeps same with dimensions of Embedding
					DistanceMetric: "COSINE",  // L2 / IP / COSINE
				},
				HNSWOptions: nil,
			},
		},
		{
			FieldName: "extra_field_number",
			FieldType: goredis.SearchFieldTypeNumeric,
		},
	}

	options := &goredis.FTCreateOptions{
		OnHash: true,
		Prefix: []any{keyPrefix},
	}

	result, err := redisClient.FTCreate(ctx, indexName, options, schemas...).Result()
	if err != nil && err.Error() != "Index already exists" {
		panic(err)
	}

	fmt.Println(result)
}

func NewArkIndexer(ctx context.Context, embedder *ark.Embedder) *ri.Indexer {
	createIndex()
	indexer, err := ri.NewIndexer(ctx, &ri.IndexerConfig{
		Client:    redisClient,
		KeyPrefix: "eino:",
		Embedding: embedder,
	})
	if err != nil {
		panic(err)
	}
	return indexer
}

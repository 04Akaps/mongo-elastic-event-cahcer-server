package main

import (
	"flag"
	"mongo-event-cacher/cmd/app"
	"mongo-event-cacher/config"
)

var configFlag = flag.String("config", "./config.toml", "configuration toml file path")

func main() {
	flag.Parse()
	cfg := config.NewConfig(*configFlag)
	app.NewListener(cfg)

	// clientOptions := options.Client().ApplyURI("mongodb://treeroot:wtdev!34@10.15.120.20:10001/")

	// // MongoDB 클라이언트 생성
	// client, err := mongo.Connect(context.Background(), clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// collection := client.Database("wemix-play-nft-market-db").Collection("nfts")
	// pipeline := []bson.M{}

	// if cursor, err := collection.Watch(ctx, pipeline, changeStreamOptions); err != nil {
	// 	fmt.Println("Err : ", err)
	// } else {
	// 	fmt.Println("Started")
	// 	for cursor.Next(ctx) {

	// 		var changeEvent types.NftChangeEvent

	// 		if err = cursor.Decode(&changeEvent); err != nil {
	// 			return
	// 		}

	// 		req := elastic.NewBulkUpdateRequest()

	// 		req.UseEasyJSON(true)
	// 		req.Id(changeEvent.FullDocument.ID.String())
	// 		req.Index("wemix-play-nft-market-db.nfts")
	// 		req.Doc(&changeEvent.FullDocument)
	// 		req.DocAsUpsert(true)

	// 		bulk.Add(req)
	// 		bulkResp, err := bulk.Do(ctx)

	// 		// Check if the Do() method returned any errors
	// 		if err != nil {
	// 			log.Fatalf("bulk.Do(ctx) ERROR:", err)
	// 		} else {
	// 			fmt.Println(bulkResp.Indexed())
	// 		}

	// 	}
	// 	defer cursor.Close(ctx)
	// }

	// time.Sleep(3 * time.Second)

}

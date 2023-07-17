
	// clientOptions := options.Client().ApplyURI("mongodb://treeroot:wtdev!34@10.15.120.20:10001/")

	// // MongoDB 클라이언트 생성
	// client, err := mongo.Connect(context.Background(), clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// collection := client.Database("wemix-play-nft-market-db").Collection("nfts")
	// pipeline := []bson.M{}

	// changeStreamOptions := options.ChangeStream()
	// changeStreamOptions.SetFullDocument(options.UpdateLookup)

	// ctx := context.TODO()

	// esClient, err := elastic.NewClient(
	// 	elastic.SetBasicAuth(
	// 		"hojin",
	// 		"testHojin",
	// 	),
	// 	elastic.SetURL("http://localhost:9200/"),
	// 	elastic.SetSniff(false),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// indexName := "some_index"

	// // Declare a string slice with the index name in it
	// indices := []string{indexName}

	// // Instantiate a new *elastic.IndicesExistsService
	// existService := elastic.NewIndicesExistsService(esClient)

	// existService.Index(indices)

	// // Have Do() return an API response by passing the Context object to the method call
	// exist, err := existService.Do(ctx)

	// if err != nil {
	// 	log.Fatalf("IndicesExistsService.Do() ERROR:", err)
	// } else if exist == false {
	// 	fmt.Println("nOh no! The index", indexName, "doesn't exist.")
	// 	fmt.Println("Create the index, and then run the Go script again")
	// } else if exist == true {
	// 	fmt.Println("Index name:", indexName, " exists!")
	// }

	// bulk := esClient.Bulk()

	// if err != nil {
	// 	panic(err)
	// }

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

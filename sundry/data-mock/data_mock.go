package datamock

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	dbUri = "mongodb://root:LS#asdfa1e3zef@" +
		"dds-2zef2ff8e79f49741.mongodb.rds.aliyuncs.com:3717," +
		"dds-2zef2ff8e79f49742.mongodb.rds.aliyuncs.com:3717" +
		"/admin?replicaSet=mgset-21366747" // 数据库地址
	dbName         = "global_storage_1040"   // 数据库名称
	collectionName = "user_server_info_data" // 集合名称

	//totalCount = int64(900000) // 需要插入元素的总量

	idGoNum    = 1      // 序列号worker pool数量
	idChanSize = 100000 // 序列号channel大小

	dataChanSize  = 1000 // 元素channel大小
	productGoNum  = 2    // 生产者worker pool数量
	consumerGoNum = 2    // 消费者worker pool数量

	collectionManyNum = 1000 // 每次插入数据量  如一次创造1000个序列号,总量和预定的总数会有千内的误差
)

var wg sync.WaitGroup
var currentIndex = int64(1) // 当前偏移

func Do(totalCount int64) {
	ctx := context.Background()
	collection := ServerCollection(ctx) // 获取集合对象

	count, err := collection.CountDocuments(ctx, bson.D{}) // 获取数据总数
	if err != nil {
		log.Fatal(count)
	}

	if count >= totalCount {
		fmt.Printf("需要%v条数据，数据库已有%v条数据，不需要再造啦~", totalCount, count)
		return
	}

	currentIndex = count + 1 // 以数据库已有数据条数为基准

	fmt.Printf("----开始mock数据：%v-----\n", time.Now().Unix())
	UniqueIdManyChan := make(chan []int64, idChanSize)
	wg.Add(1)
	go UniqueIdMany(totalCount, UniqueIdManyChan)

	dataItemListChan := make(chan []interface{}, dataChanSize)

	wg.Add(1)
	ProducerManyPool(UniqueIdManyChan, dataItemListChan)

	wg.Add(1)
	ConsumerManyPool(ctx, collection, dataItemListChan)

	wg.Wait()

	fmt.Printf("----mock数据结束：%v-----\n", time.Now().Unix())
}

// UniqueId 生成递增序列号（唯一复合索引需要）
func UniqueIdMany(totalCount int64, uniqueIdChan chan []int64) {
	var wgId sync.WaitGroup
	wgId.Add(idGoNum)
	var mutex sync.Mutex

	for i := 0; i < idGoNum; i++ {
		go func() {
			gameOver := false
			for {
				mutex.Lock()
				if currentIndex <= totalCount {
					currentIndexTemp := currentIndex
					currentIndex = currentIndex + collectionManyNum // 基数偏移更改，给其他协程使用

					var ids [collectionManyNum]int64
					for i := int64(0); i < collectionManyNum; i++ {
						ids[i] = currentIndexTemp
						currentIndexTemp++
					}
					uniqueIdChan <- ids[:]

				} else {
					gameOver = true
				}
				mutex.Unlock()

				if gameOver {
					wgId.Done()
					break
				}
			}
		}()
	}

	wgId.Wait()
	println("----序列号初始化完成----")
	close(uniqueIdChan)
	wg.Done()
}

// ProducerManyPool 生成需要插入的数据
func ProducerManyPool(UniqueIdManyChan chan []int64, dataItemListChan chan<- []interface{}) {
	var wgProduct sync.WaitGroup
	wgProduct.Add(productGoNum)
	for i := 0; i < productGoNum; i++ {
		go func(UniqueIdManyChan chan []int64, dataItemListChan chan<- []interface{}) {
			gameOver := false
		FORLOOP:
			for {
				select {
				case ids, ok := <-UniqueIdManyChan:
					if ok == false {
						gameOver = true
					} else {
						var dataItemList [collectionManyNum]interface{}
						rand.Seed(time.Now().UnixNano())
						for index, id := range ids {
							dataItem := map[string]interface{}{
								"_id":       id,
								"app_id":    rand.Int31n(100000000),
								"user_id":   id,
								"role_guid": rand.Int31n(100000000),
								"server_id": rand.Int31n(100000000),
								"data":      rand.Int31n(100000000),
							}
							dataItemList[index] = &dataItem
						}
						dataItemListChan <- dataItemList[:]
					}
				default:
					time.Sleep(time.Millisecond)
				}
				if gameOver {
					wgProduct.Done()
					break FORLOOP
				}
			}
		}(UniqueIdManyChan, dataItemListChan)
	}
	wgProduct.Wait()
	println("----模拟数据生产完成-----")
	close(dataItemListChan)
	wg.Done()
}

// ConsumerManyPool 消费入库
func ConsumerManyPool(ctx context.Context, collection *mongo.Collection, UniqueIdManyChan chan []interface{}) {
	var wgConsumer sync.WaitGroup
	wgConsumer.Add(consumerGoNum)
	for i := 0; i < consumerGoNum; i++ {
		go func(ctx context.Context, collection *mongo.Collection, UniqueIdManyChan chan []interface{}) {
			gameOver := false
			for {
				select {
				case dataItemList, ok := <-UniqueIdManyChan:
					if ok == false {
						gameOver = true
					} else {
						// 插入数据
						_, err := collection.InsertMany(ctx, dataItemList)
						if err != nil {
							fmt.Printf("err:%v\n", err)
						}
					}
				default:
					time.Sleep(time.Millisecond)
				}
				if gameOver {
					wgConsumer.Done()
					return
				}
			}
		}(ctx, collection, UniqueIdManyChan)
	}

	wgConsumer.Wait()
	println("----元素插入数据库完成-----")
	wg.Done()
}

// ServerCollection 获取mongo集合对象
func ServerCollection(ctx context.Context) *mongo.Collection {
	opts := options.Client().ApplyURI(dbUri) // 连接数据库
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	// 判断服务是不是可用
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// 获取数据库和集合
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}

//@time:2020/01/20
//@desc:

package mockdata

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golearn/util"
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	user = "root"
	//pwd         = "LS#asdfa1e3zef"
	pwd = "5elsT6KtjrqV1"
	//connections = "dds-2zef2ff8e79f49741.mongodb.rds.aliyuncs.com:3717,dds-2zef2ff8e79f49742.mongodb.rds.aliyuncs.com:3717"
	connections = "10.46.134.91:27017,10.46.148.237:27017,10.46.157.47:27017"

	//replicaSet = "mgset-21366747"

	dbName         = "global_storage_1040" // 数据库名称
	collectionName = "test"                // 集合名称

	consumerGoNum = 2 // 消费者worker pool数量
)

type mockManager struct {
	wg           sync.WaitGroup
	currentIndex int64
	coll         *mongo.Collection
	remain       int // 剩余条数
	process      int // 已入库条数
	*config
}

func NewMockManager(conf *config) *mockManager {
	return &mockManager{wg: sync.WaitGroup{}, config: conf, currentIndex: conf.beginId}
}

func (m *mockManager) Mock() {
	ctx := context.Background()
	m.initMongo(ctx) // 初始化数据库信息

	count, err := m.coll.CountDocuments(ctx, bson.D{}) // 获取数据总数
	if err != nil {
		log.Fatal(count)
	}
	if count >= m.total {
		fmt.Printf("需要%v条数据，数据库已有%v条数据，不需要再造啦~", m.total, count)
		return
	}
	m.remain = int(m.total - count)
	fmt.Printf("当前已有%d条数据,目标%d条\n", count, m.total)
	m.total = int64(m.remain)

	maxId := m.maxId(ctx)
	fmt.Printf("当前数据库最大id:%d,期望id:%d,使用较大者\n", maxId, m.currentIndex)
	if m.currentIndex < maxId {
		m.currentIndex = maxId + 1
	}

	fmt.Printf("----开始mock数据：%v-----\n", time.Now().Unix())
	uniqueIdManyChan := make(chan []int64, m.onceCount/m.insertOnce)
	dataItemListChan := make(chan []interface{}, m.onceCount/m.insertOnce)
	go util.WithRecover(func() {
		m.uniqueIdGenerator(uniqueIdManyChan)
	})
	go util.WithRecover(func() {
		m.producer(uniqueIdManyChan, dataItemListChan)
	})
	m.consumer(ctx, dataItemListChan)

	fmt.Printf("----mock数据结束：%v-----\n", time.Now().Unix())
}

// UniqueId 生成递增序列号（唯一复合索引需要）
func (m *mockManager) uniqueIdGenerator(uniqueIdChan chan []int64) {
	for {
		if m.remain <= 0 {
			close(uniqueIdChan)
			break
		}
		currentIndexTemp := m.currentIndex
		m.currentIndex = m.currentIndex + int64(m.insertOnce)

		var ids = make([]int64, m.insertOnce)
		for i := range ids {
			ids[i] = currentIndexTemp
			currentIndexTemp++
		}
		uniqueIdChan <- ids

		m.remain -= m.insertOnce
	}
	fmt.Println("----序列号初始化完毕----")
}

// ProducerManyPool 生成需要插入的数据
func (m *mockManager) producer(uniqueIdManyChan chan []int64, dataItemListChan chan<- []interface{}) {
	var count int
	for {
		ids, ok := <-uniqueIdManyChan
		if !ok {
			close(dataItemListChan)
			break
		}
		if count == m.onceCount {
			fmt.Printf("本轮数据生成结束-数据库进度%d/%d-%d\n", m.process, m.total, time.Now().Unix())
			time.Sleep(m.sleeping)
			count = 0
		}
		count += len(ids)
		var dataItemList = make([]interface{}, m.insertOnce)
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
		dataItemListChan <- dataItemList
	}
	fmt.Printf("本轮数据生成结束-数据库进度%d/%d-%d\n", m.process, m.total, time.Now().Unix())
	fmt.Println("----模拟数据生成完毕-----")
}

// ConsumerManyPool 消费入库
func (m *mockManager) consumer(ctx context.Context, UniqueIdManyChan chan []interface{}) {
	var wgConsumer sync.WaitGroup
	wgConsumer.Add(consumerGoNum)

	for i := 0; i < consumerGoNum; i++ {
		go util.WithRecover(func() {
			for {
				select {
				case dataItemList, ok := <-UniqueIdManyChan:
					if !ok {
						wgConsumer.Done()
						return
					}
					// 插入数据
					_, err := m.coll.InsertMany(ctx, dataItemList)
					if err != nil {
						fmt.Printf("err:%v\n", err)
						m.remain += len(dataItemList)
					} else {
						m.process += len(dataItemList)
					}
				default:
					time.Sleep(time.Millisecond)
				}
			}
		})
	}
	wgConsumer.Wait()
	fmt.Printf("----元素插入数据库完毕 数据库进度%d/%d-----\n", m.process, m.total)
}

// ServerCollection 获取mongo集合对象
func (m *mockManager) initMongo(ctx context.Context) {
	dbUri := fmt.Sprintf("mongodb://%s:%s@%s/admin",
		user, pwd, connections)
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
	m.coll = client.Database(dbName).Collection(collectionName)
}

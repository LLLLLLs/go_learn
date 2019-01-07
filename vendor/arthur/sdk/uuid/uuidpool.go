/*
Author      : lls
Time        : 2018/08/21
Description : Guid缓存
*/

package uuid

import (
	"arthur/conf"
	"arthur/utils/errors"
	"arthur/utils/log"
	"bytes"
	"fmt"
	"github.com/pborman/uuid"
	"strconv"
	"sync"
	"time"
)

var (
	uuidAddr string
	poolSize = 500
	GuidSrv  *guidPool
	First    = true
)

type guidPool struct {
	uuids chan string
	once  *sync.Once
}

func Init(addr string, size ...int) {
	if len(size) != 0 && size[0] != 0 {
		poolSize = size[0]
	}
	if conf.IsMode(conf.TEST) { // 减少测试时获取uuid数量
		poolSize = 10
	}
	uuidAddr = addr
	GuidSrv = &guidPool{
		uuids: make(chan string, poolSize),
		once:  &sync.Once{},
	}
	GuidSrv.once.Do(GuidSrv.getGuidFromServer)
}

func (g *guidPool) getGuidFromServer() {
	log.Debug("guid service starting...", time.Now())
	go func() {
		count := 1
		for {
			if count > 1 {
				//dclog.Error("Uuid服务异常","uuid")
				log.Error("GUID服务异常")
				g.getGuidFromLocal()
				count = 1
			}
			if ok := g.fetchGuidsFromServer(); !ok {
				count++
				continue
			}
		}
	}()
	if conf.IsMode(conf.TEST) {
		for First {
			time.Sleep(time.Millisecond)
		}
	}
}

func (g *guidPool) getGuidFromLocal() {
	for i := 0; i < poolSize; i++ {
		g.uuids <- newUuid()
	}
	if First {
		log.Debug("guid service success from local...", time.Now())
		First = false
	}
}

func (g *guidPool) fetchGuidsFromServer() (flag bool) {
	defer func() {
		if e := recover(); e != nil {
			log.Error(errors.New(fmt.Sprintf("%+v", e)))
			flag = false
		}
	}()
	ids := GetUuids(poolSize)
	if First {
		log.Debug("guid service success from server...", time.Now())
		First = false
	}
	count := 0
	for i := range ids {
		i64, err := strconv.ParseInt(ids[i], 10, 64)
		if err != nil {
			count++
			log.Debug(err)
			if count > 10 {
				panic(err)
			}
			continue
		}
		g.uuids <- strconv.FormatInt(i64, 36)
	}
	return true
}

func (g *guidPool) GetOneUuid() string {
	return <-g.uuids
}

func GetUuid() string {
	return GuidSrv.GetOneUuid()
}

func newUuid() string {
	var buffer bytes.Buffer
	for _, b := range uuid.NewRandom() {
		b1 := strconv.FormatInt(int64(b>>2&0x0f), 16)
		buffer.WriteString(b1)
	}
	str := string([]rune(buffer.String())[:15])
	i64, err := strconv.ParseInt(str, 16, 64)
	if err != nil {
		log.Error(errors.Wrap(err, "parse uuid error: "))
	}
	return strconv.FormatInt(i64, 36)
}

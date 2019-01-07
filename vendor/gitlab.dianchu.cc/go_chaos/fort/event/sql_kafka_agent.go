package event
//
//import (
//	"context"
//	"errors"
//	"fmt"
//	"time"
//
//	"gitlab.dianchu.cc/go_chaos/fort/factory"
//	"gitlab.dianchu.cc/go_chaos/fort/tools/syslog"
//	"gitlab.dianchu.cc/go_chaos/fort/utils"
//)
//
//type KafkaAgent struct {
//	SQLBaseAgent
//	Topic string
//}
//
//func (ka *KafkaAgent) produce(ctx context.Context, data []byte) error {
//	return nil
//}
//
//func (ka *KafkaAgent) consume(ctx context.Context, id string) error {
//	var (
//		deadline time.Time
//		ok       bool
//	)
//	deadline, ok = ctx.Deadline()
//	if ok && time.Now().After(deadline) {
//		errStr := fmt.Sprintf("Commands request %v timeout!", ctx.Value(utils.TRACE_ID))
//		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
//		return errors.New(errStr)
//	}
//	return nil
//}
//
//func (ka *KafkaAgent) Submit(ctx context.Context, sqlQuery *factory.SQLInfo) error {
//	return ka.submit(sqlQuery)
//}
//
//func (ka *KafkaAgent) Extend(commands *[]*factory.SQLInfo) error {
//	return ka.extend(commands)
//}
//
//func (ka *KafkaAgent) Commit(ctx context.Context) error {
//	var (
//		data []byte
//		err  error
//	)
//	data, err = utils.Marshal(ka.transaction)
//	if err != nil {
//		return err
//	}
//	err = ka.produce(ctx, data)
//	if err != nil {
//		return err
//	}
//	return ka.consume(ctx, ka.transaction[ka.Topic].Id)
//}
//
//func NewKafkaAgent(ctx context.Context, topic string) (*KafkaAgent, error) {
//	var (
//		ka  *KafkaAgent
//		err error
//	)
//	ka = new(KafkaAgent)
//	err = ka.initCommandMemory(ctx, topic)
//	if err != nil {
//		return nil, err
//	}
//	ka.Topic = topic
//	return ka, nil
//}

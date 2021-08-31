// Time        : 2019/08/07
// Description :

package string

import (
	"fmt"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	str := "123"
	([]byte)(str)[1] = 'b'
	fmt.Printf("%p\n", &str)
	fmt.Printf("%p\n", ([]byte)(str))
}

func TestSplit(t *testing.T) {
	str := "hello;"
	list := strings.Split(str, ";")
	fmt.Println(list, len(list))
}

func TestPrintStack(t *testing.T) {
	fmt.Println("subscriber name:  \nerror info: ptr is not struct type! \nstack info:goroutine 217630 [running]:\ngitlab.dianchu.cc/Taotie/event_core/base.(*EventBase).doRecover(0xc0004264e0, 0xc00fe4a2a0, 0x0, 0xc0013af938, 0x1, 0x1, 0x0, 0x0)\n\tF:/go_project/src/vikingar/vendor/gitlab.dianchu.cc/Taotie/event_core/base/eventbase.go:276 +0xee\npanic(0x1d63540, 0xc021ee8a20)\n\tC:/Go/src/runtime/panic.go:975 +0x429\ngitlab.dianchu.cc/honey-comb/util/runtime.Must(0x2670280, 0xc021ee8a20)\n\tF:/go_project/src/vikingar/vendor/gitlab.dianchu.cc/honey-comb/util/runtime/runtime.go:10 +0x82\nvikingar/boundary-context/long-house/infrastructure/repository.RoleRepository.Get(0x26ac8e0, 0xc000a5c180, 0x26a42e0, 0xc000a2d280, 0x2c84bb40, 0xc000c00748, 0xc023e28200, 0xc, 0x15, 0x43c66d, ...)\n\tF:/go_project/src/vikingar/boundary-context/long-house/infrastructure/repository/role_repository.go:26 +0x205\nvikingar/boundary-context/long-house/application.QueryService.GetRole(0x26a11e0, 0xc000a2dec0, 0x26b1f20, 0xc000baa7c0, 0x26a0f20, 0xc000baa720, 0x26a11e0, 0xc000a2dec0, 0x2c84bb40, 0xc000c00748, ...)\n\tF:/go_project/src/vikingar/boundary-context/long-house/application/query_service.go:63 +0x93\nvikingar/presentation/long-house/subscribers.RoleLoginedSubscriber.Handler(0x26b1f20, 0xc000baa7c0, 0x26a0f20, 0xc000baa720, 0x26a11e0, 0xc000a2dec0, 0x2685a60, 0xc000ad8080, 0x26a11e0, 0xc000a2dec0, ...)\n\tF:/go_project/src/vikingar/presentation/long-house/subscribers/role_logined_subscriber.go:39 +0x158\ngitlab.dianchu.cc/honey-comb/event-bus.(*bus).subscribe.func2(0x26a3a20, 0xc02362fbf0, 0x1ed6680, 0xc000cce9c0, 0x0, 0x0)\n\tF:/go_project/src/vikingar/vendor/gitlab.dianchu.cc/honey-comb/event-bus/bus_impl.go:377 +0x38e\ngitlab.dianchu.cc/Taotie/event_core/base.(*EventBase).doExec(0xc0004264e0, 0xc00fe4a2a0, 0xc0013afb01, 0x0, 0x0)\n\tF:/go_project/src/vikingar/vendor/gitlab.dianchu.cc/Taotie/event_core/base/eventbase.go:198 +0x5eb\ngitlab.dianchu.cc/Taotie/event_core/base.(*EventBase).ExecSubCallBackFunc(0xc0004264e0, 0xc00fe4a2a0, 0x0, 0x0)\n\tF:/go_project/src/vikingar/vendor/gitlab.dianchu.cc/Taotie/event_core/base/eventbase.go:86 +0x49\ngitlab.dianchu.cc/Taotie/event_core/pub.(*EventPub).doLocalPublish(0xc0004d9b00, 0xc00fe4a2a0, 0x0, 0x0, 0x0, 0x0)\n\tF:/go_project/src/vikingar/vendor/gitlab.dianchu.cc/Taotie/event_core/pub/eventpub.go:128 +0x1e2\ngitlab.dianchu.cc/Taotie/event_core/pub.(*EventPub).doPublish(0xc0004d9b00, 0xc00fe4a2a0, 0xc00fe4a240, 0x2, 0x6, 0x0, 0x0)\n\tF:/go_project/src/vikingar/vendor/gitlab.dianchu.cc/Taotie/event_core/pub/eventpub.go:192 +0x732\ngitlab.dianchu.cc/Taotie/event_core/pub.(*EventPub).Publish(0xc0004d9b00, 0x208515b, 0xc, 0x1ed6680, 0xc000cce9c0, 0xc00fe4a240, 0x2, 0x6, 0x0, 0x0)\n\tF:/go_project/src/vikingar/vendor/gitlab.dianchu.cc/Taotie/event_core/pub/eventpub.go:43 +0x123\ngitlab.dianchu.cc/honey-comb/event-bus.(*bus).publishToCore(0xc0003ea960, 0x2c84bb40, 0xc000c00748, 0x208515b, 0xc, 0x1ed6680, 0xc000cce9c0, 0xc018c8a390, 0x1f89e01, 0x0, ...)\n\tF:/go_project/src/vikingar/vendor/gitlab.dianchu.cc/honey-comb/event-bus/bus_impl.go:220 +0x482\ngitlab.dianchu.cc/honey-comb/event-bus.(*bus).Publish(0xc0003ea960, 0x2c84bb40, 0xc000c00748, 0x2680220, 0xc0326986e0, 0x0, 0x0)\n\tF:/go_project/src/vikingar/vendor/gitlab.dianchu.cc/honey-comb/event-bus/bus_impl.go:161 +0xa46\ngitlab.dianchu.cc/honey-comb/suite/middleware/mid.EventPublisher.Handle(0x26a4260, 0xc000a5ccc0, 0x26b1860, 0xc000c00748, 0x0, 0x0)\n\tF:/go_project/src/vikingar/vendor/gitlab.dianchu.cc/honey-comb/suite/middleware/mid/mid_event.go:34 +0x244\ngitlab.dianchu.cc/honey-comb/suite/middleware/adapter.handlerAdapter.Handle(0x2673f80, 0xc000a2dab0, 0x26b13e0, 0xc01dc082d0, 0x0, 0x0)\n\tF:/go_project/src/vikingar/vendor/gitlab.dianchu.cc/honey-comb/suite/middleware/adapter/util.go:35 +0xaa\ngitlab.dianchu.cc/Taotie/moat/v2/mcontext.(*Context).Next(0xc01dc082d0, 0x0, 0x0, 0x0, 0x0, 0x0)\n\tF:/go_project/src/vikingar/vendor/gitlab.dianchu.cc/Taotie/moat/v2/mcontext/context.go:114 +0x16d\ngitlab.dianchu.cc/honey-comb/suite/middleware/adapter.(*ContextAdapter).Next(0xc000c00730, 0x0, 0x0, 0x0, 0x0, 0x0)\n\tF:/go_project/src/vikingar/vendor/gitlab.dianchu.cc/honey-comb/suite/mid")
}

func TestPrint(t *testing.T) {
	fmt.Println("hello"[:2])
}

func MatchCount(r, another string) int {
	length := len(r)
	if len(another) < length {
		length = len(another)
	}
	for i := 0; i <= length; i++ {
		if r[:i] != another[:i] {
			return i - 1
		}
	}
	return length
}

func TestMatchCount(t *testing.T) {
	fmt.Println(MatchCount("abc", "abd"))
	fmt.Println(MatchCount("", ""))
	fmt.Println(MatchCount("", "abc"))
	fmt.Println(MatchCount("abc", "abc"))
}

func TestJoin(t *testing.T) {
	fmt.Println(strings.Join([]string{}, ","))
	fmt.Println(strings.Join([]string{"one"}, ","))
	fmt.Println(strings.Join([]string{"one", "two"}, ","))
}

func TestPart(t *testing.T) {
	strList := []string{"one", "two", "three"}
	fmt.Println(strList[:0])
	fmt.Println(strList[:1])
	fmt.Println(strList[:2])
	fmt.Println(strList[:3])
}

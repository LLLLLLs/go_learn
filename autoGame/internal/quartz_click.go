// 文件名：click_quartz.go
package internal

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework ApplicationServices
#include <ApplicationServices/ApplicationServices.h>

void QuartzClickWithoutMoving(int x, int y) {
    // 保存当前鼠标位置
    CGEventRef event = CGEventCreate(NULL);
    CGPoint original = CGEventGetLocation(event);
    CFRelease(event);

    // 创建点击事件（按下 + 抬起）
    CGEventRef click_down = CGEventCreateMouseEvent(NULL, kCGEventLeftMouseDown,
        CGPointMake(x, y), kCGMouseButtonLeft);
    CGEventRef click_up = CGEventCreateMouseEvent(NULL, kCGEventLeftMouseUp,
        CGPointMake(x, y), kCGMouseButtonLeft);

    // 点击
    CGEventPost(kCGHIDEventTap, click_down);
    CGEventPost(kCGHIDEventTap, click_up);

    CFRelease(click_down);
    CFRelease(click_up);

    // 恢复鼠标位置
    CGEventRef move_back = CGEventCreateMouseEvent(NULL, kCGEventMouseMoved,
        original, kCGMouseButtonLeft);
    CGEventPost(kCGHIDEventTap, move_back);
    CFRelease(move_back);
}
*/
import "C"

func QuartzClick(x, y int) {
	C.QuartzClickWithoutMoving(C.int(x), C.int(y))
}

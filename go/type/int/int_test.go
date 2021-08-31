// Time        : 2019/11/04
// Description :

package int

import (
	"fmt"
	"testing"
)

func TestSignedOverflow(t *testing.T) {
	a, b := int8(0b01111111), int8(1)
	fmt.Println((a + b) >> 1)          // -64
	fmt.Println(int8(uint8(a+b) >> 1)) // 64
	fmt.Println((a + b) / 2)           // -64
	fmt.Println(int8(uint8(a+b) / 2))  // 64
}

const (
	// 高优先级，数字越大越高
	_ uint = iota + 0 // 占位符，无意义
	DailyTimerRoleLoginPriority
	TimerRoleLoginPriority
	WorldRoleLoginedPriority
)

// 低优先级，数字越大越低
const (
	_ uint = iota + 0 // 占位符，无意义
	HonorReportRoleLoginPriority
	ExplorerVoyageRoleLoginedPriority
	RushRoleLoginPriority
	SailingSailingLoginPriority
	RuinsTreasureRoleLoginPriority
	WarriorsAxeRoleLoginedPriority
	GiftPackageRoleLoginedPriority
	PayRebateRoleLoginedPriority
	ActivityRoleLoginPriority

	VirtualGameAccountRoleLoginedPriority
)

func TestIota(t *testing.T) {
	fmt.Println(DailyTimerRoleLoginPriority)
	fmt.Println(TimerRoleLoginPriority)
	fmt.Println(WorldRoleLoginedPriority)
	fmt.Println(HonorReportRoleLoginPriority)
	fmt.Println(ExplorerVoyageRoleLoginedPriority)
	fmt.Println(RushRoleLoginPriority)
	fmt.Println(SailingSailingLoginPriority)
	fmt.Println(RuinsTreasureRoleLoginPriority)
	fmt.Println(WarriorsAxeRoleLoginedPriority)
	fmt.Println(GiftPackageRoleLoginedPriority)
	fmt.Println(PayRebateRoleLoginedPriority)
	fmt.Println(ActivityRoleLoginPriority)
	fmt.Println(VirtualGameAccountRoleLoginedPriority)
}

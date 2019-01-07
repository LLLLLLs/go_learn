/*
Created on 2018/6/21

author: WenHao Shan

Content:
*/
package mail

import (
	"arthur/conf"
	"arthur/env"
	"arthur/sdk/zookeeper"
	"arthur/utils/randomutils"
	"arthur/utils/timeutils"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func ZKConfInit() {
	zookeeper.Init(env.ZK_HOST, env.ZK_AUTH)

	// 初始化zk设置
	conf.Init(env.ZK_ROOT)
	cfg := conf.Config

	InitMail(
		cfg.Mail.Timeout,
		cfg.Mail.Method,
		cfg.Mail.GrpcAddr,
		cfg.Mail.HttpAddr,
		cfg.Mail.Version,
		cfg.Mail.IsPush,
	)
}

type TestInit struct {
	GameId   int
	AppId    int
	ServerId int
	LogFlag  int
}

func newInit() TestInit {
	testInit := new(TestInit)
	testInit.GameId = 105
	testInit.AppId = 1040
	testInit.ServerId = 5
	testInit.LogFlag = 0
	return *testInit
}

// 模拟发送邮件
func mockSendMail(toRole string, needRead bool) (string, error) {
	testInit := newInit()
	mailManager := New(testInit.GameId, testInit.AppId, testInit.ServerId, testInit.LogFlag)
	celId := strconv.Itoa(randomutils.RandomInt(1, 1000000000))
	err := mailManager.MockSendMail(celId, toRole, needRead)
	if err != nil {
		return "", err
	}
	return celId, nil
}

func newTest(sendMile bool, toRole string, needRead bool) (*Manager, []string, error) {
	testInit := newInit()
	mailManager := New(testInit.GameId, testInit.AppId, testInit.ServerId, testInit.LogFlag)
	var celId string
	var err error
	if sendMile {
		celId = strconv.Itoa(randomutils.RandomInt(1, 1000000000))
	} else {
		celId, err = mockSendMail(toRole, needRead)
		if err != nil {
			return nil, nil, err
		}
	}
	celIds := make([]string, 1, 1)
	celIds[0] = celId
	return mailManager, celIds, nil
}

// 测试从TestMain进入, 依次执行测试用例后, 最后从TestMain退出
func TestMain(m *testing.M) {
	log.Debug("Start Mail Sdk Test")
	ZKConfInit()
	exitVal := m.Run()
	os.Exit(exitVal)
	log.Debug("End Mail Sdk Test")
}

// 测试系统发送邮件
func TestManager_SystemSendMail(t *testing.T) {
	mailManager, celIds, err := newTest(true, "", false)
	award := make([][]int, 0)
	toRoles := make([]string, 0)
	celId := celIds[0]
	err = mailManager.SysSendMail(celId, timeutils.Now(), "111", "test_mail",
		"test_send_system_mail", "r01", toRoles, award, true)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
}

// 测试玩家发送邮件
func TestManager_PSendMail(t *testing.T) {
	mailManager, celIds, err := newTest(true, "", false)
	award := make([][]int, 0)
	celId := celIds[0]
	err = mailManager.PSendMail("r02", "1111", "1122", celId, timeutils.Now(),
		"1222", "test_mail", "test_send_personal_mail", "r01", award)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
}

// 测试查询已读邮件
func TestManager_QryRMail(t *testing.T) {
	toRole := "r01"
	mailManager, _, err := newTest(false, toRole, true)
	if err != nil {
		t.Errorf("Mock Send Mail Failed = %v", err)
	}
	result, err := mailManager.QryRMail(toRole, "test111", 10, timeutils.Now()+11)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) == 0 {
		t.Errorf("Query Aleady Read Mail Failed, Did Not Get Mail")
	}
}

// 测试根据邮件id列表和奖励id列表获取邮件, 以及根据邮件id查询单封 邮件
func TestManager_QryMailById(t *testing.T) {
	toRole := "r01"
	mailManager, celIds, err := newTest(false, toRole, true)
	if err != nil {
		t.Errorf("Mock Send Mail Failed = %v", err)
	}
	// 包含QryMailByCelId的测试
	result, err := mailManager.QryMailByCelId(toRole, "test111", celIds, false)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) == 0 {
		t.Errorf("Query Already Read Mail By CelId Failed, Did Not Get Mail")
	}

	// 用celId查出的结果测试QryMailById
	ids := make([]string, 1, 1)
	ids[0] = result[0].Id
	result, err = mailManager.QryMailByIds(toRole, "test111", ids)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) == 0 {
		t.Errorf("Query Already Read Mail By Id Array Failed, Did Not Get Mail")
	}

	// 测试查询QryMailById
	result, err = mailManager.QryMailById(toRole, "test111", ids[0])
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) == 0 {
		t.Errorf("Query Already Read Mail By Id Failed, Did Not Get Mail")
	}
}

// 测试查询所有未读未领取邮件
func TestManager_QryANoRMail(t *testing.T) {
	toRole := "r01"
	mailManager, _, err := newTest(false, toRole, false)
	if err != nil {
		t.Errorf("Mock Send Mail Failed = %v", err)
	}

	result, err := mailManager.QryANoRMail(toRole, "test111", 10)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) == 0 {
		t.Errorf("Query All ItemNo Read Mail Failed, Did Not Get Mail")
	}
}

// 测试查询所有未读未领取邮件数量
func TestManager_QryNrNoRMail(t *testing.T) {
	toRole := "r01"
	mailManager, _, _ := newTest(true, toRole, false)
	beforeNum, err := mailManager.QryNrNoRMail(toRole, "test111")
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	mailManager, _, err = newTest(false, toRole, false)
	if err != nil {
		t.Errorf("Mock Send Mail Failed = %v", err)
	}

	number, err := mailManager.QryNrNoRMail(toRole, "test111")
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if number != beforeNum+1 {
		t.Errorf("Query Number Of ItemNo Read And ItemNo Receive Failed")
	}
}

// 测试标记邮件已读(根据邮件id单封标记)
func TestManager_TagRMailSGL(t *testing.T) {
	toRole := "r01"
	mailManager, celIds, err := newTest(false, toRole, false)
	if err != nil {
		t.Errorf("Mock Send Mail Failed = %v", err)
	}
	// 查询删除邮件id
	result, err := mailManager.QryMailByCelId(toRole, "test111", celIds, false)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) == 0 {
		t.Errorf("Query Already Read Mail By CelId Failed, Did Not Get Mail")
	}
	mailId := result[0].Id

	err = mailManager.TagRMailSGL("test111", mailId)
	if err != nil {
		t.Errorf("Tag Mail To Read Failed")
	}

	// 校验邮件是否被标记为已读
	result, err = mailManager.QryMailByCelId(toRole, "test111", celIds, false)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) == 0 {
		t.Errorf("Query Already Read Mail By CelId Failed, Did Not Get Mail")
	}
	if result[0].IsRead != 1 {
		t.Errorf("Tag Sgl Mail Read Failed, Did Not Tag")
	}
}

// 测试标记邮件已读(根据邮件id列表多封标记)
func TestManager_TagRMailMul(t *testing.T) {
	toRole := "r01"
	mailManager, celIds, err := newTest(false, toRole, false)
	if err != nil {
		t.Errorf("Mock Send Mail Failed = %v", err)
	}
	// 查询删除邮件id
	result, err := mailManager.QryMailByCelId(toRole, "test111", celIds, false)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) == 0 {
		t.Errorf("Query Already Read Mail By CelId Failed, Did Not Get Mail")
	}
	ids := make([]string, 1, 1)
	ids[0] = result[0].Id

	err = mailManager.TagRMailMul("test111", ids)
	if err != nil {
		t.Errorf("Tag Mul Mail To Read Failed")
	}

	// 校验邮件是否被标记为已读
	result, err = mailManager.QryMailByCelId(toRole, "test111", celIds, false)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) == 0 {
		t.Errorf("Query Already Read Mail By CelId Failed, Did Not Get Mail")
	}
	if result[0].IsRead != 1 {
		t.Errorf("Tag Mul Mail Read Failed, Did Not Tag")
	}
}

// 测试标记邮件已读(标记该玩家所有邮件为已读)
func TestManager_TagRMailRole(t *testing.T) {
	toRole := "r01"
	mailManager, _, err := newTest(false, toRole, false)
	if err != nil {
		t.Errorf("Mock Send Mail Failed = %v", err)
	}

	err = mailManager.TagRMailRole("test111", toRole)
	if err != nil {
		t.Errorf("Tag Role Mail To Read Failed")
	}

	// 校验是否标记成功
	number, err := mailManager.QryNrNoRMail(toRole, "test111")
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if number != 0 {
		t.Errorf("Tag Role Mail To Read Failed")
	}
}

// 测试标记邮件已领取(根据邮件id单封标记)
func TestManager_TagGMailSGL(t *testing.T) {
	toRole := "r01"
	mailManager, celIds, err := newTest(false, toRole, false)
	if err != nil {
		t.Errorf("Mock Send Mail Failed = %v", err)
	}

	// 查询邮件id
	result, err := mailManager.QryMailByCelId(toRole, "test111", celIds, false)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) == 0 {
		t.Errorf("Query Mail By CelId Failed, Did Not Get Mail")
	}
	mailID := result[0].Id

	err = mailManager.TagGMailSGL("test111", mailID)
	if err != nil {
		t.Errorf("Tag Role Mail To Get Failed")
	}

	// 校验是否完成标记已领取
	result, err = mailManager.QryMailByCelId(toRole, "test111", celIds, false)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) == 0 {
		t.Errorf("Query Mail By CelId Failed, Did Not Get Mail")
	}
	if result[0].IsReceive != 1 {
		t.Errorf("Tag Role Mail To Get Failed")
	}
}

// 测试标记邮件已领取(多封)
func TestManager_TagGMailMul(t *testing.T) {
	toRole := "r01"
	mailManager, celIds, err := newTest(false, toRole, false)
	if err != nil {
		t.Errorf("Mock Send Mail Failed = %v", err)
	}

	// 查询邮件id
	result, err := mailManager.QryMailByCelId(toRole, "test111", celIds, false)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) == 0 {
		t.Errorf("Query Mail By CelId Failed, Did Not Get Mail")
	}
	mailID := result[0].Id
	ids := make([]string, 1, 1)
	ids[0] = mailID

	err = mailManager.TagRMailMul("test111", ids)
	if err != nil {
		t.Errorf("Tag Role Mail To Get Failed")
	}

	// 校验是否完成标记已领取
	result, err = mailManager.QryMailByCelId(toRole, "test111", celIds, false)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) == 0 {
		t.Errorf("Query Mail By CelId Failed, Did Not Get Mail")
	}
	if result[0].IsReceive != 1 {
		t.Errorf("Tag Role Mail To Get Failed")
	}
}

// 测试标记邮件已领取(标记玩家所有)
func TestManager_TagGMailRole(t *testing.T) {
	toRole := "r01"
	mailManager, _, err := newTest(false, toRole, false)
	if err != nil {
		t.Errorf("Mock Send Mail Failed = %v", err)
	}

	err = mailManager.TagGMailRole("test111", toRole)
	if err != nil {
		t.Errorf("Tag Role Mail To Get Failed")
	}

	noRAndGMail, err := mailManager.QryANoRMail(toRole, "test111", 10)
	for _, mail := range noRAndGMail {
		if mail.IsReceive != 1 {
			t.Errorf("Tag Role Mail To Get Failed, Tag Failed")
		}
	}
}

// 测试标记删除邮件(根据邮件id单封删除)
func TestManager_TagDMailSGL(t *testing.T) {
	toRole := "r01"
	mailManager, celIds, err := newTest(false, toRole, false)
	if err != nil {
		t.Errorf("Mock Send Mail Failed = %v", err)
	}

	// 查询删除邮件id
	result, err := mailManager.QryMailByCelId(toRole, "test111", celIds, false)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) == 0 {
		t.Errorf("Query Already Read Mail By CelId Failed, Did Not Get Mail")
	}
	mailId := result[0].Id

	err = mailManager.TagDMailSGL("test111", mailId)
	if err != nil {
		t.Errorf("Tag Remove One Mail Failed")
	}

	// 校验是否成功标记删除
	result, err = mailManager.QryMailByCelId(toRole, "test111", celIds, false)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Tag Remove Mul Mail Failed")
	}
}

// 测试标记删除邮件(根据邮件id列表多封删除)
func TestManager_TagDMailMul(t *testing.T) {
	toRole := "r01"
	mailManager, celIds, err := newTest(false, toRole, false)
	if err != nil {
		t.Errorf("Mock Send Mail Failed = %v", err)
	}

	// 查询除邮件id
	ids := make([]string, 1, 1)
	result, err := mailManager.QryMailByCelId(toRole, "test111", celIds, false)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) == 0 {
		t.Errorf("Query Already Read Mail By CelId Failed, Did Not Get Mail")
	}
	ids[0] = result[0].Id

	err = mailManager.TagDMailMul("test111", ids)
	if err != nil {
		t.Errorf("Tag Remove Mul Mail Failed")
	}

	// 校验是否成功标记删除
	result, err = mailManager.QryMailByCelId(toRole, "test111", celIds, false)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Tag Remove Mul Mail Failed")
	}
}

// 测试标记删除邮件(删除玩家所有邮件)
func TestManager_TagDMailRole(t *testing.T) {
	toRole := "r01"
	mailManager, _, err := newTest(false, toRole, false)
	if err != nil {
		t.Errorf("Mock Send Mail Failed = %v", err)
	}

	err = mailManager.TagDMailRole("test111", toRole)
	if err != nil {
		t.Errorf("Tag Remove Role Mail Failed")
	}

	// 校验是否全部标记删除
	result, err := mailManager.QryANoRMail(toRole, "test111", 500)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Tag Remove Role Mail Failed")
	}
}

// 测试修改邮件内容(先屏蔽邮件)
func TestManager_ChgMail(t *testing.T) {
	toRole := "r01"
	mailManager, celIds, err := newTest(false, toRole, false)
	if err != nil {
		t.Errorf("Mock Send Mail Failed = %v", err)
	}
	newTitle := "change test"

	err = mailManager.ChgMail("test111", celIds[0], newTitle, "")
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}

	// 校验是否修改成功
	result, err := mailManager.QryMailByCelId(toRole, "test111", celIds, false)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) == 0 {
		t.Errorf("Query Role Mail Failed")
	}
	if result[0].Title != newTitle {
		t.Errorf("Tag Change Role Mail Failed")
	}
}

// 测试屏蔽邮件以及解除屏蔽
func TestManager_SHLDMail(t *testing.T) {
	toRole := "r01"
	mailManager, celIds, err := newTest(false, toRole, true)
	if err != nil {
		t.Errorf("Mock Send Mail Failed = %v", err)
	}

	err = mailManager.SHLDMail("test111", celIds[0])
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}

	// 校验屏蔽邮件是否成功
	result, err := mailManager.QryMailByCelId(toRole, "test111", celIds, false)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) == 0 {
		t.Errorf("Query Role Mail Failed")
	}
	id := result[0].Id

	result, err = mailManager.QryMailById(toRole, "test111", id)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Sheld Mail Failed")
	}

	// 测试解除屏蔽
	err = mailManager.UnSHLDMail("test111", celIds[0])
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}

	// 校验解除屏蔽是否成功
	result, err = mailManager.QryMailById(toRole, "test111", id)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) == 0 {
		t.Errorf("UnSheld Mail Failed")
	}
}

// 测试物理删除邮件(按照邮件id指定删除)
func TestManager_DELMailSGL(t *testing.T) {
	toRole := "r01"
	mailManager, celIds, err := newTest(false, toRole, true)
	if err != nil {
		t.Errorf("Mock Send Mail Failed = %v", err)
	}
	result, err := mailManager.QryMailByCelId(toRole, "test111", celIds, false)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) == 0 {
		t.Errorf("Query Role Mail Failed")
	}
	id := result[0].Id

	err = mailManager.DELMailSGL("test111", id)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}

	// 校验邮件是否被删除
	result, err = mailManager.QryMailById(toRole, "test111", id)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Del Mail By Mail Id Failed")
	}
}

// 测试物理删除邮件(按照邮件id列表指定删除)
func TestManager_DELMailMul(t *testing.T) {
	toRole := "r01"
	mailManager, celIds, err := newTest(false, toRole, true)
	if err != nil {
		t.Errorf("Mock Send Mail Failed = %v", err)
	}
	result, err := mailManager.QryMailByCelId(toRole, "test111", celIds, false)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) == 0 {
		t.Errorf("Query Role Mail Failed")
	}
	id := result[0].Id
	ids := make([]string, 1, 1)
	ids[0] = id

	err = mailManager.DELMailMul("test111", ids)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}

	// 校验邮件是否被删除
	result, err = mailManager.QryMailById(toRole, "test111", id)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Del Mail By Mail Ids Failed")
	}
}

// 测试物理删除邮件(按照奖励id删除所有玩家该类邮件)
func TestManager_DELMailCelId(t *testing.T) {
	toRole := "r01"
	mailManager, celIds, err := newTest(false, toRole, true)
	if err != nil {
		t.Errorf("Mock Send Mail Failed = %v", err)
	}

	err = mailManager.DELMailCelId("test111", celIds[0])
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}

	// 校验邮件是否被删除
	result, err := mailManager.QryMailByCelId(toRole, "test111", celIds, false)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Del Mail By Mail Ids Failed")
	}
}

// 测试物理删除邮件(按照奖励id删除所有玩家该类邮件)
func TestManager_DELMailRoleSGL(t *testing.T) {
	toRole := "r01"
	mailManager, celIds, err := newTest(false, toRole, true)
	if err != nil {
		t.Errorf("Mock Send Mail Failed = %v", err)
	}

	err = mailManager.DELMailCelId("test111", celIds[0])
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}

	// 校验邮件是否被删除
	result, err := mailManager.QryMailByCelId(toRole, "test111", celIds, false)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Del ALL Player's Mail By celId Failed")
	}
}

// 测试物理删除邮件(删除个人所有邮件)
func TestManager_DELMailRoleAll(t *testing.T) {
	toRole := "r01"
	mailManager, celIds, err := newTest(false, toRole, true)
	if err != nil {
		t.Errorf("Mock Send Mail Failed = %v", err)
	}

	err = mailManager.DELMailRoleAll("test111", toRole)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}

	// 校验邮件是否被删除
	result, err := mailManager.QryMailByCelId(toRole, "test111", celIds, false)
	if err != nil {
		t.Errorf("MailAction error = %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Del ALL Player's Mail By celId Failed")
	}
}

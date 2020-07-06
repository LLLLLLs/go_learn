//@author: lls
//@time: 2020/05/26
//@desc:

package rdd

// 针对某通信产品，我们需要开发一个版本升级管理系统。该系统需要通过由Java开发的管理后台，由Telnet发起向前端基站设备的命令，
// 以获取基站设备的版本信息，并在后台比较与当前最新版本的差异，以确定执行什么样的命令对基站设备的软件文件进行操作。基站设备分为两种：
// 	- 主控板（Master Board）
// 	- 受控板（Slave Board）
//
// 基站设备允许执行的命令包括transfer、active、inactive等。
// 这些命令不仅受到设备类型的限制，还要受制于该设备究竟运行在什么样的终端。类型分为：
// 	- Shell
// 	- UShell
//
// 命令的约束如下：MB=Master Board;SB=Slave+Board;S=Shell;U=UShell
// Command	MB+S	MB+U	SB+S	SB+U
// transfer √		×		×		×
// active	×		√		×		×
// inactive ×		√		×		×
// put		×		×		√		×
// get 		×		×		√		×
// deleteF	×		×		×		√
//
// 通过登录可以连接到主控板的Shell终端，此时，若执行enterUShell命令则进入UShell终端，
// 执行enterSlaveBoard则进入受控板的Shell终端。在受控板同样可以执行enterUShell进入它的UShell终端。
// 系统还提供了对应的退出操作。整个操作引起的变迁如下图所示
//				login				enterSB
//			——————————————→		——————————————→
//	Initial					MB+S				SB+S
//			←——————————————	|  ↑←——————————————	|  ↑
//				logout		|  |	exitSB		|  |
//							|  |  				|  |
//							|  |  				|  |
//					  enterU|  |exitU	  enterU|  |exitU
//							|  |				|  |
//							|  |				|  |
//							|  |	enterSB		|  |
//							↓  |——————————————→ ↓  |
//							MB+U				SB+U
//								←——————————————
//									exitSB
// 执行升级的流程是在让基站设备处于失效状态下，获取基站设备的软件版本信息，
// 然后在后端基于最新版本进行比较。得到版本之间的差异后，
// 通过transfer命令传输新文件，put命令更新文件，deleteFiles命令删除多余的文件。
// 成功更新后，再激活基站设备。因此，一个典型的升级流程如下所示：
// - login
// - inactive
// - get
// - transfer
// - put
// - deleteFiles
// - active
// - logout

// Package main
/******************************************************************************
  作者: dmzn@163.com 2024-11-13 12:16:34
  描述: 通讯协议定义
******************************************************************************/
package main

const (
	NET_MODULE_DATA_LENGTH = 255 //与模块通信时数据区的最大长度

	// 通信命令码
	NET_MODULE_CMD_SET       = 0x01 //配置网络中的模块
	NET_MODULE_CMD_GET       = 0x02 //获取某个模块的配置
	NET_MODULE_CMD_RESET     = 0x03 //复位模块
	NET_MODULE_CMD_SEARCH    = 0x04 //搜索网络中的模块
	NET_MODULE_CMD_SET_BASE  = 0x05 //配置模块的端口基本设置
	NET_MODULE_CMD_SET_PORT1 = 0x06 //配置模块的端口1
	NET_MODULE_CMD_SET_PORT2 = 0x07 //配置模块的端口2

	NET_MODULE_ACK_SET       = 0x81 //回应配置命令码
	NET_MODULE_ACK_GET       = 0x82 //回应获取命令码
	NET_MODULE_ACK_RESEST    = 0x83 //回应复位命令码
	NET_MODULE_ACK_SEARCH    = 0x84 //回应所搜命令码
	NET_MODULE_ACK_SET_BASE  = 0x85 //配置模块的端口基本设置
	NET_MODULE_ACK_SET_PORT1 = 0x86 //配置模块的端口1
	NET_MODULE_ACK_SET_PORT2 = 0x87 //配置模块的端口2

	NET_MODULE_FLAG = `NET_MODULE_COMM\0` //用来标识通信_old
	CH9120_CFG_FLAG = "CH9120_CFG_FLAG"   //用来标识通信_new

	//模块标识
	NET_MODULE_TYPE_TCP_S = 0x00 //模块作为TCP SERVER
	NET_MODULE_TYPE_TCP_C = 0x01 //模块作为TCP CLIENT
	NET_MODULE_TYPE_UDP_S = 0x02 //模块作为UDP SERVER
	NET_MODULE_TYPE_UDP_C = 0x03 //模块作为UDP CLIENT

	//校验位标识
	NET_MODULE_VERIFY_ODD   = 0x00 //奇校验
	NET_MODULE_VERIFY_EVEN  = 0x01 //偶校验
	NET_MODULE_VERIFY_MARK  = 0x02 //mark校验
	NET_MODULE_VERIFY_SPACE = 0x03 //space校验
	NET_MODULE_VERIFY_NULL  = 0x04 //无校验

	// 是否清空串口缓冲区
	NET_MODULE_CLEAR_NO  = 0x00 //不清空缓冲区
	NET_MODULE_CLEAR_TCP = 0x01 //TCP连接时清空

	// 重连次数
	NET_MODULE_RELINK_INFINITY = 0x00 //重连次数无限
)

const (
	// 模块默认网络参数
	DEFAULT_MODULE_NAME = "NET MODULE"          //默认模块名
	DEFAULT_WORK_TYPE   = NET_MODULE_TYPE_UDP_C //默认工作模式
	DEFAULT_MODULE_IP   = "192.168.1.1"         //默认模块IP
	DEFAULT_MASK        = "255.255.255.0"       //默认子网掩码
	DEFAULT_GETWAY      = "192.168.1.2"         //默认网关
	DEFAULT_MODULE_PROT = 1124                  //默认模块端口
	DEFAULT_DEST_IP     = "192.168.1.3"         //默认目的IP
	DEFAULT_DEST_PROT   = 1124                  //默认目的端口

	// 模块默认串口参数
	DEFAULT_BOUND      = "115200"               //默认波特率
	DEFAULT_TIMEOUT    = 20                     //默认超时
	DEFAULT_DATA_BIT   = 0                      //默认数据位
	DEFAULT_STOP_BIT   = 0                      //默认停止位
	DEFAULT_VERIFY_BIT = NET_MODULE_VERIFY_NULL //默认校验位
	MODULE_CONFIG_FILE = `config.ini\0`         //配置文件名

	MODULE_DEFAULT_CONFIG_PORT = 50000 //默认的配置端口
	LOCAL_PORT                 = 60000 //默认本地配置端口
)

type (
	// unsigned char
	uchar = uint8
	// unsigned short
	ushort = uint16
	// unsigned long
	ulong = uint32

	// NET_COMM 网络通信结构体
	NET_COMM struct {
		flag    [16]uchar                     //通信标识，因为都是用广播方式进行通信的，所以这里加一个固定值
		cmd     uchar                         //命令头
		id      [6]uchar                      //标识，标识是与某个模块在通信，若与所有的模块通信，则值0XFFFFFF,目标模块mac地址+
		cfg_mac [6]uchar                      //配置软件端的MAC
		len     uchar                         //数据区长度
		data    [NET_MODULE_DATA_LENGTH]uchar //数据区缓冲区
	}

	// MODULE_CFG 模块的配置结构
	MODULE_CFG struct {
		module_name [21]uchar //模块在网络中的名字
		srv_type    uchar     //标识模块处于那模式(TCP/UDP server/client)
		src_ip      [4]uchar  //模块本身的IP地址
		mask        [4]uchar  //模块本身的子网掩码
		getway      [4]uchar  //模块对应的网关地址
		baud        ulong     //模块的串口波特率
		other       [3]uchar  //模块的串口其他配置(校验位、数据位、停止位)
		time_out    ulong     //模块的串口超时
		src_port    ushort    //模块源端口
		dest_ip     [4]uchar  //目的IP地址
		dest_port   ushort    //目的端口
		relink      uchar     //重连次数
		clear_buff  uchar     //是否清楚串口缓冲区
	}

	//  DEVICEHW_CONFIG
	DEVICEHW_CONFIG struct {
		DevType        uchar     /* 设备类型,具体见设备类型表 */
		AuxDevType     uchar     /* 设备子类型 */
		Index          uchar     /* 设备序号 */
		DevHardwareVer uchar     /* 设备硬件版本号 */
		DevSoftwareVer uchar     /* 设备软件版本号 */
		Modulename     [21]uchar /* 模块名*/
		DevMAC         [6]uchar  /* 模块网络MAC地址 */
		DevIP          [4]uchar  /* 模块IP地址*/
		DevGWIP        [4]uchar  /* 模块网关IP */
		DevIPMask      [4]uchar  /* 模块子网掩码 */
		DhcpEnable     uchar     /* DHCP 使能，是否启用DHCP,1:启用，0：不启用*/
		WebPort        uchar     /* WEB网页地址 */
		Username       [8]uchar  /* 用户名同模块名*/
		PassWordEn     uchar     /*密码使能 1：使能 0： 禁用*/
		zPassWord      [8]uchar  /* 密码*/
		UpdateFlag     uchar     /* 固件升级标志，1：升级 0：不升级*/
		ComcfgEn       uchar     /*串口协商进入配置模式使能，1：使能 0:不使能 */
		Reserved       [8]uchar  /* 保留*/
	}

	DeviceHWConfigS = DEVICEHW_CONFIG

	DEVICEPORT_CONFIG struct {
		Index           uchar     /* 端口序号 */
		PortEn          uchar     /* 端口启用标志 1：启用后 ；0：不启用 */
		NetMode         uchar     /* 网络工作模式: 0: TCP SERVER;1: TCP CLENT; 2: UDP SERVER 3：UDP CLIENT; */
		RandSportFlag   uchar     /* TCP 客户端模式下随即本地端口号，1：随机 0: 不随机*/
		NetPort         ushort    /* 网络通讯端口号 */
		DesIP           [4]uchar  /* 目的IP地址 */
		DesPort         ushort    /* 工作于TCP Server模式时，允许外部连接的端口号 */
		BaudRate        ulong     /* 串口波特率: 300---921600bps */
		DataSize        uchar     /* 串口数据位: 5---8位 */
		StopBits        uchar     /* 串口停止位: 1表示1个停止位; 2表示2个停止位 */
		Parity          uchar     /* 串口校验位: 0表示奇校验; 1表示偶校验; 2表示标志位(MARK,置1); 3表示空白位(SPACE,清0);  */
		PHYChangeHandle uchar     /* PHY断开，Socket动作，1：关闭Socket 2、不动作*/
		RxPktlength     ulong     /* 串口RX数据打包长度，最大1024 */
		RxPktTimeout    ulong     /* 串口RX数据打包转发的最大等待时间,单位为: 10ms,0则表示关闭超时功能 */
		ReConnectCnt    uchar     /* 工作于TCP CLIENT时，连接TCP SERVER的最大重试次数*/
		ResetCtrl       uchar     /* 串口复位操作: 0表示不清空串口数据缓冲区; 1表示连接时清空串口数据缓冲区 */
		DNSFlag         uchar     /* 域名功能启用标志，1：启用 2：不启用*/
		Domainname      [20]uchar /* 域名*/
		DNSHostIP       [4]uchar  /* DNS 主机*/
		DNSHostPort     ushort    /* DNS 端口*/
		Reserved        [8]uchar  /* 保留*/
	}

	DevicePortConfigS = DEVICEPORT_CONFIG

	NET_DEVICE_CONFIG struct {
		HWCfg   DeviceHWConfigS      /*从硬件处获取的配置信息*/
		PortCfg [2]DevicePortConfigS /*网络设备所包含的子设备的配置信息*/
	}
)

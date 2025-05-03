import (
	"context"
	"encoding/json"
)

var baseConf *BaseConfig
var initOnce sync.Once

type LimiterConfFunc func() string

var limiterConfFn LimiterConfFunc

type BaseConf struct {
	XMLName xml.Name `xml:"fsm_conf"`
	TransactionList []*TransactionConf `xml:"transaction"`
	TransactionMap map[string]*TransactionConf
	NodeMap map[string]*INode
	EHandlerMap map[string]*FSMLimiter
	GlobalLimiterMap map[stirng]GlobalFSMLimiter

	PushWorkerId string // 推单workerId
	BizPSM string // 业务PSM
	LocalAutoPushSwitch bool // 本地自动推单开关
	LocalMaxPushNum int // 本地推单携程池大小
	LocalMaxQueueNum int // 本地推单缓存队列大小
	GlobalAutoPushSwitch bool // 全局自动推单开关
	GlobalPushInterval int64 // 全局自动推单间隔
	GlobalMaxPushNum int // 全局自动推单携程池大小
	GlobalMaxQueueNum int // 全局自动推单缓存队列大小
	GlobalRecordScanStep int64 //全局扫表步长
	AlertSwitch bool // 告警开关
	AlertAppId string // 告警appId
	AlertAppSecret string // 告警appSecret
	AlertGroup string // 告警分组
	AlertViceGroupID string // 报警副通知群
	AlertTestGroupID string // 报警测试群
	StressCluster string // 压测集群名，为空代表不开启
	StressKey string // 压测key.StressCluster不为空生效
	StopCheckInc chan int
	StopGlobalC chan int
	StopLocalC chan int
	LarkbotHttpPort string
	SkipRetryIntervalCheck bool //是否跳过重试间隔校验
	DataReportingOpen bool //是否开启数据上报
	GloablLimitRedisPsm string //全局限流redis
}


package form

import "Dexun/utils"

type CloudEyeListInfo struct {
	CEId        int64      `json:"ce_id"`
	UUID        string     `json:"uuid"`
	TcName      string     `json:"tc_name"`
	KsMoney     int64      `json:"ks_money"`
	Content     string     `json:"content"`
	TaskType    utils.JSON `json:"task_type"`
	MonitorType utils.JSON `json:"monitor_type"`
	StartCount  int64      `json:"start_count"`
	OrderMoney  int64      `json:"order_money"`
	ZkMoney     int64      `json:"zk_money"`
	SellStatus  int        `json:"sell_status"`
}

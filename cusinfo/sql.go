package main

import (
	"database/sql"
	"errors"
	"log"

	"github.com/dilfish/tools"
)

type Service struct {
	db *sql.DB
}

func NewService(conf *tools.DBConfig) *Service {
	db, err := tools.InitDB(conf)
	if err != nil {
		log.Println("init db error:", conf, err)
		return nil
	}
	var s Service
	s.db = db
	return &s
}

// table info
// id
// order_time 下单时间
// pay_time 付款时间
// submit_time 提交时间
// ship_time 发货时间
// refund_time 退款时间
// print_time 面单打印时间
// pick_order_time 拣货单打印时间
// cus_acc 买家账号
// cus_name 买家姓名
// cus_email 买家Email
// recv_name 收货人姓名
// recv_company 收货人公司
// recv_tax_no 收货人税号
// recv_addr_no 收货人门牌号
// addr_detail 详细地址
// addr1 地址1
// addr2 地址2
// addr1plus2 地址1+地址2
// recv_city 收货人城市
// recv_state 收货人州/省
// postcode 邮编
// country 收货人国家
// country_cn 中文国家名
// country_code 国家二字码
// phone 收货人电话
// cellphone 收货人手机
// sku SKU
// prod_id 产品ID
// prod_name 产品名称
// prod_price 产品售价
// prod_num 产品数量
// prod_mod 产品规格
// pic_url 图片网址
// source_url 来源URL
// sale_url 销售链接
// multi_prod_name 多品名
// pay_method 付款方式
// currenty 币种缩写
// order_price 订单金额
// ship_fee 买家支付运费
// refund 退款金额
// est_profit 预估利润
// cost_profit_rate 成本利润率
// sale_profit_rate 销售利润率
// est_ship_fee 预估运费
// pkg_no 包裹号
// order_no 订单号
// tx_no 交易号
// order_status 订单状态
// platform 平台渠道
// shop_acc 店铺账号
// order_comment 订单备注
// pick_comment 拣货备注
// cus_service_comment 客服备注
// refund_reason 退款理由
// order_tag 订单标识
// order_label 订单标记
// appoint_ship 买家指定物流
// ship_method 物流方式
// ship_no 运单号
// ship_order 物流订单号
// weight 称重重量
// cn_clearance_name 中文报关名
// en_clearance_name 英文报关名
// clearance_price 申报单价
// clearance_weight 报关重量
// clearance_no 海关编码

func (s *Service) InsertDB(data []string) error {
	if len(data) != 67 {
		return errors.New("bad data, we need 67 rows")
	}
	ssql := "insert into `cus_info` (`id`,`order_time`,`pay_time`,"
	ssql = ssql + "`submit_time`,`ship_time`,`refund_time`,`print_time`,"
	ssql = ssql + "`pick_order_time`,`cus_acc`,`cus_name`,`cus_email`,"
	ssql = ssql + "`recv_name`,`recv_company`,`recv_tax_no`,"
	ssql = ssql + "`recv_addr_no`,`addr_detail`,`addr1`,`addr2`,"
	ssql = ssql + "`addr1plus2`,`recv_city`,`recv_state`,`postcode`,"
	ssql = ssql + "`country`,`country_cn`,`country_code`,`phone`,"
	ssql = ssql + "`cellphone`,`sku`,`prod_id`,`prod_name`,`prod_price`,"
	ssql = ssql + "`prod_num`,`prod_mod`,`pic_url`,`source_url`,"
	ssql = ssql + "`sale_url`,`multi_prod_name`,`pay_method`,`currency`,"
	ssql = ssql + "`order_price`,`ship_fee`,`refund`,`est_profit`,"
	ssql = ssql + "`cost_profit_rate`,`sale_profit_rate`,`est_ship_fee`,"
	ssql = ssql + "`pkg_no`,`order_no`,`tx_no`,`order_status`,`platform`,"
	ssql = ssql + "`shop_acc`,`order_comment`,`pick_comment`,"
	ssql = ssql + "`cus_service_comment`,`refund_reason`,`order_tag`,"
	ssql = ssql + "`order_label`,`appoint_ship`,`ship_method`,"
	ssql = ssql + "`ship_no`,`ship_order`,`weight`,`cn_clearance_name`,"
	ssql = ssql + "`en_clearance_name`,`clearance_price`,"
	ssql = ssql + "`clearance_weight`,`clearance_no`) value (0,"
	for i := 0; i < 66; i++ {
		ssql = ssql + "?,"
	}
	ssql = ssql + "?);"
	// log.Println("ssql is", ssql)
	// time == 0, string == 1, float == 2
	typeList := []int{
		0, 0, 0, 0, 0, 0, 0, // 7 time
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, // 11 string, until addr1plus2
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, // 11 string, until prod_name
		1, 1, 1, 1, 1, 1, 1, 1, 1, // 9 string, until order_price
		2, 2, 2, 2, 2, 2, 2, // 7 float, until pkg_no
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, // 11 string, until order label
		1, 1, 1, 1, 1, // 5 string, until weight
		2,       // 1 float until cn_clearance_name
		1, 1, 1, // 3 string until clearance_price
		2, // 1 float, until clearn_no
		1, // 1 string
	}
	args := []interface{}{}
	for i := 0; i < 67; i++ {
		if typeList[i] == TypeTime {
			tm, err := HandleTime(data[i])
			if err != nil {
				log.Println("handle time error:", data[i], err)
				return err
			}
			args = append(args, tm)
		}
		if typeList[i] == TypeString {
			str := HandleString(data[i])
			args = append(args, str)
		}
		if typeList[i] == TypeFloat {
			fl, err := HandleFloat(data[i])
			if err != nil {
				log.Println("handle float error:", data[i], err)
				return err
			}
			args = append(args, fl)
		}
	}
	_, err := s.db.Exec(ssql, args...)
	if err != nil {
		log.Println("db.insert error:", err)
		return err
	}
	return nil
}

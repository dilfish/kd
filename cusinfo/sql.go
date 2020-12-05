package main

import (
	"database/sql"
	"errors"
	"log"

	"github.com/dilfish/tools"
)

var ErrBadNameLen = errors.New("名字长度不合法")
var ErrBadLabel = errors.New("名字包含特殊字符")

type Service struct {
	db        *sql.DB
	tableName string
}

func NewService(conf *tools.DBConfig, t string) *Service {
	db, err := tools.InitDB(conf)
	if err != nil {
		log.Println("连接数据库错误:", conf, err)
		return nil
	}
	var s Service
	s.db = db
	s.tableName = t
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
		return errors.New("数据格式错误，我们要求数据必须是67列")
	}
	ssql := "insert into `"
	ssql = ssql + s.tableName
	ssql = ssql + "` (`id`,`order_time`,`pay_time`,"
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
	args := []interface{}{}
	for i := 0; i < 67; i++ {
		args = append(args, data[i])
	}
	_, err := s.db.Exec(ssql, args...)
	if err != nil {
		log.Println("插入数据错误:", err)
		return err
	}
	return nil
}

func (s *Service) CreateTable() error {
	start := "CREATE TABLE `"
	end := "` (`id` int unsigned not null AUTO_INCREMENT,  `order_time` varchar(32) not null default '',`pay_time` varchar(32) not null default '',`submit_time` varchar(32) not null default '', `ship_time` varchar(32) not null default '', `refund_time` varchar(32) not null default '', `print_time` varchar(32) not null default '', `pick_order_time` varchar(32) not null default '', `cus_acc` varchar(64) not null default '', `cus_name` varchar(64) not null default '',`cus_email` varchar(64) not null default '', `recv_name` varchar(64) not null default '', `recv_company` varchar(64) not null default '', `recv_tax_no` varchar(64) not null default '', `recv_addr_no` varchar(64) not null default '', `addr_detail` text,`addr1` text, `addr2` text, `addr1plus2` text, `recv_city` varchar(64) not null default '', `recv_state` varchar(64) not null default '', `postcode` varchar(64) not null default '', `country` varchar(64) not null default '',`country_cn` varchar(64) not null default '',`country_code` varchar(64) not null default '', `phone` varchar(64) not null default '',`cellphone` varchar(64) not null default '', `sku` text, `prod_id` text, `prod_name` text,  `prod_price` text, `prod_num` varchar(64) not null default '', `prod_mod` text, `pic_url` text, `source_url` text, `sale_url` text,`multi_prod_name` text,`pay_method` varchar(128) not null default '',`currency` varchar(128) not null default '', `order_price` float not null default 0, `ship_fee` varchar(32) not null default '', `refund` varchar(32) not null default '', `est_profit` varchar(32) not null default '', `cost_profit_rate` varchar(32) not null default '', `sale_profit_rate` varchar(32) not null default '', `est_ship_fee` varchar(32) not null default '', `pkg_no` varchar(64) not null default '', `order_no` varchar(64) not null default '', `tx_no` text, `order_status` varchar(64) not null default '', `platform` varchar(64) not null default '',`shop_acc` varchar(64) not null default '', `order_comment` text, `pick_comment` varchar(64) not null default '',`cus_service_comment` varchar(64) not null default '',`refund_reason` varchar(64) not null default '',`order_tag` varchar(64) not null default '', `order_label` varchar(64) not null default '',`appoint_ship` text, `ship_method` varchar(64) not null default '', `ship_no` varchar(64) not null default '', `ship_order` varchar(64) not null default '', `weight` varchar(32) not null default '', `cn_clearance_name` varchar(64) not null default '', `en_clearance_name` varchar(128) not null default '', `clearance_price` varchar(64) not null default '', `clearance_weight` varchar(32) not null default '', `clearance_no` varchar(64) not null default '', PRIMARY KEY (`id`)) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;"

	name := s.tableName
	if len(name) < 1 || len(name) > 50 {
		log.Println("名字长度不允许:", len(name))
		return ErrBadNameLen
	}
	for _, n := range name {
		if checkNumLetter(n) == false {
			log.Println("名字含有特殊字符:", name)
			return ErrBadLabel
		}
	}
	ssql := start + name + end
	_, err := s.db.Exec(ssql)
	if err != nil {
		log.Println("创建表失败:", name, err)
		return err
	}
	return nil
}

func checkNumLetter(n rune) bool {
	if n <= '9' && n >= '0' {
		return true
	}
	if n <= 'z' && n >= 'a' {
		return true
	}
	if n <= 'Z' && n >= 'A' {
		return true
	}
	return false
}

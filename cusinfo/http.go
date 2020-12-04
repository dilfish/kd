package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type SearchArgs struct {
	OrderNo string `json:"order_no"`
}

type SearchRet struct {
	Id                uint64  `json:"id"`
	OrderTime         string  `json:"order_time"`
	PayTime           string  `json:"pay_time"`
	SubmitTime        string  `json:"submit_time"`
	ShipTime          string  `json:"ship_time"`
	RefundTime        string  `json:"refund_time"`
	PrintTime         string  `json:"print_time"`
	PickOrderTime     string  `json:"pick_order_time"`
	CusAcc            string  `json:"cus_acc"`
	CusName           string  `json:"cus_name"`
	CusEmail          string  `json:"cus_email"`
	RecvName          string  `json:"recv_name"`
	RecvCompany       string  `json:"recv_company"`
	RecvTaxNo         string  `json:"recv_tax_no"`
	RecvAddrNo        string  `json:"recv_addr_no"`
	AddDetail         string  `json:"addr_detail"`
	Addr1             string  `json:"addr1"`
	Addr2             string  `json:"addr2"`
	Addr1plus2        string  `json:"addr1plus2"`
	RecvCity          string  `json:"recv_city"`
	RecvState         string  `json:"recv_state"`
	PostCode          string  `json:"postcode"`
	Country           string  `json:"country"`
	CountryCn         string  `json:"country_cn"`
	CountryCode       string  `json:"country_code"`
	Phone             string  `json:"phone"`
	Cellphone         string  `json:"cellphone"`
	Sku               string  `json:"sku"`
	ProdId            string  `json:"prod_id"`
	ProdName          string  `json:"prod_name"`
	ProdPrice         float32 `json:"prod_price"`
	PordNum           int     `json:"prod_num"`
	ProdMod           string  `json:"prod_mod"`
	PicUrl            string  `json:"pic_url"`
	SourceUrl         string  `json:"source_url"`
	SaleUrl           string  `json:"sale_url"`
	MultiProdName     string  `json:"multi_prod_name"`
	PayMethod         string  `json:"pay_method"`
	Currency          string  `json:"currency"`
	OrderPrice        string  `json:"order_price"`
	ShipFee           string  `json:"ship_fee"`
	Refund            string  `json:"refund"`
	EstProfit         string  `json:"est_profit"`
	CostProfitRate    string  `json:"cost_profit_rate"`
	SaleProfitRate    string  `json:"sale_profit_rate"`
	EstShipFee        string  `json:"est_ship_fee"`
	PkgNo             string  `json:"pkg_no"`
	OrderNo           string  `json:"order_no"`
	TxNo              string  `json:"tx_no"`
	OrderStatus       string  `json:"order_status"`
	Platform          string  `json:"platform"`
	ShopAcc           string  `json:"shop_acc"`
	OrderComment      string  `json:"order_comment"`
	PickComment       string  `json:"pick_comment"`
	CusServiceComment string  `json:"cus_service_comment"`
	RefundReason      string  `json:"refund_reason"`
	OrderTag          string  `json:"order_tag"`
	OrderLabel        string  `json:"order_label"`
	AppointShip       string  `json:"appoint_ship"`
	ShipMethod        string  `json:"ship_method"`
	ShipNo            string  `json:"ship_no"`
	ShipOrder         string  `json:"ship_order"`
	Weight            string  `json:"weight"`
	CnClearanceName   string  `json:"cn_clearance_name"`
	EnClearanceName   string  `json:"en_clearance_name"`
	ClearancePrice    string  `json:"clearance_price"`
	ClearanceWeight   string  `json:"clearance_weight"`
	ClearanceNo       string  `json:"clearance_no"`
}

func (s *Service) queryDB(args *SearchArgs) (*SearchRet, error) {
	var ret SearchRet
	log.Printf("args is %+v", args)
	ssql := "select `id`,`order_time`,`pay_time`,`submit_time`,`ship_time`,`refund_time`,`print_time`,`pick_order_time`,`cus_acc`,`cus_name`,`cus_email`,`recv_name`,`recv_company`,`recv_tax_no`,`recv_addr_no`,`addr_detail`,`addr1`,`addr2`,`addr1plus2`,`recv_city`,`recv_state`,`postcode`,`country`,`country_cn`,`country_code`,`phone`,`cellphone`,`sku`,`prod_id`,`prod_name`,`prod_price`,`prod_num`,`prod_mod`,`pic_url`,`source_url`,`sale_url`,`multi_prod_name`,`pay_method`,`currency`,`order_price`,`ship_fee`,`refund`,`est_profit`,`cost_profit_rate`,`sale_profit_rate`,`est_ship_fee`,`pkg_no`,`order_no`,`tx_no`,`order_status`,`platform`,`shop_acc`,`order_comment`,`pick_comment`,`cus_service_comment`,`refund_reason`,`order_tag`,`order_label`,`appoint_ship`,`ship_method`,`ship_no`,`ship_no`,`weight`,`cn_clearance_name`,`en_clearance_name`,`clearance_price`,`clearance_weight`,`clearance_no` from `cus_info` where"
	where := ""
	qs := []interface{}{}
	if args.OrderNo != "" {
		where = where + " `order_no`=?"
		qs = append(qs, args.OrderNo)
	}
	rows, err := s.db.Query(ssql+where, qs...)
	if err != nil {
		log.Println("query error:", where, err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&ret.Id, &ret.OrderTime, &ret.PayTime, &ret.SubmitTime, &ret.ShipTime, &ret.RefundTime, &ret.PrintTime, &ret.PickOrderTime, &ret.CusAcc, &ret.CusName, &ret.CusEmail, &ret.RecvName, &ret.RecvCompany, &ret.RecvTaxNo, &ret.RecvAddrNo, &ret.AddDetail, &ret.Addr1, &ret.Addr2, &ret.Addr1plus2, &ret.RecvCity, &ret.RecvState, &ret.PostCode, &ret.Country, &ret.CountryCn, &ret.CountryCode, &ret.Phone, &ret.Cellphone, &ret.Sku, &ret.ProdId, &ret.ProdName, &ret.ProdPrice, &ret.PordNum, &ret.ProdMod, &ret.PicUrl, &ret.SourceUrl, &ret.SaleUrl, &ret.MultiProdName, &ret.PayMethod, &ret.Currency, &ret.OrderPrice, &ret.ShipFee, &ret.Refund, &ret.EstProfit, &ret.CostProfitRate, &ret.SaleProfitRate, &ret.EstShipFee, &ret.PkgNo, &ret.OrderNo, &ret.TxNo, &ret.OrderStatus, &ret.Platform, &ret.ShopAcc, &ret.OrderComment, &ret.PickComment, &ret.CusServiceComment, &ret.RefundReason, &ret.OrderTag, &ret.OrderLabel, &ret.AppointShip, &ret.ShipMethod, &ret.ShipNo, &ret.ShipOrder, &ret.Weight, &ret.CnClearanceName, &ret.EnClearanceName, &ret.ClearancePrice, &ret.ClearanceWeight, &ret.ClearanceNo)
		if err != nil {
			log.Println("scan error:", err)
			return nil, err
		}
	}
	log.Println("ret is", ret)
	return &ret, nil
}

func (s *Service) hget(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		log.Println("method is not post:", req.Method)
		w.Write([]byte("bad request method"))
		return
	}
	mp := req.URL.Query()
	orderNo := mp["order_no"]
	if len(orderNo) != 1 {
		log.Println("bad args", orderNo)
		w.Write([]byte("bad args"))
		return
	}
	var args SearchArgs
	args.OrderNo = orderNo[0]
	ret, err := s.queryDB(&args)
	if err != nil {
		log.Println("query error", args, err)
		w.Write([]byte("query error"))
		return
	}
	bt, _ := json.Marshal(ret)
	w.Write(bt)
}

func (s *Service) h(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		log.Println("method is not post:", req.Method)
		w.Write([]byte("bad request method"))
		return
	}
	defer req.Body.Close()
	bt, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("read all error:", err)
		w.Write([]byte("bad read"))
		return
	}
	var args SearchArgs
	err = json.Unmarshal(bt, &args)
	if err != nil {
		log.Println("unmarshal error:", err)
		w.Write([]byte("unjosn error"))
		return
	}
	ret, err := s.queryDB(&args)
	if err != nil {
		log.Println("query error:", err)
		w.Write([]byte("query error"))
		return
	}
	bt, err = json.Marshal(ret)
	if err != nil {
		log.Println("marshal error:", err)
		w.Write([]byte("marshal error"))
		return
	}
	w.Write(bt)
	return
}

func (s *Service) Srv(port int) error {
	http.HandleFunc("/v1/search/info", s.h)
	http.HandleFunc("/v1/search/msg", s.hget)
	addr := ":" + strconv.FormatInt(int64(port), 10)
	log.Println("listen on:", addr)
	return http.ListenAndServe(addr, nil)
}

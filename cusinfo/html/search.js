$(document).ready(function () {
    $('#id-btn-search').click(function () {
        let d = {
            order_no: $('#id-order-no').val(),
            table_name: $('#id-table-name').val()
        };
        $.ajax({
            type: "POST",
            url: "/v1/search/info",
            data: JSON.stringify(d),
            success: function (data) {
                if (data.code != 0) {
                    alert('查询失败!');
                } else {
                    setTable(data);
                }
            },
            dataType: 'json'
        });
    });
});

function setTable(data) {
    let table = $('#id-table');
    for (var i = 0; i < data.list.length; i++) {
        let ht = "<tr>";
        ht = ht + "<td>" + data.list[i].id + "</td>";
        ht = ht + "<td>" + data.list[i].order_time + "</td>";
        ht = ht + "<td>" + data.list[i].pay_time + "</td>";
        ht = ht + "<td>" + data.list[i].submit_time + "</td>";
        ht = ht + "<td>" + data.list[i].ship_time + "</td>";
        ht = ht + "<td>" + data.list[i].refund_time + "</td>";
        ht = ht + "<td>" + data.list[i].print_time + "</td>";
        ht = ht + "<td>" + data.list[i].pick_order_time + "</td>";
        ht = ht + "<td>" + data.list[i].cus_acc + "</td>";
        ht = ht + "<td>" + data.list[i].cus_name + "</td>";
        ht = ht + "<td>" + data.list[i].cus_email + "</td>";
        ht = ht + "<td>" + data.list[i].recv_name + "</td>";
        ht = ht + "<td>" + data.list[i].recv_company + "</td>";
        ht = ht + "<td>" + data.list[i].recv_tax_no + "</td>";
        ht = ht + "<td>" + data.list[i].recv_addr_no + "</td>";
        ht = ht + "<td>" + data.list[i].addr_detail + "</td>";
        ht = ht + "<td>" + data.list[i].addr1 + "</td>";
        ht = ht + "<td>" + data.list[i].addr2 + "</td>";
        ht = ht + "<td>" + data.list[i].addr1plus2 + "</td>";
        ht = ht + "<td>" + data.list[i].recv_city + "</td>";
        ht = ht + "<td>" + data.list[i].recv_state + "</td>";
        ht = ht + "<td>" + data.list[i].postcode + "</td>";
        ht = ht + "<td>" + data.list[i].country + "</td>";
        ht = ht + "<td>" + data.list[i].country_cn + "</td>";
        ht = ht + "<td>" + data.list[i].country_code + "</td>";
        ht = ht + "<td>" + data.list[i].phone + "</td>";
        ht = ht + "<td>" + data.list[i].cellphone + "</td>";
        ht = ht + "<td>" + data.list[i].sku + "</td>";
        ht = ht + "<td>" + data.list[i].prod_id + "</td>";
        ht = ht + "<td>" + data.list[i].prod_name + "</td>";
        ht = ht + "<td>" + data.list[i].prod_price + "</td>";
        ht = ht + "<td>" + data.list[i].prod_num + "</td>";
        ht = ht + "<td>" + data.list[i].prod_mod + "</td>";
        ht = ht + "<td>" + data.list[i].pic_url + "</td>";
        ht = ht + "<td>" + data.list[i].source_url + "</td>";
        ht = ht + "<td>" + data.list[i].sale_url + "</td>";
        ht = ht + "<td>" + data.list[i].multi_prod_name + "</td>";
        ht = ht + "<td>" + data.list[i].pay_method + "</td>";
        ht = ht + "<td>" + data.list[i].currenty + "</td>";
        ht = ht + "<td>" + data.list[i].order_price + "</td>";
        ht = ht + "<td>" + data.list[i].ship_fee + "</td>";
        ht = ht + "<td>" + data.list[i].refund + "</td>";
        ht = ht + "<td>" + data.list[i].est_profit + "</td>";
        ht = ht + "<td>" + data.list[i].cost_profit_rate + "</td>";
        ht = ht + "<td>" + data.list[i].sale_profit_rate + "</td>";
        ht = ht + "<td>" + data.list[i].est_ship_fee + "</td>";
        ht = ht + "<td>" + data.list[i].pkg_no + "</td>";
        ht = ht + "<td>" + data.list[i].order_no + "</td>";
        ht = ht + "<td>" + data.list[i].tx_no + "</td>";
        ht = ht + "<td>" + data.list[i].order_status + "</td>";
        ht = ht + "<td>" + data.list[i].platform + "</td>";
        ht = ht + "<td>" + data.list[i].shop_acc + "</td>";
        ht = ht + "<td>" + data.list[i].order_comment + "</td>";
        ht = ht + "<td>" + data.list[i].pick_comment + "</td>";
        ht = ht + "<td>" + data.list[i].cus_service_comment + "</td>";
        ht = ht + "<td>" + data.list[i].refund_reason + "</td>";
        ht = ht + "<td>" + data.list[i].order_tag + "</td>";
        ht = ht + "<td>" + data.list[i].order_label + "</td>";
        ht = ht + "<td>" + data.list[i].appoint_ship + "</td>";
        ht = ht + "<td>" + data.list[i].ship_method + "</td>";
        ht = ht + "<td>" + data.list[i].ship_no + "</td>";
        ht = ht + "<td>" + data.list[i].ship_order + "</td>";
        ht = ht + "<td>" + data.list[i].weight + "</td>";
        ht = ht + "<td>" + data.list[i].cn_clearance_name + "</td>";
        ht = ht + "<td>" + data.list[i].en_clearance_name + "</td>";
        ht = ht + "<td>" + data.list[i].clearance_price + "</td>";
        ht = ht + "<td>" + data.list[i].clearance_weight + "</td>";
        ht = ht + "<td>" + data.list[i].clearance_no + "</td>";
        ht = ht + "</tr>";
        table.append(ht);
    }
}

$(document).ready(function () {
    load_table_name();
    $('#id-btn-search').click(function () {
        let d = {
            order_no: $('#id-order-no').val(),
            table_name: $('#id-select').val()
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

function set_table_name(data) {
    var list;
    for (var i = 0; i < data.list.length; i++) {
        list = list + "<option value=" + data.list[i] + ">";
        list = list + data.list[i] + "</option>";
    }
    $("#id-select").html(list);
    return;
}

function load_table_name() {
    $.ajax({
        type: "GET",
        url: "/v1/table/list",
        success: function (data) {
            if (data.code != 0) {
                alert('获取列表失败!');
            } else {
                set_table_name(data);
            }
        },
        dataType: 'json'
    });
}

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
        /*
      
        */
        ht = ht + "</tr>";
        table.append(ht);
    }
}

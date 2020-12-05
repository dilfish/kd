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
        /*
      
        */
        ht = ht + "</tr>";
        table.append(ht);
    }
}

$(document).ready(function(){
    $('#id-btn-search').click(function() {
	let d = {
		order_no : $('#id-order-no').val()
	};
        $.ajax({
            type: "POST",
            url: "/v1/search/info",
            data: JSON.stringify(d),
            success: function(data) {
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
	$('#id-email').html(data.cus_email);
	$('#id-recv-name').html(data.recv_name);
}

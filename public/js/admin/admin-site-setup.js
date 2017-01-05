
$(function () {
    var status = $("#status").val();
    var reg = $("#reg").val();

    if (reg == "3001") {
        $("#site-reg-1").attr("checked", 'checked')
    }else if (reg == "3002") {
        $("#site-reg-2").attr("checked", 'checked')
    }

    if (status == "2001") {
        $("#site-status-1").attr("checked", 'checked')
    } else if (status == "2002") {
        $("#site-status-2").attr("checked", 'checked')
    } else if (status == "2003") {
        $("#site-status-3").attr("checked", 'checked')
    }

});



var paraseData = function () {

    var array = {
        "title": $("#site-title").val(),
        "subtitle": $("#site-subtitle").val(),
        "url": $("#site-url").val(),
        "seo": $("#site-seo").val(),
        "reg": $("input[name='site-reg']:checked").val(),
        "foot": $("#site-foot").val(),
        "statistics": $("#site-statistics").val(),
        "status": $("input[name='site-status']:checked").val()
    };

    return array
};

var update = function () {

    var json = paraseData()
    console.log(json);
    $.post("/admin/site", json, function (data) {
        alert(data);
    });

};
var Install = {
    step: 0,
    totalStep: 2,
    init: function () {

    },
    submit: function () {
        switch (Install.step) {
            case 0:
                Install.addDB();
                break;
            case 1:
                Install.addAdmin();
                break;
            case 2:
                Install.finish();
                break;
        }
    },
    addDB: function () {
        var data = {
            "info.Db_host": $("#db-host").val(),
            "info.Db_port": $("#db-port").val(),
            "info.Db_user": $("#db-user").val(),
            "info.Db_pass": $("#db-pass").val(),
            "info.Db_name": $("#db-name").val(),
            "info.Db_prefix": $("#db-prefix").val()
        }
        if (data["info.Db_host"] == "") {
            alertify.alert("Error", "数据库地址不能为空", null);
        }
        if (data["info.Db_port"] == "") {
            alertify.alert("Error", "数据库端口不能为空", null);
        }
        if (data["info.Db_user"] == "") {
            alertify.alert("Error", "数据库用户不能为空", null);
        }
        if (data["info.Db_name"] == "") {
            alertify.alert("Error", "数据库名称不能为空", null);
        }

        var $button = $("#next-button");
        $button.attr("disabled", true);
        $button.text("正在验证");
        $.post("/install/adddb", data, function (data) {
            if (data.Success) {
                $button.attr("disabled", false).text("下一步");
                Install.renderNode(Install.step + 1);
                if (data.Data > 0) {
                    $("#msg").text("有超级管理员存在，这一步可以跳过");
                    $("#skip").show();
                }
            } else {
                alertify.alert("Error", data.Msg, null);
                $button.attr("disabled", false).text("下一步");
            }
        })

    },
    addAdmin: function () {
        var data = {
            "info.Admin_user": $("#admin-user").val(),
            "info.Admin_pass": $("#admin-pass").val(),
            "info.Admin_email": $("#admin-email").val(),
        }
        if (data["info.Admin_user"] == "") {
            alertify.alert("Error", "用户名不能为空", null);
        }
        if (data["info.Admin_pass"] == "") {
            alertify.alert("Error", "密码不能为空", null);
        }
        if (data["info.Admin_email"] == "") {
            alertify.alert("Error", "邮箱不能为空", null);
        }

        var $button = $("#next-button");
        $button.attr("disabled", true);
        $button.text("正在验证");
        $.post("/install/addadmin", data, function (data) {
            if (data.Success) {
                // $button.attr("disabled", false).text("下一步");
                // Install.renderNode(Install.step + 1);
                Install.finish();
            } else {
                alertify.alert("Error", data.Msg, null);
                $button.attr("disabled", false).text("下一步");
            }
        })
    },
    finish: function(){
        location.href="/main/";
    },
    skip: function(){
        Install.renderNode(Install.step + 1);
    },
    renderNode: function (stepNum) {
        if (stepNum == Install.totalStep){
            Install.finish();
            return;
        }
        var $dom = $("#route");
        var lenth = $dom.find(".pointer").lenth;
        var $doms = $dom.find(".pointer");
        var temp = null;
        $("#msg").text("");
        $("#skip").hide();
        Install.step = stepNum;
        $doms.each(function (index, item) {
            temp = $(item);
            if (index < stepNum) {
                temp.removeClass("ing");
                temp.addClass("finish");
                $("#step-" + index).hide();
            }
            if (index == stepNum) {
                $("#step-" + stepNum).show();
                temp.addClass("ing");
            }
            if (index > stepNum) {
                $("#step-" + index).hide();
                temp.removeClass("ing");
                temp.removeClass("finish");
            }
        })
    },
    getData: function () {

    }
}
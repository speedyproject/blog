var CategoryArea = {
    save: function () {
        var data = {
            ca_name: $("#ca-name").val(),
            ca_ident: $("#ca-ident").val(),
            p_ca: $("#p-ca").val(),
            ca_desc: $("#ca-desc").val()
        }
        $.post("/admin/category/add", data, function (data) {
            if(data.Success){
                alert("添加成功");
            }else{
                alert(data.Msg);
            }
        })
    }
}
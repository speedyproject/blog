var CategoryArea = {
    save: function () {
        var data = {
            ca_id: $("#ca-id").val(),
            ca_name: $("#ca-name").val(),
            ca_ident: $("#ca-ident").val(),
            ca_p: $("#p-ca").val(),
            ca_desc: $("#ca-desc").val()
        }
        $.post("/admin/category/add", data, function (data) {
            if(data.Success){
                alertify.success("添加成功");
                setInterval(function(){location.href="/admin/category";},300);
            }else{
                alertify.alert("Error",data.Msg);
            }
        })
    },
    del: function(id){
        alertify.confirm("确认删除这个分类？",function(){
            $.post("/admin/category/del",{"id":id},function(data){
                if(data.Success){
                    alertify.success("删除成功");
                    $("#category-col-"+id).remove();
                }else{
                    alertify.alert("Error",data.Msg);
                }
            })})
    }

}
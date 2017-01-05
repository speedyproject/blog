var adminTag = {
    init: function(){
        adminTag.bindTag();
    },
    bindTag: function(){
        $("#tag-area").on("click",".item",function(event){
            var $dom = $(event.target);
            $dom.toggleClass("active");
            adminTag.edit($dom);
        })
    },
    edit: function($dom){
        $("#edit-tag-id").val($dom.attr("data-id"));
        $("#tag-name").val($dom.text());
        $("#tag-ident").val($dom.attr("data-ident"));
    },
    editTag: function(){
        var data = {
            tagID: $("#edit-tag-id").val(),
            tagName: $("#tag-name").val(),
            tagIdent: $("#tag-ident").val()
        };
        $.post("/admin/tag/index",data,function(data){
            if(data.Success){
                location.reload();
            }else{
                alert(data.Msg);
            }
        });
    },
    delete: function(){
        var idStr = $(".item.active").map(function(){return $(this).attr("data-id")}).get().join(",");
        $.post("/admin/tag/del",{ids:idStr},function(data){
            if(data.Success){
                location.reload();
            }else{
                alert(data.Msg);
            }
        })
    }
}

$(function(){
    adminTag.init();
});

$(function () {
    if ($("#editormd").length > 0) {
        Post.init();
    } else {
        Manager.init();
    }
})

var Post = {
    editor: null,
    ms: null,
    init: function () {
    Post.editor = editormd({
            id: "editormd",
            height: 640,
            syncScrolling: "single",
            saveHTMLToTextarea: true,
            path: "/public/third/mdeditor/lib/",
            watch:false,
            onload: function(){
                this.setMarkdown($("#blog-src").val());
            }
        });

        Post.ms = $('#magicsuggest').magicSuggest({
            placeholder:"请输入标签"
        });
        var tags = $("#blog-tags").text();
        if(tags.length > 0){
            Post.ms.setSelection(JSON.parse(tags));
        }
    },
    submit: function () {
        var tags = Post.ms.getSelection(),
            length = tags.length,
            allTags = Post.ms.getData().map(function(item){return item.name})
            tagIds = [],
            newTag = [];
        for(var i = 0;i<length;i++){
            var item = tags[i];
            if(allTags.indexOf(item.name) >= 0){
                tagIds.push(item.id);
            }else{
                newTag.push(item.name);
            }
        }
        var data = {
            "data.Id":$("#blog-id").val(),
            "data.Title": $("#blog-title").val(),
            "data.Ident": $("#blog-ident").val(),
            "data.ContentMD": Post.editor.getMarkdown(),
            "data.ContentHTML": Post.editor.getHTML(),
            "data.Type": 0,
            "data.Summary": $("#blog-summary").val(),
            "data.Category": $("input[name=category]:checked").val(),
            "data.Tag": tagIds.join(","),
            "data.NewTag": newTag.join(","),
            "data.Createtime":$("#blog-createtime").val()
        }
        if(data["data.Title"] == ""){
            alertify.alert("Error","标题不能为空");
        }
        if(data["data.ContentHTML"] == ""){
            alertify.alert("Error","内容不能为空");
        }
        $.post("/admin/post/index", data, function (data) {
            if(data.Success){
                alertify.success("发布成功");
                location.href="/admin/manage-post";
            }else{
                alertify.alert("Error",data.Msg,null);
            }
        })
    },
}

var Manager = {
    selectCount: 0,
    totalCount: 0,
    init: function(){
        Manager.totalCount = $(".select-for-js").length;
        Manager.bindClick();
    },
    selectAll: function(event){
        var $dom = $(event.target);
        if($dom.prop("checked")){
            $("#post-list-table").find(".select-for-js").prop("checked",true);
            Manager.selectCount = Manager.totalCount;
        }else{
            $("#post-list-table").find(".select-for-js").prop("checked",false);
            Manager.selectCount = 0;
        }
    },
    bindClick: function(){
        $("#post-list-table").on("click",".select-for-js",function(event){
            var $dom = $(event.target);
            if($dom.prop("checked")){
                Manager.addCount(true);
            }else{
                Manager.addCount(false);
            }
        })
    },
    addCount: function(add){
        if(add){
            Manager.selectCount++;
        }else{
            Manager.selectCount--;
        }
        if(Manager.selectCount == Manager.totalCount){
            $("#select-all").prop("checked",true);
        }else{
            $("#select-all").prop("checked",false);
        }

    },
    delete: function(){
        var idarr = $("#post-list-table").find(".select-for-js:checked").map(function(){return $(this).attr("data-id")}).get();
        $.post("/admin/manage-post/delete",{ids:idarr.join(",")},function(data){
            if(data.Success){
                for(var i in idarr){
                    $("#blog-"+idarr[i]).remove();
                }
                alertify.success('删除成功！');
            }else{
                alertify.alert("注意",data.Msg);
            }
        })
    }
}

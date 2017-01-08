$(function () {
    if ($("#editormd").length > 0) {
        Post.init();
    } else {
        Manager.init();
    }
})

var Post = {
    editor: null,
    init: function () {
        Post.editor = editormd({
            id: "editormd",
            height: 640,
            syncScrolling: "single",
            saveHTMLToTextarea: true,
            path: "/public/mdeditor/lib/",
            watch:false
        });
    },
    submit: function () {
        var data = {
            "data.Title": $("#blog-title").val(),
            "data.ContentMD": Post.editor.getMarkdown(),
            "data.ContentHTML": Post.editor.getHTML(),
            "data.Type": 0,
        }
        $.post("/admin/post/index", data, function (d) {
            console.log(d);
        })
    }
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
            }else{
                alert(data.Msg);
            }
        })
    }
}
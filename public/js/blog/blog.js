var Comment = {
    post: function(){
        var content = $("#comment-content").val(),
            name = "L",
            blogID = $("#blog-id").val()
        $.post("/ajax/new-comment",{
            "content": content,
            "name": name,
            "blogid": blogID
        },function(data){
            if (data.Success){
                alert("ok")
            }else{
                alert(data.Msg)
            }

        })    

    }
}
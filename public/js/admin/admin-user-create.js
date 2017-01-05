var submitit = function(){
    var param = {
        username:$("#input-username").val(),
        nickname:$("#input-nickname").val(),
        password:$("#input-password").val(),
        email:$("#input-email").val(),
        group:$("#input-group").val(),
    }
    $.post("/admin/user/create",param,function(data){
        if(data.Success){
            location.href="/admin/user"
        }else{
            console.log(data.Msg);
        }
    })
    return false;
}

var updateit = function(){
    var param = {
        username:$("#input-username").val(),
        nickname:$("#input-nickname").val(),
        password:$("#input-password").val(),
        email:$("#input-email").val(),
        group:$("#input-group").val(),
        id:$("#user-id").val()
    }
    $.post("/admin/user/edit",param,function(data){
        if(data.Success){
            cosole.log("add successful");
        }else{
            console.log(data.Msg);
        }
    })
    return false;
}

var deletethem = function(){
    var ids = $(".mul-check:checked").map(function(){return $(this).val()}).get().join(",");
    $.post("/admin/user/del",{ids:ids},function(data){
        location.reload();
    })
}
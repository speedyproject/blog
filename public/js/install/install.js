var Install = {
    data: {},
    submit: function(){
        Install.checkParams();
        Install.getData();
        $.post("/install/index",Install.data,function(data){
            if(data.Success){

            }
            console.log(data);
        })    
    },
    checkParams: function(){

    },
    getData: function(){
        Install.data={
            "info.Db_host": $("#db-host").val(),
            "info.Db_port": $("#db-port").val(),
            "info.Db_user": $("#db-user").val(),
            "info.Db_pass": $("#db-pass").val(),
            "info.Db_name": $("#db-name").val(),
            "info.Admin_user": $("#admin-user").val(),
            "info.Admin_pass": $("#admin-pass").val(),
            "info.Admin_email": $("#admin-email").val(),
        }
    }
}
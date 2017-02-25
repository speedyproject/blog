var Common ={
    init: function(){
        Common.initSearchKey();
    },
    initSearchKey: function(){
        $("#search-key").keypress(function(e){
            var k = e.which;
            if(k == 13){
                Common.toSearch();
            }
        })
    },
    toSearch: function(){
        location.href="/search?q="+$("#search-key").val();
    }
}

$(function(){
    Common.init();
})
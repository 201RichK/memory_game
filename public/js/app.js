/*
var c = [37];
for (var i=1;i<37;i++) {
    (function(i){
        c[i] = document.getElementById('c'+i);
        c[i].addEventListener("click",function(){
           var  url = "/data";
            $.ajax({
                url: url,
                type: 'post',
                dataType: 'html',
                data : { ajax_post_data: i},
                success : function(data) {
                    console.log("success" + i)
                },
            });
        });
    }(i));
}
*/
/*
function activeCard(id) {
    //document.getElementById("c"+id).addClass("active")
    //c[id].classList.add("active");

}
*/

$(document).ready(function(){
    $("#div1 h3").click(function(){
        $(this).next().toggle(500).siblings("ul").hide(500)
    })
})
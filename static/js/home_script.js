$(".nav1 a").on("click", function () {
    var position = $(this).parent().position();
    var width = $(this).parent().width();
    $(".nav1 .slide1").css({opacity: 1, left: +position.left, width: width});
});

$(".nav1 a").on("mouseover", function () {
    var position = $(this).parent().position();
    var width = $(this).parent().width();
    $(".nav1 .slide2").css({
        opacity: 1,
        left: +position.left,
        width: width
    });
});

$(".nav1 a").on("mouseout", function () {
    $(".nav1 .slide2").css({opacity: 0});
});

var currentWidth = $(".nav1").find("li:nth-of-type(1) a").parent("li").width();
var current = $("li:nth-of-type(1) a").position();
$(".nav1 .slide1").css({left:+current.left, width:currentWidth});
// $(".nav .slide1").css({left: +current.left, width: currentWidth});
 $(function () {
    var thisTime;
    $('.nav2-ul li').mouseleave(function (even) {
        thisTime = setTimeout(thisMouseOut, 1000);
    })

    $('.nav2-ul li').mouseenter(function () {
        clearTimeout(thisTime);
        var thisUB = $('.nav2-ul li').index($(this));
        if ($.trim($('.nav2-slide-o').eq(thisUB).html()) != "") {
            $('.nav2-slide').addClass('hover');
            $('.nav2-slide-o').hide();
            $('.nav2-slide-o').eq(thisUB).show();
        } else {
            $('.nav2-slide').removeClass('hover');
        }

    });

    function thisMouseOut() {
        $('.nav2-slide').removeClass('hover');
    }

    $('.nav2-slide').mouseenter(function () {
        clearTimeout(thisTime);
        $('.nav2-slide').addClass('hover');
    });
    $('.nav2-slide').mouseleave(function () {
        $('.nav2-slide').removeClass('hover');
    })

})
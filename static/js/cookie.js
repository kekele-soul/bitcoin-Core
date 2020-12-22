var input = document.getElementsByTagName("input");
var id = document.getElementById("btn");

if (getCookie("username")) {
    input[0].value = getCookie("username");
    input[1].value = getCookie("password");
}

id.onclick = function () {
    var username = input[0].value;
    var password = input[1].value;

    setCookie("username", username, 1);
    setCookie("password", password, 1);
};


/*
  cookie机制：
1、  客户端发送一个请求到服务器
2、  服务器发送一个HttpResponse响应到客户端，其中包含Set-Cookie的头部
3、  客户端保存cookie，之后向服务器发送请求时，HttpRequest请求中会包含一个Cookie的头部
4、  服务器返回响应数据
 */

/**
 * 该方法用于生成cookie实例
 * @param cname 要传入的数据类型
 * @param cvalue 要传入的参数
 * @param day 设置该方法生命周期
 */
function setCookie(name, value, day) {
    var date = new Date();
    date.setDate(date.getDate() + day);
    // var expires = "expires=" + date.toDateString();
    // document.cookie = cname + "=" + cvalue + ";" + expires;
    document.cookie = name + "=" + value + ";expires=" + date;
}

/**
 * 该方法用于获取cookie实例
 * @param cname 要持久保存的数据类型
 * @returns {string} 返回获取的长度，如未获取到则返回空
 */
function getCookie(name) {
    var str = document.cookie;
    var arr = str.split("; ");

    for (var i = 0; i < arr.length; i++) {
        var arr1 = arr[i].split("=");
        if (arr1[0] = name) {
            return arr1[1];
        }
    }
    // var name = cname + "=";
    // var can = document.cookie.split(';');
    // for (var i = 0; i < can.length; i++) {
    //     var c = can[i].trim();
    //     if (c.indexOf(name) == 0) {
    //         return c.substring(name.length, c.length);
    //     }
    //     return "";
    // }
}

/**
 * 该方法用于停止cookie的生命周期
 */
function removecookie(name) {
    setCookie(name, 1, -1)

}

/**
 * 该方法用于检索cookie的值
 */
// function checkCookie() {
//     var user = getCookie("username");
//     if (user != "") {
//         alert("检索cookie成功" + user)
//     } else {
//         user = prompt("请输入你的名字：", "");
//         if (user != "" && user != null) {
//             setCookie("username", user, 7)
//         }
//     }
// }
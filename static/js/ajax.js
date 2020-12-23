//js实现ajax
/**
 * 该方法用于创建ajax的实例
 * @param type 请求类型
 * @param url 要请求信息的路径
 * @param success 请求成功
 * @param error 错误
 * @constructor
 */
function Ajax(type, url, success, error) {
    //判断浏览器版本
    var xhr;
    if (window.XMLHttpRequest) {
        //
        xhr = new XMLHttpRequest();
    } else {
        xhr = new ActiveXObject("Microsoft.XMLHTTP");
    }
    //判断请求方式
    if (type === "GET") {
        xhr.open(type, url, true);
        xhr.send()
    } else {
        xhr.open(type, url, true);
        xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        xhr.send()
    }

    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                success(xhr);
            } else {
                error(xhr)
            }
        }
    }
}
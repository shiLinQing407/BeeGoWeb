var $c = {};


$c.ajax = function (url, params, callback, method, type = "json") {
    if (!method) {
        method = 'POST'
    }
    $.ajax({
        url: url,
        data: params,
        method: method,
        // type: type,
        success: function (ret) {
            if (typeof callback == 'function') {
                callback(ret)
            }
        },
        error: function (ret) {
            if (typeof callback == 'function') {
                callback("{'code':-1,'message':'请求失败'}")
            }
        }
    })
}

$c.isEmpty = function (data) {
    if (data && typeof data != "undefined" && data != null && data != '' && ((Array.isArray(data) && data.length > 0) || !Array.isArray(data)) && ((typeof data == "object" && JSON.stringify(data) != "{}") || typeof data != "object")) {
        return false;
    } else {
        return true;
    }
}

/**
 *  判断值是否在数组中
 * 存在则返回键值  不存在返回null
 * @param search
 * @param array
 * @returns {boolean}
 */
$c.inArray = function (search, array) {
    for (let key in array) {
        if (array[key] == search) {
            return true;
        }
    }
    return false;
}

/**
 * 返回输入数组中某个单一列的值。
 * @param array
 * @param column_key   需要返回值的列。 可以是索引数组的列的整数索引，或者是关联数组的列的字符串键值。
 * 该参数也可以是 NULL，此时将返回整个数组（配合 index_key 参数来重置数组键的时候，非常有用）。
 * @param index_key   可选。用作返回数组的索引/键的列。
 * @returns {array}  返回数组，此数组的值为输入数组中某个单一列的值。
 */
$c.arrayColumn = function (array, column_key, index_key = null) {
    let data = [];
    for (let i = 0; i < array.length; i++) {
        let item;
        if (!empty(column_key)) {
            item = array[i][column_key];
        } else {
            item = array[i]
        }
        if (!empty(index_key)) {
            data[index_key] = item;
        } else {
            data[i] = item;
        }
    }
    return data;
}

/**
 * 搜索数组中给定的值并返回键名。
 * @param search
 * @param array
 * @returns {boolean}
 */
$c.arraySeach = function (search, array) {
    for (let key in array) {
        if (array[key] == search) {
            return key;
        }
    }
    return false;
}

/**
 * 获取当前时间
 */
$c.setDate = function(){
    var date = new Date();
    var year = date.getFullYear();
    nowDate1 = year + "-" + addZero((date.getMonth() + 1)) + "-" + addZero(date.getDate()) + "  ";
    nowDate1 += $c.addZero(date.getHours()) + ":" + addZero(date.getMinutes()) + ":" + addZero(date.getSeconds());
    return nowDate1;
}

/**
 * 年月日是分秒为10以下的数字则添加0字符串
 * @param time
 * @returns {number | *}
 */
$c.addZero = function(time) {
    var i = parseInt(time);
    if (i / 10 < 1) {
        i = "0" + i;
    }
    return i;
}

$c.error = function (vue, msg, duration = 2, onClose = '') {
    vue.$Message.error(msg, duration, onClose);
}

$c.success = function (vue, msg, duration = 2, onClose = '') {
    vue.$Message.success(msg, duration, onClose);
}

$c.warning = function (vue, msg, duration = 2, onClose = '') {
    vue.$Message.warning(msg, duration, onClose);
}

$c.loading = function (vue, duration = 0) {
    vue.$Message.loading('正在加载中...', duration);
}
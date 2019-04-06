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

$c.getGridHeight = function() {
    var win_height = $(window).height();
    main_header_height = $("#js_main_header").outerHeight(true);
    if (!main_header_height)
        main_header_height = 0;
    var height = win_height - main_header_height - 10;
    return height;
}

$c.parseToDate = function(value) {
    if (value == null || value == '') {
        return undefined;
    }
    var dt;
    if (value instanceof Date) {
        dt = value;
    } else {
        if (!isNaN(value)) {
            dt = new Date(parseInt(value) * 1000);
            //dt = new Date(value);
        } else if (value.indexOf('/Date') > -1) {
            value = value.replace(/\/Date(−?\d+)\//, '$1');
            dt = new Date();
            dt.setTime(value);
        } else if (value.indexOf('/') > -1) {
            dt = new Date(Date.parse(value.replace(/-/g, '/')));
        } else {
            dt = new Date(value);
        }
    }
    return dt;
}

$c.formatDate = function(value) {
    if (value == null || value == '') {
        return '';
    }
    var dt = $c.parseToDate(value);//关键代码，将那个长字符串的日期值转换成正常的JS日期格式
    return dt.format("yyyy-MM-dd"); //这里用到一个javascript的Date类型的拓展方法，这个是自己添加的拓展方法，在后面的步骤3定义
}

$c.formatDateTime = function(value){
    if (value == null || value == '') {
        return '';
    }
    var dt = $c.parseToDate(value);
    return dt.format("yyyy-MM-dd hh:mm:ss");
}

/**
 * layer弹出层
 * @param title 标题
 * @param width_par 宽度%
 * @param height_par 高度%
 * @param url 请求地址
 * @param callback 回调
 * @param btn 按钮设置 默认:['保存', '保存并关闭', '取消']
 */
$c.openBarWin = function(title, width_par, height_par, url, callback) {
    // if (typeof (btn) == 'undefined') {
    //     btn = ['保存', '保存并关闭', '取消'];
    //     flag = false;
    // }
    // var flag = true;
    // if (btn.length > 2)
    //     flag = false;
    title = '<i class="icon wb-order"></i>' + title;
    layer.open({
        type: 2,
        title: title,
        // btn: btn,
        shade: false,
        maxmin: true,
        shade: 0.5,
        anim: 4,
        area: [width_par + '%', height_par + '%'],
        maxmin: false,
        //skin: 'layui-layer-rim', //加上边框
        content: url,
        // btn2: function (index) {
        //     if (flag) {
        //         this.close(index);//编辑时 该按钮是取消的功能
        //         return;
        //     }
        //     var iframe = window.frames["layui-layer-iframe" + index];
        //     if (!iframe)
        //         iframe = window.parent.frames["layui-layer-iframe" + index];
        //     if (typeof (iframe.saveData) != "undefined") {
        //         iframe.saveData.call(this, callback, index, true);
        //     } else if (typeof (callback) != "undefined") {
        //         callback(index);
        //         this.close(index); //一般设定yes回调，必须进行手工关闭
        //     } else {
        //         this.close(index); //一般设定yes回调，必须进行手工关闭
        //     }
        //     return false;
        // },
        // yes: function (index, layero) {
        //     var iframe = window.frames["layui-layer-iframe" + index];
        //     if (!iframe)
        //         iframe = window.parent.frames["layui-layer-iframe" + index];
        //     if (typeof (iframe.saveData) != "undefined") {
        //         iframe.saveData.call(this, callback, index, flag);
        //     } else if (typeof (callback) != "undefined" && callback != "") {
        //         callback(index);
        //         this.close(index); //一般设定yes回调，必须进行手工关闭
        //     } else {
        //         this.close(index); //一般设定yes回调，必须进行手工关闭
        //     }
        // },
        end: function () {
            $(".layui-laypage-btn")[0].click();
        }
    });
}

/**
 * todo vue.js 的created 和 mounted 中layer还未创建
 */
$c.alert = function (msg) {
    swal({
            title: msg,
            text: "",
            type: "info",
            closeOnConfirm: true,
            // showCancelButton: true,
            // showLoaderOnConfirm: true,
     });
};

$c.showLoading = function(msg) {
    swal({
        title: "Loading...",
        type: "",
        imageUrl: '/static/imgs/loading.gif',
        allowEscapeKey: false,
        allowOutsideClick: false,
        showConfirmButton: false
    });
};

$c.hideLoading = function() {
    swal.close();
};

$c.successMsg = function(msg, time=3000) {
    swal({
        title: msg,
        type: "success",
        timer: time,
        showConfirmButton: false
    });
};
$c.errorMsg = function(msg, time=3000) {
    swal({
        title: msg,
        type: "error",
        timer: time,
        showConfirmButton: false
    });
};

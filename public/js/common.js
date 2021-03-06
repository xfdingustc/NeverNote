/**
 * Created by Xiaofei on 2016/1/19.
 */

var Notebook = {
    cache: {},
}

var trimTitle = function(title) {
    if (!title || typeof title != 'string') {
        return '';
    }

    return title.replace(/</g, "&lt;").replace(/>/g, "&gt;");
}

function getObjectId() {
    return ObjectId();
}

//function ajaxGet(url, param, successFunc, failureFunc, async) {
//    return _ajax("GET", url, param, successFunc, failureFunc, async);
//}

function ajaxGet(url, param, successFunc, failureFunc, async) {
    return _ajax("GET", url, param, successFunc, failureFunc, async);
}

function ajaxPost(url, param, successFunc, failureFunc, async) {
    _ajax("POST", url, param, successFunc, failureFunc, async);
}

function _ajaxCallback(ret, successFunc, failureFunc) {
    if (ret === true || ret == "true" || typeof ret == "object") {

        if (typeof successFunc == "function") {
            successFunc(ret);
        }
    } else {
        if (typeof failureFunc == "function") {
            failureFunc(ret);
        } else {
            alert("error!");
        }
    }
}

function _ajax(type, url, param, successFunc, failureFunc, async) {
    if (typeof async == "undefined") {
        console.log("async");
        async = true;
    } else {
        console.log("sync");
        async = false;
    }

    console.log("Ajax");
    return $.ajax({
        type: type,
        url: url,
        data: param,
        async: async,
        success: function(ret) {
            _ajaxCallback(ret, successFunc, failureFunc);
        },
        error: function(ret) {
            _ajaxCallback(ret, successFunc, failureFunc);
        }
    })
}

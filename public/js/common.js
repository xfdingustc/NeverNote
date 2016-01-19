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

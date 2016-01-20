/**
 * Created by Xiaofei on 2016/1/19.
 */
Notebook.allNotebookId = "0";
Notebook.trashNotebookId = "-1";

Notebook.isAllNotebookId = function(notebookId) {
    return notebookId == Notebook.allNotebookId;
}

Notebook.getTreeSetting = function(isSearch, isShare) {
    function addDiyDom(treeId, treeNode) {
        var spaceWidth = 5;
        var switchObj = $("#" + treeId + " #" + treeNode.tId + "_switch");
        var icoObj = $("#" + treeId + " #" + treeNode.tId + "_ico");
        switchObj.remove();
        icoObj.before(switchObj);


        if (!isShare) {
            if (!Notebook.isAllNotebookId(treeNode.NotebookId)) {
                icoObj.after($('<span class="notebook-number-notes" id="numberNotes ' + treeNode.NotebookId +'">' + (treeNode.NumberNotes || 0) + '</span>'));
                icoObj.after($('<span class="fa notebook-setting" title="setting"></span>'))
            }
        }
    }
    var setting = {
        view: {
            showTitle: true,
            showLine: false,
            showIcon: false,
            selectedMulti: false,
            dblClickExpand: false,
            //addDiyDom: addDiyDom
        },
        data: {
            key: {
                name: "Title",
                children: "Sub",
            }
        }
    };



    return setting;
}

Notebook.renderNotebooks = function(notebooks) {
    var self = this;

    if (!notebooks || typeof notebooks != "object" || notebooks.length < 0) {
        notebooks = [];
    }

    for (var i = 0, len = notebooks.length; i < len; ++i) {
        var notebook = notebooks[i];
        notebook.Title = trimTitle(notebook.Title);
    }

    notebooks = [
        {
            NotebookId: Notebook.allNotebookId,
            Title: getMsg("all"),
        }
    ].concat(notebooks)

    self.tree = $.fn.zTree.init($("#notebookList"), self.getTreeSetting(), notebooks)
}
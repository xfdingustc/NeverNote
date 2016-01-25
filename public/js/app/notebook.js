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
            },

        },
        edit: {
            enable: true,
            showRenameBtn: false,
            showRemoveBtn: false,
        },
        callback: {
            beforeRename: function(treeId, treeNode, newName, isCancel) {
                if (treeNode.IsNew) {
                    var parentNode = treeNode.getParentNode();
                    var parentNotebookId = parentNode ? parentNode.NotebookId : "";
                    self.doAddNotebook(treeNode.NotebookId, newName, parentNotebookId);
                }
            }
        }
    };



    return setting;
}

Notebook.addNotebook = function() {
    var self = Notebook;

    var newNotebook = {
        Title: "",
        NotebookId: getObjectId(),
        IsNew: true,
    }
    self.tree.addNodes(null, newNotebook, true, true);
}

Notebook.postAddNotebook = function(parentNotebookId) {
    var self = Notebook;
    var newNotebook = {
        parentNotebookId: parentNotebookId,
    }
    var successFunc = function(ret) {
        if (ret.Title == "") {
            ret.Title = "New Notebook";
        }
        self.tree.addNodes(null, ret, true, true);
    }
    ajaxPost("/notebook/addNotebook", newNotebook, successFunc);
}

Notebook.doAddNotebook = function(notebookId, title, parentNotebookId) {
    var self = Notebook;
    var newNotebook = {
        notebookId: notebookId,
        title: title,
        parentNotebookId: parentNotebookId,
    }
    var successFunc = function(ret) {
        if (ret.NotebookId) {
            var notebook = self.tree.getNodeByTId(notebookId)
            notebook.IsNew = false;
        }
    }
    ajaxPost("/notebook/addNotebook", newNotebook, successFunc);
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


$(function() {
   $("#addNotebookPlus").click(function(e) {
        e.stopPropagation();
        Notebook.postAddNotebook();
    });
});
/**
 * Created by Xiaofei on 2016/1/19.
 */
Notebook.allNotebookId = "0";
Notebook.trashNotebookId = "-1";

Notebook.isAllNotebookId = function(notebookId) {
    return notebookId == Notebook.allNotebookId;
}

Notebook.getTreeSetting = function(isSearch, isShare) {
    var self = Notebook;
    function addDiyDom(treeId, treeNode) {
        var spaceWidth = 5;
        var switchObj = $("#" + treeId + " #" + treeNode.tId + "_switch");
        var icoObj = $("#" + treeId + " #" + treeNode.tId + "_ico");
        switchObj.remove();
        icoObj.before(switchObj);


        if (!isShare) {
            if (!Notebook.isAllNotebookId(treeNode.NotebookId)) {
                icoObj.after($('<span class="notebook-number-notes" id="numberNotes ' + treeNode.NotebookId +'">' + (treeNode.NumberNotes || 0) + '</span>'));
                icoObj.after($('<span class="fa notebook-setting" title="setting"></span>'));
            }
        } else {
            icoObj.after($('<span class="fa notebook-setting" title="setting"></span>'));
        }
    }

    if (!isShare) {
        var onDblClick = function(e) {
            var notebookId = $(e.target).attr("notebookId");
            if (!Notebook.isAllNotebookId(notebookId)) {
                self.updateNotebookTitle(e.target);
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
            addDiyDom: addDiyDom
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
            onDblClick: onDblClick,
            beforeRename: function(treeId, treeNode, newName, isCancel) {
                self.doUpdateNotebookTitle(treeNode.NotebookId, newName);
            }
        }
    };



    return setting;
}

Notebook.updateNotebookTitle = function(target) {
    var self = Notebook;

    var notebookId = $(target).attr("notebookId");

    self.tree.editName(self.tree.getNodeByTId(notebookId));
}

Notebook.doUpdateNotebookTitle = function(notebookId, newTitle) {
    var self = Notebook;


    ajaxPost("/notebook/updateNotebookTitle", {
        notebookId: notebookId,
        title: newTitle,
    }, function(e) {
        if (e.Ok) {

        }
    });
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


Notebook.deleteNotebook = function(target) {
    var self = Notebook;


    var notebookId = $(target).attr("notebookId");
    if (!notebookId) {
        return;
    }

    var deletedNotebook = {
        notebookId: notebookId,
    };

    var successFunc = function(ret) {
        if (ret.Ok) {
            console.log("Ok");
            self.tree.removeNode(self.tree.getNodeByTId(notebookId));
        } else {
            console.log("not ok");
        }
    };

    var failFunc = function(ret) {
        console.log("faile");
    }

    ajaxGet("/notebook/DeleteNotebook", deletedNotebook, successFunc, failFunc);
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
   var notebookListMenu = {
       width: 180,
       items: [
           {
               text: getMsg("delete"),
               icon: "",
               alias: 'delete',
               faIcon: "fa-trash-o",
               action: Notebook.deleteNotebook
           }
       ],
       onShow: applyrule,
       parent: "#notebookList ",
       children: "li a"
   }

    function applyrule(menu) {

    }

    $("#addNotebookPlus").click(function(e) {
        e.stopPropagation();
        Notebook.postAddNotebook();
    });


    Notebook.contextmenu = $("#notebookList li a").contextmenu(notebookListMenu);

    $("#notebookList").on("click", ".notebook-setting", function(e) {
        e.preventDefault();
        e.stopPropagation();
        var $p = $(this).parent();
        Notebook.contextmenu.showMenu(e, $p);
    });

});
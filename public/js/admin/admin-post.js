$(function () {
    Post.init();
})

var Post = {
    editor: null,
    init: function () {
        Post.editor = editormd({
            id: "editormd",
            height: 640,
            syncScrolling: "single",
            saveHTMLToTextarea: true,
            path: "/public/mdeditor/lib/"
        });
    },
    submit: function () {
        var data = {
            "data.Title": $("#blog-title").val(),
            "data.ContentMD": Post.editor.getMarkdown(),
            "data.ContentHTML": Post.editor.getHTML(),
            "data.Type": 0,
        }
        $.post("/admin/post/index", data, function (d) {
            console.log(d);
        })
    }
}
layui.use(['form', 'layedit', 'laydate'], function () {
    var token = getCookie("token");
    var form = layui.form
        , layer = layui.layer
        , layedit = layui.layedit
        , laydate = layui.laydate;
    //监听提交
    form.on('submit(save)', function (data) {
        $.ajax({
            type: "post",
            url: "/config/addbalance",
            data: data.field,
            headers: {'AUTHORIZATION': token},
            success: function (data) {
                if (data.status === 200) {
                    layer.msg('添加成功');
                    setTimeout(function () {
                        var index = parent.layer.getFrameIndex(window.name); //先得到当前iframe层的索引
                        window.parent.location.reload();
                        parent.layer.close(index); //再执行关闭
                    }, 1000);
                }
            }
        });
        return false;
    });
});

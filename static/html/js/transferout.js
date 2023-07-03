function initRole(res) {
    $("#id").val(res.id);
    $("#member").html(res.username);
    $("#balance").html(res.balance);
    $("#transfer_in").html(res.transfer_in);
    $("#transfer_out").html(res.transfer_out);
    $("#interest").html(res.interest);
}
layui.use(['form', 'laydate'], function () {
    var layer = layui.layer;
    var form = layui.form;
    var index = parent.layer.getFrameIndex(window.name); //获取窗口索引;
    /* 自定义验证规则 */
    form.verify({
        intnum: [/^[1-9]\d*$/, '必须为整数'],
        content: function (value) {
            layedit.sync(editIndex);
        }
    });
    /* 监听提交 */
    form.on('submit(editUserSubmit)', function (data) {
        var paramsData = data.field;
        paramsData.type = 1;
        $.ajax({
            type: 'post',
            url: '/transfer',
            data: paramsData,
            success: function (result) {
                if (result.code == 200) {
                    layer.alert("修改成功", {
                        icon: 6,
                        title: "提示"
                    }, function (index) {
                        layer.close(index);
                        var index = parent.layer.getFrameIndex(window.name);  //先得到当前iframe层的索引
                        parent.layer.close(index);  //再执行关闭
                        window.parent.location.reload();
                    });
                } else {
                    layer.alert(result.message, {
                        icon: 5,
                        title: "提示"
                    }, function (index) {
                        layer.close(index);
                        var index = parent.layer.getFrameIndex(window.name);  //先得到当前iframe层的索引
                        parent.layer.close(index);  //再执行关闭
                        window.parent.location.reload();
                    });
                }
            }
        });
        return false;
    });
    var active = {
        cancel: function (set) {
            parent.layer.close(index);
        }
    };
    $('.layui-btn').on('click', function () {
        var othis = $(this),
            type = othis.data('type');
        active[type] && active[type].call(this);
    });
});

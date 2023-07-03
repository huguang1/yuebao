function initRole(res) {
    $("#id").val(res.id);
}
layui.use(['form', 'laydate'], function () {
    var layer = layui.layer;
    var form = layui.form;
    var index = parent.layer.getFrameIndex(window.name); //获取窗口索引;
    /* 自定义验证规则 */
    form.verify({
        userName: function (value) {
            if (value.length < 3) {
                return '名称至少得3个字';
            }
        },
        pwd: [/(.+){3,12}$/, '密码必须3到12位'],
        content: function (value) {
            layedit.sync(editIndex);
        }
    });
    /* 监听提交 */
    form.on('submit(editUserSubmit)', function (data) {
        if (data.field.password != data.field.repassword) {
            layer.alert('两次输入的密码不一致，请确认！');
            return false;
        }
        var paramsData = data.field;
        $.ajax({
            type: 'post',
            url: '/changepassword',
            data: paramsData,
            success: function (result) {
                if (result.code == 200) {
                    layer.alert("修改成功", {
                        icon: 6,
                        title: "提示"
                    }, function (index) {
                        layer.close(index);
                        var index = parent.layer.getFrameIndex(window.name); //先得到当前iframe层的索引
                        parent.layer.close(index); //再执行关闭
                        window.parent.location.reload();
                    });
                } else {
                    layer.alert(result.message, {
                        icon: 5,
                        title: "提示"
                    }, function (index) {
                        layer.close(index);
                        var index = parent.layer.getFrameIndex(window.name); //先得到当前iframe层的索引
                        parent.layer.close(index); //再执行关闭
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

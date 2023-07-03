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
            url: "/config/updatebalance",
            data: data.field,
            headers: {'AUTHORIZATION': token},
            success: function (data) {
                if (data.status === 200) {
                    layer.msg("修改成功");
                    setTimeout(function () {
                        var index = parent.layer.getFrameIndex(window.name); //先得到当前iframe层的索引
                        parent.layer.close(index); //再执行关闭
                    }, 1000);

                }
            }
        });
        return false;
    });

    function renderForm() {
        layui.use('form', function () {
            var form = layui.form;
            form.render();
        });
    }
});

function inputMemberInfo(data) {
    $('input[name="member"]').val(data.member);
    $('input[name="balance"]').val(data.balance);
    $('input[name="transfer_in"]').val(data.transfer_in);
    $('input[name="transfer_out"]').val(data.transfer_out);
    $('input[name="interest"]').val(data.interest);
    $("#id").val(data.id);
    layui.form.render('select');
}

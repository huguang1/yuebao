layui.use(['table', 'layer',], function () {
    var token = getCookie("token");
    var $ = layui.jquery, layer = layui.layer;
    var table = layui.table;
    var form = layui.form;
    var tableIns = table.render({
        elem: '#userList',
        url: '/config/userlist',
        method: 'get',
        headers: {'AUTHORIZATION': token},
        cols: [[{
            field: 'id',
            title: 'ID',
            sort: true,
            align: 'center'
        }, {
            field: 'username',
            title: '用户名',
            align: 'center'
        }, {field: 'is_superuser', title: '超级用户', align: 'center', templet: function (d) {
            if (d.is_superuser === 1) {
                return "<div class=\"layui-form\"> <input type=\"checkbox\" onclick=\"changeStatus();\" status=" + d.is_superuser + " value=\"" + d.id + "\" name=\"isValid\" lay-filter=\"isValid\" checked lay-skin=\"switch\" lay-text=\"是|否\"></div>";
            } else {
                return "<div class=\"layui-form\"> <input type=\"checkbox\"  status=" + d.is_superuser + " value=\"" + d.id + "\" name=\"isValid\" lay-filter=\"isValid\" lay-skin=\"switch\" lay-text=\"是|否\"></div>";
            }
        }}
        , {
            field: 'last_login',
            title: '最近登陆',
            align: 'center'
        }, {
            title: '操作',
            width: 250,
            align: 'center',
            fixed: 'right',
            toolbar: '#barDemo'
        }]],
        page: true //是否显示分页
        , parseData: function (res) {
            return {
                "code": 0,
                'msg': '',
                "count": res.count,
                "data": res.results
            }
        }
        , limit: 10,
        limits: [5, 10, 100]
        //添加权限控制
    });

    $('#selectbtn').on('click', function () {
        active.reload();
    });
    //监听工具条
    table.on('tool(demo)', function (obj) {
        var data = obj.data;
        if (obj.event === 'del') {
            layer.confirm('真的删除行么', function (index) {
                $.ajax({
                    type: "post",
                    url: "/config/deleteuser",
                    data: {"id": data.id},
                    dataType: "json",
                    headers: {'AUTHORIZATION': token},
                    success: function (data) {
                        if (data.status === 200){
                            layer.msg('删除成功');
                            setTimeout(function () {
                                active.reload();
                            }, 1000);
                        }
                    }
                });
                layer.close(index);
            });
        } else if (obj.event === 'edit') {
            var index = layer.open({
                type: 2,
                content: '/static/view/user/editUser.html?id=' + data.id,
                area: ['45%', '50%'],
                maxmin: true,
                success: function (layero, index) {
                    var body = layer.getChildFrame('body', index);//确定页面间的父子关系，没有这句话数据传递不了
                    var iframeWin = window[layero.find('iframe')[0]['name']];
                    iframeWin.inputMemberInfo(data);
                },
                end: function () {
                    active.reload();
                }
            });
        }
    });
    active = {
        reload: function () {
            var searchAccount = $('#searchAccount');
            // 执行重载
            table.reload('userList', {
                page: {
                    curr: 1
                },
                where: {
                    search_account: searchAccount.val()
                }
            });
        }
    };
    // 添加
    $('#addbtn').on('click', function () {
        layer.open({
            type: 2,
            title: false,
            content: '/static/view/user/addUser.html',
            area: ['45%', '45%']
        });
    });
    // 超级用户设置
    form.on('switch(isValid)', function(obj){
        var id = $(this).val();
        var status = $(this).attr("status");
        if (status === 'false') {
            $(this).attr("status", "true");
            status = 1;
        } else if (status === 'true') {
            $(this).attr("status", "false");
            status = 0;
        }
        $.ajax({
            type: 'patch',
            url: '/config/user/' + id + "/",
            data: {"is_superuser": status},
            async: false,
            headers: {'AUTHORIZATION': "JWT " + token},
            success: function (result) {
                if (result) {
                    layer.msg("修改成功");
                    active.reload();
                } else {
                    layer.msg("修改失败");
                }
            }
        })

    });
});

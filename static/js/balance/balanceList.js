layui.use(['table', 'layer',], function () {
    var token = getCookie("token");
    var $ = layui.jquery, layer = layui.layer;
    var table = layui.table;
    var tableIns = table.render({
        elem: '#balanceList',
        url: '/config/balancelist',
        method: 'get',
        headers: {'AUTHORIZATION': token},
        cols: [[{
            field: 'id',
            title: 'ID',
            sort: true,
            align: 'center'
        }, {
            field: 'member',
            title: '会员账号',
            align: 'center'
        }, {
            field: 'balance',
            title: '当前余额',
            align: 'center'
        }, {
            field: 'transfer_in',
            title: '今日转入',
            align: 'center'
        }, {
            field: 'transfer_out',
            title: '今日转出',
            align: 'center'
        }, {
            field: 'interest',
            title: '利息',
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
                    url: "/config/deletebalance",
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
                content: '/static/view/balance/editBalance.html?id=' + data.id,
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
            table.reload('balanceList', {
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
            title: '添加用户',
            content: '/static/view/balance/addBalance.html',
            area: ['75%', '75%']
        });
    });


});

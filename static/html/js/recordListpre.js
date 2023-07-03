layui.use(['table', 'layer'], function () {
    var username = cookie.get('username');
    var $ = layui.jquery, layer = layui.layer;
    var table = layui.table;
    var tableIns = table.render({
        elem: '#recordList',
        url: '/recordlist?member=' + username,
        method: 'get',
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
            field: 'amount',
            title: '金额',
            align: 'center'
        }, {
            field: 'type', title: '转账类型', align: 'center', templet: function (obj) {
                if (obj.type == 0) {
                    return '转入';
                } else if (obj.type == 1) {
                    return '转出';
                } else if (obj.type == 2) {
                    return '存入利息';
                } else {
                    return '';
                }
            }
        }]],
        page: true //是否显示分页
        , parseData: function (res) {
            return {
                "code": 0,
                'msg': '',
                "count": res.count,
                "data": res.data
            }
        }
        , limit: 10,
        limits: [5, 10, 100]
        //添加权限控制
    });

    $('#selectbtn').on('click', function () {
        active.reload();
    });

    active = {
        reload: function () {
            var searchAccount = $('#searchAccount');
            // 执行重载
            table.reload('recordList', {
                page: {
                    curr: 1
                },
                where: {
                    search_account: searchAccount.val()
                }
            });
        }
    };
});

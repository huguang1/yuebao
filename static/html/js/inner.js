/**
 * Created by python on 19-7-29.
 */
layui.use(['table', 'layer'], function () {
    var $ = layui.jquery;
    var layer = layui.layer;
    var username = document.getElementById("membername").innerHTML;
    // 更改密码
    $('.changePWD').on('click', function () {
        var index = layer.open({
            title: '修改密码',
            type: 2,
            content: '/static/html/changepwd.html',
            area: ['45%', '52%'],
            maxmin: true,
            success: function (layero, index) {
                var body = layer.getChildFrame('body', index);//确定页面间的父子关系，没有这句话数据传递不了
                var iframeWin = window[layero.find('iframe')[0]['name']];
                $.ajax({
                    type: 'get',
                    url: '/changepassword',
                    data: {"username": username},
                    success: function (result) {
                        if (result.code == 200) {
                            iframeWin.initRole(result.data);
                        } else {
                            layer.alert("查询失败");
                        }
                    }
                })
            }
        });
    });

    // 查看记录
    $('.recordlist').on('click', function () {
        var index = layer.open({
            title: '修改密码',
            type: 2,
            content: '/static/html/recordListpre.html',
            area: ['85%', '90%'],
            maxmin: true,
            success: function (layero, index) {
                var body = layer.getChildFrame('body', index);//确定页面间的父子关系，没有这句话数据传递不了
                var iframeWin = window[layero.find('iframe')[0]['name']];
            }
        })
    });

    // 转入
    $('.transferin').on('click', function () {
        var index = layer.open({
            title: '转入',
            type: 2,
            content: '/static/html/transferin.html',
            area: ['85%', '60%'],
            maxmin: true,
            success: function (layero, index) {
                var body = layer.getChildFrame('body', index);//确定页面间的父子关系，没有这句话数据传递不了
                var iframeWin = window[layero.find('iframe')[0]['name']];
                $.ajax({
                    type: 'get',
                    url: '/transfer',
                    data: {"username": username},
                    success: function (result) {
                        if (result.code == 200) {
                            iframeWin.initRole(result.data);
                        } else {
                            layer.alert("查询失败");
                        }
                    }
                })
            }
        })
    });

    // 转出
    $('.transferout').on('click', function () {
        var index = layer.open({
            title: '转出',
            type: 2,
            content: '/static/html/transferout.html',
            area: ['85%', '60%'],
            maxmin: true,
            success: function (layero, index) {
                var body = layer.getChildFrame('body', index);//确定页面间的父子关系，没有这句话数据传递不了
                var iframeWin = window[layero.find('iframe')[0]['name']];
                $.ajax({
                    type: 'get',
                    url: '/transfer',
                    data: {"username": username},
                    success: function (result) {
                        if (result.code == 200) {
                            iframeWin.initRole(result.data);
                        } else {
                            layer.alert("查询失败");
                        }
                    }
                })
            }
        })
    });
});

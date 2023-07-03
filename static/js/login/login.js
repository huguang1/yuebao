var box = {
    content: ''
    , btn: ['确定']
    , yes: function () {
        window.location.href = '/static/templates/login.html';
    }
};

$(function () {
    //  跳转顶层页面
    if (window !== top) {
        top.location.href = location.href;
    }
    //  更新验证码
    $.ajax({
        url: '/captcha',
        type: 'get',
        dataType: 'json',
        data: {},
        success: function (data) {
            if (data.status === 200) {
                var CaptchaId = data.CaptchaId;
                $("#captcha").attr("src", "/captcha/" + CaptchaId +".png?t=" + new Date().getTime());
            } else {
                box.content = data.message;
                layer.open(box)
            }
        },
        error: function () {
            box.content = '服务器错误';
            layer.open(box)
        }
    });
    $("#captcha").on("click", function () {
        setSrcQuery(document.getElementById('captcha'), "reload=" + (new Date()).getTime());
        return false;
    });
    function setSrcQuery(e, q) {
        var src  = e.src;
        var p = src.indexOf('?');
        if (p >= 0) {
            src = src.substr(0, p);
        }
        e.src = src + "?" + q
    }
    //  更新验证码
    $.ajax({
        url: '/logintoken',
        type: 'get',
        dataType: 'json',
        data: {},
        success: function (data) {
            if (data.status === 200) {
                console.log("成功")
            } else {
                console.log(data.message)
            }
        },
        error: function () {
            console.log("服务器错误")
        }
    });
});

// 用户登陆
layui.use('layer', function () {
    var layer = layui.layer;
    $('#login_button').click(function () {
        var layer = layui.layer;
        var username = $('#username').val();
        var password = $('#password').val();
        var text = $('#verificationcode').val();
        var captchaId = cookie.get('captchaId');
        var loginToken = cookie.get('loginToken');
        if (username && password && text && captchaId && loginToken) {
            $.ajax({
                url: '/login',
                type: 'post',
                dataType: 'json',
                data: {
                    'username': username,
                    'password': password,
                    'text': text,
                    'captchaId': captchaId
                },
                headers: {'AUTHORIZATION': loginToken},
                success: function (data) {
                    if (data.status === 200) {
                        window.location.href = '/static/view/index.html';
                    } else {
                        box.content = data.message;
                        layer.open(box)
                    }
                },
                error: function () {
                    box.content = '用户名或密码错误';
                    layer.open(box)
                }
            })
        } else {
            box.content = '请输入完整参数';
            layer.open(box)
        }
    });
});

$(function () {
    initUser();
    /**
     * 获取用户信息
     */
    function initUser() {
        var userID = cookie.get("userID");
        var token = cookie.get("token");
        console.log(token);
        $.ajax({
            url: "/config/user",
            type: "post",
            data: {'userID': userID},
            headers: {'AUTHORIZATION': token},
            success: function (response) {
                if (response.status !== 200) {
                    window.location.href = "/static/templates/login.html";
                }
                $("#nameh3").html(response.data);
                $("#logout").attr("href","/logout?name="+response.data);
            },
            error: function (e) {
                console.log(e);
            }
        })
    }
});
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>开源CRM</title>
    <link rel="stylesheet" href="/static/layui/css/layui.css">
    <link rel="stylesheet" href="/static/layui/css/main.css">
    <script type="application/javascript" src="/static/layui/layui.js"></script>
    <script type="application/javascript" src="/static/layui/jquery.js"></script>
    <script type="application/javascript" src="/static/layui/jquery.validate.js"></script>
    <script type="application/javascript" src="/static/layui/jquery.form.js"></script>
</head>
<body class="manage-login-body">
<div class="manage-loginpage layui-anim">
    <img class="manage-login-log" src="">
    <h1>CRM模板</h1>
    <form class="layui-form" method="post" action="/UserAccount/Login" id="loginform">
        <div class="layui-form-item manage-loginpage-input">
            <input type="text" name="username" required lay-verify="required" placeholder="请输入用户名" autocomplete="off" class="layui-input">
        </div>

        <div class="layui-form-item manage-loginpage-input">
            <input type="password" name="password" required lay-verify="required" placeholder="请输入密码" autocomplete="off" class="layui-input">
        </div>

        <div class="layui-form-item manage-loginpage-input">
            <button type="submit" class="layui-btn">立即登录</button>
        </div>
    </form>
</div>
<script type="text/javascript">
    var layer;
    $().ready(function() {
        layui.use('layer',
            function() {
                layer = layui.layer;
            })
        $('#loginform').validate({
            rules: {
                username: {
                    required: true
                },
                password: {
                    required: true
                }
            },
            messages: {
                username: {
                    required: "请输入用户名"
                },
                password: {
                    required: "请输入密码"
                }
            },
            errorPlacement: function(error, element) {
                error.insertAfter(element);
            },
            submitHandler: function(form) {
                layer.load(0, { shade: false });
                $(form).ajaxSubmit(function(res) {
                    if (res.code == 1) {
                        window.location = '/EnterCustom/ConsoleIndex';
                    } else {
                        layer.msg('登录失败', { icon: 5 });
                    }
                });
            }
        });
    });
</script>
</body>
</html>

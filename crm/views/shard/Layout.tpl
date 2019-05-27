<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>首页--星辰科技</title>
    <meta name="renderer" content="webkit">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
    {{.header}}

</head>
<body class="layui-layout-body" leftmargin=0 topmargin=0 oncontextmenu='return false' ondragstart='return false' onselectstart ='return false' onselect='document.selection.empty()' oncopy='document.selection.empty()' onbeforecopy='return false' onmouseup='document.selection.empty()'>
<div class="layui-layout layui-layout-admin">
    <!-- 头部区域 -->
    <div class="layui-header">
        <div class="layui-logo">开源CRM</div>

        <ul class="layui-nav layui-layout-right">
            <li class="layui-nav-item">
                <a href="/Account/LoginOut">退出</a>
            </li>
            <li class="layui-nav-item">
                <a href="/Account/LoginOut">退出</a>
            </li>
            <li class="layui-nav-item">
                <a href="/Account/LoginOut">退出</a>
            </li>
        </ul>
    </div>
    <!-- 侧边栏 -->
    <div class="layui-side layui-bg-black">
        <div class="layui-side-scroll">
            <ul class="layui-nav layui-nav-tree" lay-filter="test">
                <li class="layui-nav-item">
                <a href="javascript:;">邮件营销</a>
                <dl class="layui-nav-child">
                    <dd><a href="javascript:;">邮箱资源管理</a></dd>
                    <dd><a >目标邮箱管理</a></dd>
                    <dd><a >邮件模板</a></dd>
                    <dd><a >邮件发送任务</a></dd>
                </dl>
                </li>
            </ul>
        </div>
    </div>

    <!-- 内容区域 -->
    <div class="layui-body">
{{.LayoutContent}}
    </div>

    <!-- 底部固定区域 -->
    <div class="layui-footer">
        © 开源CRM - @DateTime.Now.Year
    </div>
</div>
{{.Sidebar}}
</body>
</html>
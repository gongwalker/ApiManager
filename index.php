<?php
    include './MinPHP/run/init.php';
    $act = $_GET['act'];
    $act = empty($act) ? 'index' : $_GET['act'];
    $menu = '';
    switch($act){
        //接口分类
        case 'cate':
            $menu = ' - 分类管理';
            $file = './MinPHP/run/cate.php';
            break;
        //登录退出
        case 'login':
            $menu = ' - 登录';
            $file = './MinPHP/run/login.php';
            break;
        //首页
        case 'index':
            $menu = ' - 欢迎';
            $file ='./MinPHP/run/hello.php';
            break;
        //接口详细页
        case 'api':
            $sql = "select cname from cate where aid='{$_GET['tag']}' and isdel=0";
            $menu = find($sql);
            $menu = ' - ' . $menu['cname'];
            $file ='./MinPHP/run/info.php';
            break;
        //接口排序页
        case 'sort':
            $sql = "select cname from cate where aid='{$_GET['tag']}' and isdel=0";
            $menu = find($sql);
            $menu = ' - ' . "<a href='".U(array('act'=>'api','tag'=>$_GET['tag']))."'>{$menu['cname']}</a>";
            $file ='./MinPHP/run/sort.php';
            break;
        //导出静态文件
        case 'export':
            die(include('./MinPHP/run/export.php'));
            break;
        //ajax请求
        case 'ajax':
            die(include('./MinPHP/run/ajax.php'));
            break;
        //修改密码
        case 'modpwd':
            $menu = ' - 修改密码';
            $file = './MinPHP/run/modpwd.php';
            break;
        default :
            $menu = ' - 欢迎';
            $file = './MinPHP/run/hello.php';
            break;
    }
    include './MinPHP/run/main.php';

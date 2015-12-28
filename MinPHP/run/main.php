<?php defined('API') or exit('http://gwalker.cn');?>
<!DOCTYPE html>
<html lang="zh-CN" style="height:100%">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
    <title>API接口管理工具</title>
    <link rel="icon" type="image/x-icon" href="./MinPHP/res/favicon.ico">
    <link href="./MinPHP/res/bootstrap-3.3.4-dist/css/bootstrap.min.css" rel="stylesheet">
    <link href="./MinPHP/res/style.css" rel="stylesheet">
    <!--[if lt IE 9]>
    <script src="./MinPHP/res/html5shiv.min.js"></script>
    <script src="./MinPHP/res/respond.min.js"></script>
    <![endif]-->
</head>
<body style="height:100%">
<div class="container-fluid" style="background:white;height:100%;">
    <div class="row" style="height:100%;">
        <!--左侧导航start-->
        <div id="navbar" class="col-md-3" style="position:relative;background:#f5f5f5;padding:10px;height:100%;border-right:#ddd 1px solid;overflow-y:auto;<?php if($_COOKIE[C('cookie->navbar')]==1){?>display:none<?php }?>">
            <div style="height:50px;font-size:30px;line-height:50px;">
                <a class="home" style="color:#000000;text-shadow:1px 0px 1px #666;text-decoration: none" href="<?php echo U()?>">
                    <span class="glyphicon glyphicon-random" aria-hidden="true" style="width:40px;"></span>&nbsp;
        <span style="position: relative;top:-3px;">API Manager <span style="font-size:12px;position:relative;top:-13px;">&nbsp;<?php echo C('version->no')?></span>
                </a>
                </span>
            </div>
            <?php
            include('./MinPHP/run/menu.php');
            ?>
        </div>
        <!--左侧导航end-->
        <div id="mainwindow" <?php if($_COOKIE[C('cookie->navbar')]==1){?>class="col-md-12"<?php }else{?>class="col-md-9" <?php }?>  style="height:100%;background:white;margin:0px;overflow-y:auto;padding:0px;">
            <!--顶部导航start-->
            <div class="textshadow" style="font-size:16px;widht:100%;height:60px;line-height:60px;padding:0 16px 0 16px;;border-bottom:#ddd 1px solid">
                <span> <a class="home" href="<?php echo U() ?>">Home</a><?php echo $menu;?></span>
        <span id="topbutton" style="float:right">
            <?php
            if(is_login()){
                echo '欢迎您：' . session('nice_name') . '&nbsp;&nbsp;';
                //如果是接口详情页的话,就显示【导出】按钮 与 【排序】按钮
                if($_GET['act']=='api' && isset($_GET['tag']) && !isset($_GET['op'])){
                    echo '<a href="?act=sort&tag='.$_GET['tag'].'">排序&nbsp;&nbsp;</a>';
                    echo '<a href="?act=export&tag='.$_GET['tag'].'">导出&nbsp;&nbsp;</a>';
                }
                echo '<a href="?act=modpwd">修改密码</a>&nbsp;&nbsp;<a href="?act=login&type=quit">退出&nbsp;&nbsp;<span class="glyphicon glyphicon-log-out" aria-hidden="true"></span></a>';
            }else{
                echo '<a href="?act=login">登录&nbsp;&nbsp;<span class="glyphicon glyphicon-log-in" aria-hidden="true"></span></a>';
            }
            ?>
        </span>
            </div>
            <!--顶部导航end-->
            <!--主窗口start-->
            <div style="padding:16px;">
                <?php
                if(!empty($file)){
                    include($file);
                }
                ?>
            </div>
            <!--主窗口end-->
        </div>
    </div>
</div>
<script src="./MinPHP/res/jquery.min.js"></script>
<script src="./MinPHP/res/jquery.cookie.js"></script>
<script src="./MinPHP/res/bootstrap-3.3.4-dist/js/bootstrap.min.js"></script>
</body>
</html>

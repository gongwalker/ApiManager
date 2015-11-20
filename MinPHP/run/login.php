<!--登录与退出start-->
<?php defined('API') or exit('http://gwalker.cn');?>
<?php
    $type= I($_GET['type']);
    //登录
    if($type  == 'do'){
        $_VAL = I($_POST);
        $login_name = $_VAL['name'];
        $login_pwd = md5($_VAL['pwd']);
        $sql = "select * from user where login_name = '{$login_name}' and login_pwd = '{$login_pwd}' and isdel = '0'";
        $info = find($sql);
        if(!empty($info)){
            session('id',$info['id']); //用户id
            session('nice_name',$info['nice_name']); //昵称
            session('login_name',$info['login_name']); //登录名
            session('issupper',$info['issuper']); //是否为超级管理员
            $time = time();
            $sql = "update user set last_time = '{$time}' where id = {$info['id']}";
            update($sql);
            go(U());
        }else{
            echo '<div class="alert alert-danger" role="alert"><span class="glyphicon glyphicon-info-sign" aria-hidden="true"></span> 登录失败,帐号或密码错误</div>';
        }
    //退出
    }if($type == 'quit'){
        session('login_name','');
        session('issupper','');
        go(U());
    }
?>
<div style="border:1px solid #ddd">
    <div style="background:#f5f5f5;padding:20px;position:relative">
        <h4>登录</h4>
        <div>
            <form action="?act=login&type=do" method="post">
                <div class="form-group">
                    <input type="text" class="form-control" name="name" placeholder="登录名" required="required">
                </div>
                <div class="form-group">
                    <input type="password" class="form-control" name="pwd" placeholder="密码" required="required">
                </div>
                <button class="btn btn-success">Submit</button>
            </form>
        </div>
    </div>
</div>
<!--登录与退出end-->

<!--修改密码start-->
<?php defined('API') or exit('http://gwalker.cn');?>
<?php
    $type= I($_GET['type']);

    if($type  == 'do'){
        $_VAL = I($_POST);
        $ord_pwd = md5($_VAL['ord_pwd']);
        $new_pwd = md5($_VAL['new_pwd']);
        $new_pwd2 = md5($_VAL['new_pwd2']);
        $login_name = session('login_name');
        if ($new_pwd != $new_pwd2) {//判断新密码和确认密码是否一致
            echo '<div class="alert alert-danger" role="alert"><span class="glyphicon glyphicon-info-sign" aria-hidden="true"></span> 新密码和确认密码不一致</div>';
        } else {
            $sql = "select * from user where login_name = '{$login_name}' and login_pwd = '{$ord_pwd}' and isdel = '0'";
            $info = find($sql);
            if (!$info) {//判断旧密码和当前登录账号是否匹配
                echo '<div class="alert alert-danger" role="alert"><span class="glyphicon glyphicon-info-sign" aria-hidden="true"></span> 修改密码失败,帐号或密码错误</div>';
            } else {
                $sql = "update user set login_pwd='{$new_pwd}' where login_name='{$login_name}'";
                $re = update($sql);
                if($re !== false){
                    go(U());
                }else{
                    echo '<div class="alert alert-danger" role="alert"><span class="glyphicon glyphicon-info-sign" aria-hidden="true"></span> 密码修改失败</div>';
                }
            }            
        }
    }
?>
<div style="border:1px solid #ddd">
    <div style="background:#f5f5f5;padding:20px;position:relative">
        <h4>修改密码</h4>
        <div>
            <form action="?act=modpwd&type=do" method="post">
                <div class="form-group">
                    <input type="text" class="form-control" name="ord_pwd" placeholder="旧密码" required="required">
                </div>
                <div class="form-group">
                    <input type="text" class="form-control" name="new_pwd" placeholder="新密码" required="required">
                </div>
                <div class="form-group">
                    <input type="password" class="form-control" name="new_pwd2" placeholder="确认新密码" required="required">
                </div>
                <button class="btn btn-success">Submit</button>
            </form>
        </div>
    </div>
</div>
<!--修改密码end-->

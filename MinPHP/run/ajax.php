<?php
    defined('API') or exit('http://gwalker.cn');
    if(!is_supper()){die('只有超级管理员才可进行ajax操作');}
    //得到ajax操作
    $op = I($_GET['op']);
    //执行ajax操作
    $op();
    //删除某个接口
    function apiDelete(){
        //接口id
        $id = I($_POST['id']);
        $sql = "update api set isdel='1' where id='{$id}'";
        $re = update($sql);
        die ($re ? '1' : '0');
    }
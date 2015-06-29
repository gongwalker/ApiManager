<?php defined('API') or exit('http://gwalker.cn');?>
<?php
    /**
     * @dec 得到配置文件的配置项
     * @param null $name
     * @return mixed
     * 使用方法,例子
     * C('db') 或 C('version->no')
     */
    function C($name = null){
        static $_config = array();
        if(empty($_config)){
            $_config = include_once './MinPHP/core/config.php';
        }
        if(is_null($name)){
            return $_config;
        }else{
            if(strpos($name,'->')){
                $arr = explode('->',$name);
                $tmp = $_config;
                foreach($arr as $v){
                    $tmp = $tmp[$v];
                }
                return $tmp;
            }
            return $_config[$name];
        }
    }

    //得到数据库连接资源
    function M(){
        static $_model = null;
        if(is_null($_model)){
            $db=C('db');
            try {
                $_model = new PDO("mysql:host={$db['host']};dbname={$db['dbname']}","{$db['user']}","{$db['passwd']}");
            } catch ( PDOException $e ) {
                die ( "Connect Error Infomation:" . $e->getMessage () );
            }
            //设置数据库编码
            $_model->exec('SET NAMES utf8');
        }
        return $_model;
    }

    //返回一条记录集
    function find($sql){
        $rs = M()->query($sql);
        $row = $rs->fetch(PDO::FETCH_ASSOC);
        return $row;
    }

    //返回多条记录
    function select($sql){
        $rs = M()->query($sql);
        $rows = array();
        while($row = $rs->fetch(PDO::FETCH_ASSOC)){
            $rows[] = $row;
        }
        return $rows;
    }

    //insert
    function insert($sql){
        return M()->exec($sql);
    }

    //update
    function update($sql){
        return M()->exec($sql);
    }

    //设置和获取session值
    function session($key = null,$value = null){
        $pre = C('session->prefix');  //session前缀
        if(is_null($key)){
            return $_SESSION[$pre];
        }else{
            if(is_null($value)){
                return $_SESSION[$pre][$key];
            }else{
                $_SESSION[$pre][$key] = $value;
            }
        }
    }

    //判断是否登录
    function is_lgoin(){
        $login_name = session('login_name');
        return empty($login_name) ? false : true;
    }

    //判断是否为超级管理员
    function is_supper(){
        return session('issupper') == 1 ? true : false;
    }

    //跳转
    function go($url){
        $gourl = '<script language="javascript" type="text/javascript">window.location.href="'.$url.'"</script>';
        die($gourl) ;
    }

    //生成url
    function U($array = null){
        if(is_null($array)){
            $url = '';
        }else{
            $url = '?'.http_build_query($array);
            $url = str_replace('%23','#',$url);
        }
        return 'index.php'.$url;
    }

    //安全过滤
    function I($val){
        if(is_array($val)){
            foreach($val as $k => $v){
                $val[$k] = I($v);
            }
            return $val;
        }else{
            if(is_numeric($val)){
                return intval($val);
            }else if(is_string($val)){
                return htmlspecialchars(trim($val),ENT_QUOTES);
            }else{
                return $val;
            }
        }
    }
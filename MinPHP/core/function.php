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

    //网站基础路径baseUrl
    function baseUrl(){
        $currentPath = $_SERVER['SCRIPT_NAME'];
        $pathInfo = pathinfo($currentPath);
        $hostName = $_SERVER['HTTP_HOST'];
        $protocol = strtolower(substr($_SERVER["SERVER_PROTOCOL"],0,5))=='https://' ? 'https://' : 'http://';
        return $protocol.$hostName.$pathInfo['dirname']."/";
    }

    //下载html
    function downfile($fileName){
        $fileName = '路径+实际文件名';
        //文件的类型
        header('Content-type: application/pdf');
        //下载显示的名字
        header('Content-Disposition: attachment; filename="保存时的文件名.pdf"');
        readfile("$fileName");
        exit();
    }

/**
 * @dec 下载文件 指定了content参数，下载该参数的内容
 * @access public
 * @param string $showname 下载显示的文件名
 * @param string $content  下载的内容
 * @param integer $expire  下载内容浏览器缓存时间
 * @return void
 */
function download($showname='',$content='',$expire=180) {
    $type	=	"application/octet-stream";
    //发送Http Header信息 开始下载
    header("Pragma: public");
    header("Cache-control: max-age=".$expire);
    //header('Cache-Control: no-store, no-cache, must-revalidate');
    header("Expires: " . gmdate("D, d M Y H:i:s",time()+$expire) . "GMT");
    header("Last-Modified: " . gmdate("D, d M Y H:i:s",time()) . "GMT");
    header("Content-Disposition: attachment; filename=".$showname);
    header("Content-type: ".$type);
    header('Content-Encoding: none');
    header("Content-Transfer-Encoding: binary" );
    die($content);
}

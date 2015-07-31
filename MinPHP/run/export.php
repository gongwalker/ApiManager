<?php
    defined('API') or exit('http://gwalker.cn');
    if(!is_supper()){die('只有超级管理员才可进行导出操作');}
    define('BASEURL',baseUrl());
    //接口分类id
    $tag = I($_GET['tag']);
    //下载的文件名
    $filename = find("select cname from cate where aid='{$tag}'");
    $version = date('YmdHis');
    $filename = $filename['cname'].$version.'.html';
    //要抓取的接口分类url
    $url = BASEURL.U(array('act'=>'api','tag'=>$tag));
    //分类详情页的内容
    $content = file_get_contents($url);

    //========js与css静态文件替换start=======================================
    //css文件替换--start
    $pattern = '/<link href="(.+?\.css)" rel="stylesheet">/is';
    function getCssFileContent($matches){
        $filepath = BASEURL.ltrim($matches[1],'./');
        $content = file_get_contents($filepath);
        return "<style>".$content."</style>";
    }
    $content =  preg_replace_callback($pattern,'getCssFileContent',$content);
    //css文件替换--end

    //js文件替换--start
    $pattern = '/<script src="(.+?\.js)"><\/script>/is';
    function getJSFileContent($matches){
        $filepath = BASEURL.ltrim($matches[1],'./');
        $content = file_get_contents($filepath);
        return "<script>".$content."</script>";
    }
    $content =  preg_replace_callback($pattern,'getJSFileContent',$content);
    //js文件替换--end
    //========js与css静态文件替换end=======================================

    //=======页面锚点连接替换start=======================================
    $pattern = '/<a href=".+?act=api&tag=\d#(\w+).+?">(.+?)<\/a>/is';
    function changeLink($matches){
        return "<a href='#{$matches[1]}'>$matches[2]</a>";
    }
    $content =  preg_replace_callback($pattern,'changeLink',$content);
    $tag = C('version->no');
    //=======页面锚点连接替换end=========================================
$headhtml=<<<START
<!--
=======================================================================
导出时间:{$version}
=======================================================================
此文档由API Manager {$tag} 导出
=======================================================================
github : https://github.com/gongwalker/ApiManager.git
=======================================================================
作者 : 路人庚
=======================================================================
QQ : 309581329
=======================================================================
-->
START;
$appendhtml=<<<END
<script>
$('.glyphicon').remove();
$('#topbutton').html('版本号:{$version}');
$('.home').attr('href','#');
</script>
END;
$content=$headhtml.$content.$appendhtml;
download($filename,$content);
exit;
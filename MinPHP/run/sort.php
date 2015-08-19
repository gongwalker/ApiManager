<?php
    defined('API') or exit();
    if(!is_supper()){die('只有超级管理员才可进行排序操作');}
?>

<!--接口排序管理start-->
<?php
    //操作类型{type}
    $type = $_GET['type'];
    if(empty($type)){
        //已经分类下的所有接口start
        $sql = "select id,num,name from api where aid='{$_GET['tag']}' and isdel=0 order by ord desc,id desc";
        $list = select($sql);
    }else if($type == 'do'){
        $ord = count($_POST['api']);
        foreach($_POST['api'] as $v){
            $sql = "update api set ord = '{$ord}' where id='{$v}' and aid='{$_GET['tag']}'";
            update($sql);
            $ord--;
        }
        $url = U(array('act'=>'api','tag'=>$_GET['tag']));
        go($url);
    }
?>
<div style="border:1px solid #ddd;margin-bottom:20px;">
    <div style="background:#ffffff;padding:20px;">
        <h5 class="textshadow" >接口列表</h5>
        <form action="<?php echo U(array('act'=>'sort','tag'=>$_GET['tag'],'type'=>'do')) ?>" method="post">
            <table class="table">
                <thead>
                <tr>
                    <th class="col-md-2">接口编号</th>
                    <th class="col-md-9">接口名</th>
                    <th class="col-md-1">操作</th>
                </tr>
                </thead>
                <tbody>
                <?php foreach($list as $v){?>
                    <tr>
                        <td><input name="api[]" type="hidden" value="<?php echo $v['id'];?>"><?php echo $v['num'];?></td>
                        <td><?php echo $v['name'];?></td>
                        <td>
                            <span onclick="up(this)" style="color:red;cursor: pointer" class="glyphicon glyphicon-arrow-up" aria-hidden="true"></span>
                            &nbsp;
                            <span  onclick="down(this)" style="color:green;cursor: pointer" class="glyphicon glyphicon-arrow-down" aria-hidden="true"></span>
                        </td>
                    </tr>
                <?php } ?>
                </tbody>
            </table>
            <button class="btn btn-success">Submit</button>
        </form>
    </div>
</div>
<script>
    //上移
    function up(obj){
        var $TR = $(obj).parents('tr');
        var prevTR = $TR.prev();
        prevTR.insertAfter($TR);
        $('tr.info').removeClass('info');
        $TR.addClass('info');
        $TR.hide();
        $TR.show(300);
    }
    //下移
    function down(obj){
        var $TR = $(obj).parents('tr');
        var nextTR = $TR.next();
        nextTR.insertBefore($TR);
        $('tr.info').removeClass('info');
        $TR.addClass('info');
        $TR.hide();
        $TR.show(300);
    }
</script>
<!--接口排序管理end-->
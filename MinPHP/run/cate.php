<?php defined('API') or exit('http://gwalker.cn');?>
<!--接口分类管理-->
<?php
if(!is_supper()){die('只有超级管理员才可对分类进行操作');}
//操作类型{add,delete,edit}
$op = I($_POST['op']);
//是否执行操作(如果为do的话,则为执行添加,删除,编辑操作)
$type = I($_GET['type']);
switch($op){
    //添加
    case 'add':
        if($type  == 'do'){
            $_VAL = I($_POST);
            $cname = $_VAL['cname'];
            $cdesc = $_VAL['cdesc'];
            $time = time();
            if(!empty($cname) && !empty($cdesc)){
                $sql = "insert into cate (cname,cdesc,addtime) values('{$cname}','{$cdesc}','{$time}')";
                $re = insert($sql);
                if($re){
                    go(U());
                }else{
                    echo '<div class="alert alert-danger" role="alert"><span class="glyphicon glyphicon-info-sign" aria-hidden="true"></span> 分类添加失败</div>';
                }
            }else{
                echo '<div class="alert alert-danger" role="alert"><span class="glyphicon glyphicon-info-sign" aria-hidden="true"></span> 分类名与描述不能为空</div>';
            }
        }
    break;
    //删除
    case 'delete':
        $aid = I($_POST['aid']);
        $sql = "update cate set isdel=1 where aid='{$aid}'";
        $re = update($sql);
        if($re){
            go(U());
        }else{
            echo '<div class="alert alert-danger" role="alert"><span class="glyphicon glyphicon-info-sign" aria-hidden="true"></span> 分类删除失败</div>';
        }
    break;
    //编辑
    case 'edit';
        $_VAL = I($_POST);
        if($type == 'do'){
            $sql = "update cate set cname='{$_VAL['cname']}',cdesc='{$_VAL['cdesc']}' where aid='{$_VAL['aid']}'";
            $re = update($sql);
            if($re !== false){
                go(U());
            }else{
                echo '<div class="alert alert-danger" role="alert"><span class="glyphicon glyphicon-info-sign" aria-hidden="true"></span> 分类修改失败</div>';
            }
        }else{
            $aid = $_VAL['aid'];
            $sql = "select * from cate where aid='{$aid}'";
            $info = find($sql);
        }
    break;
}
?>
<?php if($op == 'add'){ ?>
    <div style="border:1px solid #ddd">
        <div style="background:#f5f5f5;padding:20px;position:relative">
            <h4>添加分类</h4>
            <div>
                <form action="?act=cate&type=do" method="post">
                    <div class="form-group">
                        <input type="text" class="form-control" name="cname" placeholder="分类名">
                    </div>
                    <div class="form-group">
                        <input type="text" class="form-control" name="cdesc" placeholder="描述">
                    </div>
                    <button class="btn btn-success" name="op" value="add">Submit</button>
                </form>
            </div>
        </div>
    </div>
<?php }if($op == 'edit'){ ?>
    <div style="border:1px solid #ddd">
        <div style="background:#f5f5f5;padding:20px;position:relative">
            <h4>编辑分类</h4>
            <div>
                <form action="?act=cate&type=do" method="post">
                    <div class="form-group">
                        <input type="hidden" name="aid" value="<?php echo $info['aid'] ?>">
                        <input type="text" class="form-control" name="cname" placeholder="分类名" value="<?php echo $info['cname'] ?>">
                    </div>
                    <div class="form-group">
                        <input type="text" class="form-control" name="cdesc" placeholder="描述" value="<?php echo $info['cdesc'] ?>">
                    </div>
                    <button class="btn btn-success" name="op" value="edit">Submit</button>
                </form>
            </div>
        </div>
    </div>
<?php } ?>
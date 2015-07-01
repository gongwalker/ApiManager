<?php defined('API') or exit();?>
<!--接口详情列表与接口管理start-->
<?php
   $_VAL = I($_POST);
   //操作类型{add,delete,edit}
   $op = $_GET['op'];
   $type = $_GET['type'];
   //添加接口
   if($op == 'add'){
        if($type == 'do'){
            if(!is_supper()){die('只有超级管理员才可对接口进行操作');}
            $aid = I($_GET['tag']);    //所属分类
            if(empty($aid)){
                die('<span class="glyphicon glyphicon-info-sign" aria-hidden="true"></span> 所属分类不能为空');
            }
            $num = htmlspecialchars($_POST['num'],ENT_QUOTES);   //接口编号(为了导致编号的前导0去过滤掉。不用用I方法过滤)
            $name = $_VAL['name'];  //接口名称
            $memo = $_VAL['memo']; //备注
            $des = $_VAL['des'];    //描述
            $type = $_VAL['type'];  //请求方式
            $url = $_VAL['url'];

            $parameter = serialize($_VAL['p']);
            $re = $_VAL['re'];  //返回值
            $lasttime = time(); //最后操作时间
            $lastuid = session('id'); //操作者id
            $isdel = 0; //是否删除的标识
            $sql = "insert into api (
            `aid`,`num`,`name`,`des`,`url`,
            `type`,`parameter`,`re`,`lasttime`,
            `lastuid`,`isdel`,`memo`
            )values (
            '{$aid}','{$num}','{$name}','{$des}','{$url}',
            '{$type}','{$parameter}','{$re}','{$lasttime}',
            '{$lastuid}','{$isdel}','{$memo}'
            )";
            $re = insert($sql);
            if($re){
                go(U(array('act'=>'api','tag'=>$_GET['tag'])));
            }else{
                echo '<div class="alert alert-danger" role="alert"><span class="glyphicon glyphicon-info-sign" aria-hidden="true"></span> 添加失败</div>';
            }
        }
   //修改接口
   }else if($op == 'edit'){
       if(!is_supper()){die('只有超级管理员才可对接口进行操作');}
       //执行编辑
       if($type == 'do'){
           $id = $_VAL['id'];   //接口id
           $num = htmlspecialchars($_POST['num'],ENT_QUOTES);   //接口编号(为了导致编号的前导0去过滤掉。不用用I方法过滤)
           $name = $_VAL['name'];  //接口名称
           $memo = $_VAL['memo']; //备注
           $des = $_VAL['des'];    //描述
           $type = $_VAL['type'];  //请求方式
           $url = $_VAL['url']; //请求地址

           $parameter = serialize($_VAL['p']);
           $re = $_VAL['re'];  //返回值
           $lasttime = time(); //最后操作时间
           $lastuid = session('id'); //操作者id

           $sql ="update api set num='{$num}',name='{$name}',
           des='{$des}',url='{$url}',type='{$type}',
           parameter='{$parameter}',re='{$re}',lasttime='{$lasttime}',lastuid='{$lastuid}',memo='{$memo}'
           where id = '{$id}'";
           $re = update($sql);
           if($re){
               go(U(array('act'=>'api','tag'=>($_GET['tag'].'#info_api_'.md5($id)))));
           }else{
               echo '<div class="alert alert-danger" role="alert"><span class="glyphicon glyphicon-info-sign" aria-hidden="true"></span> 修改失败</div>';
           }
       }
       //编辑界面
       if(empty($id)){$id = I($_GET['id']);}
       $aid = I($_GET['tag']);
       //得到数据的详情信息start
       $sql = "select * from api where id='{$id}' and aid='{$aid}'";
       $info = find($sql);
       //得到数据的详情信息end
       if(!empty($info)){
           $info['parameter'] = unserialize($info['parameter']);
           $count = count($info['parameter']['name']);
           $p = array();
           for($i = 0;$i < $count; $i++){
               $p[$i]['name']=$info['parameter']['name'][$i];
               $p[$i]['type']=$info['parameter']['type'][$i];
               $p[$i]['default']=$info['parameter']['default'][$i];
               $p[$i]['des']=$info['parameter']['des'][$i];
           }
           $info['parameter'] = $info['parameter'];
       }
   //此分类下的接口列表
   }else{
        $sql = "select api.id,aid,num,url,name,des,parameter,memo,re,lasttime,lastuid,type,login_name
        from api
        left join user
        on api.lastuid=user.id
        where aid='{$_GET['tag']}' and api.isdel=0
        order by ord desc,api.id desc";
        $list = select($sql);
   }
?>
<?php if($op == 'add'){ ?>
    <!--添加接口 start-->
    <div style="border:1px solid #ddd">
        <div style="background:#f5f5f5;padding:20px;position:relative">
            <h4>添加接口<span style="font-size:12px;padding-left:20px;color:#a94442">注:"此色"边框为必填项</span></h4>
            <div style="margin-left:20px;">
                <form action="?act=api&tag=<?php echo $_GET['tag']?>&type=do&op=add" method="post">
                    <h5>基本信息</h5>
                    <div class="form-group has-error">
                        <div class="input-group">
                            <div class="input-group-addon">接口编号</div>
                            <input type="text" class="form-control" name="num" placeholder="接口编号" required="required">
                        </div>
                    </div>
                    <div class="form-group has-error">
                        <div class="input-group">
                            <div class="input-group-addon">接口名称</div>
                            <input type="text" class="form-control" name="name" placeholder="接口名称" required="required">
                        </div>
                    </div>
                    <div class="form-group has-error">
                        <div class="input-group">
                            <div class="input-group-addon">请求地址</div>
                            <input type="text" class="form-control" name="url" placeholder="请求地址" required="required">
                        </div>
                    </div>
                    <div class="form-group">
                        <textarea name="des" class="form-control" placeholder="描述"></textarea>
                    </div>
                    <div class="form-group" required="required">
                        <select class="form-control" name="type">
                            <option value="GET">GET</option>
                            <option value="POST">POST</option>
                        </select>
                    </div>
                    <div class="form-group">
                        <h5>请求参数</h5>
                        <table class="table">
                            <thead>
                            <tr>
                                <th class="col-md-3">参数名</th>
                                <th class="col-md-2">必传</th>
                                <th class="col-md-2">缺省值</th>
                                <th class="col-md-4">描述</th>
                                <th class="col-md-1">
                                    <button type="button" class="btn btn-success" onclick="add()">新增</button>
                                </th>
                            </tr>
                            </thead>
                            <tbody id="parameter">
                            <tr>
                                <td class="form-group has-error">
                                    <input type="text" class="form-control" name="p[name][]" placeholder="参数名" required="required">
                                </td>
                                <td>
                                    <select class="form-control" name="p[type][]">
                                        <option value="Y">Y</option>
                                        <option value="N">N</option>
                                    </select>
                                </td>
                                <td><input type="text" class="form-control" name="p[default][]" placeholder="缺省值"></td>
                                <td><textarea name="p[des][]" rows="1" class="form-control" placeholder="描述"></textarea></td>
                                <td><button type="button" class="btn btn-danger" onclick="del(this)">删除</button></td>
                            </tr>
                            </tbody>
                        </table>
                    </div>
                    <div class="form-group">
                        <h5>返回结果</h5>
                        <textarea name="re" rows="3" class="form-control" placeholder="返回结果"></textarea>
                    </div>
                    <div class="form-group">
                        <h5>备注</h5>
                        <textarea name="memo" rows="3" class="form-control" placeholder="备注"></textarea>
                    </div>
                    <button class="btn btn-success">Submit</button>
                </form>
            </div>
        </div>
    </div>
    <script>
        function add(){
            var $html ='<tr>' +
                '<td class="form-group has-error" ><input type="text" class="form-control has-error" name="p[name][]" placeholder="参数名" required="required"></td>' +
                '<td>' +
                '<select class="form-control" name="p[type][]">' +
                '<option value="Y">Y</option> <option value="N">N</option>' +
                '</select >' +
                '</td>' +
                '<td>' +
                '<input type="text" class="form-control" name="p[default][]" placeholder="缺省值"></td>' +
                '<td>' +
                '<textarea name="p[des][]" rows="1" class="form-control" placeholder="描述"></textarea>' +
                '</td>' +
                '<td>' +
                '<button type="button" class="btn btn-danger" onclick="del(this)">删除</button>' +
                '</td>' +
                '</tr >';
            $('#parameter').append($html);
        }
        function del(obj){
            $(obj).parents('tr').remove();
        }
    </script>
    <!--添加接口 end-->
<?php }else if($op == 'edit'){ ?>
    <!--修改接口 start-->
    <div style="border:1px solid #ddd">
        <div style="background:#f5f5f5;padding:20px;position:relative">
            <h4>修改接口<span style="font-size:12px;padding-left:20px;color:#a94442">注:"此色"边框为必填项</span></h4>
            <div style="margin-left:20px;">
                <form action="?act=api&tag=<?php echo $_GET['tag']?>&type=do&op=edit" method="post">
                    <h5>基本信息</h5>
                    <div class="form-group has-error">
                        <div class="input-group">
                            <div class="input-group-addon">接口编号</div>
                            <input type="hidden" name="id" value="<?php echo $info['id']?>"/>
                            <input type="text" class="form-control" name="num" placeholder="接口编号" value="<?php echo $info['num']?>" required="required">
                        </div>
                    </div>
                    <div class="form-group has-error">
                        <div class="input-group">
                            <div class="input-group-addon">接口名称</div>
                            <input type="text" class="form-control" name="name" placeholder="接口名称" value="<?php echo $info['name']?>" required="required">
                        </div>
                    </div>
                    <div class="form-group has-error">
                        <div class="input-group">
                            <div class="input-group-addon">请求地址</div>
                            <input type="text" class="form-control" name="url" placeholder="请求地址" value="<?php echo $info['url']?>" required="required">
                        </div>
                    </div>
                    <div class="form-group">
                        <textarea name="des" class="form-control" placeholder="描述"><?php echo $info['des']?></textarea>
                    </div>
                    <div class="form-group" required="required">
                        <select class="form-control" name="type">
                            <?php
                                $selected[0] = ($info['type'] == 'GET') ? 'selected' : '';
                                $selected[1] = ($info['type'] == 'POST') ? 'selected' : '';
                            ?>
                            <option value="GET"  <?php echo $selected[0]?>>GET</option>
                            <option value="POST" <?php echo $selected[1]?>>POST</option>
                        </select>
                    </div>
                    <div class="form-group">
                        <h5>请求参数</h5>
                        <table class="table">
                            <thead>
                            <tr>
                                <th class="col-md-3">参数名</th>
                                <th class="col-md-2">必传</th>
                                <th class="col-md-2">缺省值</th>
                                <th class="col-md-4">描述</th>
                                <th class="col-md-1">
                                    <button type="button" class="btn btn-success" onclick="add()">新增</button>
                                </th>
                            </tr>
                            </thead>
                            <tbody id="parameter">

                            <?php $count = count($info['parameter']['name']);?>
                            <?php for($i=0;$i<$count;$i++){ ?>
                            <tr>
                                <td class="form-group has-error">
                                    <input type="text" class="form-control" name="p[name][]" placeholder="参数名" value="<?php echo $info['parameter']['name'][$i]?>" required="required">
                                </td>
                                <td>
                                    <?php
                                        $selected[0] = ($info['parameter']['type'][$i] == 'Y') ? 'selected' : '';
                                        $selected[1] = ($info['parameter']['type'][$i] == 'N') ? 'selected' : '';
                                    ?>
                                    <select class="form-control" name="p[type][]">
                                        <option value="Y" <?php echo $selected[0]?>>Y</option>
                                        <option value="N" <?php echo $selected[1]?>>N</option>
                                    </select>
                                </td>
                                <td><input type="text" class="form-control" name="p[default][]" placeholder="缺省值" value="<?php echo $info['parameter']['default'][$i]?>"></td>
                                <td><textarea name="p[des][]" rows="1" class="form-control" placeholder="描述"><?php echo $info['parameter']['des'][$i]?></textarea></td>
                                <td><button type="button" class="btn btn-danger" onclick="del(this)">删除</button></td>
                            </tr>
                            <?php } ?>

                            </tbody>
                        </table>
                    </div>
                    <div class="form-group">
                        <h5>返回结果</h5>
                        <textarea name="re" rows="3" class="form-control" placeholder="返回结果"><?php echo $info['re']?></textarea>
                    </div>
                    <div class="form-group">
                        <h5>备注</h5>
                        <textarea name="memo" rows="3" class="form-control" placeholder="备注"><?php echo $info['memo']?></textarea>
                    </div>
                    <button class="btn btn-success">Submit</button>
                </form>
            </div>
        </div>
    </div>
    <script>
        function add(){
            var $html ='<tr>' +
                '<td class="form-group has-error" ><input type="text" class="form-control has-error" name="p[name][]" placeholder="参数名" required="required"></td>' +
                '<td>' +
                '<select class="form-control" name="p[type][]">' +
                '<option value="Y">Y</option> <option value="N">N</option>' +
                '</select >' +
                '</td>' +
                '<td>' +
                '<input type="text" class="form-control" name="p[default][]" placeholder="缺省值"></td>' +
                '<td>' +
                '<textarea name="p[des][]" rows="1" class="form-control" placeholder="描述"></textarea>' +
                '</td>' +
                '<td>' +
                '<button type="button" class="btn btn-danger" onclick="del(this)">删除</button>' +
                '</td>' +
                '</tr >';
            $('#parameter').append($html);
        }
        function del(obj){
            $(obj).parents('tr').remove();
        }
    </script>
    <!--修改接口 end-->
<?php }else{ ?>
    <!--接口详细列表start-->
    <?php if(count($list)){ ?>
        <?php foreach($list as $v){ ?>
        <div class="info_api" style="border:1px solid #ddd;margin-bottom:20px;" id="info_api_<?php echo md5($v['id'])?>">
            <div style="background:#f5f5f5;padding:20px;position:relative">
                <div class="textshadow" style="position: absolute;right:0;top:4px;right:8px;">
                    最后修改者: <?php echo $v['login_name']?> &nbsp;<?php echo date('Y-m-d H:i:s',$v['lasttime'])?>&nbsp;
                    <?php if(is_supper()){?>
                    <button class="btn btn-danger btn-xs " onclick="deleteApi(<?php echo $v['id']?>,'<?php echo md5($v['id'])?>')">delete</button>&nbsp;
                    <button class="btn btn-info btn-xs " onclick="editApi('<?php echo U(array('act'=>'api','op'=>'edit','id'=>$v['id'],'tag'=>$_GET['tag']))?>')">edit</button>
                    <?php } ?>
                </div>
                <h4 class="textshadow"><?php echo $v['name']?></h4>
                <p>
                    <b>编号&nbsp;&nbsp;:&nbsp;&nbsp;<span style="color:red"><?php echo $v['num']?></span></b>
                </p>
                <div>
                    <?php
                        $color = 'green';
                        if($v['type']=='POST'){
                            $color = 'red';
                        }
                    ?>
                    <kbd style="color:<?php echo $color?>"><?php echo $v['type']?></kbd> - <kbd><?php echo $v['url']?></kbd>
                </div>
            </div>
            <?php if(!empty($v['des'])){ ?>
            <div class="info">
                <?php echo $v['des']?>
            </div>
            <?php } ?>
            <div style="background:#ffffff;padding:20px;">
                <h5 class="textshadow" >请求参数</h5>
                <table class="table">
                    <thead>
                    <tr>
                        <th class="col-md-3">参数名</th>
                        <th class="col-md-2">必传</th>
                        <th class="col-md-2">缺省值</th>
                        <th class="col-md-5">描述</th>
                    </tr>
                    </thead>
                    <tbody>
                    <?php
                        $parameter = unserialize($v['parameter']);
                        $pnum = count($parameter['name']);
                    ?>
                    <?php for( $i=0; $i<$pnum; $i++ ) {?>
                    <tr>
                        <td><?php echo $parameter['name'][$i]?></td>
                        <td><?php if($parameter['type'][$i]=='Y'){echo '<span style="color:red">Y<span>';}else{echo '<span style="color:green">N<span>';}?></td>
                        <td><?php echo $parameter['default'][$i]?></td>
                        <td><?php echo $parameter['des'][$i]?></td>
                    </tr>
                    <?php } ?>

                    </tbody>
                </table>
            </div>
            <?php if(!empty($v['re'])){ ?>
            <div style="background:#ffffff;padding:20px;">
                <h5 class="textshadow" >返回值</h5>
                <pre><?php echo $v['re']?></pre>
            </div>
            <?php } ?>
            <?php if(!empty($v['memo'])){ ?>
            <div style="background:#ffffff;padding:20px;">
                <h5 class="textshadow">备注</h5>
                <pre style="background:honeydew"><?php echo $v['memo']?></pre>
            </div>
            <?php } ?>
        </div>
        <!--接口详细列表end-->
        <?php } ?>
    <?php } else{ ?>
        <div style="font-size:16px;">
            <span class="glyphicon glyphicon-info-sign" aria-hidden="true"></span> 此分类下还没有任何接口
        </div>
    <?php }?>
    <script>
        //删除某个接口
        var $url = '<?php echo U(array('act'=>'ajax','op'=>'apiDelete'))?>';
        function deleteApi(apiId,divId){
            if(confirm('是否确认删除此接口?')){
                $.post($url,{id:apiId},function(data){
                    if(data == '1'){
                        $('#api_'+divId).remove();//删除左侧菜单
                        $('#info_api_'+divId).remove();//删除接口详情
                    }
                })
            }
        }
        //编辑某个接口
        function editApi(gourl){
            window.location.href=gourl;
        }
    </script>
<?php } ?>
<!--接口详情列表与接口管理end-->
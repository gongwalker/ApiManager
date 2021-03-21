$(function () {
    window.role = localStorage.getItem('USER_ROLE');
    initNanvar();
    initControlBar({export: 0, sort: 0});
    let curUrl = layui.url();
    let module = curUrl.pathname[0] == 'm' ? curUrl.pathname[1] : 'main'
    let search = curUrl.search;
    switch (module) {
        case 'main':
            // 面包屑导航
            render('header-control-nav', 'controlNav', {nav: ' - welcome'});
            render('welcome', 'modelWindow', {})
            getCollectionList();
            break;

        case 'cate' :
            if (search['aid']) {
                getCollectionApi(search['aid'], function (cname) {
                    render('header-control-nav', 'controlNav', {nav: ` - ${cname}`});
                    initControlBar({export: 0, sort: 1});
                });
            }
            break;

        case 'addapi':
            if (search['aid']) {
                getCollectionMenuApi(search['aid']);
                render('addapi', 'modelWindow', {aid: search['aid']});
            }
            break;

        case 'editapi':
            if (search['aid'] && search['id']) {
                getCollectionMenuApi(search['aid']);
                editApi(search['id'])
            }
            break;


        default:
    }


});

// html转义
function html_encode(str) {
    return $('<div>').text(str).html()
}

// html解转义
function html_decode(str) {
    $('<div>').html(str).text()
}

// 模板渲染
function render(tplId, viewId, data, callback) {
    if (data == null) data = []
    let layer
    layui.use('layer', function () {
        layer = layui.layer;
    });
    layui.use('laytpl', function () {
        let tpl = document.getElementById(tplId),
            view = document.getElementById(viewId);
        layui.laytpl(tpl.innerHTML).render(data, function (html) {
            view.innerHTML = html;
        });

        // 解决跳描点问题
        if (callback) {
            callback()
        }
    });


}

// ajax请求
function ajaxReq(dataReq, successFun) {
    let defaultReq = {
        headers: {QuestType: "ajax"},
        method: 'post',
        data: {},
    };
    dataReq = Object.assign({}, defaultReq, dataReq);
    $.ajax({
        headers: dataReq.headers,
        method: dataReq.method,
        data: dataReq.data,
        url: dataReq.url,
        success: function (data) {
            successFun(data);
        },
        error: function (data) {
            layui.use(['layer', 'table', 'element', 'form'], function () {
                if (data.status == 401) {
                    layer.msg(data.responseJSON.msg, {
                        icon: 5
                        , time: 1000,
                    }, function () {
                        location.href = "/login"
                    })
                } else {
                    layer.msg(data.responseJSON.msg, {
                        icon: 5
                    })
                }

            })

        }
    });
}

// 退出
function doExit() {
    //alert(3);
    $.ajax({
        headers: {QuestType: "ajax"},
        method: 'post',
        url: "/login/exit",
        success: function (data) {
            localStorage.removeItem("USER_ID");
            localStorage.removeItem("USER_NAME");
            localStorage.removeItem("USER_ROLE");
            localStorage.removeItem("window_status")
            location.href = "/login"
        }
    });
}

// 新建分类
function CreateCollection() {
    render('addcate', 'modelWindow', {});
}

// 新建分类
function doCreateCollection() {
    let data = {
        cname: $.trim($('#createCollection input[name="cname"]').val()),
        cdesc: $.trim($('#createCollection input[name="cdesc"]').val())
    }

    let req = {
        method: 'post',
        data: data,
        url: "/cate/add",
    }

    ajaxReq(req, function (data) {
        layer.msg(data.msg, {
            icon: 6,
            time: 2000
        }, function () {
            location.href = "/"
        });
    });
}

// 编辑分类界面
function editCollection(aid) {
    let data = {
        url: '/cate/info/' + aid,
        method: 'get',
    }
    ajaxReq(data, function (data) {
        render('editcate', 'modelWindow', data.data);
    });
}

// 编辑分类
function doEditCollection(obj) {
    let that = $(obj).parents('.editCollection');
    let data = {
        aid: that.find('input[name="aid"]').val(),
        cname: $.trim(that.find('input[name="cname"]').val()),
        cdesc: $.trim(that.find('input[name="cdesc"]').val()),
        csort: that.find('input[name="csort"]').val()
    }

    let req = {
        method: 'post',
        data: data,
        url: "/cate/edit",
    }

    ajaxReq(req, function (data) {
        layer.msg(data.msg, {
            icon: 6,
            time: 2000
        }, function () {
            location.href = "/"
        });
    });
}

// 删除分类
function doDeleteCollection(obj) {
    let that = $(obj);
    let aid = that.attr('data-aid');
    let cname = that.attr('data-cname');
    layer.confirm(`${cname} <br>Are you sure delete this collection?`, {
        title: 'Tip',
        btn: ['Yes', 'No']
    }, function () {

        let req = {
            method: 'delete',
            url: "/cate/del/" + aid,
        }
        ajaxReq(req, function (data) {
            layer.msg(data.msg, {
                icon: 6,
                time: 2000
            }, function () {
                location.href = "/"
            });
        });
    });
}

// 分类列表
function getCollectionList() {
    let req = {
        method: 'get',
        url: "/cate/list",
    }
    ajaxReq(req, function (data) {
        render('catelist', 'menuList', data.data);
    });
}

// 某个分类的接口详情
function getCollectionApi(aid, callback) {
    let req = {
        method: 'get',
        url: "/cate/api/" + aid,
    }
    ajaxReq(req, function (data) {
        let cateName = data.data && data.data.info && data.data.info.cname ? data.data.info.cname : '';
        let cateApiList = data.data && data.data.apis ? data.data.apis : []
        callback(cateName)
        // 渲染菜单表
        render('apimenu', 'menuList', cateApiList);
        // 渲染详情列表
        render('apiinfo', 'modelWindow', cateApiList, function () {
            let thisId = window.location.hash;
            if (thisId != "" && thisId != undefined) {
                location.href = thisId;
            }
        });
    });
}

// 某个分类的接品列表
function getCollectionMenuApi(aid) {
    let req = {
        method: 'get',
        url: "/cate/api/" + aid,
    }

    ajaxReq(req, function (data) {
        // 渲染菜单表
        render('apimenu', 'menuList', data.data.apis);
    });
}

// 搜索框搜索事件
function search(obj) {
    let $find = $.trim($(obj).val());
    let $type = $('#menutype').attr('data-type');
    $(".keyword:contains('" + $find + "')")
    if ($find != '') {
        $(".menu").hide();
        if ($type == 'api') {
            $(".info_api").hide();
        }
        let $keywordobj = $(".keyword:contains('" + $find + "')")
        $keywordobj.each(function (i) {
            let o = $($keywordobj[i]);
            let mid = o.attr('nid');
            $("#" + $type + '_' + mid).show();//左侧导航菜单 部份 隐藏
            if ($type == 'api') {
                $("#info_api_" + mid).show();//接口详情 部份 隐藏
            }
        });
    } else {
        $(".menu").show();
        if ($type == 'api') {
            $(".info_api").show();
        }
    }
}

// 新建api
function CreateApi() {
    let curUrl = layui.url();
    let search = curUrl.search;
    let aid = search['aid']
    location.href = "/m/addapi?aid=" + aid;
}

// 添加参数
function addParam() {
    let $html = '<tr>' +
        '<td class="form-group has-error" ><input type="text" class="form-control has-error" name="p[name][]" placeholder="name" requiredgongwen="required"></td>' +
        '<td class="form-group has-error">' +
        '<input type="text" class="form-control" name="p[paramType][]" placeholder="type" requiredgongwen="required"></td>' +
        '<td>' +
        '<select class="form-control" name="p[type][]">' +
        '<option value="Y">Y</option> <option value="N">N</option>' +
        '</select >' +
        '</td>' +
        '<td>' +
        '<input type="text" class="form-control" name="p[default][]" placeholder="default value"></td>' +
        '<td>' +
        '<textarea name="p[des][]" rows="1" class="form-control" placeholder="description"></textarea>' +
        '</td>' +
        '<td>' +
        '<i onclick="sort_api(this,\'up\')" class="layui-icon layui-icon-up op-button"></i>' +
        ' <i onclick="sort_api(this,\'down\')" class="layui-icon layui-icon-down op-button"></i>' +
        ' <i onclick="delParam(this)" class="layui-icon layui-icon-delete op-button op-delete"></i>' +
        '</td>' +
        '</tr >';
    $('#parameter').append($html);
}

// 删除参数
function delParam(obj) {
    $(obj).parents('tr').remove();
}

// 新建api
function doCreateApi() {
    let data = $('#createApiFrom').serialize()
    let req = {
        method: 'post',
        url: "/api/create",
        data: data,
    }
    ajaxReq(req, function (data) {
        layer.msg('success', {icon: 1, time: 2000}, function () {
            let curUrl = layui.url();
            let search = curUrl.search;
            let aid = search['aid']
            location.href = "/m/cate?aid=" + aid;
        })
    });
    return false;
}

// 修改api
function doEditApi() {
    let data = $('#editApiFrom').serialize()
    let req = {
        method: 'post',
        url: "/api/edit",
        data: data,
    }
    ajaxReq(req, function (data) {
        layer.msg('success', {icon: 1, time: 2000}, function () {
            let curUrl = layui.url();
            let search = curUrl.search;
            let aid = search['aid']
            location.href = "/m/cate?aid=" + aid;
        })
    });
    return false;
}

// 修改api界面
function editApi(id) {
    let data = {
        url: '/api/info/' + id,
        method: 'get',
    }
    ajaxReq(data, function (data) {
        render('editapi', 'modelWindow', data.data);
    });
}

// 删除api
function doDeleteApi(obj) {
    let that = $(obj);
    let name = that.attr('data-tit');
    let id = that.attr('data-id');

    layer.confirm(`${name} <br>Are you sure you want to delete?`, {
        title: 'Tip',
        btn: ['Yes', 'No']
    }, function () {
        let req = {
            method: 'delete',
            url: `/api/del/${id}`,
        }
        ajaxReq(req, function (data) {
            layer.msg(data.msg, {
                icon: 6,
                time: 2000
            }, function () {
                location.reload();
            });
        });
    })
}

// 复制api
function duplicateApi(obj) {
    let that = $(obj);
    let id = that.attr('data-id');
    layer.prompt(
        {
            title: 'Please enter a new API name'
            , btn: ['Yes', 'No']
        },
        function (val) {
            let req = {
                method: 'post',
                url: `/api/duplicate/${id}/${val}`,
            }
            ajaxReq(req, function (data) {
                layer.msg(data.msg, {
                    icon: 6,
                    time: 2000
                }, function () {
                    location.reload();
                });
            });
        });
}

function showParamsTable(params) {
    params = JSON.parse(params);
    return params.param_name ? params.param_name.length : 0;
}

function getApiParamsView(params) {
    params = JSON.parse(params);
    let str = '';
    let l = params.param_name ? params.param_name.length : 0;
    let color = ''
    for (let i = 0; i < l; i++) {
        color = '#5FB878';
        params.param_name[i] = html_encode(params.param_name[i]);
        params.param_type[i] = html_encode(params.param_type[i]);
        params.param_default[i] = html_encode(params.param_default[i]);
        params.param_des[i] = html_encode(params.param_des[i]);

        if (params.param_cate[i] == 'Y') {
            color = '#FF5722';
        }

        str += `
        <tr>
        <td>${params.param_name[i]}</td>
        <td>${params.param_type[i]}</td>
        <td><span style="color:${color}">${params.param_cate[i]}<span></td>
        <td>${params.param_default[i]}</td>
        <td>${params.param_des[i]}</td>
        </tr>
        `
    }
    return str;
}

function getApiParamsViewForEdit(params) {
    params = JSON.parse(params);
    let str = '';
    let l = params.param_name ? params.param_name.length : 0;
    let cateSelected = {Y: '', N: ''};
    for (let i = 0; i < l; i++) {

        cateSelected[params.param_cate[i]] = "selected";

        str += `
        <tr>
        <td>
            <input value="${params.param_name[i]}" type="text" class="form-control" name="p[name][]" placeholder="name" requiredgongwen="required">
        </td>
        <td>
            <input value="${params.param_type[i]}" type="text" class="form-control" name="p[paramType][]" placeholder="type" requiredgongwen="required">
        </td>
        <td>
            <select class="form-control" name="p[type][]">
                <option ${cateSelected['Y']} value="Y">Y</option>
                <option ${cateSelected['N']} value="N">N</option>
            </select>
        </td>
        <td>
            <input value="${params.param_default[i]}" type="text" class="form-control" name="p[default][]" placeholder="default value">
        </td>
        <td>
            <textarea name="p[des][]" rows="1" class="form-control" placeholder="description">${params.param_des[i]}</textarea></td>
        </td>
        <td>
            <i onclick="sort_api(this,'up')" class="layui-icon layui-icon-up op-button"></i>
            <i onclick="sort_api(this,'down')" class="layui-icon layui-icon-down op-button"></i>
            <i onclick="delParam(this)" class="layui-icon layui-icon-delete op-button op-delete"></i>
        </td>
        </tr>
        `
    }
    return str;
}

function getSelectedType(type) {
    let types = [
        'GET',
        'POST',
        'PUT',
        'PATCH',
        'DELETE',
        'HEAD',
        'OPTIONS',
        'UNLINK',
        'PURGE',
        'UNLOCK',
        'PROPFIND',
        'VIEW',
    ]
    let l = types.length;
    let s = '';
    let e = '';
    for (let i = 0; i < l; i++) {
        e = '';
        if (type == types[i]) {
            e = 'selected';
        }
        s += `<option ${e} value="${types[i]}">${types[i]}</option>`;
    }
    return s;
}


// 返回至顶部
function goTop() {
    $('#mainWindow').animate(
        {scrollTop: '0px'}, 200
    );
}

// 全屏和normal
function navbar() {
    let obj = '#navbarhanlder';
    if ($('#mainWindow').hasClass('col-md-9')) {
        $(obj).html('&gt;');
        $(obj).css("cursor", "e-resize");
        $('#mainWindow').removeClass('col-md-9').addClass('col-md-12');
        $('#navbar').hide();
        localStorage.setItem('window_status', '1');
    } else {
        $(obj).html('&lt;');
        $(obj).css("cursor", "w-resize");
        $('#mainWindow').removeClass('col-md-12').addClass('col-md-9');
        $('#navbar').show();
        localStorage.setItem('window_status', '0');
    }
}

function initNanvar() {
    let obj = '#navbarhanlder';
    if (localStorage.getItem('window_status') == 1) {
        $(obj).html('&gt;');
        $(obj).css("cursor", "e-resize");
        $('#mainWindow').removeClass('col-md-9').addClass('col-md-12');
        $('#navbar').hide();
    } else {
        $(obj).html('&lt;');
        $(obj).css("cursor", "w-resize");
        $('#mainWindow').removeClass('col-md-12').addClass('col-md-9');
        $('#navbar').show();
    }
}

function initControlBar(con) {
    let data = {
        name: localStorage.getItem('USER_NAME'),
        id: localStorage.getItem('USER_ID'),
        export: con.export,
        sort: con.sort,
    }
    render('header-control-bar', 'controlBar', data);
}

function boolRole(condition) {
    return condition;
}

// 导出 todo
function doExport() {
}

// 排序
function doSort() {
    let apis = $('#menuList .api-term');
    let apis_html = '';
    $(apis).each(function (i, item) {
        let that = $(item)
        let name = that.attr('data-cname');
        let num = that.attr('data-num');
        let id = that.attr('data-id');
        apis_html +=
            `<tr id="sort_api_${id}" class="sort_api" data-id="${id}">
                  <td><kbd style="">${num}</kbd></td>
                  <td>${name}</td>
                  <td class="sort_op_container"> 
                      <i onclick="sort_api(this,'up')" class="layui-icon layui-icon-up" ></i>&nbsp;&nbsp;
                      <i onclick="sort_api(this,'down')" class="layui-icon layui-icon-down"></i>
                  </td>
            </tr>`;
    })

    if (apis_html == "") {
        apis_html = `<tr ><td colspan="3"> <i class="layui-icon layui-icon-face-surprised"></i>   No data for sorting !</td></tr>`;
    }

    let html = `
<div id="sort-container">  
    <div style="padding:0 10px;">
    <table class="layui-table" lay-size="sm" >
      <colgroup>
        <col width="100">
        <col>
        <col width="46">
      </colgroup>
      <thead>
        <tr>
          <th>Api No</th>
          <th>Name</th>
          <th>Operation</th>
        </tr> 
      </thead>
      <tbody>
        ${apis_html}
      </tbody>
    </table>
    </div>
</div>
`;
    layui.use(['layer', 'table', 'element', 'form'], function () {
        layer.open({
            skin: 'manage-class',
            title: 'Sort',
            type: 1,
            resize: false,
            area: ['856px', '500px'],
            btn: ['submit', 'cancel'],
            yes: function () {
                let trs = $('#sort-container .sort_api');
                let ids = [];
                $(trs).each(function (i, obj) {
                    ids.push($(obj).data('id'));
                })

                let data = {
                    url: '/api/sort',
                    method: 'post',
                    data: {ids: ids}
                }

                ajaxReq(data, function (data) {
                    layer.msg('success', {icon: 1, time: 1000}, function () {
                        layer.closeAll();
                    });
                });


            },
            content: html,
            success: function () {
            }
        });
    });
}

function doSet() {
    let html = `
<div id="manager-container">
    <div class="layui-tab layui-tab-card">
      <ul class="layui-tab-title">
        <li class="layui-this">User Management</li>
      </ul>
      <div class="layui-tab-content">
        <div class="layui-tab-item layui-show">
            <!-- search begin -->
            <div id="search-container" class="layui-form" style="margin-top:5px;">
                <div class="layui-input-inline" style="width:300px;">
                  <input type="text" name="login_name"  placeholder="login_name" class="layui-input">
                </div>
                 
                <div class="layui-input-inline">
                    <select name="role">
                        <option value="">role</option>
                        <option value="1">Super Administrator</option>
                        <option value="2">Administrator</option>
                        <option value="3">Guest</option>
                    </select>
                </div>
                
                <div class="layui-input-inline">
                    <button type="button" style="height:30px;line-height: 30px;" class="layui-btn" onclick="SearchUser()">Search</button>
                </div>
                
                <div class="layui-input-inline" style="float:right">
                    <button type="button" style="height:30px;line-height: 30px;" class="layui-btn layui-btn-normal" onclick="CreateUser()">Create</button>
                </div>
            </div>
            <!-- search end -->
            <table id="user_table" lay-filter="user_table"></table>
        </div>
        <div class="layui-tab-item">2</div>
      </div>
    </div>
</div>
`;
    layui.use(['layer', 'table', 'element', 'form'], function () {
        let table = layui.table;
        let form = layui.form;
        layer.open({
            skin: 'manage-class',
            title: 'Manage',
            type: 1,
            resize: false,
            area: ['800px', '500px'],
            content: html,
            success: function () {
                form.render('select');
                table.render({
                    elem: '#user_table'
                    , height: 370
                    , url: '/user/list'
                    , parseData: function (res) {
                        return {
                            "code": res.code,
                            "msg": res.msg,
                            "count": res.data.count,
                            "data": res.data.list
                        }
                    }
                    , where: getSearchWhere()
                    , response: {
                        statusCode: 200
                    }
                    , page: true
                    , cols: [[
                        {field: 'id', title: 'ID', width: 80, sort: true}
                        , {field: 'login_name', title: 'Login Name'}
                        , {field: 'role', title: 'Role', sort: true, width: 160, templet: '#role_tpl'}
                        , {field: 'isdel', title: 'Available Status', width: 134, templet: '#status_tpl'}
                        , {field: 'isdel', title: 'Operation', templet: '#opration_tpl', width: 196}
                    ]]
                });
            }
        });
    });
}

function getRole(d) {
    if (d.role == 1) {
        return 'Super Administrator';
    } else if (d.role == 2) {
        return 'Administrator';
    } else if (d.role == 3) {
        return 'Guest';
    }
    return 'Unknown';
}

function getStatus(d) {
    if (d.isdel == 1) {
        return 'Disabled';
    } else if (d.isdel == 0) {
        return 'Enable';
    }
    return 'Unknown';
}

function getButton(d) {
    let html = '';
    if (d.isdel == 0) {
        html += `<button data-id="${d.id}" data-login_name="${d.login_name}" data-role="${d.role}" data-to-status="1" onclick="changeUserStatus(this)" style="width:50px;" class="layui-btn layui-btn-xs layui-btn-danger">Disable</button>`;
    } else {
        html += `<button data-id="${d.id}" data-login_name="${d.login_name}" data-role="${d.role}" data-to-status="0" onclick="changeUserStatus(this)" style="width:50px;" class="layui-btn layui-btn-xs layui-btn-base">Enable</button>`;
    }
    html += `<button data-id="${d.id}" data-login_name="${d.login_name}" data-role="${d.role}" onclick="editUser(this)" class="layui-btn layui-btn-xs layui-btn-normal">Edit</button>`;

    html += `<button data-id="${d.id}" data-login_name="${d.login_name}" onclick="resetPwd(this)" class="layui-btn layui-btn-xs layui-btn-warm">Reset Pwd</button>`;
    return html;
}

// 搜索用户
function SearchUser() {
    layui.use(['layer', 'table', 'element', 'form'], function () {
        let table = layui.table;
        table.reload('user_table', {
            url: '/user/list'
            , where: getSearchWhere()
            , page: {
                curr: 1
            }
        });
    })
}

function getSearchWhere() {
    let login_name = $.trim($('#search-container input[name=login_name]').val());
    let role = $.trim($('#search-container select[name=role]').val());
    return {login_name: login_name, role: role}
}

// 新增用户
function CreateUser() {

    let html = `
        <div id="create-user" class="layui-form" style="padding:10px;">
  
            <div class="layui-form-item">
                <div class="layui-input-inline" style="width:300px;">
                     <input type="text" name="login_name" placeholder="login_name" class="layui-input">
                </div>
            </div>
            
            <div class="layui-form-item">    
                <div class="layui-input-inline" style="width:300px;">
                    <input type="text" name="password" placeholder="password" class="layui-input">
                </div>
           </div>
           
           <div class="layui-form-item">    
                <div class="layui-input-inline" style="width:480px;">
                <input type="radio" name="role" value="1" title="Super Administrator">
                <input type="radio" name="role" value="2" title="Administrator">
                <input type="radio" name="role" value="3" title="Guest">
                </div>
            </div>
              
        </div>
    `;
    layer.closeAll();
    layui.use(['layer', 'table', 'element', 'form'], function () {
        let form = layui.form;
        layer.open({
            skin: 'manage-class'
            , title: 'Create Account'
            , type: 1
            , resize: false
            , content: html
            , btn: ['submit', 'cancel']
            , area: ['438px', '240px']
            , yes: function (index, layero) {
                let login_name = $.trim($('#create-user input[name=login_name]').val());
                let password = $.trim($('#create-user input[name=password]').val());
                let role = $.trim($('#create-user input[name=role]:checked').val());
                if (!role) {
                    layer.msg('please select a role for the account', {icon: 5});
                    return;
                }

                let uPattern = /^[a-zA-Z0-9_-]{2,20}$/; //用户名正则，3到20位（字母，数字，下划线，减号）
                if (!uPattern.test(login_name)) {
                    layer.msg('login name 2 to 20 digits (letters, numbers, underscores, minus)', {icon: 5});
                    return;
                }

                let uPwd = /^[a-zA-Z0-9_-]{6,16}$/; //密码正则，6到16位（字母，数字，下划线，减号）
                if (!uPwd.test(password)) {
                    layer.msg('password 6 to 16 digits (letters, numbers, underscores, minus)', {icon: 5});
                    return;
                }

                let data = {
                    url: '/user/add',
                    method: 'post',
                    data: {login_name: login_name, password: password, role: role}
                }

                ajaxReq(data, function (data) {
                    layer.msg('success', {icon: 1, time: 1000}, function () {
                        layer.closeAll();
                    });
                });
            }
            , success: function () {
                form.render('radio');
            }
        });
    })
}

// 禁用启用用户
function changeUserStatus(obj) {
    let that = $(obj)
    let id = that.attr('data-id')
    let isdel = that.attr('data-to-status')
    let login_name = that.attr('data-login_name')
    let role = that.attr('data-role')

    let data = {
        url: '/user/changeStatus',
        method: 'post',
        data: {id: id, isdel: isdel}
    }
    ajaxReq(data, function (data) {
        $('#user_status_' + id).html(getStatus({isdel: isdel, id: id}));
        $('#user_button_' + id).html(getButton({isdel: isdel, id: id, role: role, login_name: login_name}));
    });
}

// 编辑用户
function editUser(obj) {
    let that = $(obj)
    let id = that.attr('data-id')
    let login_name = that.attr('data-login_name');
    let role = that.attr('data-role');
    let selected = {1: '', 2: '', 3: ''};
    selected[role] = 'checked';

    let html = `
        <div id="create-user" class="layui-form" style="padding:10px;">
  
            <div class="layui-form-item">
                <div class="layui-input-inline" style="width:300px;">
                     <input type="text" name="login_name" value="${login_name}" placeholder="login_name" class="layui-input">
                </div>
            </div>
            
           <div class="layui-form-item">    
                <div class="layui-input-inline" style="width:480px;">
                <input type="radio" name="role" value="1" ${selected[1]} title="Super Administrator">
                <input type="radio" name="role" value="2" ${selected[2]} title="Administrator">
                <input type="radio" name="role" value="3" ${selected[3]} title="Guest">
                </div>
            </div>
              
        </div>
    `;

    layer.closeAll();
    layui.use(['layer', 'table', 'element', 'form'], function () {
        let form = layui.form;
        layer.open({
            skin: 'manage-class'
            , title: 'Edit Account'
            , type: 1
            , resize: false
            , content: html
            , btn: ['submit', 'cancel']
            , area: ['438px', '200px']
            , yes: function (index, layero) {
                let login_name = $.trim($('#create-user input[name=login_name]').val());
                let role = $.trim($('#create-user input[name=role]:checked').val());
                if (!role) {
                    layer.msg('please select a role for the account', {icon: 5});
                    return;
                }

                let uPattern = /^[a-zA-Z0-9_-]{2,20}$/; //用户名正则，3到20位（字母，数字，下划线，减号）
                if (!uPattern.test(login_name)) {
                    layer.msg('login name 3 to 20 digits (letters, numbers, underscores, minus)', {icon: 5});
                    return;
                }

                let data = {
                    url: '/user/edit',
                    method: 'post',
                    data: {login_name: login_name, id: id, role: role}
                }

                ajaxReq(data, function (data) {
                    layer.msg('success', {icon: 1, time: 1000}, function () {
                        layer.closeAll();
                    });
                });
            }
            , success: function () {
                form.render('radio');
            }
        });
    })
}

// 重置密码
function resetPwd(obj) {
    let that = $(obj)
    let id = that.attr('data-id')
    let login_name = that.attr('data-login_name');
    let role = that.attr('data-role');
    let selected = {1: '', 2: '', 3: ''};
    selected[role] = 'checked';

    let html = `
        <div id="create-user" class="layui-form" style="padding:10px;">
  
            <div class="layui-form-item">
                <div class="layui-input-inline" style="width:300px;">
                     <input type="text" readonly name="login_name" value="${login_name}" placeholder="login_name" class="layui-input">
                </div>
            </div>
            
           <div class="layui-form-item">    
                <div class="layui-input-inline" style="width:300px;">
                    <input type="text" name="password" placeholder="password" class="layui-input">
                </div>
           </div>
              
        </div>
    `;

    layer.closeAll();
    layui.use(['layer', 'table', 'element', 'form'], function () {
        let form = layui.form;
        layer.open({
            skin: 'manage-class'
            , title: 'Reset Password'
            , type: 1
            , resize: false
            , content: html
            , btn: ['submit', 'cancel']
            , area: ['330px', '190px']
            , yes: function (index, layero) {
                let login_name = $.trim($('#create-user input[name=login_name]').val());
                let password = $.trim($('#create-user input[name=password]').val());
                let uPwd = /^[a-zA-Z0-9_-]{6,16}$/; //密码正则，6到16位（字母，数字，下划线，减号）
                if (!uPwd.test(password)) {
                    //layer.msg('password 6 to 16 digits (letters, numbers, underscores, minus)', {icon: 5});
                    //return;
                }

                let data = {
                    url: '/user/resetpwd',
                    method: 'post',
                    data: {id: id, password: password}
                }

                ajaxReq(data, function (data) {
                    layer.msg('success', {icon: 1, time: 1000}, function () {
                        layer.closeAll();
                    });
                });
            }
            , success: function () {
                form.render('radio');
            }
        });
    })
}

// 排序
function sort_api(obj, type) {
    let current = $(obj).parent().parent(); //获取当前<tr>
    switch (type) {
        case 'up':
            let prev = current.prev();  //获取当前<tr>前一个元素
            if (current.index() > 0) {
                current.insertBefore(prev); //插入到当前<tr>前一个元素前
                current.hide(0).fadeIn(666)
            } else {
                layer.msg("It's at the top", {icon: 5});
            }
            break;
        case 'down':
            let next = current.next(); //获取当前<tr>后面一个元素
            if (next.index() > 0) {
                current.insertAfter(next);//插入到当前<tr>后面一个元素后面
                current.hide(0).fadeIn(666)
            } else {
                layer.msg("It's at the bottom", {icon: 5});
            }
            break;
    }
}




<?php /* Smarty version Smarty-3.1.8, created on 2019-06-01 11:55:23
         compiled from "/root/devspace/gowork/src/github.com/longlong202/desertstar/application/views/qixi/love/lovejessie.tpl" */ ?>
<?php /*%%SmartyHeaderCode:13534934795cf1f72b33db76-47327722%%*/if(!defined('SMARTY_DIR')) exit('no direct access allowed');
$_valid = $_smarty_tpl->decodeProperties(array (
  'file_dependency' => 
  array (
    'd096e72c1c1d2e246772fc9dea3f0d2b01ca7950' => 
    array (
      0 => '/root/devspace/gowork/src/github.com/longlong202/desertstar/application/views/qixi/love/lovejessie.tpl',
      1 => 1534585915,
      2 => 'file',
    ),
  ),
  'nocache_hash' => '13534934795cf1f72b33db76-47327722',
  'function' => 
  array (
  ),
  'has_nocache_code' => false,
  'version' => 'Smarty-3.1.8',
  'unifunc' => 'content_5cf1f72b33ec34_81543440',
),false); /*/%%SmartyHeaderCode%%*/?>
<?php if ($_valid && !is_callable('content_5cf1f72b33ec34_81543440')) {function content_5cf1f72b33ec34_81543440($_smarty_tpl) {?><!DOCTYPE HTML>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, user-scalable=no, minimum-scale=1.0, maximum-scale=1.0">
<title>抽奖咯</title>
<style type="text/css">
.demo{
    position:absolute;
    left:50%;
    top:50%;
    transform:translate(-50%,-50%);
}
#disk{
    width:auto; 
    height:auto; 
}
#start{
    width:163px; 
    height:320px; 
    position:absolute; 
    left:50%;
    top:50%;
    transform:translate(-50%,-50%);
}
#start img{
    cursor:pointer
}
</style>
<script type="text/javascript" src="/js/qixi/jquery.min.js"></script>
<script type="text/javascript" src="/js/qixi/jQueryRotate.2.2.js"></script>
<script type="text/javascript" src="/js/qixi/jquery.easing.min.js"></script>
<script type="text/javascript">
$(function(){
	$("#startbtn").rotate({
		bind:{
			click:function(){
				var a = Math.floor(Math.random() * 360);
				 $(this).rotate({
					 	duration:3000,
					 	angle: 0, 
            			animateTo:1440+a,
						easing: $.easing.easeOutSine,
						callback: function(){
							alert('恭喜Jessie，七夕中奖了！');
						}
				 });
			}
		}
	});
});
</script>
</head>

<body>

<div id="main">
   <div class="demo">
        <div id="disk"><img src="/img/qixi/disk2.jpg" id="diskbtn"></div>
        <div id="start"><img src="/img/qixi/start.png" id="startbtn"></div>
   </div>
</div>


</body>
</html>
<?php }} ?>
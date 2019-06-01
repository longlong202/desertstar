<?php /* Smarty version Smarty-3.1.8, created on 2019-06-01 15:33:26
         compiled from "/root/devspace/gowork/src/github.com/longlong202/desertstar/application/views/qr/genqr/index.tpl" */ ?>
<?php /*%%SmartyHeaderCode:1783064265cf22893a7f984-87636993%%*/if(!defined('SMARTY_DIR')) exit('no direct access allowed');
$_valid = $_smarty_tpl->decodeProperties(array (
  'file_dependency' => 
  array (
    'dee59e15049953800474599071f4b4ad0b472464' => 
    array (
      0 => '/root/devspace/gowork/src/github.com/longlong202/desertstar/application/views/qr/genqr/index.tpl',
      1 => 1559374390,
      2 => 'file',
    ),
  ),
  'nocache_hash' => '1783064265cf22893a7f984-87636993',
  'function' => 
  array (
  ),
  'version' => 'Smarty-3.1.8',
  'unifunc' => 'content_5cf22893a82ad3_48765845',
  'variables' => 
  array (
    'pic' => 0,
  ),
  'has_nocache_code' => false,
),false); /*/%%SmartyHeaderCode%%*/?>
<?php if ($_valid && !is_callable('content_5cf22893a82ad3_48765845')) {function content_5cf22893a82ad3_48765845($_smarty_tpl) {?><html>
<head>
    <title>我的二维码</title>
    <meta name="viewport" content="width=device-width, user-scalable=no, minimum-scale=1.0, maximum-scale=1.0">
    <link rel="stylesheet" type="text/css" href="/style/qrcode/main.css">
</head>
<body>
    <div class='main' align='center' position='absolute'><img src="<?php echo htmlspecialchars($_smarty_tpl->tpl_vars['pic']->value, ENT_QUOTES, 'UTF-8');?>
" alt="欢迎Jessie"></div>
</body>
</html>
<?php }} ?>
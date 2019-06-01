<?php
//定义常量
#define('ROOTPATH',realpath(dirname(__DIR__)));
define('ROOTPATH',dirname(dirname(__FILE__)));
define('APPPATH',ROOTPATH . '/application');
define('WWWPATH',ROOTPATH . '/public');
define('TPLPATH',APPPATH . '/views');
define('MODULEPATH',APPPATH . '/modules');
define('SDKPATH', '/root/devspace/');

//xhprof_enable();//开始

//自动加载类
ini_set('yaf.use_spl_autoload', 1);
error_reporting(E_ALL & ~E_STRICT & ~E_NOTICE & ~E_WARNING);//忽略php报告
$app = new Yaf_Application(ROOTPATH."/conf/application.ini");
define('ENV', ini_get("yaf.environ"));
$app->bootstrap()->run();

/*
$xhprof_data = xhprof_disable();

//print_r($xhprof_data);

include_once "xhprof_lib.php";
include_once "xhprof_runs.php";

$xhprof_runs = new XHProfRuns_Default();
$run_id = $xhprof_runs->save_run($xhprof_data,"xhprof_foo");
*/

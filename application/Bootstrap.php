<?php

/**
 *yaf框架全局配置入口
 *@author longlong
 *@version 2018-08-04
 */

#class Bootstrap extends Yaf\Bootstrap_Abstract
class Bootstrap extends Yaf_Bootstrap_Abstract
{/*{{{*/
    //加载类包，以_init开头的方法都会被自动加载
    public function _initLoader(Yaf_Dispatcher $dispatcher)
    {/*{{{*/

        //前端tpl相关
        //Yaf\Loader::import(APPPATH . '/library/Front/*');
        $add = SDKPATH . 'jacob_bus/autoload.php';//TODO-为什么SDKPATH不对！！！
//        echo $add."\n";
        $add2 ="/home/q/php/jacob_bus/autoload.php";
        Yaf_Loader::import($add2);
//        Yaf_Loader::import(SDKPATH . 'jacob_bus/autoload.php');

    }/*}}}*/

    //初始化配置文件
    public function _initConfig(Yaf_Dispatcher $dispatcher)
    {/*{{{*/
        Yaf_Registry::set("config",Yaf_Application::app()->getConfig());
    }/*}}}*/

    //初始化日志 Yaf\Dispatcher $dispatcher
    public function _initLog(Yaf_Dispatcher $dispatcher)
    {/*{{{*/
        Common\Logger::init(\Utils\Conf::getConf('common.log.name'),\Utils\Conf::getConf('common.log.path'),\Utils\Conf::getConf('common.log.option'));
        //$var_dump($config);
    }/*}}}*/

    //初始化smarty模板
    public function _initSmarty(Yaf_Dispatcher $dispatcher)
    {/*{{{*/
        $config = Yaf_Registry::get("config");
        //var_dump($config);
        $smarty = new Smarty_Adapter(null,$config["smarty"]);
        $smarty->enableEscape();
        $dispatcher->setView($smarty);

        //关闭自动渲染，手动调用display
        $dispatcher->disableView();
    }/*}}}*/

}/*}}}*/

?>

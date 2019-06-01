<?php
/**
 * Created by PhpStorm.
 * User: longlong
 * Date: 2018/9/5
 * Time: 11:36
 */

namespace Common;
use Utils\Conf;
use \Jacob\Utils\Lib\VUtilsLog;
use \Jacob\Utils\Lib\VUtilsLogFile;
use Component\Point\Performance;
use Component\Point\Config;

class Logger
{
    private static $_logName = '';
    private static $_logPath = '';
    private $_msg = '';
    static protected $instance = null;
    protected $logger = null;

    private $_logName2 = null;
    private $_logPath2 = null;
    private $_writer = null;

    static public function init($logName,$logPath,$option = array('logLevel'=>15,'buffering'=>0))
    {
        self::$_logName = !empty($logName) ? $logName : Conf::getConf('common.log.name');
        self::$_logPath = !empty($logPath) ? $logPath : Conf::getConf('common.log.path');
        VUtilsLog::reg($logName,$logPath,$option);
        VUtilsLog::setBizname('toffee');
        self::$instance = new Logger();
        self::$instance->logger = VUtilsLog::ins($logName);//logger对象
        if(strlen($_SERVER['REQUEST_UUID']) >= 16) {
            VUtilsLog::setLogId($_SERVER['REQUEST_UUID']);
        }

        if (self::$instance->_writer == null)
        {
            self::$instance->_writer = new VUtilsLogFile();
        }
        self::$instance->_logPath2 = self::$_logPath;
        self::$instance->_logName2 = self::$_logName;

        //point init
        list($path, $query) = explode('?', $_SERVER['REQUEST_URI']);
        Performance::init();
        $straceFile = Conf::getConf('common.log.path') . '/point.log';
        Config::setLogFile($straceFile);
        //TODO 追踪开关
//        Config::disablePoint();
//        Config::setValue('uuid', (string)$_SERVER['REQUEST_UUID']);
//        Config::setValue('uidshow', (string)$_REQUEST['uidshow']);
//        Config::setValue('mid', (string)$_REQUEST['mid']);
        Config::setValue('path', (string)$path);
        Config::setValue('reqtime', date("Y-m-d H:i:s"));
    }


    private function __construct(){}


    static public function getLogger()
    {
        return self::$instance;
    }

    public function error($msg)
    {
        $this->_msg = $msg;
        self::$instance->logger->error($msg);
        return $this;
    }
    public function biz($msg)
    {
        $this->_msg = $msg;
        self::$instance->logger->biz($msg);
        return $this;
    }

    public function debug($msg)
    {
        $this->_msg = $msg;
        self::$instance->logger->debug($msg);
        return $this;
    }

    public function __call($name, $arguments)
    {
        return call_user_func(array($this->logger, $name), $arguments);
    }
}

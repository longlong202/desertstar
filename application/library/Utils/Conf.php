<?php
namespace Utils;

class Conf
{
    private static $configerGlobal = array(); //全局配置
    private static $configerProduct = array(); //模块配置
    protected static $configerGlobalDir = '';
    protected static $configerProductDir = '';

    static public function init($product)
    {
        if ($product){
            self::$configerProductDir = ROOTPATH . '/conf/'.$product. '/';
        }

        if (!is_dir(self::$configerProductDir)){
            throw new \Exception(self::$configerProductDir . "is not dir");
        }
    }

    public static function getConf($key = '')
    {
        if (empty($key)){
            return $key;
        }

        if (strpos($key,'.') !== false){
            list($fileName,$path) = explode('.',$key,2);
        }else{
            $fileName = $key;
        }

        if(!isset(self::$configerGlobal[$fileName])) {
            self::$configerGlobal[$fileName] = new \Yaf_Config_Ini(ROOTPATH . '/conf/'. $fileName . '.ini', ENV);
        }

        if (isset($path)){
            $ret = self::$configerGlobal[$fileName]->get($path);//这样也可以获取数组的值？
            //is_a()如果对象属于该类或者该类是此对象的父类则返回true
            if (is_a($ret,'Yaf_Config_Ini')){
                return $ret->toArray();
            }else{
                return $ret;
            }
        }else{
            return self::$configerGlobal[$fileName]->toArray();
        }
    }

    //获取产品个性配置
    public static function getProductConf($key='')
    {
        if (empty($key)){
            return $key;
        }

        if(strpos($key, '.') !== false) {
            list($fileName, $path) = explode('.', $key, 2);
        } else {
            $fileName = $key;
        }

        if(!isset(self::$configerProduct[$fileName])) {
            self::$configerProduct[$fileName] = new \Yaf_Config_Ini(self::$configerProductDir . $fileName . '.ini', ENV);
        }

        if(isset($path)) {
            $ret = self::$configerProduct[$fileName]->get($path);
            if(is_a($ret, 'Yaf_Config_Ini')) {
                return $ret->toArray();
            } else {
                return $ret;
            }
        } else {
            return self::$configerProduct[$fileName]->toArray();
        }
    }

}
?>
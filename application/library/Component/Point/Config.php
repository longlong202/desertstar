<?php 

namespace Component\Point;

class Config
{
    /*
     * Create a config instance
     */
    private static $config;

    private static function instance()
    {
        if (!self::$config)
            self::$config = Performance::instance()->config;
        return self::$config;
    }

    private static function enableTool()
    {
        $config = self::instance();

        // Check DISABLE_TOOL
        if( ! $config->isEnableTool())
            return false;

        return true;
    }
    public static function disablePoint()
    {
        self::instance()->disableTool();
    }

    /*
     * Set config item point label LTrim
     * @param bool $mask
     * return void
     */
    public static function setPointLabelLTrim($mask)
    {
        if( ! self::enableTool())
            return;

        self::$config->setPointLabelLTrim($mask);
    }

    /*
     * Set config item point label RTrim
     * @param bool $mask
     * return void
     */
    public static function setPointLabelRTrim($mask)
    {
        if( ! self::enableTool())
            return;

        self::$config->setPointLabelRTrim($mask);
    }

    
    /*
     * Set config item logfile
     * @param mixed $value
     * @param string $viewType
     * return void
     */
    public static function setLogFile($logFile)
    {
        if( ! self::enableTool())
            return;

        self::$config->setLogFile($logFile);
    }
    
    /*
     * Set config item presenter
     * @param mixed $value
     * return void
     */
    public static function setPresenter($type)
    {
        if( ! self::enableTool())
            return;

        self::$config->setPresenter($type);
    }

    /*
     * Set config point label nice
     * @param bool $status
     * return void
     */
    public static function setPointLabelNice($status)
    {
        if( ! self::enableTool())
            return;

        self::$config->setPointLabelNice($status);
    }

    /*
     * Reset
     */
    public static function instanceReset()
    {
        self::$config = null;
    }

    public static function setValue($key, $val)
    {
        if( ! self::enableTool())
            return;
        self::$config->$key = $val;
    }

}
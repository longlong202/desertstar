<?php 

namespace Component\Point;

use Component\Point\Handlers\PerformanceHandler;

class Performance
{
    /*
     * Create a performance instance
     */
    private static $performance;
    private static $bootstrap = false;

    public static function init()
    {
        self::$performance = new PerformanceHandler();
        Config::instanceReset();
    }
    
    public static function instance()
    {
        if( ! self::$performance)
            self::$performance = new PerformanceHandler();
        return self::$performance;
    }

    private static function enableTool()
    {
        $performance = self::instance();

        // Check DISABLE_TOOL
        if( ! $performance->config->isEnableTool())
            return false;

        // Check bootstrap
        if( ! self::$bootstrap)
        {
            $performance->bootstrap();
            self::$bootstrap = true;
        }

        return true;
    }

    /*
     * Set measuring point X
     *
     * @param string|null   $label
     * @return void
     */
    public static function point($label = null, $line = 0, $file = '', $msg = '')
    {
        if( ! self::enableTool() )
            return;
        $label = $label.microtime();
        // Run
        self::$performance->point($label, $line, $file, $msg);
    }

    /*
     * Set a message associated with the config
     *
     * @param string|null   $message
     * @return void
     */
    public static function message($message = null)
    {
        if( ! self::enableTool() or ! $message)
            return;
        // Run
        self::$performance->message($message);
    }


    /*
     * Finish measuring point X
     *
     * @param string|null   $label
     * @return void
     */
    public static function finish()
    {
        if( ! self::enableTool() )
            return;

        // Run
        self::$performance->finish();
    }

    /*
     * Export helper
     *
     * @return Component\Handlers\ExportHandler
     */
    public static function export()
    {
        if( ! self::enableTool() )
            return;

        // Run
        return self::$performance->export();
    }

    /*
     * Return test results
     *
     * @return Component\Handlers\ExportHandler
     */
    public static function results()
    {
        if( ! self::enableTool() )
            return;

        // Run
        return self::$performance->results();
    }

    /*
     * Reset
     */
    public static function instanceReset()
    {
        // Run
        Config::instanceReset();
        self::$performance = null;
        self::$bootstrap = false;

    }
}
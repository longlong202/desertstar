<?php 

namespace Component\Point;

use Component\Point\Handlers\ConfigHandler;
use Component\Point\Interfaces\ExportInterface;

class Point implements ExportInterface
{
    const POINT_PRELOAD = '__POINT_PRELOAD';
    const POINT_CALIBRATE = 'Calibrate point';

    private $config;

    private $active;
    private $label;
    private $startTime;
    private $startMemoryUsage;
    private $stopTime;
    private $stopMemoryUsage;
    private $memoryPeak;
    private $differenceTime;
    private $differenceMemory;
    private $queryLog = [];
    private $newLineMessage = [];
    private $lineNum;
    private $scriptFile;
    private $msg;

    /**
     * Point constructor.
     * @param $name
     * @param $startTime
     * @param $startMemory
     */
    public function __construct(ConfigHandler $config, $name, $line, $file, $msg)
    {
        // Set items
        $this->config = $config;
        $this->setLabel($name);
        $this->lineNum = $line;
        $this->scriptFile = $file;
        $this->msg = $msg;
    }

    /*
     * Start point
     *
     * return void
     */
    public function start()
    {
        $this->setActive(true);
        $this->setStartTime();
        $this->setStartMemoryUsage();
    }

    /*
     * Finish point
     *
     * return void
     */
    public function finish()
    {
        $this->setActive(false);
        $this->setStopTime();
        $this->setStopMemoryUsage();
        $this->setMemoryPeak();
        $this->setDifferenceTime();
        $this->setDifferenceMemory();
    }

    /*
     * Simple export function
     */
    public function export()
    {
        $vars = get_object_vars($this);
        unset($vars['config']); // skip config

        return $vars;
    }

    // Get and set

    /**
     * @return mixed
     */
    public function getMemoryPeak()
    {
        return $this->memoryPeak;
    }

    /**
     * @param mixed $stopMemoryPeak
     */
    public function setMemoryPeak($memoryPeak = null)
    {
        $this->memoryPeak = ( $memoryPeak ) ? $memoryPeak : memory_get_peak_usage(true);
    }

    /**
     * @return mixed
     */
    public function isActive()
    {
        return $this->active;
    }

    /**
     * @param mixed $active
     */
    public function setActive($active)
    {
        $this->active = $active;
    }

    /**
     * @return mixed
     */
    public function getStopTime()
    {
        return $this->stopTime;
    }

    /**
     * @param mixed $stopTime
     */
    public function setStopTime($stopTime = null)
    {
        $this->stopTime = ($stopTime) ? $stopTime : microtime(true);
    }

    /**
     * @return mixed
     */
    public function getStopMemoryUsage()
    {
        return $this->stopMemoryUsage;
    }

    /**
     * @param mixed $stopMemory
     */
    public function setStopMemoryUsage($stopMemoryUsage = null)
    {
        $this->stopMemoryUsage = ($stopMemoryUsage) ? $stopMemoryUsage : memory_get_usage(true);
    }


    /**
     * @return mixed
     */
    public function getLabel()
    {
        return $this->label;
    }

    /**
     * @param mixed $label
     */
    public function setLabel($label)
    {
        // Run ltrim
        $configLTrim = $this->config->getPointLabelLTrim();
        if($configLTrim)
            $label = ltrim($label, $configLTrim);

        // Run Rtrim
        $configRTrim = $this->config->getPointLabelRTrim();
        if($configRTrim)
            $label = rtrim($label, $configRTrim);

        // Set nice label
        if($label !== self::POINT_PRELOAD and $this->config->isPointLabelNice())
            $label = ucfirst(str_replace('  ', ' ', preg_replace('/(?<!^)[A-Z]/', ' $0', $label)));

        $this->label = $label;
    }
    

    public function setLineNum($num)
    {
        $this->lineNum = $num;
    }
    
    public function getLineNum()
    {
        return $this->lineNum;
    }
    
    public function setScriptFile($file)
    {
        $this->scriptFile = $file;
    }
    
    public function getScriptFile()
    {
        return $this->scriptFile;
    }
    
    public function setMsg($msg)
    {
        $this->msg = $msg;
    }
    
    public function getMsg()
    {
        return $this->msg;
    }
    
    

    /**
     * @return mixed
     */
    public function getStartTime()
    {
        return $this->startTime;
    }

    /**
     * @param mixed $startTime
     */
    public function setStartTime($startTime = null)
    {
        $this->startTime = ( $startTime ) ? $startTime : microtime(true);
    }

    /**
     * @return mixed
     */
    public function getStartMemoryUsage()
    {
        return $this->startMemoryUsage;
    }

    /**
     * @param mixed $startMemory
     */
    public function setStartMemoryUsage($startMemoryUsage = null)
    {
        $this->startMemoryUsage = ( $startMemoryUsage ) ? $startMemoryUsage : memory_get_usage(true);
    }

    /**
     * @return mixed
     */
    public function getDifferenceTime()
    {
        return $this->differenceTime;
    }

    /**
     * @param mixed $differenceTime
     */
    public function setDifferenceTime()
    {
        $this->differenceTime = $this->stopTime - $this->startTime;
    }

    /**
     * @return mixed
     */
    public function getDifferenceMemory()
    {
        return $this->differenceMemory;
    }

    /**
     * @param mixed $differenceMemory
     */
    public function setDifferenceMemory()
    {
        $difference = $this->stopMemoryUsage - $this->startMemoryUsage;
        $this->differenceMemory = ($difference > 0 ) ? $difference: 0;
    }

    /**
     * @return array
     */
    public function getQueryLog()
    {
        return $this->queryLog;
    }

    /**
     * @param array $queryLog
     */
    public function setQueryLog($queryLog)
    {
        $this->queryLog = $queryLog;
    }

    /**
     * @return array
     */
    public function getNewLineMessage()
    {
        return $this->newLineMessage;
    }

    /**
     * @param string $newLineMessage
     */
    public function addNewLineMessage($newLineMessage)
    {
        $this->newLineMessage[] = $newLineMessage;
    }

}
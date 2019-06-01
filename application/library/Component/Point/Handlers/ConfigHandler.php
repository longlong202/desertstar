<?php namespace Component\Point\Handlers;

use Component\Point\Interfaces\ExportInterface;
use Component\Point\Presenters\Presenter;

class ConfigHandler implements ExportInterface
{
    // Config items
    
    private $enableTool = true;
    private $pointLabelLTrim = false;
    private $pointLabelRTrim = false;
    private $pointLabelNice = false;
    private $presenter;
    private $logFile = '';
    private $startTime;

    /*
     * Hold state of the query log
     * null = not set
     * false = config is false
     * true = query log is set
     */
    public $queryLogState;

    public function __construct()
    {
        // Set default
        $this->setDefaultPresenter();
        $this->startTime = microtime(true);
    }

    /*
     * Simple export function
     */
    public function export()
    {
        return get_object_vars($this);
    }

    public function getAllItemNames()
    {
        return array_keys($this->configItems);
    }


    // Print format
    private function setDefaultPresenter()
    {
        $this->setPresenter(Presenter::PRESENTER_LOG);
    }

    /**
     * @return bool
     */
    public function isEnableTool()
    {
        return $this->enableTool;
    }

    /**
     * @param void 
     */
    public function disableTool()
    {
        $this->enableTool = false;
    }

    

    /**
     * @param bool $logFile
     */
    public function setLogFile($logFile)
    {
        $this->logFile = $logFile;
    }

    /**
     * @return string
     */
    public function getLogFile()
    {
        return $this->logFile;
    }



    /**
     * @return bool
     */
    public function isPointLabelTrim()
    {
        return $this->pointLabelTrim;
    }

    /**
     * @return mixed
     */
    public function getPointLabelLTrim()
    {
        return $this->pointLabelLTrim;
    }

    /**
     * @param string $pointLabelLTrim
     */
    public function setPointLabelLTrim($status)
    {
        $this->pointLabelLTrim = $status;
    }

    /**
     * @return mixed
     */
    public function getPointLabelRTrim()
    {
        return $this->pointLabelRTrim;
    }

    /**
     * @param string $pointLabelRTrim
     */
    public function setPointLabelRTrim($status)
    {
        $this->pointLabelRTrim = $status;
    }

    /**
     * @return mixed
     */
    public function getPresenter()
    {
        return $this->presenter;
    }

    /**
     * @param mixed $presenter
     */
    public function setPresenter($mixed)
    {
        if(is_int($mixed))
            $this->presenter = $mixed;
        else
            $this->presenter = Presenter::PRESENTER_LOG;
    }

    /**
     * @return bool
     */
    public function isPointLabelNice()
    {
        return $this->pointLabelNice;
    }

    /**
     * @param bool $pointLabelNice
     */
    public function setPointLabelNice($pointLabelNice)
    {
        $this->pointLabelNice = (bool) $pointLabelNice;
    }

    public function __set($name, $value)
    {
        $this->$name = $value;
    }
    
    public function __get($name)
    {
        return isset($this->$name) ? $this->$name : null;
    }

}
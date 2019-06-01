<?php 

namespace Component\Point\Presenters;

use Component\Point\Handlers\ConfigHandler;
use Component\Point\Point;

abstract class Presenter {

    // Set print format
    const PRESENTER_LOG = 1;

    // Config
    protected $config;
    protected $formatter;
    protected $calculate;
    protected $pointStack;

    public function __construct(ConfigHandler $config)
    {
        $this->config = $config;
        $this->formatter = new Formatter();
        $this->calculate = new Calculate();


        // Choose display format
        $this->bootstrap();
    }

    /*
     * Bootstrap sub class
     */
    abstract public function bootstrap();

    /*
     * Passed trigger to results to display
     */
    abstract public function displayResultsTrigger($pointStack);

    /*
     * Passed trigger finish point to display
     */
    abstract public function finishPointTrigger(Point $point);
}
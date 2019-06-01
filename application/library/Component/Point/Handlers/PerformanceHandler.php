<?php 

namespace Component\Point\Handlers;

use Component\Point\Point;
use Component\Point\Presenters\LogPresenter;

class PerformanceHandler
{
    /*
     * Version
     */
    const VERSION = '2.0.2';

    /*
     * Hold point stack
     */
    private $pointStack = [];

    /*
     *  Hold presenter
     */
    private $presenter;

    /*
     * Hold the config class
     */
    public $config;

    /*
     *
     */
    private $messageToLabel = null;

    public function __construct()
    {
        // Set config
        $this->config = new ConfigHandler();
    }

    public function bootstrap()
    {
        // Set display
        $this->bootstrapDisplay();

        // Preload class point
        $this->preload();
    }

    /*
     * Set measuring point X
     *
     * @param string|null   $label
     * @return void
     */
    public function point($label = null, $line, $file, $msg)
    {
        // Check if point already exists
        $this->finishLastPoint();
        $this->checkIfPointLabelExists($label);

        // Set label
        if(is_null($label))
            $label = 'Task ' . (count($this->pointStack) - 1);

        // Create point
        $point = new Point($this->config, $label, $line, $file, $msg);

        // Create and add point to stack
        $this->addPointToStack($point);

        // Start point
        $point->start();
    }

    /*
     * Set message
     *
     * @param string|null   $message
     * @param boolean|null   $newLine
     * @return void
     */
    public function message($message)
    {
        $this->config->msg=$message;
    }

    /*
     * Finish measuring point X
     *
     * @param string|null   $label
     * @return void
     */
    public function finish()
    {
        $this->finishLastPoint();
    }

    /*
     * Return test results
     *
     * @return Component\Point\Handlers\ExportHandler
     */
    public function results()
    {
        // Finish all
        $this->finishLastPoint();

        // Add results to presenter
        $this->presenter->displayResultsTrigger($this->pointStack);
        // Return export
        return $this->export();
    }

    public function export()
    {
        return new ExportHandler($this);
    }

    public function getPoints()
    {
        return $this->pointStack;
    }

//
// PRIVATE
//

    private function bootstrapDisplay()
    {
        $this->presenter = new LogPresenter($this->config);
    }

    /*
     * Add point to stack
     *
     * @param Point   $point
     * @return void
     */
    private function addPointToStack(Point $point)
    {
        $this->pointStack[] = $point;
    }

    /*
     * Finish all point in the stack
     *
     * @return void
     */
    private function finishLastPoint()
    {
        if(count($this->pointStack))
        {
            // Get point
            $point = end($this->pointStack);

            if($point->isActive())
            {
                // Check if message in label text
                $this->checkAndSetMessageInToLabel($point);

                // Finish point
                $point->finish();

                // Trigger presenter listener
                $this->presenter->finishPointTrigger($point);
            }
        }
    }

    private function checkIfPointLabelExists($label)
    {
        $labelExists = false;
        foreach ($this->pointStack as $point)
        {
            if($point->getLabel() == $label)
            {
                $labelExists = true;
                break;
            }
        }

        if($labelExists)
            throw new \Exception(" Label " . $label . " already exists! Choose new point label." );
    }

    /*
     * Preload wil setup te point class
     */
    private function preload()
    {
        $this->point( Point::POINT_PRELOAD,0,'','' );
        $this->point( Point::POINT_CALIBRATE ,0,'','');
    }

   
    /*
     * Update point label with message
     */
    private function checkAndSetMessageInToLabel(Point $point)
    {
        if( ! $this->messageToLabel)
            return;

        // Update label
        $point->setLabel( $point->getLabel() . " - " . $this->messageToLabel);

        // Reset
        $this->messageToLabel = '';

    }

}
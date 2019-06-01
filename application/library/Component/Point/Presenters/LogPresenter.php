<?php 

namespace Component\Point\Presenters;

use Component\Point\Point;

class LogPresenter extends Presenter
{
    public function bootstrap()
    {}

    public function finishPointTrigger(Point $point)
    {}

    public function displayResultsTrigger($pointStack)
    {
        $this->pointStack = $pointStack;
        $this->writeToLogFile();
    }

    // Private
    private function writeToLogFile()
    {
        $config = $this->config->export();
        $out = [];
        $out['uuid'] = $config['uuid'];
        $out['uidshow'] = $config['uidshow'];
        $out['mid'] = $config['mid'];
        $out['path'] = $config['path'];
        $out['reqtime'] = $config['reqtime'];
        $out['msg'] = $config['msg'];
        $out['mem'] = 0;
        $out['cost'] = 0;
        foreach($this->pointStack as $key => $point) {
            if($key > 1) {
                $out['points'][] = [
                    $point->getLabel(),
                    $this->formatter->timeToHuman($point->getDifferenceTime()),
                    $this->formatter->memoryToHuman($point->getDifferenceMemory()),
                    $this->formatter->memoryToHuman($point->getStopMemoryUsage()),
                    $point->getLineNum(),
                    $point->getScriptFile(),
                    $point->getMsg(),
                ];
            }
        }
        $out['mem']  = $this->formatter->memoryToHuman($point->getStopMemoryUsage());
        $out['cost'] = $this->formatter->timeToHuman(microtime(true) - $config['startTime']);
        if($config['logFile']) {
            file_put_contents($config['logFile'], json_encode($out) . "\n", FILE_APPEND);
        }
    }
    
    
}
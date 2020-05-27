<?php

//class IndexController extends Yaf_Controller_Abstract
class IndexController extends BaseController
{
    public function indexAction()
    {
//        $this->getView()->assign("content","Hello XXX");
        $a=array(
            'name'=>"张三",
            'age'=>18,
        );

        $this->renderJson($a);//输出json
    }

    public function zeusAction()
    {
//        $this->getView()->assign("content","Hello XXX");
        $a=array(
            'zeus'=>"jacob",
        );

        $this->renderJson($a);//输出json
    }
}

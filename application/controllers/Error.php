<?php
class ErrorController extends BaseController
{/*{{{*/

    protected function accessRules()
    {
        return ['error'];
    }
    public function errorAction($exception)
    {/*{{{*/
        if($exception instanceof Yaf_Exception_LoadFailed){
            $msg = "url loading error";
            print_r($exception);
            exit($msg);
        }elseif($exception instanceof Yaf_Exception){
            $msg = "system error";
            exit($msg);
        }else{
            echo "unknow error";
            print_r($exception);
        }
        exit;
    }/*}}}*/ 
}/*}}}*/

?>

<?php
class ErrorController extends BaseController
{/*{{{*/

    protected function accessRules()
    {
        return ['error'];
    }
    public function errorAction($exception)
    {/*{{{*/
        echo "<pre><h3>出错提示：</h3><br></pre>";
        if($exception instanceof Yaf_Exception_LoadFailed){
            $msg = "加载异常，您访问的地址不存在";
            print_r($exception);
            exit($msg);
        }elseif($exception instanceof Yaf_Exception){
            $msg = "系统繁忙";
            exit($msg);
        }else{
            echo "!!!调用至此!!!";
            print_r($exception);
        }
        exit;
    }/*}}}*/ 
}/*}}}*/

?>

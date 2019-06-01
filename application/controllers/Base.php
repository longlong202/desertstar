<?php
/**
 * controller抽象基类，其他controller均继承自该类
 */

abstract class BaseController extends Yaf_Controller_Abstract
{/*{{{*/
    protected $moduleName       = '';
    protected $actionName       = '';
    protected $controllerName   = '';
    protected $curViewPath      = '';
    protected $curAccessPath    = '';
    
    //初始化
    public function init()
    {/*{{{*/
        $this->moduleName       = $this->getModuleName();
        $this->actionName       = $this->getRequest()->getActionName();
        $this->controllerName   = $this->getRequest()->getControllerName();
        $this->curAccessPath    = $this->moduleName."/".$this->controllerName."/".$this->actionName;
        $this->path             = $this->getRequest()->getRequestUri(); 

        Yaf_Registry::set("mycontroller",$this);
        //以下是手动加载模板的路径名
        $this->curViewPath = TPLPATH . '/' . strtolower($this->moduleName) . '/' .strtolower($this->controllerName) . '/';
        $this->getView()->setScriptPath(array($this->curViewPath));

        //以下是自动加载模板的路径名
//        $this->curViewPath = TPLPATH . '/' . strtolower($this->moduleName);
//        $this->setViewpath($this->curViewPath);

        return true;
    }/*}}}*/

    //渲染模板
    public function renderTpl($data=null,$tplFile=null)
    {/*{{{*/
        if(empty($tplFile))
        {
            $tplFile = strtolower($this->actionName).'.tpl';
        }
        if (!empty($data))
        {
            $this->getView()->assign($data);
        }
        $this->getView()->display($tplFile);
    }/*}}}*/

    //接口输出json串
    public function renderJson($data=array(),$s='')
    {/*{{{*/
        if (empty($data)) {
            $data = array(
                'errno' => 0,
                'msg' => 'success',
            );
        }
        if (!empty($s))
        {
            echo ' '.$s.'('.json_encode($data).');';
        }
        else{
            echo json_encode($data);
        }
        exit;

    }/*}}}*/
}/*}}}*/

?>

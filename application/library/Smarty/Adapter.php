<?php
class Smarty_Adapter implements Yaf_View_Interface
{
    public $_smarty;
 
    public function __construct($tmplPath = null, $extraParams = array()) {

        require_once "Smarty.class.php";
        $this->_smarty = new Smarty;

        if (null !== $tmplPath) {
            $this->setScriptPath($tmplPath);
        }
 
        foreach ($extraParams as $key => $value) {
            $this->_smarty->$key = $value;
        }
    }
    public function clear() {
        return $this->_smarty->clearAllAssign();
    }
    public function assign($spec, $value = null) {
        if (is_array($spec)) {
            $this->_smarty->assign($spec);
            return;
        }
 
        $this->_smarty->assign($spec, $value);
    }
    public function render($tpl_name, $tpl_vars = NULL) {
        if($tpl_vars)
        {
            $this->assign($tpl_vars);
        }
        return $this->_smarty->fetch($tpl_name);
    }
    public function display($tpl_name, $tpl_vars = NULL) {
        if($tpl_vars)
        {
            $this->assign($tpl_vars);
        }
        $this->_smarty->display($tpl_name);
    }
    public function setScriptPath($tpl_dir) {
        $this->_smarty->setTemplateDir($tpl_dir);
    }
    public function addScriptPath($tpl_dir) {
        $this->_smarty->addTemplateDir($tpl_dir);
    }
    public function enableEscape(){
        $this->_smarty->escape_html=true;
    }
    public function getScriptPath() {
        return $this->_smarty->template_dir;
    }
    public function openDebug() {
        $this->_smarty->force_compile = true;
        $this->_smarty->debugging = false; 
        $this->_smarty->debugging_ctrl ='URL' ;
    }
    public function setPluginsPath($dirs) {
        if(is_string($dirs))
        {
            $dirs = array($dirs);
        } 
        $this->_smarty->addPluginsDir($dirs );

    }
}

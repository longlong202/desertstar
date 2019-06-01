<?php
class ShoppingController extends BaseController
{/*{{{*/

    public function indexAction()
    {/*{{{*/
        $content = 'Jessie & Jacob';
        $this->getView()->assign("content",$content); 
    }/*}}}*/
}/*}}}*/
?>

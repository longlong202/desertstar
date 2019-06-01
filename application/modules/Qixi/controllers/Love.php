<?php
class LoveController extends BaseController
{/*{{{*/

    public function loveJessieAction()
    {/*{{{*/
        $content="Jessie & Jacob";
//        $this->getView()->assign("content",'Jessie & Jacob');
        $data =array(
            'content'=>$content,
        );
        $this->renderTpl($data);
//        $this->getView()->assign($data);
//        $this->getView()->assign("content",$content);
    }/*}}}*/

    public function indexAction()
    {/*{{{*/
        $arr        = array(49,38,65,97,76,13,27,50);
        $arrSize    = count($arr);

        //从第一次排序抽出来
        Utils\Algorithm::buildHeap($arr,$arrSize);

        for($i=$arrSize-1;$i>0;$i--)
        {
            Utils\Algorithm::swap($arr,$i,0);
            $arrSize--;
            Utils\Algorithm::buildHeap($arr,$arrSize);
        }

//        $json_str="{'service': 'X_API', 'no': [1531451295, 1530930561, 1529743579, 1528969927, 1528967462, 1528944986, 1528888504, 1528878566], 'pass': [64, 88, 85, 100, 100, 100, 100, 100], 'fail': [14, 12, 12, 0, 0, 0, 0, 0], 'err': [21, 0, 2, 0, 0, 0, 0, 0]}";
//        $this->getView()->assign("res","hello");
//        $data['res']="xxx";
        $data =array(
            'content'=>"desertstar",
        );
//        $this->renderTpl($data);
        $this->renderJson($data);//输出json
    }/*}}}*/
}/*}}}*/
?>

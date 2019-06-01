<?php
class IndexController extends BaseController
{/*{{{*/

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
        $json_str = array(
            'service'   => 'X_API',
            'title'        => array('北京', '郑州','深圳'),
            'pass'      => array(60, 88, 85),
            'fail'      => array(14, 12, 12, 0, 0, 0, 0, 0),
            'err'       => array(21, 0, 2, 0, 0, 0, 0, 0),
        );
//        $this->getView()->assign("res",$json_str);
//        $data['res']="xxx";
        $data =array(
            'res'=>"-smarty模板",
            "data"=>$json_str,
        );
        $data2 =array(
            'res'=>"-smarty模板",
            "data"=> 'test',
        );
        $this->renderTpl($data);
//        $this->renderJson($data);
    }/*}}}*/
}/*}}}*/
?>

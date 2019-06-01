<?php
namespace Utils;
use Common\Logger;

class Algorithm
{/*{{{*/
    
    //用数组建立最小堆
    static function buildHeap(&$arr,$arrSize)
    {/*{{{*/
        //计算出最开始的下标$index,最大值，比较每一个子树的父结点和子结点，将最小值存入父结点
        //从$index处对一个树进行循环比较，形成最小堆
        for($index=intval($arrSize/2)-1;$index>=0;$index--)
        {/*{{{*/
            //如果有左节点，将其下标存进最小值$min
            if($index*2+1<$arrSize)
            {/*{{{*/
                $min = $index*2+1;
                //如果右子节点比左结点更小，将结点下标记录为最小值$min
                if($index*2+2<$arrSize)
                {/*{{{*/
                    if($arr[$index*2+2]<$arr[$min])
                    {
                        $min = $index*2+2;
                    }
                }/*}}}*/
                //将子节点中较小的与父结点比较，若子结点较小，与父结点交换位置，同事更新较小
                if($arr[$min]<$arr[$index])
                {
                    self::swap($arr,$min,$index);
                }
            }/*}}}*/
        }/*}}}*/
//        Logger::getLogger()->debug("Test Logger Success");
        Logger::getLogger()->error("Test Logger error");
    }/*}}}*/

    //交换数组中下标为$one和$another的数据
    static function swap(&$arr,$one,$another)
    {/*{{{*/
        $tmp            = $arr[$one];
        $arr[$one]      = $arr[$another];
        $arr[$another]  = $tmp;
    }/*}}}*/
}/*}}}*/
?>

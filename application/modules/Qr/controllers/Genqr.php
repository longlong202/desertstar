<?php
/*
 *modules/controller/action 三级目录就会进入modules
 *controller/action 二级目录就会进入controllers
 */
class GenqrController extends BaseController
{/*{{{*/

    public function indexAction()
    {/*{{{*/
        $content = self::genQrcode("http://116.196.106.8/qixi/love/loveJessie");
        //$content = self::genQrcode("http://116.196.106.8/hk/shopping/index");

        $data =array(
            'content'=>$content,
        );
        //$this->renderJson($data);//输出json
        //var_dump($this->getView());
        //$this->getView()->assign($data);
        $this->getView()->assign('pic',$content);
        $this->getView()->display('index.tpl');
//        $this->getView()->assign("content",$content);

    }/*}}}*/

    static public function genQrcode($url='')
    {/*{{{*/
        //require_once APPPATH.'/library/Utils/phpqrcode.php';
        require_once 'phpqrcode.php';
        $value = $url;//二维码内容
        $errorCorrectionLevel = 'H';//容错级别
        $martrixPointSize = 6;//生成图片大小
        //生成的二维码图片
        $filename = WWWPATH.'/img/qrcode.png';
        QRcode::png($value,$filename,$errorCorrectionLevel,$martrixPointSize,2);
        //Utils\QRcode::png($value,$filename,$errorCorrectionLevel,$martrixPointSize,2);
        $logo = WWWPATH.'/img/wechat.jpg';
        $QR = $filename;

        if(file_exists($logo))
        {
            $QR = imagecreatefromstring(file_get_contents($QR));
            $logo = imagecreatefromstring(file_get_contents($logo));
            $QR_width = imagesx($QR);//二维码图片宽度
            $QR_height = imagesy($QR);//二维码图片高度
            $logo_width = imagesx($logo);//logo图片宽度
            $logo_height = imagesy($logo);//logo图片高度
            $logo_qr_width = $QR_width/4;//组合之后的logo宽度占二维码1/5
            $scale = $logo_width/$logo_qr_width;//logo宽度缩放比，本身宽度/组合后的宽度
            $logo_qr_height = $logo_height/$scale;//组合之后的logo高度
            $from_width = ($QR_width-$logo_qr_width)/2;//组合之后logo左上角所在坐标点

            imagecopyresampled($QR,$logo,$from_width,$from_width,0,0,$logo_qr_width,$logo_qr_height,$logo_width,$logo_height);//将一幅图像中的一块正方形区域拷贝到另一个图像中

        }

        //输出图片
        imagepng($QR,WWWPATH.'/img/qrcode.png');
        imagedestroy($QR);
        imagedestroy($logo);
        return '../../img/qrcode.png';

    }/*}}}*/
}/*}}}*/
?>

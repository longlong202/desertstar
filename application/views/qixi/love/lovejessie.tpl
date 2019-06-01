<!DOCTYPE HTML>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, user-scalable=no, minimum-scale=1.0, maximum-scale=1.0">
<title>抽奖咯</title>
<style type="text/css">
.demo{
    position:absolute;
    left:50%;
    top:50%;
    transform:translate(-50%,-50%);
}
#disk{
    width:auto; 
    height:auto; 
}
#start{
    width:163px; 
    height:320px; 
    position:absolute; 
    left:50%;
    top:50%;
    transform:translate(-50%,-50%);
}
#start img{
    cursor:pointer
}
</style>
<script type="text/javascript" src="/js/qixi/jquery.min.js"></script>
<script type="text/javascript" src="/js/qixi/jQueryRotate.2.2.js"></script>
<script type="text/javascript" src="/js/qixi/jquery.easing.min.js"></script>
<script type="text/javascript">
$(function(){
	$("#startbtn").rotate({
		bind:{
			click:function(){
				var a = Math.floor(Math.random() * 360);
				 $(this).rotate({
					 	duration:3000,
					 	angle: 0, 
            			animateTo:1440+a,
						easing: $.easing.easeOutSine,
						callback: function(){
							alert('恭喜Jessie，七夕中奖了！');
						}
				 });
			}
		}
	});
});
</script>
</head>

<body>

<div id="main">
   <div class="demo">
        <div id="disk"><img src="/img/qixi/disk2.jpg" id="diskbtn"></div>
        <div id="start"><img src="/img/qixi/start.png" id="startbtn"></div>
   </div>
</div>


</body>
</html>

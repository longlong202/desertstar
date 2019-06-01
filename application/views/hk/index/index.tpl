<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Jessie&Jacob</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <!-- Bootstrap  core CSS-->
    <link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css">
    <link rel="stylesheet" href="/style/common/dataTables.bootstrap.min.css">
    <link rel="stylesheet" href="/style/common/jquery.dataTables.min.css">
    <!--<script src="/js/common/bootstrap.min.js"></script>-->
    <!--<script src="/js/qixi/jquery.min.js"></script>-->
    <script src="/js/common/echarts.common.min.js"></script>
    <!-- jQuery (Bootstrap 插件需要引入) -->
    <script src="https://cdn.bootcss.com/jquery/2.1.1/jquery.min.js"></script>
    <!-- 包含了所有编译插件 先引用jquery,再引用bootstrap-->
    <script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>

    <script src="/js/common/jquery.dataTables.min.js"></script>
    <script src="/js/common/dataTables.bootstrap.min.js"></script>
    <script src="/js/hk/index/index.main.js"></script>
    <script type="text/javascript" src="/js/hk/index/detail.main.js?_v=2.0.0"></script>
    <script >
        var title= [{%foreach $data.title as $row%}'{%$row%}',{%/foreach%}];
        var val= [{%foreach $data.pass as $row%}{%$row%},{%/foreach%}];
    </script>

<style>
 body {
    padding-top: 50px;
    padding-left: 50px;
}
</style>
</head>
<body>
    
<div class="navbar navbar-fixed-top navbar-inverse" role="navigation">
    <div class="container">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="#">Jessie & Jacob Website</a>
        </div>
        <div class="collapse navbar-collapse">
            <ul class="nav navbar-nav">
                <li class="active"><a href="#"><span class="glyphicon glyphicon-home"></span> Home</a></li>
                <li><a href="#shop"><span class="glyphicon glyphicon-shopping-cart"></span> Shop</a></li>
                <li><a href="#support"><span class="glyphicon glyphicon-headphones"></span> Support</a></li>
            </ul>
        </div><!-- /.nav-collapse -->
        
    </div><!-- /.container -->
</div>
<div>
        <!-- -->
        <br>
        <!--加入自己的内容 ,pre-scrollable内容可滚动 pre适合显示代码，首行会有空格-->
        <!--
        <pre class="pre-scrollable">
        <?php echo $content; ?>
        </pre>
        -->
        <!-- 
        justify 根据屏幕对超出的文字进行换行
        nowrap  不会根据屏幕大小对超出的文字进行换行
        -->

        <!-- 
        <p class="text-center">
        <?php var_dump($content); ?>
        </p>
        -->
        <!-- Smarty调用一定要加$符号-->
        <h2 style="text-align:center">图-demo{%$res%}</h2>
        <div id="main" style="width:50%;height:400px;border:0px solid #999;padding:3px;margin:0 auto">
            Long
        </div>




        <h2 class="sub-header">表格-demo</h2>
        <div class="table-responsive">
            <!-- 表格 begin-->
            <table id="datatable_sheet" class="table table-striped table-bordered" cellspacing="0" width="100%">
                <thead>
                <!-- 表格只需要自定义好thead-->
                    <tr>
                        <th>运行批次</th>
                        <th>平台</th>
                        <th>服务名</th>
                        <th>接口路径</th>
                        <th>用例名</th>
                        <th>代码路径</th>
                        <th>结果</th>
                        <th>耗时</th>
                        <th>详细</th>
                        <th>异常信息</th>
                        <th>操作</th>
                    </tr>
                </thead>
                <!-- tbody由插件补充-->
            </table>
            <!-- 表格 end-->
        </div>

</div>
    
</body>
</html>

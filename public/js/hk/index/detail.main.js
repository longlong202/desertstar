//配置数据列表
$(document).ready(function() {
    // datatable start
    var wageNowTable;
    var initAjaxsource= "/demo.json";//默认调用最新批次执行结果
    var initParams={
        "ajax": {
            "url": initAjaxsource,
        },
        //该配置确保datatable是可以刷新的，不能少
        retrieve: true,
        "columns": [
            { "data": "no"},
            { "data": "platform" },
            { "data": "service" },
            { "data": "uri" },
            { "data": "casename" },
            { "data": "casepath" },
            { "data": "result" },
            { "data": "cost" },
            { "data": "consoleinfo", "visible": false },
            { "data": "errinfo", "visible": false },
            { "data": null , "defaultContent": "<button>查看</button>"}
        ],
        "order":[[6,"asc"]]
    };
    wageNowTable = $('#datatable_sheet').DataTable(initParams);//初始化数据表
    // datatable end

    //添加datatable'查看'按钮点击事件
    $('#datatable_sheet tbody').on( 'click', 'button', function () {
        wageNowTable = $('#datatable_sheet').DataTable(initParams);
        var data = wageNowTable.row( $(this).parents('tr') ).data();//父级tr数据是json格式
        alert( data["consoleinfo"] +"\n\n"+ data["errinfo"]);//获取json格式进行输出
    });

});
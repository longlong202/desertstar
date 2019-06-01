$(document).ready(function() {
    // 基于准备好的dom，初始化echarts实例
    var myChart = echarts.init(document.getElementById('main'));


    // 指定图表的配置项和数据
    var option={
        title:{
            text:'demo'
        },
        tooltip:{},
        legend:{
            data:['销量']
        },
        xAxis:{
            data:[title[0],title[1],title[2]]
        },
        yAxis:{},
        series:[{
            name:'销量',
            type:'bar',
            data:[val[0],val[1],val[2]]
            // data:[20,30,50]
        }]
    };
    // 使用刚指定的配置项和数据显示图表。
    myChart.setOption(option);
   // myChart.on('click', function (params) {
    //    window.open('/report/detail?kw=' + encodeURIComponent(params.name));
   // });
});

// import THREE from 'three'
window.onload = function () {
    console.log("this Root verion:", THREE.REVISION)
    const R = 200;
    var Root = {
        W: window.innerWidth,
        H: window.innerHeight,
        Root: document.getElementById('root'),
    }
    var renderer = new THREE.WebGLRenderer({ antialias: true })
    //建议设置大小，否则会出现锯齿，也可能是因为我没有设置div具体宽高
    renderer.setSize(window.innerWidth, window.innerHeight);
    Root.Root.appendChild(renderer.domElement)

    //创建场景
    var scene = new THREE.Scene();

    //添加光源
    var light = new THREE.DirectionalLight(0xffffff, 1);
    light.position.set(0, 150, 0);
    scene.add(light);
    //  添加辅助线
    var axisHelper = new THREE.AxisHelper(800);
    scene.add(axisHelper);

    //添加相机
    var camera = new THREE.PerspectiveCamera(60, Root.W / Root.H, 0.1, 2000);
    camera.position.set(0, 0, 1000);
    camera.lookAt(scene.position);
    scene.add(camera);


    //添加组
    var groupSphere = new THREE.Group();



    function animate() {
        renderer.setSize(window.innerWidth, window.innerHeight);
        //重复渲染
        requestAnimationFrame(animate);
        renderer.render(scene, camera);
    }
    animate();
}
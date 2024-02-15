import * as THREE from 'three';
import { GLTFLoader } from 'three/addons/loaders/GLTFLoader.js';

const scene = new THREE.Scene();
const camera = new THREE.PerspectiveCamera(10, window.innerWidth / window.innerHeight, 0.1, 1000);

const canvas = document.getElementById('canvas');

const renderer = new THREE.WebGLRenderer({ canvas: canvas, alpha: false });
renderer.setClearColor(0x000000, 1);
renderer.setSize(window.innerWidth, window.innerHeight);

const loader = new GLTFLoader();

let hoth;
let corruscant;
let korriban;
let mustafar;
let xwing;

loader.load("assets/3dModels/hoth/hoth.gltf", function (gltf) {
    hoth = gltf.scene;
    hoth.position.set(-7, 0, 0);
    scene.add(hoth);
})

loader.load("assets/3dModels/corruscant/corruscant.gltf", function (gltf) {
    corruscant = gltf.scene
    corruscant.position.set(-3.5, 0, 0)
    scene.add(corruscant)
})

loader.load("assets/3dModels/korriban/korriban.gltf", function (gltf) {
    korriban = gltf.scene
    korriban.position.set(0, 0, 0)
    scene.add(korriban)
})

loader.load("assets/3dModels/mustafar/mustafar.gltf", function (gltf) {
    mustafar = gltf.scene
    mustafar.position.set(3.5, 0, 0)
    scene.add(mustafar)
});

loader.load("assets/3dModels/x-wing/x-wing.gltf", function (gltf) {
    xwing = gltf.scene
    xwing.position.set(0, -1, 0)
    xwing.rotation.y = Math.PI/2
    xwing.scale.set(0.5, 0.5, 0.5)
    scene.add(xwing)
})

camera.position.x = 0;
camera.position.y = -0.5;
camera.position.z = 53;

const topLight = new THREE.DirectionalLight(0xffffff, 1);
topLight.position.set(0, 0, 1);
topLight.castShadow = false;
scene.add(topLight);

function animate() {
    requestAnimationFrame(animate);

    hoth.rotation.y += 0.002;
    corruscant.rotation.y += 0.002;
    korriban.rotation.y += 0.002;
    mustafar.rotation.y += 0.002;
    xwing.rotation.y += 0.005;

    renderer.render(scene, camera);
}

animate();




import * as THREE from 'three';
import { GLTFLoader } from 'three/addons/loaders/GLTFLoader.js';

const scene = new THREE.Scene();
const camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000);
const renderer = new THREE.WebGLRenderer({ alpha: true });
renderer.setSize(window.innerWidth, window.innerHeight);
document.getElementById("containerHoth").appendChild(renderer.domElement);

const loader = new GLTFLoader();

let hoth;
let corruscant;
let korriban;
let mustafar;
let xwing;

loader.load("assets/3dModels/hoth/hoth.gltf", function (gltf) {
    hoth = gltf.scene
    hoth.position.set(-3, 0, 0)
    scene.add(hoth)
})

loader.load("assets/3dModels/corruscant/corruscant.gltf", function (gltf) {
    corruscant = gltf.scene
    corruscant.position.set(0, 0, 1)
    scene.add(corruscant)
})

loader.load("assets/3dModels/korriban/korriban.gltf", function (gltf) {
    korriban = gltf.scene
    korriban.position.set(3, 0, 0)
    scene.add(korriban)
})

loader.load("assets/3dModels/mustafar/mustafar.gltf", function (gltf) {
    mustafar = gltf.scene
    mustafar.position.set(6, 0, -1)
    scene.add(mustafar)
});


loader.load("assets/3dModels/x-wing/x-wing.gltf", function (gltf) {
    xwing = gltf.scene
    xwing.position.set(0, 1, 0)
    xwing.rotation.y = Math.PI/2
    scene.add(xwing)
})

camera.position.x = 0;
camera.position.y = 1;
camera.position.z = 6;

camera.rotation.x = -0.15;

const topLight = new THREE.DirectionalLight(0xffffff, 1);
topLight.position.set(0, 0, 1)
topLight.castShadow = false;
scene.add(topLight);

function animate() {
    requestAnimationFrame(animate);

    corruscant.rotation.y += 0.002;
    hoth.rotation.y += 0.002;
    korriban.rotation.y += 0.002;
    mustafar.rotation.y += 0.002;
    xwing.rotation.y += 0.01;

    renderer.render(scene, camera);
}

animate();

function onDocumentMouseDown() {
    console.log("aaaaaaaaaaaaa");
}

hoth.addEventListener("click", onDocumentMouseDown, true);


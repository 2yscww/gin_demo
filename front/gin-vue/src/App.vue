<template>
  <div id="app">
    <b-cont>
      <NavBar></NavBar>
    </b-cont>
    <!-- 渲染路由匹配的组件 -->
     
    <b-container>
      <router-view msg="Welcome to Your Vue.js App"></router-view>
    </b-container> 
  </div>

  <!-- <img alt="Vue logo" src="./assets/logo.png"> -->
  <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
</template>

<script setup>
// import HelloWorld from './components/HelloWorld.vue'
import NavBar from './components/NavBar.vue';


</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  /* margin-top: 60px; */
}
</style>

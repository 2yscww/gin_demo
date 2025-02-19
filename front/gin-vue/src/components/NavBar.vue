<template>
    <!-- <nav>
      <ul>
        <li><router-link to="/">首页</router-link></li>
        <li><router-link to="/register">注册</router-link></li>
        <li><router-link to="/login">登录</router-link></li>
      </ul>
    </nav> -->
    <nav class="navbar navbar-expand-lg navbar-dark ">
  <a class="navbar-brand" href="#">Navbar</a>
  <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
    <span class="navbar-toggler-icon"></span>
  </button>
  <div class="collapse navbar-collapse" id="navbarNav">
    <ul class="navbar-nav">
      <li class="nav-item active">
        <router-link class="nav-link" to="/">Home <span class="sr-only"></span></router-link>
      </li>
      <li class="nav-item">
        <router-link class="nav-link" to="/login">login</router-link>
      </li>
      <li class="nav-item">
        <router-link class="nav-link" to="/register">register</router-link>
      </li>
      <!-- <li class="nav-item">
        <a class="nav-link disabled" href="#">Disabled</a>
      </li> -->
    </ul>
  </div>
    </nav>

    <div class="test-block">

    </div>

</template>


<script setup>

</script>


<style lang="scss" scoped>
    .navbar{
        background-color: #2c3e50 !important; /* 设置导航栏背景色 */
    }
    .navbar-brand {
        
        padding-left: 10px;
    }   

    .test-block{
        padding-bottom: 20px;
    }

    
//     nav {
//   background-color: #2c3e50; /* 设置导航栏背景色 */
//   padding: 10px 0;
// }

// ul {
//   display: flex; /* 让列表横向排列 */
//   justify-content: center; /* 让内容居中 */
//   list-style: none; /* 移除默认的列表样式 */
//   padding: 0;
//   margin: 0;
// }

// li {
//   margin: 0 15px; /* 设置列表项之间的间距 */
// }

// a {
//   text-decoration: none; /* 去掉下划线 */
//   color: white; /* 设置文字颜色 */
//   font-size: 18px;
//   padding: 8px 12px;
//   border-radius: 5px;
//   transition: background 0.3s;
// }

// a:hover {
//   background: rgba(255, 255, 255, 0.2); /* 添加悬停效果 */
// }</style>
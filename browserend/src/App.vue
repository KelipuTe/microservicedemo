<template>
  <div class="app">
    <h2>app</h2>
    <div class="app-navigate">
      <RouterLink to="/home" active-class="active">home</RouterLink>
      <RouterLink to="/user" active-class="active">user</RouterLink>
    </div>
    <div class="app-content">
      <RouterView></RouterView>
    </div>
  </div>
</template>

<script lang="ts" setup name="App">
import { RouterLink, RouterView } from 'vue-router';

</script>

<style>
.app {
  background-color: gray;
}

.app-navigate a {
  width: 100px;
  height: 50px;
  border: 1px solid black;
}

.app-navigate a.active {
  border: 1px solid red;
}
</style>

<template>
  <div class="user">
    <h2>user</h2>
    <div class="user-navigate">
      <RouterLink to="/user/signup" active-class="active">signup</RouterLink>
      <RouterLink to="/user/login" active-class="active">login</RouterLink>
      <RouterLink to="/user/profile" active-class="active">profile</RouterLink>
    </div>
    <div class="user-content">
      <RouterView></RouterView>
    </div>
  </div>
</template>

<script lang="ts" setup name="User">
import { RouterLink, RouterView } from 'vue-router';
</script>

<style>
.user {
  border: 1px solid gray;
}

.user-navigate a {
  width: 100px;
  height: 50px;
  border: 1px solid black;
}

.user-navigate a.active {
  border: 1px solid red;
}
</style>
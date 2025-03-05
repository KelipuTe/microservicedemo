<template>
  <div class="profile">
    <h2>profile</h2>
    <p>id:{{ user.id }}</p>
    <p>id:{{ user.username }}</p>
    <p><button @click="logout">退出登录</button></p>
  </div>
</template>

<script lang="ts" setup name="Profile">
import { ref, reactive } from 'vue';
import { useRouter } from 'vue-router';
import axios from 'axios';

import { type User } from '@/types/user'

const router = useRouter()

let user = ref<User>({
  id: 0,
  username: ''
})

axios.get('http://localhost:8080/user/profile',
  { withCredentials: true })
  .then(function (response) {
    console.log(response);
    if (response.status == 200) {
      user.value = response.data
    }
  }).catch(function (error) {
    console.log(error);
  })

function logout(){
  axios.post('http://localhost:8080/user/logout',
        {},
        { withCredentials: true }
    ).then(function (response) {
        if (response.status == 200) {
            router.push({ path: '/home' });
        }
    }).catch(function (error) {
        console.log(error);
    })
}

</script>

<style>
.profile {
  border: 1px solid gray;
}
</style>
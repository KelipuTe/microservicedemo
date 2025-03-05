<template>
    <div class="login">
        <h2>login</h2>
        <p>
            <span>用户名：<input v-model="username"></span>
        </p>
        <p>
            <span>密码：<input v-model="password"></span>
        </p>
        <p>
            <button @click="login">登录</button>
        </p>
    </div>
</template>

<script lang="ts" setup name="Login">

import { ref, reactive } from 'vue'
import axios from 'axios';
import { useRouter } from 'vue-router';

const router = useRouter()

let username = ref("")
let password = ref("")

function login() {
    axios.post('http://localhost:8080/user/login',
        {
            username: username.value,
            password: password.value
        }, { withCredentials: true })
        .then(function (response) {
            if (response.status == 200) {
                router.push({ path: '/home' });
            }
        }).catch(function (error) {
            console.log(error);
        })
}
</script>

<style>
.login {
    border: 1px solid gray;
}
</style>
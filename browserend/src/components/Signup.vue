<template>
    <div class="signup">
        <h2>signup</h2>
        <p>
            <span>用户名：</span>
            <input v-model="username">
        </p>
        <p>
            <span>密码：</span>
            <input v-model="password">
        </p>
        <p>
            <span>确认密码：</span>
            <input v-model="confirmPassword">
        </p>
        <p>
            <button @click="signup">注册</button>
        </p>
    </div>
</template>

<script lang="ts" setup name="Signup">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router';
import axios from 'axios';

const router = useRouter()

let username = ref("")
let password = ref("")
let confirmPassword= ref("")

function signup(){
    axios.post('http://localhost:8080/user/signup',
        {
            username: username.value,
            password: password.value,
            confirmPassword: confirmPassword.value
        },
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
.signup {
    border: 1px solid gray;
}
</style>
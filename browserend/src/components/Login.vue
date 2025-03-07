<template>
    <div class="login">
        <h2>login</h2>
        <p>
            <span>用户名：</span>
            <input v-model="username">
        </p>
        <p>
            <span>密码：</span>
            <input v-model="password">
        </p>
        <p>
            <button @click="login">登录</button>
        </p>
        <p>
            <span>邮箱：</span>
            <input v-model="email">
            <button @click="signupEmail">发送邮箱验证码</button>
        </p>
        <p>
            <span>手机：</span>
            <input v-model="phone">
            <button @click="signupSms">发送手机验证码</button>
        </p>
        <p>
            <span>验证码：</span>
            <input v-model="code">
        </p>
        <p>
            <button @click="loginEmail">邮箱登录</button>
            <button @click="loginSms">手机登录</button>
        </p>
    </div>
</template>

<script lang="ts" setup name="Login">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router';
import axios from 'axios';

const router = useRouter()

let username = ref("")
let password = ref("")

let email= ref("")
let phone= ref("")
let code= ref("")

function login() {
    axios.post('http://localhost:8080/user/login',
    {
            username: username.value,
            password: password.value
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

function signupEmail(){
    axios.post('http://localhost:8080/user/signup_email',
        {
            email: email.value
        },
        {  }
    ).then(function (response) {
        console.log(response);
    }).catch(function (error) {
        console.log(error);
    })
}

function loginEmail(){
    axios.post('http://localhost:8080/user/login_email',
        {
            email: email.value,
            code:code.value
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

function signupSms(){
    axios.post('http://localhost:8080/user/signup_sms',
        {
            phone: phone.value
        },
        {  }
    ).then(function (response) {
        console.log(response);
    }).catch(function (error) {
        console.log(error);
    })
}

function loginSms(){
    axios.post('http://localhost:8080/user/login_sms',
        {
            phone: phone.value,
            code:code.value
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
.login {
    border: 1px solid gray;
}
</style>
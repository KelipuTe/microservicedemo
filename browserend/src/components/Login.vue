<template>
    <div class="login">
        <p>
            <span>用户名：{{ username }}</span>
        </p>
        <p>
            <span>密码：</span>
        </p>
        <button @click="login">登录</button>

        <p>
            <span>aa:{{ car.price }}</span>
        </p>
    </div>
</template>

<script lang="ts" setup name="Login">
import {ref,reactive} from 'vue'

let username = ref("xhy")
let car = reactive({price:100})


function login(){
            alert("aaa");
username.value="bbb"
            Object.assign(car,{price:200})
        }
</script>

<style>
.login{
    background-color:aqua
}
</style>
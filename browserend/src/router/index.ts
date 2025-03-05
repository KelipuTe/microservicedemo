import { createRouter, createWebHistory } from 'vue-router'

import Home from '@/views/Home.vue'
import User from '@/views/User.vue'

import Login from '@/components/Login.vue'
import Profile from '@/components/Profile.vue'
import Signup from '@/components/Signup.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            redirect: '/home'
        },
        {
            path: '/home',
            component: Home
        },
        {
            path: '/user',
            component: User,
            children: [
                {
                    path: 'login',
                    component: Login
                },
                {
                    path: 'profile',
                    component: Profile
                },
                {
                    path: 'signup',
                    component: Signup
                }
            ]
        }
    ]
})

export default router
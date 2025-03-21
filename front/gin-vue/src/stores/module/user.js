import { defineStore } from 'pinia';
import storageService from "@/utils/storageService";
import router from '@/router/router';

export const useUserStore = defineStore('user', {
    state: () => ({
        token: storageService.get(storageService.USER_TOKEN),
        userName: storageService.get(storageService.USER_INFO),
    }),
    actions: {
        setToken(token) {
            //更新本地缓存
            storageService.set(storageService.USER_TOKEN, token);
            // 更新state
            this.token = token;
        },
        setUserInfo(userName) {
            //更新本地缓存
            storageService.set(storageService.USER_INFO, userName);
            // 更新state
            this.userName = userName;
        },
        clearUserInfo() {
            // 移除本地缓存服务的记录
            storageService.remove(storageService.USER_INFO);
            storageService.remove(storageService.USER_TOKEN);

            // 移除pinia中的记录
            this.token = '';
            this.userName = '';

            // 先跳转到登录页，再刷新整个页面
            router.push('/').then(() => {
                location.reload(); // 立即刷新页面
            });
        }
    }
});



// const userModule = {
//     namespaced: true,
//     state: {
//         token: storageService.get(storageService.USER_TOKEN),
//         userName: storageService.get(storageService.USER_INFO),
//     },
//     mutations: {
//         SET_TOKEN(state, token) {
//             //更新本地缓存
//             storageService.set(storageService.USER_TOKEN, token);
//             // 更新state
//             state.token = token;
//         },
//         SET_USER_INFO(state, userName) {
//             //更新本地缓存
//             storageService.set(storageService.USER_INFO, userName);
//             // 更新state
//             state.userName = userName;
//         },
//     }
// };
import { defineStore } from 'pinia';
import storageService from "@/utils/storageService";

export const useUserStore = defineStore('user', {
    state: () => ({
        token: storageService.get(storageService.USER_TOKEN),
        userName: storageService.get(storageService.USER_INFO),
    }),
    actions:{
        setToken(token){
            //更新本地缓存
            storageService.set(storageService.USER_TOKEN, token);
            // 更新state
            this.token = token;
        },
        setUserInfo(userName){
            //更新本地缓存
            storageService.set(storageService.USER_INFO, userName);
            // 更新state
            this.userName = userName;
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
import axios from "axios";
import storageService from "@/utils/storageService.js";


export const instance = axios.create({
    baseURL: "http://127.0.0.1:8081/api", // 你的 API 地址
    timeout: 5000, // 请求超时时间
    headers: { "Content-Type": "application/json" }
})

export const userService = axios.create({
    baseURL: "http://127.0.0.1:8081/api", // 你的 API 地址
    timeout: 5000, // 请求超时时间
    // headers: { Authorization: 'Bearer ' + `${storageService.get(storageService.USER_TOKEN)}` }
    
    // 写在这里可能会导致这个接口始终使用的是旧的token
})

// 添加请求拦截器，确保请求时使用最新 Token
userService.interceptors.request.use(
    (config) => {
        const latestToken = storageService.get(storageService.USER_TOKEN); // 获取最新 Token
        if (latestToken) {
            config.headers.Authorization = 'Bearer ' + `${latestToken}`; // 设置最新 Token
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);


import axios from "axios";

const instance = axios.create({
    baseURL: "http://127.0.0.1:8081/api", // 你的 API 地址
    timeout: 5000, // 请求超时时间
    headers: { "Content-Type": "application/json" }
})

export default instance
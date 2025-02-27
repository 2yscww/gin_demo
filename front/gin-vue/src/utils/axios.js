import instance from "@/utils/request"



// 用户登录
export const loginIntface = (data) => {
    return instance.post("/auth/login", data)
};

// 用户注册
export const registerInt = (data) => {
    return instance.post("/auth/register", data)
};


// 获取用户信息

// export const userInfo = () => {
//     return service.get("/auth/info")
// }
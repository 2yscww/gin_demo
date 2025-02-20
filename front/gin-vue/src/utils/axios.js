import instance from "@/utils/request"



// 用户登录
export const loginInt =(data) => {
    return instance.post("/auth/login",data)
};

// 用户注册
export const registerInt =(data) => {
    return instance.post("/auth/register",data)
};

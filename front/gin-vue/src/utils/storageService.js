// 本地缓存服务

// ? 便于替换方法实现，例如以后想要使用cookies，只需要修改代理方法

const PREFIX = 'gin_vue_';

// user 模块

const USER_PREFIX = `${PREFIX}user_`;
const USER_TOKEN = `${USER_PREFIX}token`;
const USER_INFO = `${USER_PREFIX}info`;

// 储存

const set = (key, data) => {
    localStorage.setItem(key, data)
}


// 读取
const get = (key) => {
    return localStorage.getItem(key)
}


export default {
    set, 
    get,
    USER_TOKEN,
    USER_INFO
};
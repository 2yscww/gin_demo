// 本地缓存服务

// ? 便于替换方法实现，例如以后想要使用cookies，只需要修改代理方法

const PREFIX = 'gin_vue_';

// user 模块

const USER_PREFIX = `${PREFIX}user_`;
const USER_TOKEN = `${USER_PREFIX}token`;
const USER_INFO = `${USER_PREFIX}info`;

// 储存

const set = (key, data) => {
    console.log(`存储的数据: ${data}`); // 检查存储的数据
    localStorage.setItem(key, data)
}


// 读取
const get = (key) => {
    return localStorage.getItem(key)
}

const remove = (key) => {
    try {
        localStorage.removeItem(key);
    } catch (e) {
        console.error(`删除存储数据失败: ${key}`, e);
    }
}


export default {
    set, 
    get,
    remove,
    USER_TOKEN,
    USER_INFO
};
<template>
    <div class="register">
        <b-row>
            <b-col md="8" offset-md="2" lg="6" offset-lg="3">
                <b-card title="注册">
                    <form @submit.prevent="register">
                        <div class="mb-3">
                            <label for="exampleInputEmail1" class="form-label">名称(选填)</label>
                            <input type="text" v-model="user.username" class="form-control" id="username1">
                        </div>
                        <div class="mb-3">
                            <label for="exampleInputTelephone1" class="form-label">电话</label>
                            <input type="tel" class="form-control" v-model="user.telephone" id="exampleInputTelephone1"
                                placeholder="输入您的电话号码">
                            <b-form-text class="text-danger" v-if="telephoneNumRed">手机号必须为11位</b-form-text>
                            <b-form-text class="text-danger" v-if="telephoneHasExistRed">{{
                                telephoneErrorMsg}}</b-form-text>
                        </div>
                        <div class="mb-3">
                            <label for="exampleInputPassword1" class="form-label">密码</label>
                            <input type="password" v-model="user.password" class="form-control"
                                id="exampleInputPassword1" placeholder="输入您的密码">
                            <b-form-text class="text-danger" v-if="passwdRed">密码不少于8位</b-form-text>
                        </div>
                        <div class="mb-3 form-check">
                            <input type="checkbox" class="form-check-input" id="exampleCheck1">

                        </div>
                        <button type="submit" class="btn btn-primary">注册</button>
                    </form>
                </b-card>
            </b-col>
        </b-row>

    </div>
</template>


<script setup>
import { reactive } from 'vue';
import { ref } from 'vue';
import { registerInt, userInfo } from '@/utils/axios';
import router from '@/router/router';
// import storageService from '@/utils/storageService';
import { useUserStore } from '@/stores/module/user';

const telephoneNumRed = ref(false);

const passwdRed = ref(false);

const telephoneHasExistRed = ref(false)

const telephoneErrorMsg = ref(""); // 存储后端返回的错误信息

const user = reactive({
    username: "",
    telephone: "",
    password: ""
});

// 检查电话号码是否符合前端要求
const telephoneCheckFunc = () => {
    if (user.telephone.length !== 11) {
        telephoneNumRed.value = true;
    } else {
        telephoneNumRed.value = false;
    }
}

// 检查密码是否符合前端要求
const passwdCheckFunc = () => {
    if (user.password.length < 8) {
        passwdRed.value = true;
    } else {
        passwdRed.value = false;
    }
}

//  代码还原

// const telephoneHasExistFunc = (mesg) => {
//     if (mesg === "该手机号已注册用户") {
//         telephoneHasExistRed.value = true
//         telephoneErrorMsg.value = mesg
//     } else {
//         telephoneHasExistRed.value = false
//         telephoneErrorMsg.value = ""
//     }
// }

const telephoneHasExistFunc = (mesg) => {
    if (mesg != "") {
        telephoneHasExistRed.value = true
        telephoneErrorMsg.value = mesg
    } else {
        telephoneHasExistRed.value = false
        telephoneErrorMsg.value = ""
    }
}

// 在页面中定义了 userStroe
const userStore = useUserStore();

const register = async () => {
    console.log("点击注册按钮，开始执行注册逻辑");
    // 先执行检查
    telephoneCheckFunc();
    passwdCheckFunc();

    //电话号码或密码前端要求不符合
    if (telephoneNumRed.value == true || passwdRed.value == true) {
        console.log("电话或者密码不符合要求");
        return;
    }

    // 验证数据
    try {
        // 发送数据到后端
        const response = await registerInt(user);


        if (response.data && response.data.code === 200) {

            // 保存token
            // 假设后端返回了 token，保存它到 localStorage 或 pinia
            // localStorage.setItem("token", response.data.data.token);
            // storageService.set(storageService.USER_TOKEN, JSON.stringify(response.data.data.token));

            
            // 调用 storageService 来保存信息
            // storageService.set(storageService.USER_TOKEN, response.data.data.token);

            // 由pinia接管
            userStore.setToken(response.data.data.token)


            // 获取到最新的token
            // const newToken = storageService.get(storageService.USER_TOKEN);
            const newToken = userStore.token

            // TODO 完善功能后要将log输出的代码都删除 

            //TODO 完善用户信息的校验,比如校验token

            console.log("注册保存的token------------");
            console.log("Token  storage:", newToken);
            console.log("注册保存的token------------");

            // 保存用户信息

            const infoResponse = await userInfo();

            if (infoResponse.data && infoResponse.data.code === 200) {
                // * 将用户名保存
                console.log(infoResponse.data.data.user.username); // 确认这个值是否正确
                // storageService.set(storageService.USER_INFO, JSON.stringify(response.data.data.user.username));

                // 调用 storageService 来保存信息
                // storageService.set(storageService.USER_INFO, infoResponse.data.data.user.username);

                // pinia接管
                userStore.setUserInfo(infoResponse.data.data.user.username);

                console.log(infoResponse)

            } else {
                console.log("获取用户信息失败:", infoResponse);
            }

            

            


            console.log("注册成功:", response.data);

            // 跳转主页
            router.push("/")


        } else {
            console.log("注册失败：", response.data.message);
        }


    } catch (error) {
        // console.error("请求错误：", error);
        if (error.response) {
            const { msg } = error.response.data;


            telephoneHasExistFunc(msg);

            // if (msg === "该手机号已注册用户") {
            //     alert("该手机号已注册用户")
            // }
        }
    }



}




</script>


<style lang="scss" scoped></style>
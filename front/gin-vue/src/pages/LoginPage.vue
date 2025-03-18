-- Active: 1696057084513@@127.0.0.1@3306@go_test
-- Active: 1742305285737@@127.0.0.1@49161
<template>
    <div class="login">
        <b-row>
            <b-col md="8" offset-md="2" lg="6" offset-lg="3">
                <b-card title="登录">
                    <form @submit.prevent="login">
                        <div class="mb-3">
                            <label for="exampleInputTelephone1" class="form-label">电话</label>
                            <input type="tel" class="form-control" v-model="user.telephone" id="exampleInputTelephone1"
                                placeholder="输入您的电话号码">
                            <b-form-text class="text-danger" v-if="telephoneNumRed">手机号必须为11位</b-form-text>
                            <b-form-text class="text-danger" v-if="telephoneHasExistRed">{{
                                telephoneErrorMsg }}</b-form-text>
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
                        <button type="submit" class="btn btn-primary">登录</button>
                    </form>
                </b-card>
            </b-col>
        </b-row>

    </div>
</template>


<script setup>
import { reactive } from 'vue';
import { ref } from 'vue';
import { loginIntface, userInfo } from '@/utils/axios';
import { useUserStore } from '@/stores/module/user';
import storageService from '@/utils/storageService';
import router from '@/router/router';


const telephoneNumRed = ref(false);

const passwdRed = ref(false);

const errorMsgRed = ref(false);

// 存储后端返回的错误信息
const errorMsg = ref("");

// 定义 userStore pinia
const userStore = useUserStore(); 


const user = reactive({
    telephone: "",
    password: ""
});


// 检查电话是否符合前端要求
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


// 后端返回的错误信息
const errorMsgFunc = (mesg) => {
    if (mesg != "") {
        errorMsgRed.value = true
        errorMsg.value = mesg
    } else {
        errorMsgRed.value = false
        errorMsg.value = ""
    }
}




const login = async () => {
    // TODO 登录逻辑需要完善
    
    // 先执行前端检查
    telephoneCheckFunc();
    passwdCheckFunc();

    //电话号码或密码前端要求不符合
    if (telephoneNumRed.value == true || passwdRed.value == true) {
        console.log("电话或者密码不符合要求");
        return;
    }

    // 验证数据

    try {
        const response = await loginIntface(user);

        if (response.data && response.data.code === 200) {

            // 由pinia接管
            userStore.setToken(response.data.data.token)

            console.log(storageService.get(storageService.USER_TOKEN)); // 确认这个值是否正确

            const infoResponse = await userInfo();

            if (infoResponse.data && infoResponse.data.code === 200) {
                // 将用户名保存
                
                // pinia接管
                userStore.setUserInfo(infoResponse.data.data.user.username);

                console.log(storageService.get(storageService.USER_INFO)); // 确认这个值是否正确

                console.log(infoResponse)

            } else {
                console.log("获取用户信息失败:", infoResponse);
            }

            // 跳转主页
            router.push("/")


        } else {
            console.log("登录失败:",response);
        }

    } catch (error) {
        if (error.response) {
            const { msg } = error.response.data;


            errorMsgFunc(msg);

            // if (msg === "该手机号已注册用户") {
            //     alert("该手机号已注册用户")
            // }
        }
    }























    //登录失败的逻辑


    // if (user.telephone.length !== 11) {
    //     telephoneNumRed.value = true;
    //     return;
    // } else {
    //     telephoneNumRed.value = false;
    // }

    // if (user.password.length < 8) {
    //     passwdRed.value = true;
    // return;
    // } else {
    //     passwdRed.value = false;
    // }
}


</script>


<style lang="scss" scoped></style>

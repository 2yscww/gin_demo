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
                        <button type="submit" class="btn btn-primary" >登录</button>
                    </form>
                </b-card>
            </b-col>
        </b-row>

    </div>
</template>


<script setup>
import { reactive } from 'vue';
import { ref } from 'vue';

const telephoneNumRed = ref(false);


const passwdRed = ref(false);


const user = reactive({
    username: "",
    telephone: "",
    password: ""
});

const telephoneCheckFunc = () => {
    if (user.telephone.length !== 11) {
        telephoneNumRed.value = true;
    } else {
        telephoneNumRed.value = false;
    }
}

const passwdCheckFunc = () => {
    if (user.password.length < 8) {
        passwdRed.value = true;
    } else {
        passwdRed.value = false;
    }
}

// TODO 登录逻辑需要完善

const login = () => {
    // 先执行检查
    telephoneCheckFunc();
    passwdCheckFunc();

    //登录失败的逻辑
    if (telephoneNumRed.value == true || passwdRed.value == true){
        console.log("不通过!");
        return;
    }

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


<style lang="scss" scoped>
</style>
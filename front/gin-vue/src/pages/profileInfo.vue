<template>
    <h3>您好 <span v-if="localUserName">{{ localUserName }}</span></h3>
</template>

<script setup>
import { useUserStore } from '@/stores/module/user';
import { computed } from 'vue';

const userStore = useUserStore();

const localUserName = computed(() => {
    return userStore.userName
});

</script>

<style lang="scss" scoped></style>
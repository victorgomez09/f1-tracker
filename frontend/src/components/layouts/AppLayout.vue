<script setup lang="ts">
import Navbar from '../navbar/Navbar.vue';

</script>

<template>
    <div class="flex flex-col w-full h-full">
        <Navbar></Navbar>

        <div class="overflow-auto">
            <router-view></router-view>
        </div>
    </div>
</template>
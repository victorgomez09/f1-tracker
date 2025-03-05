<script setup lang="ts">
import { onMounted } from "vue";
import Navbar from "./components/navbar/Navbar.vue";
import { themeChange } from "theme-change";

onMounted(() => {
  themeChange(false)
})
</script>

<template>
  <div class="flex flex-col overflow w-full h-full bg-base-300">
    <div class="flex p-2 h-full w-full">
      <router-view></router-view>
      <!--<router-view
      v-if="connected && dataUpdated"
      class="flex flex-1 w-full"
      ></router-view>-->
    </div>
  </div>
</template>

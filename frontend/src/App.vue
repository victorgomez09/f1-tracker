<script setup lang="ts">
import Navbar from "./components/navbar/Navbar.vue";
</script>

<template>
  <div class="flex flex-col overflow w-full h-full">
    <Navbar />

    <div class="flex p-2 h-full w-full">
      <router-view></router-view>
      <!--<router-view
      v-if="connected && dataUpdated"
      class="flex flex-1 w-full"
      ></router-view>-->
    </div>
  </div>
</template>

<script setup lang="ts">
import Navbar from "@/components/navbar/Navbar.vue"
</script>

<template>
  <div class="flxe flex-col h-full w-full">
    <Navbar />
    <div class="hero bg-base-200 h-full p-0">
      <div class="hero-content text-center p-0">
        <div class="max-w-md">
          <h1 className="my-20 text-center text-5xl font-bold">
            Real-time Formula 1 <br />
            telemetry and timing
          </h1>

          <div className="flex flex-wrap gap-4">
            <RouterLink to="/dashboard">
              <button class="btn btn-primary">Go to live</button>
            </RouterLink>

            <RouterLink to="/historical">
              <button class="btn btn-outline">Historical races</button>
            </RouterLink>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import Navbar from "@/components/navbar/Navbar.vue"
</script>

<template>
  <Navbar />
  <div className="container mx-auto max-w-screen-lg px-4">
    <section className="flex h-full w-full flex-col items-center pt-20 sm:justify-center sm:pt-0">
      <!-- <Image src={icon} alt="f1-dash tag logo" width={200} /> -->

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
    </section>
  </div>
</template>
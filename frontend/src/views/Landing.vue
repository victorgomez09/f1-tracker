<script setup lang="ts">
import { Button } from '@/components/ui/button';
</script>

<template>
  <div className="container mx-auto max-w-screen-lg px-4">
    <section className="flex h-full w-full flex-col items-center pt-20 sm:justify-center sm:pt-0">
      <!-- <Image src={icon} alt="f1-dash tag logo" width={200} /> -->

      <h1 className="my-20 text-center text-5xl font-bold">
        Real-time Formula 1 <br />
        telemetry and timing
      </h1>

      <div className="flex flex-wrap gap-4">
        <RouterLink to="/dashboard">
          <Button>Go to Dashboard</Button>
        </RouterLink>

        <RouterLink to="/historical">
          <Button variant="outline">Historical races</Button>
        </RouterLink>
      </div>
    </section>
  </div>
</template>
<script setup lang="ts">
import { Historical } from '@/models/historical.model';
import { computed, onMounted, ref, Ref } from 'vue';
import { HistoricalSession } from '@/models/session.model';
import moment from 'moment';
import { useRouter } from 'vue-router';

const router = useRouter()
const data: Ref<Array<Historical>> = ref([]);

onMounted(async () => {
  // const result = await fetch("https://stunning-system-j4wxj4p5v4j3555p-3000.app.github.dev/historical")
  const result = await fetch("http://localhost:3000/historical")

  data.value = await result.json()
})

const parsedEvents = computed(() => {
  return data.value.filter(event => event.Type !== HistoricalSession.PreSeasonSession)
})
</script>

<template>
  <div class="flex flex-col gap-2 p-2">
    <input type="text" placeholder="Search event" class="input">
    <div class="overflow-x-auto w-full">
      <table class="table table-zebra w-full">
        <!-- head -->
        <thead>
          <tr>
            <th>Name</th>
            <th>Country</th>
            <th>Event type</th>
            <th>Date</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="event in parsedEvents">
            <th v-on:click="router.push({ name: 'historical', params: { eventName: event.Name } })"
              class="cursor-pointer underline hover:font-bold">
              {{ event.Name }}
            </th>
            <td>{{ event.Country }}</td>
            <td>{{ HistoricalSession[event.Type] }}</td>
            <td>{{ moment(event.EventTime).format("YYYY-MM-DD") }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
<script setup lang="ts">
import { Button } from '@/components/ui/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuPortal,
  DropdownMenuSeparator,
  DropdownMenuShortcut,
  DropdownMenuSub,
  DropdownMenuSubContent,
  DropdownMenuSubTrigger,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { Drs, SafeftyCar, Session, Status } from '@/models/session.model'
import { useEventStore, useInformationStore, usePausedStore, useTimeStore } from '@/store/data.store'
import momentTz from 'moment-timezone'

// const ws = getWs()

const {isReplay} = defineProps({
  isReplay: Boolean,
  ws: WebSocket
})

const parseEventType = (eventType: number) => {
  return Session[eventType]
}

const parseStatus = (status: number) => {
  switch (status) {
    case Status.UnknownState:
      return "Unknown";
    case Status.Inactive:
      return "Inactive";
    case Status.Started:
      return "Started";
    case Status.Aborted:
      return "Aborted";
    case Status.Finished:
      return "Finished";
    case Status.Finalised:
      return "Finalised";
    case Status.Ended:
      return "Ended";
  }
}

const parseStatusColor = (status: number) => {
  switch (status) {
    case Status.UnknownState:
    case Status.Inactive:
    case Status.Finished:
    case Status.Finalised:  
    case Status.Ended:
      return "text-white";
    case Status.Started:
      return "text-green-500";
    case Status.Aborted:
      return "text-red-500"
  }
}

const parseDrs = (drs: number) => {
  switch (drs) {
    case Drs.DRSUnknown:
      return "Unknown"
    case Drs.DRSEnabled:
      return "Enabled"
    case Drs.DRSDisabled:
      return "Disabled"
  }
}

const parseSafeftyCar = (safetyCar: number) => {
  switch (safetyCar) {
    case SafeftyCar.Clear:
      return "Clear";
    case SafeftyCar.VirtualSafetyCar:
      return "VSC Deployed";
    case SafeftyCar.VirtualSafetyCarEnding:
      return "VSC Ending"
    case SafeftyCar.SafetyCar:
      return "Deployed"
    case SafeftyCar.SafetyCarEnding:
      return "Ending"
  }
}

const parseSafeftyCarColor = (safetyCar: number) => {
  switch (safetyCar) {
    case SafeftyCar.VirtualSafetyCar:
    case SafeftyCar.VirtualSafetyCarEnding:
      return "text-yellow-500";
    case SafeftyCar.SafetyCar:
    case SafeftyCar.SafetyCarEnding:
      return "text-red-500"
    default:
      return "text-green-500"
  }
}

const parseRemainingTime = (remaininTime: number) => {
	// hour := int(w.remainingTime.Seconds() / 3600)
	// minute := int(w.remainingTime.Seconds()/60) % 60
	// second := int(w.remainingTime.Seconds()) % 60
  const time = remaininTime / 1000
  const hours = time / 3600
  const minutes = (time / 60) % 60
  const seconds = time % 60

  return `${Math.trunc(hours)}:${Math.trunc(minutes)}:${seconds.toFixed(2)}`
}

const goToStart = async () => {
  await fetch(`https://stunning-system-j4wxj4p5v4j3555p-3000.app.github.dev/actions`, {
    headers: {
      'Content-type': 'application/json'
    },
    method: 'POST',
    body: JSON.stringify({skipToStart: true})
  })
}
const skip5Secs = async () => {
  await fetch(`https://stunning-system-j4wxj4p5v4j3555p-3000.app.github.dev/actions`, {
    headers: {
      'Content-type': 'application/json'
    },
    method: 'POST',
    body: JSON.stringify({skip5Secs: true})
  })
}
const skipMinute = async () => {
  await fetch(`https://stunning-system-j4wxj4p5v4j3555p-3000.app.github.dev/actions`, {
    headers: {
      'Content-type': 'application/json'
    },
    method: 'POST',
    body: JSON.stringify({skipMinute: true})
  })
}
const skip10Minutes = async () => {
  await fetch(`https://stunning-system-j4wxj4p5v4j3555p-3000.app.github.dev/actions`, {
    headers: {
      'Content-type': 'application/json'
    },
    method: 'POST',
    body: JSON.stringify({skip10Minutes: true})
  })
}
const pause = async () => {
  const result = await fetch(`https://stunning-system-j4wxj4p5v4j3555p-3000.app.github.dev/actions`, {
    headers: {
      'Content-type': 'application/json'
    },
    method: 'POST',
    body: JSON.stringify({pause: true})
  })

  usePausedStore.setPaused(await result.json())
}
</script>

<template>
  <div class="flex items-center gap-2" v-if="useInformationStore.information.CircuitTimezone">
    <!-- Name -->
    <span>{{ useEventStore.event.Name }}: {{ parseEventType(useEventStore.event.Type) }}</span>

    <!-- Track time -->
    <span>- Track time: {{
      momentTz(useTimeStore.time.Timestamp)
      .tz(useInformationStore.information.CircuitTimezone).format("YYYY-MM-DD HH:mm:ss") }}</span>

    <!-- Track status -->
    <span>- Track status: <span :class="parseStatusColor(useEventStore.event.Status)">{{
        parseStatus(useEventStore.event.Status) }}</span></span>

    <!-- DRS enabled -->
    <span>- DRS enabled: {{ parseDrs(useEventStore.event.DRSEnabled) }}</span>

    <!-- Safefty car -->
    <span>- Safefty car: <span :class="parseSafeftyCarColor(useEventStore.event.Status)">{{
        parseSafeftyCar(useEventStore.event.SafetyCar) }}</span></span>

    <!-- Current lap -->
    <span v-if="useEventStore.event.Type === Session.Race">- Laps: {{ useEventStore.event.CurrentLap }} / {{
      useEventStore.event.TotalLaps }}</span>
    <span>- Remaining: {{ parseRemainingTime(useTimeStore.time.Remaining) }}</span>

    <div v-if="isReplay" class="flex ml-auto">
      <DropdownMenu>
        <DropdownMenuTrigger as-child>
          <Button variant="outline">
            Actions
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent class="w-56">
          <DropdownMenuLabel>Session actions</DropdownMenuLabel>
          <DropdownMenuSeparator />
          <DropdownMenuGroup>
            <DropdownMenuItem v-on:click="skip5Secs()">
              <span>Skip 5 seconds</span>
              <DropdownMenuShortcut>⇧⌘P</DropdownMenuShortcut>
            </DropdownMenuItem>
            <DropdownMenuItem v-on:click="skipMinute()">
              <span>Skip 1 minute</span>
              <DropdownMenuShortcut>⌘B</DropdownMenuShortcut>
            </DropdownMenuItem>
            <DropdownMenuItem v-on:click="skip10Minutes()">
              <span>Skip 10 minutes</span>
              <DropdownMenuShortcut>⌘S</DropdownMenuShortcut>
            </DropdownMenuItem>
            <DropdownMenuItem v-on:click="goToStart()">
              <span>Skip to race start</span>
              <DropdownMenuShortcut>⌘K</DropdownMenuShortcut>
            </DropdownMenuItem>
          </DropdownMenuGroup>
          <DropdownMenuSeparator />
          <DropdownMenuGroup>
            <DropdownMenuItem v-on:click="pause()">
            <span>{{ usePausedStore.paused ? 'Resume' : 'Pause' }}</span>
              <DropdownMenuShortcut>⌘+T</DropdownMenuShortcut>
            </DropdownMenuItem>
          </DropdownMenuGroup>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  </div>
</template>

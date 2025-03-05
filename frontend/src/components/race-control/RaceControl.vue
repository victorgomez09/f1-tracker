<script setup lang="ts">
import momentTz from 'moment-timezone'

import { useRaceControlStore, useInformationStore } from '@/store/data.store';
import { RaceControl } from '@/models/race-control.model';
import { SessionMessage } from '@/models/session.model';

const parseRaceControlMessage = (raceControl: RaceControl) => {
    let prefix = ""
    let color = "text-white"

    switch (raceControl.Flag) {
        case SessionMessage.GreenFlag:
            color = "text-green-500"
            if (raceControl.Msg.startsWith("GREEN LIGHT")) {
                prefix = "●"
            } else {
                prefix = "⚑"
            }

            break;
        case SessionMessage.YellowFlag:
            color = "text-yellow-500"
            prefix = "⚑"

            break;
        case SessionMessage.DoubleYellowFlag:
            color = "text-yellow-500"
            prefix = "⚑" + "⚑"

            break;
        case SessionMessage.BlueFlag:
            color = "text-cyan-500"
            prefix = "⚑"

            break;
        case SessionMessage.RedFlag:
            color = "text-red-500"
            if (raceControl.Msg.startsWith("RED LIGHT")) {
                prefix = "●"
            } else {
                prefix = "⚑"
            }

            break;
        case SessionMessage.BlackAndWhite:
            color = "text-white"
            prefix = "⚑" + "⚑"

            break;
    }

    if (prefix.length != 0) {
        return {
            color: color,
            msg: `${momentTz(raceControl.Timestamp).tz(useInformationStore.information.CircuitTimezone).format("HH:mm:ss")} ${prefix} - ${raceControl.Msg}`
        }
    } else {
        return {
            color: color,
            msg: `${momentTz(raceControl.Timestamp).tz(useInformationStore.information.CircuitTimezone).format("HH:mm:ss")} - ${raceControl.Msg}`
        }
    }
}
</script>

<template>
    <div class="overflow-auto">
        <div v-for="raceControl in useRaceControlStore.raceControl" class="flex items-center mb-2">
            <span :class="parseRaceControlMessage(raceControl).color" class="leading-none">
                {{ parseRaceControlMessage(raceControl).msg }}
            </span>
        </div>
    </div>
</template>
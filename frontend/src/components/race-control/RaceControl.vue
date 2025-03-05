<script setup lang="ts">
import momentTz from 'moment-timezone'

import { useRaceControlStore, useInformationStore } from '@/store/data.store';
import { RaceControl } from '@/models/race-control.model';
import { SessionMessage } from '@/models/session.model';

const parseRaceControlMessage = (raceControl: RaceControl) => {
    // func (f FlagState) String() string {
    // 	return [...]string{"None", "Green", "Yellow", "Double Yellow", "Red", "Chequered", "Blue", "Black and White"}[f]
    // }
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
        // msgs = append(msgs,
        //     giu.Style().SetColor(giu.StyleColorText, color).To(giu.Label(fmt.Sprintf("%s %s - %s", r.rcMessages[x].Timestamp.In(r.dataSrc.CircuitTimezone()).
        //         Format("15:04:05"), prefix, r.rcMessages[x].Msg)).Wrapped(true)))
    } else {
        return {
            color: color,
            msg: `${momentTz(raceControl.Timestamp).tz(useInformationStore.information.CircuitTimezone).format("HH:mm:ss")} - ${raceControl.Msg}`
        }
        // msgs = append(msgs,
        //     giu.Label(
        //         fmt.Sprintf("%s - %s",
        //             r.rcMessages[x].Timestamp.In(r.dataSrc.CircuitTimezone()).
        //                 Format("15:04:05"), r.rcMessages[x].Msg)).Wrapped(true))
    }
}
</script>

<template>
    <div class="overflow-auto">
        <div v-for="raceControl in useRaceControlStore.raceControl" class="flex items-center mb-1">
            <span :class="parseRaceControlMessage(raceControl).color">{{ parseRaceControlMessage(raceControl).msg
                }}</span>
        </div>
    </div>
</template>
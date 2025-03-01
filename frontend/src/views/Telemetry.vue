<script setup lang="ts">
import { Button } from '@/components/ui/button'
import { valueUpdater } from '@/lib/utils'
import type {
  ColumnDef,
  ColumnFiltersState,
  ExpandedState,
  SortingState,
  VisibilityState,
} from '@tanstack/vue-table'

import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import { Driver, DriverLocation, PitStop } from '@/models/driver.model'
import { useDriverStore, useGeneralStore, useTelemetryStore } from '@/store/data.store'
import { parseData } from '@/utils/parse-data.utils'
import {
  FlexRender,
  getCoreRowModel,
  getExpandedRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  useVueTable,
} from '@tanstack/vue-table'
import { ChevronDown } from 'lucide-vue-next'
import { computed, h, onMounted, ref } from 'vue'
import moment from 'moment'
import DriverTag from '@/components/driver/DriverTag.vue'

const socket = ref();
const retry = ref();

const delayMs = ref(0);
const connected = ref(false);
const blocking = ref(false);

const dataUpdated = ref(false);

const initWebsocket = (handleMessage: any) => {
  if (retry.value) {
    clearTimeout(retry.value);
    retry.value = undefined;
  }

  const wsUrl = "ws://localhost:3001";

  const ws = new WebSocket(wsUrl);

  ws.addEventListener("open", () => {
    connected.value = true;
  });

  ws.addEventListener("close", () => {
    connected.value = false;
    blocking.value = true;
    () => {
      if (!retry.value && !blocking.value)
        retry.value = window.setTimeout(() => {
          initWebsocket(handleMessage);
        }, 1000);
    };
  });

  ws.addEventListener("error", () => {
    ws.close();
  });

  ws.addEventListener("message", ({ data }) => {
    setTimeout(() => {
      handleMessage(data);
      console.log("message", data);
    }, delayMs.value);
  });

  socket.value = ws;
};

onMounted(() => {
  const ws = new WebSocket(
    "ws://localhost:3000/ws"
  );

  ws.addEventListener("open", () => {
    console.log("open");
    connected.value = true;
  });

  ws.onmessage = (data) => {
    try {
      parseData(JSON.parse(data.data));
      dataUpdated.value = true;
    } catch (e) {
      console.error(`could not process message: ${e}`);
    }
  };
});


const columns: ColumnDef<Driver>[] = [
  {
    accessorKey: 'Position',
    cell: ({ row }) => h(DriverTag, { short: row.original.ShortName, color: row.original.HexColor, position: row.original.Position }),
    header: 'Position',
  },
  {
    accessorKey: 'Gap',
    header: 'Gap',
    cell: ({ row }) => h('div', { class: 'capitalize' }, row.getValue('FastestLap')),
  },
  {
    accessorKey: 'Segments',
    header: 'Sectors',
    cell: ({ row }) => h('div', { class: 'capitalize' }, row.getValue('Segments')),
  },
  {
    accessorKey: 'S1',
    header: 'S1',
    cell: ({ row }) => h('div', { class: 'capitalize' }, row.getValue('FastestLap')),
  },
  {
    accessorKey: 'S2',
    header: 'S2',
    cell: ({ row }) => h('div', { class: 'capitalize' }, row.getValue('FastestLap')),
  },
  {
    accessorKey: 'S3',
    header: 'S3',
    cell: ({ row }) => h('div', { class: 'capitalize' }, row.getValue('FastestLap')),
  },
  {
    accessorKey: 'FastestLap',
    header: 'Fastest lap',
    cell: ({ row }) => h('div', { class: 'capitalize' }, row.getValue('FastestLap')),
  },
  {
    accessorKey: 'LastLap',
    header: 'Last lap',
    cell: ({ row }) => h('div', { class: 'capitalize' }, row.getValue('FastestLap')),
  },
  {
    accessorKey: 'DRS',
    header: 'DRS',
    cell: ({ row }) => h('div', { class: 'capitalize' }, row.getValue('FastestLap')),
  },
  {
    accessorKey: 'Tire',
    header: 'Tire/Laps',
    cell: ({ row }) => h('div', { class: 'capitalize' }, row.getValue('FastestLap')),
  },
  {
    accessorKey: 'Pits',
    header: 'Sectors',
    cell: ({ row }) => h('div', { class: 'capitalize' }, row.getValue('FastestLap')),
  },
  {
    accessorKey: 'LastPit',
    header: 'Last pit stop',

    cell: ({ row }) => h('div', { class: 'capitalize' }, row.getValue('FastestLap')),
  },
  {
    accessorKey: 'Telemetry',
    header: 'Sectors',
    cell: ({ row }) => h('div', { class: 'capitalize' }, row.getValue('FastestLap')),
  },
  {
    accessorKey: 'Speed trap',
    header: 'Speed trap',
    cell: ({ row }) => h('div', { class: 'capitalize' }, row.getValue('FastestLap')),
  },
  {
    accessorKey: 'Location',
    header: 'Location',
    cell: ({ row }) => h('div', { class: 'capitalize' }, row.getValue('FastestLap')),
  },
]
const sorting = ref<SortingState>([])
const columnFilters = ref<ColumnFiltersState>([])
const columnVisibility = ref<VisibilityState>({})
const rowSelection = ref({})
const expanded = ref<ExpandedState>({})

const sorted = computed(() =>
  useDriverStore.drivers.sort((a, b) => a.Position - b.Position)
);

const table = useVueTable({
  data: sorted.value,
  columns,
  getCoreRowModel: getCoreRowModel(),
  getPaginationRowModel: getPaginationRowModel(),
  getSortedRowModel: getSortedRowModel(),
  getFilteredRowModel: getFilteredRowModel(),
  getExpandedRowModel: getExpandedRowModel(),
  onSortingChange: updaterOrValue => valueUpdater(updaterOrValue, sorting),
  onColumnFiltersChange: updaterOrValue => valueUpdater(updaterOrValue, columnFilters),
  onColumnVisibilityChange: updaterOrValue => valueUpdater(updaterOrValue, columnVisibility),
  onRowSelectionChange: updaterOrValue => valueUpdater(updaterOrValue, rowSelection),
  onExpandedChange: updaterOrValue => valueUpdater(updaterOrValue, expanded),
  state: {
    get sorting() { return sorting.value },
    get columnFilters() { return columnFilters.value },
    get columnVisibility() { return columnVisibility.value },
    get rowSelection() { return rowSelection.value },
    get expanded() { return expanded.value },
  },
})

const getMinisectorColor = (segment: number) => {
  // None SegmentType = iota
  // YellowSegment
  // GreenSegment
  // InvalidSegment // Doesn't get displayed, cut corner/boundaries or invalid segment time?
  // PurpleSegment
  // RedSegment     // After chequered flag/stopped on track
  // PitlaneSegment // In pitlane
  // Mystery
  // Mystery2 // ??? 2021 - Turkey Practice_2
  // Mystery3 // ??? 2020 - Italy Race
  switch (segment) {
    case 1:
      return "bg-yellow-500";
    case 2:
      return "bg-green-500";
    case 3:
      return "bg-black-500";
    case 4:
      return "bg-violet-500";
    case 5:
      return "bg-red-500";
    case 6:
      return "bg-white";
    case 7:
    case 8:
    case 9:
      return "bg-blue-500";
    default:
      return "bg-white";
  }
};

const parseDuration = (duration: number = 0): string => {
  if (duration == 0) {
    return "";
  }

  const parsed = moment(duration)
  if (parsed.minutes() <= 0) {
    return moment(duration).format('ss.SSS');
  }

  return moment(duration).format('mm:ss.SSS');
};

const getCarPosition = (carPosition: number) => {
  // NoLocation CarLocation = iota
  // Pitlane
  // PitOut
  // OutLap
  // OnTrack
  // OutOfRace
  // Stopped
  switch (carPosition) {
    case 1:
      return "Pit lane";
    case 2:
      return "Pit Exit";
    case 3:
      return "Out Lap";
    case 4:
      return "On Track";
    case 5:
      return "Out";
    case 6:
      return "Stopped";
    default:
      break;
  }
};

const getDriverTelemetry = (driver: Driver) => {
  return useTelemetryStore.telemetry.find(
    (t) => t.DriverNumber === driver.Number
  );
};

const parseTimeColor = (personalFastest: boolean, overallFastest: boolean) => {
  if (overallFastest) {
    return "text-purple-500";
  } else if (personalFastest) {
    return "text-green-500";
  } else {
    return "text-yellow-500";
  }
};

const parseTireColor = (tire: number) => {
  switch (tire) {
    case 1:
      return "text-red-500";
    case 2:
      return "text-yellow-500";
    case 3:
      return "text-white";
    case 4:
      return "text-green-500";
    case 5:
      return "text-blue-500";
    default:
      return "text-orange-500";
  }
};

const parseTireType = (tire: number) => {
  // Unknown TireType = iota
  // Soft
  // Medium
  // Hard
  // Intermediate
  // Wet
  // Test
  // HYPERSOFT
  // ULTRASOFT
  // SUPERSOFT
  switch (tire) {
    case 0:
      return "Unknown"
    case 1:
      return "Soft"
    case 2:
      return "Medium"
    case 3:
      return "Hard"
    case 4:
      return "Intermediate"
    case 5:
      return "Wet"
    case 6:
      return "Test"
    case 7:
      return "HYPERSOFT"
    case 8:
      return "ULTRASOFT"
    case 9:
      return "SUPERSOFT"
    default:
      break;
  }
}

const parseLastPitTime = (pitStopTimes: PitStop[]) => {
  if (pitStopTimes && pitStopTimes.length > 0) {
    let lastPitlane = pitStopTimes[pitStopTimes.length - 1];

    if (lastPitlane.PitlaneTime != 0) {
      return parseDuration(lastPitlane.PitlaneTime);
    }
  }
};

const getPilanePosition = (driver: Driver) => {
  let potentialPositionChange: number = 0;
  let positionColor: string = "text-green-500";

  if (
    driver.Location != DriverLocation.Pitlane &&
    driver.Location != DriverLocation.PitOut
  ) {
    let newPosition = driver.Position;
    let pitTimeLost =
      useGeneralStore.general.TimeLostInPitlane +
      useGeneralStore.general.TimeLostInPitlane;
    // Default value for the last car because we won't enter the loop
    let timeToCarAhead = pitTimeLost;
    let timeToCarBehind = 0;

    let timeToClearPitlane = 1000 * 1000 * 10;
    let gapToCar = 1000 * 1000 * 0;
    let minGap = pitTimeLost - timeToClearPitlane;

    for (
      let driverBehind = 1;
      driverBehind < useDriverStore.drivers.length;
      driverBehind++
    ) {
      //timeToCarAhead = pitTimeLost

      // Can't drop below stopped cars
      if (
        useDriverStore.drivers[driverBehind].Location ==
        DriverLocation.Stopped ||
        useDriverStore.drivers[driverBehind].Location ==
        DriverLocation.OutOfRace
      ) {
        break;
      }

      newPosition++;

      gapToCar += useDriverStore.drivers[driverBehind].TimeDiffToPositionAhead;

      // If the gap to the prediction car is less than the time to drive past the pitlane then keep looking
      if (gapToCar < minGap) {
        continue;
      }

      timeToCarBehind = gapToCar - minGap;
      timeToCarAhead =
        useDriverStore.drivers[driverBehind].TimeDiffToPositionAhead -
        timeToCarBehind;
      break;
    }

    if (newPosition == driver.Position) {
      timeToCarAhead += driver.TimeDiffToPositionAhead;
    }

    if (
      driver.Location != DriverLocation.Stopped &&
      driver.Location != DriverLocation.OutOfRace
    ) {
      // potentialPositionChange = fmt.Sprintf("%02d", newPosition)
      potentialPositionChange = newPosition;
    }

    if (newPosition != driver.Position) {
      // positionColor = colornames.Red
      positionColor = "text-red-500";
    }
  } else {
    positionColor = "text-white";
    potentialPositionChange = 0;
  }

  return { potentialPositionChange, positionColor };
};
</script>

<template>
  <div class="w-full">
    <div class="flex gap-2 items-center py-4">
      <!-- <Button @click="randomize">
        Randomize
      </Button> -->
      <DropdownMenu>
        <DropdownMenuTrigger as-child>
          <Button variant="outline" class="ml-auto">
            Columns
            <ChevronDown class="ml-2 h-4 w-4" />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuCheckboxItem v-for="column in table.getAllColumns().filter((column) => column.getCanHide())"
            :key="column.id" class="capitalize" :model-value="column.getIsVisible()" @update:model-value="(value) => {
              column.toggleVisibility(!!value)
            }">
            {{ column.id }}
          </DropdownMenuCheckboxItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
    <div class="rounded-md border">
      <Table>
        <TableHeader>
          <TableRow v-for="headerGroup in table.getHeaderGroups()" :key="headerGroup.id">
            <TableHead v-for="header in headerGroup.headers" :key="header.id">
              <FlexRender v-if="!header.isPlaceholder" :render="header.column.columnDef.header"
                :props="header.getContext()" />
            </TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <template v-if="table.getRowModel().rows?.length">
            <template v-for="row in table.getRowModel().rows" :key="row.id">
              <TableRow :data-state="row.getIsSelected() && 'selected'">
                <TableCell v-for="cell in row.getVisibleCells()" :key="cell.id">
                  <FlexRender :render="cell.column.columnDef.cell" :props="cell.getContext()" />
                </TableCell>
              </TableRow>
              <TableRow v-if="row.getIsExpanded()">
                <TableCell :colspan="row.getAllCells().length">
                  {{ JSON.stringify(row.original) }}
                </TableCell>
              </TableRow>
            </template>
          </template>

          <TableRow v-else>
            <TableCell :colspan="columns.length" class="h-24 text-center">
              No results.
            </TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </div>

    <div class="flex items-center justify-end space-x-2 py-4">
      <div class="flex-1 text-sm text-muted-foreground">
        {{ table.getFilteredSelectedRowModel().rows.length }} of
        {{ table.getFilteredRowModel().rows.length }} row(s) selected.
      </div>
      <div class="space-x-2">
        <Button variant="outline" size="sm" :disabled="!table.getCanPreviousPage()" @click="table.previousPage()">
          Previous
        </Button>
        <Button variant="outline" size="sm" :disabled="!table.getCanNextPage()" @click="table.nextPage()">
          Next
        </Button>
      </div>
    </div>
  </div>
</template>
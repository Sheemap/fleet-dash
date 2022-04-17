<script setup lang="ts">
import SummaryShipCard from './SummaryShipCard.vue';
import {useFleetStore} from "../js/fleetStore";
import {computed, getCurrentInstance, Ref, ref} from "vue";
import {useUserStore} from "../js/userStore";
import SummaryTextItemWithRef from "./SummaryTextItemWithRef.vue";

const keyStr = getCurrentInstance()?.vnode.key;
const id = typeof keyStr === "string" ? keyStr.split("_")[1] : "";

type SummaryShip = {
  id: number;
  count: number;
}

type SummaryLocation = {
  id: number;
  count: number;
  location_name: string;
}

const fleetStore = useFleetStore();
const userStore = useUserStore();

let fleetBoss = computed(() : boolean => fleetStore.able_to_fetch_members);

let shipSummary = computed(() : SummaryShip[] => {
  let ships = fleetStore.members.map(m => m.ship_type_id);
  let counts = ships.reduce((acc, shipId) => {
    acc[shipId] = (acc[shipId] || 0) + 1;
    return acc;
  }, {});
  return Object.entries(counts).map(([shipId, count]) => ({id: parseInt(shipId), count} as SummaryShip)).sort((a, b) => b.count - a.count);
});

let locationSummary = computed(() => {
  let locations = (fleetStore.members || []).map(m => m.solar_system_id);
  let counts = locations.reduce((acc, locationId) => {
    acc[locationId] = (acc[locationId] || 0) + 1;
    return acc;
  }, {});

  return Object.entries(counts).map(([locationId, count]) => ({id: parseInt(locationId), count} as SummaryLocation)).sort((a, b) => b.count - a.count);
});

function getSystemName(systemId: number) : Ref<string> {
  let systemName = ref('');

  userStore.getActiveToken().then(token =>{
    fleetStore.fetchSystemName(systemId, token).then(name => {
      systemName.value = name;
    });
  })

  return systemName;
}

function removeFromGrid(): void {
  userStore.dash_layout = userStore.dash_layout.filter((item) => !item.i.endsWith(id));
  localStorage.removeItem("widget_settings_" + id);
}
</script>

<template>
    <div class="border rounded drop-shadow bg-zinc-800 h-full">
      <svg v-if="userStore.widget_drawer_open && id !== ''"
           @click="removeFromGrid"
           class="fixed z-50 top-0 right-0 mr-1 config-button cursor-pointer transition ease-in-out opacity-70 hover:opacity-100"
           xmlns="http://www.w3.org/2000/svg"
           height="30px"
           viewBox="0 0 24 24"
           width="30px"
           fill="#DC2626">
        <path d="M0 0h24v24H0z" fill="none"/>
        <path d="M19 6.41L17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 13.41 12z"/>
      </svg>

      <div v-if="fleetBoss">
      <div class="w-full h-7 px-2 flex flex-row overflow-clip">
        <span class="max-h-2 text-xl flex-grow">Fleet Overview</span>
      </div>
      <hr/>

      <div class="h-full grid grid-rows-2 w-full px-2 gap-2 pt-2 pb-8 min-h-0 min-w-0 md:fixed">
        <div>
          <div class="h-full w-full flex flex-col overflow-auto flex-nowrap xl:flex-wrap">
            <SummaryShipCard v-for="ship in shipSummary" :key="ship.id" :ship-id=ship.id :count=ship.count />
          </div>
        </div>

        <div>
            <span>Locations</span>
            <hr/>
            <div class="overflow h-5/6 flex flex-col overflow-auto flex-nowrap lg:flex-wrap">
              <SummaryTextItemWithRef v-for="location in locationSummary" :key=location.id :text=getSystemName(location.id) :count=location.count />
            </div>
          </div>
        </div>
        </div>
      <div v-else class="py-5 text-zinc-400 italic text-center">Only the FleetBoss can access the overview</div>
    </div>
</template>

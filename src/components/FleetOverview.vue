<script setup lang="ts">
import SummaryShipCard from './SummaryShipCard.vue';
import SummaryTextItem from './SummaryTextItem.vue';
</script>

<template>
    <div class="border rounded drop-shadow bg-zinc-800 h-full">
      <div class="w-full h-7 px-2 flex flex-row overflow-clip">
        <span class="max-h-2 text-xl flex-grow">Fleet Overview</span>
        <span class="max-h-2 text-l">Docked: </span>

      </div>
      <hr/>

      <div class="max-h-full grid grid-rows-2 w-full px-2 gap-2 pt-2 pb-8 min-h-0 min-w-0 md:fixed">
        <div>
          <div class="h-full w-full flex flex-col overflow-auto flex-nowrap xl:flex-wrap">
            <SummaryShipCard ship-id="587" count="7" />
            <SummaryShipCard ship-id="587" count="7" />
            <SummaryShipCard ship-id="587" count="7" />
            <SummaryShipCard ship-id="587" count="7" />
            <SummaryShipCard ship-id="587" count="7" />
            <SummaryShipCard ship-id="587" count="7" />
            <SummaryShipCard ship-id="587" count="7" />
            <SummaryShipCard ship-id="587" count="7" />
            <SummaryShipCard ship-id="587" count="7" />
            <SummaryShipCard ship-id="587" count="7" />
            <SummaryShipCard ship-id="587" count="7" />
            <SummaryShipCard ship-id="587" count="7" />
            <SummaryShipCard ship-id="587" count="7" />
            <SummaryShipCard ship-id="587" count="7" />
            <SummaryShipCard ship-id="587" count="7" />
          </div>
        </div>

        <div class="">
            <span class="fixed sticky">Locations</span>
            <hr/>
            <div class="overflow h-5/6 flex flex-col overflow-auto flex-nowrap lg:flex-wrap">
              <SummaryTextItem count="3" text="LX-Q" />
              <SummaryTextItem count="3" text="LX-Q" />
              <SummaryTextItem count="3" text="LX-Q" />
              <SummaryTextItem count="3" text="LX-Q" />
              <SummaryTextItem count="3" text="LX-Q" />
              <SummaryTextItem count="3" text="LX-Q" />
              <SummaryTextItem count="3" text="LX-Q" />
              <SummaryTextItem count="3" text="LX-Q" />
              <SummaryTextItem count="3" text="LX-Q" />
              <SummaryTextItem count="3" text="LX-Q" />
              <SummaryTextItem count="3" text="LX-Q" />
              <SummaryTextItem count="3" text="LX-Q" />
            </div>
          </div>
        </div>
    </div>
</template>

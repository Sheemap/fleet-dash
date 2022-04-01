<script setup lang="ts">
import { useUserStore } from "../js/userStore";
import { useFleetStore } from "../js/fleetStore";
import { useEventStore} from "../js/eventStore";
import {ref} from "vue";

const userStore = useUserStore();
const fleetStore = useFleetStore();
const eventStore = useEventStore();

let running = ref(false);
let errored = ref(false);

async function updateFleet(failCount: number) {
  if (!eventStore.active) {
    running.value = false;
    return;
  }

  running.value = true;

  try {
    let token = await userStore.getActiveToken();
    await fleetStore.fetchFleet(userStore.character_id, token);
    await fleetStore.fetchMembers(token);

    // 5 seconds is the lowest cache time for these APIs
    setTimeout(updateFleet, 1000 * 5, 0);

  } catch (e) {
    console.error(e);
    if (failCount > 5) {
      running.value = false;
      errored.value = true;
      alert('Failed to update fleet info! Please refresh the page to retry. If the problem persists, please contact the developer.');
      return;
    }
    // Exponential backoff
    setTimeout(updateFleet, Math.pow(2, failCount) * 5000, failCount + 1);
  }
}

while (!errored.value) {
  if (!running.value && eventStore.active) {
    updateFleet(0);
  }

  await new Promise(resolve => setTimeout(resolve, 1000));
}

</script>
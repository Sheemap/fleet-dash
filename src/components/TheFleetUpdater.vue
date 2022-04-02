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
let nextTryFetchMember = ref(0);

async function updateFleet(failCount: number) {
  if (!eventStore.active) {
    running.value = false;
    return;
  }

  running.value = true;

  try {
    let token = await userStore.getActiveToken();
    await fleetStore.fetchFleet(userStore.character_id, token);

    // Can only fetch members if we are a fleet commander
    if(nextTryFetchMember.value < Date.now()) {
      try{
        await fleetStore.fetchMembers(token);
        fleetStore.able_to_fetch_members = true;
      } catch(e) {
        // If we fail to fetch members, we will try again in 1 minute
        nextTryFetchMember.value = Date.now() + 60000;
        fleetStore.able_to_fetch_members = false;
      }
    }

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
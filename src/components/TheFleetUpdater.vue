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

let lock = ref(false);
async function updateFleet() {
  if (lock.value){
    return;
  }

  lock.value = true;

  if (!eventStore.active) {
    running.value = false;
    lock.value = false;
    return;
  }

  running.value = true;
  try{
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
  } catch (e) {
    // Ignore, likely because we are not in a fleet
  }

  lock.value = false;
}
let interval;
while (!errored.value) {
  if (!running.value && eventStore.active) {
    updateFleet();
    interval = setInterval(updateFleet, 5000);
  }
  if (!eventStore.active && interval) {
    clearInterval(interval);
  }

  await new Promise(resolve => setTimeout(resolve, 1000));
}

</script>
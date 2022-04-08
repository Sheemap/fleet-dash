<script setup lang="ts">
import { useUserStore } from "../js/userStore";
import { useFleetStore } from "../js/fleetStore";
import { useEventStore} from "../js/eventStore";
import useToast from "../js/toast";
import {ref} from "vue";

const userStore = useUserStore();
const fleetStore = useFleetStore();
const eventStore = useEventStore();

const toast = useToast();

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
    let fleet = await fleetStore.fetchFleet(userStore.character_id, token);

    // If fleet differs from stored fleet, shutdown session
    // This will trigger a new session to be joined (if it exists)
    if (fleet.fleet_id !== eventStore.session_fleet_id){
      eventStore.shutdownLocalStream();
      toast.info("Fleet has changed, rejoin session to continue.", {
        duration: 10000
      });
      lock.value = false;
      return;
    }

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
    if (e === "Character is not in a fleet"){
      eventStore.shutdownLocalStream();
      toast.info("Fleet has changed, rejoin session to continue.", {
        duration: 10000
      });
    }
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
<script setup>
  import {useUserStore} from "../js/userStore";
  import {ref} from "vue";

  const userStore = useUserStore();

  let authUrl = ref('');
  let authenticated = ref(false);
  let portrait_url = ref('');

  userStore.getActiveToken().then(_ =>{
      authenticated.value = true;
      userStore.getPortraitUrl().then(res =>{
        portrait_url.value = res;
      })
  })
  .catch(_ =>{
    // No token stored
    authenticated.value = false;
    let auth_state = userStore.setState();
    authUrl.value =`https://login.eveonline.com/v2/oauth/authorize?response_type=code` +
        `&redirect_uri=${import.meta.env.VITE_EVE_REDIRECT_URI}` +
        `&client_id=${import.meta.env.VITE_EVE_CLIENT_ID}` +
        `&scope=${import.meta.env.VITE_EVE_SCOPES}` +
        `&state=${auth_state}`;
  })
</script>

<template>
  <div v-if="authenticated">
    <img v-if="authenticated" v-on:click="this.toggleContextMenu" class="my-3 mr-5 rounded-full cursor-pointer" v-bind:src="portrait_url" alt="Character Portrait"/>
    <div v-if="contextMenu" class="bg-white text-black z-50 list-none divide-y divide-gray-100 rounded shadow my-4 fixed right-0 top-14 mx-5">
      <div class="px-4 py-3">
        <span class="block text-sm cursor-default">{{ this.character_name }}</span>
      </div>
      <ul class="py-1" aria-labelledby="dropdown">
        <li>
          <a v-on:click="this.signOut" class="text-sm hover:bg-gray-200 text-gray-700 block px-4 py-2 cursor-pointer">Sign out</a>
        </li>
      </ul>
    </div>
  </div>
  <a v-else v-bind:href="authUrl" target="_blank"><img class="m-5" src="../assets/eve-login.png" alt="Login with EVE Online" /></a>

</template>

<script>
import { ref } from 'vue';
import { useUserStore } from "../js/userStore";

export default {
  data(){
    return {
      contextMenu: false,
    }
  },
  methods: {
    toggleContextMenu(){
      this.contextMenu = !this.contextMenu;
    },
    signOut(){
      this.userStore.signOut();
      window.location.href = '/';
    }
  },
  computed: {
    character_name(){
      return this.userStore.character_name;
    }
  }
}
</script>

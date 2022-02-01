<template>
  <img v-if="authenticated" class="my-auto mr-5 rounded-full" v-bind:src="portrait_url" alt="Character Portrait"/>
  <a v-else v-bind:href="authUrl" target="_blank"><img class="m-5" src="../assets/eve-login.png" alt="Login with EVE Online" /></a>
</template>

<script>
import { ref } from 'vue';
import { useUserStore } from "../js/store";

export default {
  name: "TheAuthenticator",
  setup(props, context){
    const userStore = useUserStore();

    let authUrl = ref('');
    let authenticated = ref(false);
    let portrait_url = ref('');

    userStore.getActiveToken().then(res =>{
      if (Object.keys(res).length === 0){
        authenticated.value = false;
        let auth_state = userStore.setState();
        authUrl.value =`https://login.eveonline.com/v2/oauth/authorize?response_type=code` +
            `&redirect_uri=${import.meta.env.VITE_EVE_REDIRECT_URI}` +
            `&client_id=${import.meta.env.VITE_EVE_CLIENT_ID}` +
            `&scope=${import.meta.env.VITE_EVE_SCOPES}` +
            `&state=${auth_state}`;
      }else{
        authenticated.value = true;
        userStore.getPortraitUrl().then(res =>{
          portrait_url.value = res;
        })
      }
    })
    return { userStore, authenticated, authUrl, portrait_url };
  }
}
</script>

<style scoped>

</style>

<template>
  <h1 v-if="errorMessage">Error has occurred. Please try again.</h1>
</template>

<script>
import { useUserStore } from "../js/store";
import { inject, ref } from 'vue';

export default {
  name: "EveSSOCallback",
  props:{
    queryParams: {
      type: Object,
      default: () => {
        return {};
      }
    }
  },
  setup(props, context) {
    const userStore = useUserStore();
    const axios = inject('axios')

    let authState = props.queryParams.find(item => typeof(item.state) !== 'undefined');
    if (typeof(authState) === 'undefined' ||
        userStore.auth_state !== authState.state){
        const errorMessage = 'Could not validate state parameter.'
        console.error(errorMessage);
        return { errorMessage };
    }
    userStore.resetState();

    let authCode = props.queryParams.find(item => typeof(item.code) !== 'undefined');
    if (typeof(authCode) === 'undefined'){
        const errorMessage = 'Did not receive validate code parameter.'
        console.error(errorMessage);
        return { errorMessage };
    }

    let errorMessage = ref(false);
    const params = new URLSearchParams();
    params.append('client_id', import.meta.env.VITE_EVE_CLIENT_ID);
    params.append('client_secret', import.meta.env.VITE_EVE_CLIENT_SECRET);
    params.append('grant_type', 'authorization_code');
    params.append('code', authCode.code);
    params.append('redirect_uri', import.meta.env.VITE_EVE_REDIRECT_URI);


    axios.post('https://login.eveonline.com/v2/oauth/token',params).then(response =>{
      let tokenSet = userStore.setToken(response.data)


    }).catch(error => {
      console.error(error);
      errorMessage.value = true;
    });

    return { errorMessage }
  },
}
</script>

<style scoped>

</style>

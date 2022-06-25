<template>
  <h1 v-if="errorMessage">Error has occurred. Please try again.</h1>
</template>

<script lang="ts">
import { useUserStore } from "../js/userStore";
import { defineComponent, ref } from 'vue';

const component = defineComponent({
  props:{
    queryParams: {
      type: Object,
      default: () => {
        return {};
      }
    }
  },
  setup(props, _) {
    const userStore = useUserStore();

    let authState = props.queryParams.find(item => typeof(item.state) !== 'undefined');
    if (typeof(userStore.auth_state) === 'undefined' ||
        typeof(userStore.auth_state.state) === 'undefined' ||
        userStore.auth_state.state !== authState.state){

        const errorMessage = 'Could not validate state parameter.'
        console.error(errorMessage);
        return { errorMessage };
    }

    let authCode = props.queryParams.find(item => typeof(item.code) !== 'undefined');
    if (typeof(authCode) === 'undefined'){
        const errorMessage = 'Did not receive validate code parameter.'
        console.error(errorMessage);
        return { errorMessage };
    }

    let errorMessage = ref(false);
    const params = new URLSearchParams();
    params.append('client_id', import.meta.env.VITE_EVE_CLIENT_ID);
    params.append('code_verifier', userStore.auth_state.code_verifier);
    params.append('grant_type', 'authorization_code');
    params.append('code', authCode.code);
    params.append('redirect_uri', import.meta.env.VITE_EVE_REDIRECT_URI);

    const options = {
      method: 'POST',
      body: params,
    }
    fetch('https://login.eveonline.com/v2/oauth/token', options)
        .then(response => response.json())
        .then(response =>{
          userStore.setToken(response)
          window.location.href = '/';
        }).catch(error => {
          console.error(error);
          errorMessage.value = true;
        })
        .finally(() => {
          userStore.resetState();
        });

    return { errorMessage }
  },
});

export default component;
</script>
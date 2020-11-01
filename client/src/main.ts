import Vue from 'vue'
import App from './App.vue'
import router from './router'

import { Auth0Plugin } from "./auth";
import { domain, clientId, redirectUri } from '../auth0_config.json'

// Install the authentication plugin here
Vue.use(Auth0Plugin, {
  domain,
  clientId,
  redirectUri,
  onRedirectCallback: appState => {
    router.push(
      appState && appState.targetUrl
        ? appState.targetUrl
        : window.location.pathname
    );
  }
});

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')

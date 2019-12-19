// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'
import i18n from './localization'

import VueI18n from 'vue-i18n';
Vue.use(VueI18n);

import BootstrapVue from 'bootstrap-vue'
Vue.use(BootstrapVue);
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

import VueResorce from 'vue-resource'
Vue.use(VueResorce);

import Vuex from 'vuex'
Vue.use(Vuex);

import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { library } from '@fortawesome/fontawesome-svg-core'

import { faTrash } from '@fortawesome/free-solid-svg-icons'
import { faEdit } from '@fortawesome/free-solid-svg-icons'
import { faEye } from '@fortawesome/free-solid-svg-icons'
import { faPassport } from '@fortawesome/free-solid-svg-icons'
import { faTimesCircle } from '@fortawesome/free-solid-svg-icons'

library.add(faTrash,faEdit,faEye,faPassport,faTimesCircle);


Vue.component('font-awesome-icon', FontAwesomeIcon);

Vue.config.productionTip = false;

const API = 'http://localhost:5000';
//const API = 'http://178.62.14.228:5001';

//const API = process.env.ROOT_API;
const API_URL = API+'/api/';
//Vue.use(API);
//Vue.use(API_URL);

Vue.http.options.root = API_URL;

Vue.http.interceptors.unshift((request, next) => {
  let url = request.url;

  request.url = API_URL+url;

  console.log(request.url);
  let token = localStorage.getItem("token");
  if (token) {
    request.headers.set('Authorization', 'Bearer '+ token);
    //request.headers.set('Content-Type', 'application/json');
  }else {
    router.push('/login');
  }

  console.log(request);
  next((response)=>{
    if (url !== "login"){
      if (response.status === 401){
        store.dispatch("logout").then(res => {
          router.push('/login');
        })
      }
    }
  })
});

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  i18n,
  components: { App },
  template: '<App/>'
});

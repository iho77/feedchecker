// main.js

import Vue from 'vue'
import VueRouter from 'vue-router';
import VueAxios from 'vue-axios';
import axios from 'axios';
import NProgress from 'nprogress';


import VueMaterial from 'vue-material'
import 'vue-material/dist/vue-material.min.css'
import 'vue-material/dist/theme/default.css'

import TreeView from "vue-json-tree-view"


import App from './App.vue';
import Edit from './components/Edit.vue';
import Search from './components/Search.vue';
import Ip from './components/Ip.vue';
import Worker from './components/Worker.vue';
import Domain from './components/Domain.vue';
import URL from './components/URL.vue';

//import '../node_modules/bootstrap/dist/css/bootstrap.min.css';
import '../node_modules/nprogress/nprogress.css';

Vue.use(VueMaterial);
Vue.use(VueRouter);
Vue.use(VueAxios, axios);
Vue.use(TreeView)

Vue.config.productionTip = false;

const routes = [
  {
    name: 'Worker',
    path: '/worker',
    component: Worker
  },
  {
    name: 'Domain',
    path: '/domain',
    component: Domain
  },
  {
    name: 'Ip',
    path: '/index',
    component: Ip
  },
  {
    name: 'URL',
    path: '/URL',
    component: URL
  },
  {
    name: 'Edit',
    path: '/edit/:id',
    component: Edit
  },
  {
    name: 'Search',
    path: '/search',
    component: Search
  },
];

const router = new VueRouter({ mode: 'history', routes: routes });

router.beforeResolve((to, from, next) => {
  if (to.name) {
      NProgress.start()
  }
  next()
});

router.afterEach(() => {
  NProgress.done()
});

new Vue({
  render: h => h(App),
  router
}).$mount('#app')
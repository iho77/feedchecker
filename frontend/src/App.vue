// App.vue
 
<template>

<div class="page-container">
    <md-app md-waterfall md-mode="fixed">

      <md-app-toolbar class="md-primary">
        <div class="md-toolbar-row">
          <div class="md-toolbar-section-start">
            <md-field md-inline>
             <label>Search for IOC</label>
             <md-input v-model="ioc"></md-input>
           </md-field>
          </div>  
       
       <div class="md-layout md-gutter">
      <div class="md-layout-item">
          <md-field>
         <label for="feed">Source feed</label>
           <md-select v-model="feed" name="feed" id="feed" placeholder="feed collection">
            <md-option value="dns">dns</md-option>
            <md-option value="ip">ip</md-option>
            <md-option value="url">url</md-option>
          </md-select>
          </md-field>
      
      </div>
       </div>
     
      <md-button class="md-icon-button" @click="$router.push({name: 'Search', params: { feed: feed, name: ioc  },})">
            <md-icon>refresh</md-icon>
           
          </md-button>
         </div>
      
       
        
      </md-app-toolbar>

      <md-app-drawer md-permanent="full">
        <md-toolbar class="md-transparent" md-elevation="1">
          Navigation
        </md-toolbar>

        <md-list>
          <md-list-item>
            <router-link :to="{ name: 'Ip' }" class="md-list-item-text">IP IOCs</router-link>
          </md-list-item>

          <md-list-item>
            <router-link :to="{ name: 'Domain' }" class="md-list-item-text">DNS IOCs</router-link>
          </md-list-item>

          <md-list-item>
            <router-link :to="{ name: 'URL' }" class="md-list-item-text">URL IOCs</router-link>
          </md-list-item>

          <md-list-item>
           <router-link :to="{ name: 'Worker' }" class="md-list-item-text">Workers list</router-link>
          </md-list-item>
        </md-list>
      </md-app-drawer>
      
      <md-app-content>
          <router-view></router-view>     
          <nprogress-container></nprogress-container>  
      </md-app-content>
      
    </md-app>
  </div>

</template>

<style lang="scss" scoped>
  .md-app {
    max-height: 800px;
    border: 1px solid rgba(#000, .12);
  }

   // Demo purposes only
  .md-drawer {
    width: 130px;
    max-width: calc(100vw - 125px);
  }

  .search {
    max-width: 500px;
  }

</style>

<script>

export default {
  name: 'Normal'
}
</script>

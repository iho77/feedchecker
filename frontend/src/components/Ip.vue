// Index.vue



<template>
    <md-table v-model="items" md-sort="item.feed[0].score.total" md-sort-order="desc" md-card>
      <md-table-toolbar>
        <h1 class="md-title">IP IOC Alert list</h1>
      </md-table-toolbar>

      <md-table-row slot="md-table-row" slot-scope="{ item }">
        <md-table-cell md-label="IP" md-sort-by="domain">{{ item.domain }}</md-table-cell>
        <md-table-cell md-label="Count" md-sort-by="count">{{ item.count }}</md-table-cell>
        <md-table-cell md-label="FirstSeen" md-sort-by="firstseen ">{{ item.firstseen }}</md-table-cell>
        <md-table-cell md-label="LastSeen" md-sort-by="lastseen">{{ item.lastseen }}</md-table-cell>
        <md-table-cell md-label="Description" md-sort-by="feed[0].description">{{ item.feed[0].description }}</md-table-cell>
        <md-table-cell md-label="Score" md-sort-by="feed[0].score.total">{{ item.feed[0].score.total }}</md-table-cell>
        <md-table-cell md-label="CVE" md-sort-by="feed[0].cve">{{ item.feed[0].cve }}</md-table-cell>
        <md-table-cell md-label="Related" md-sort-by="feed[0].related.domains">{{ item.feed[0].related.domains }}</md-table-cell>
        <md-table-cell><router-link :to="{name: 'Edit', params: { id: item.feed[0]._id, feed: 'ip' }}" class="md-raised md-primary">info</router-link></md-table-cell>
      </md-table-row>
     </md-table>
</template>

<script>

 

    export default {
        data(){
            return{
                items: []
            }
        },

        created: function()
        {
            this.fetchItems();
        },
 
        methods: {
            fetchItems()
            { 
                
              let uri = '/api/iocs/ip';
              this.axios.get(uri).then((response) => {
                  this.items = response.data;
              });
            }
        }
    }
</script>
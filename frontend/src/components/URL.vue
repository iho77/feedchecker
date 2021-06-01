// Index.vue

<template>

<md-table v-model="items" md-sort="item.score.total" md-sort-order="desc" md-card>
      <md-table-toolbar>
        <h1 class="md-title">Domain Alarms list</h1>
      </md-table-toolbar>

      <md-table-row slot="md-table-row" slot-scope="{ item }">
        <md-table-cell md-label="URL" md-sort-by="domain">{{ item.domain }}</md-table-cell>
        <md-table-cell md-label="Count" md-sort-by="count">{{ item.count }}</md-table-cell>
        <md-table-cell md-label="FirstSeen" md-sort-by="firstseen ">{{ item.firstseen }}</md-table-cell>
        <md-table-cell md-label="LastSeen" md-sort-by="lastseen">{{ item.lastseen }}</md-table-cell>
        <md-table-cell v-if="item.feed[0].res != 'no data'" md-label="URL" md-sort-by="url">{{ item.url }}</md-table-cell>
         <md-table-cell v-else>none</md-table-cell>
        <md-table-cell v-if="item.feed[0].res != 'no data'" md-label="Description" md-sort-by="description">{{ item.description }}</md-table-cell>
         <md-table-cell v-else>none</md-table-cell>
        <md-table-cell v-if="item.feed[0].res != 'no data'" md-label="Score" md-sort-by="score.total">{{ item.score.total }}</md-table-cell>
         <md-table-cell v-else>none</md-table-cell>
        <md-table-cell v-if="item.feed[0].res != 'no data'" md-label="CVE" md-sort-by="cve">{{ item.cve }}</md-table-cell>
         <md-table-cell v-else>none</md-table-cell>
        <md-table-cell v-if="item.feed[0].res != 'no data'" md-label="Domain" md-sort-by="parsed.domain">{{ item.parsed.domain }}</md-table-cell>
         <md-table-cell v-else>none</md-table-cell>
        <md-table-cell v-if="item.feed[0].res != 'no data'" ><router-link :to="{name: 'Edit', params: { id: item._id, feed: 'url' }}" class="md-raised md-primary">info</router-link></md-table-cell>
         <md-table-cell v-else>none</md-table-cell>
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
              let uri = '/api/iocs/url';
              this.axios.get(uri).then((response) => {
                  this.items = response.data;
              });
            }
        }
    }
</script>
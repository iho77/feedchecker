// Index.vue

<template>

<md-table v-model="items"  md-card>
      <md-table-toolbar>
        <h1 class="md-title">Workerks list</h1>
      </md-table-toolbar>

      <md-table-row slot="md-table-row" slot-scope="{ item }">
        <md-table-cell md-label="Name" >{{ item.name }}</md-table-cell>
        <md-table-cell md-label="URI" >{{ item.uri }}</md-table-cell>
        <md-table-cell md-label="Tag" >{{ item.tag }}</md-table-cell>
        <md-table-cell md-label="Filter" >{{ item.filter }}</md-table-cell>
        <md-table-cell md-label="InTopic" >{{ item.intopic }}</md-table-cell>
        <md-table-cell md-label="OutTopic" >{{ item.outtopic }}</md-table-cell>
        <md-table-cell md-label="Lag" >{{ item.lag }}</md-table-cell>
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
              let uri = '/api/workers';
              this.axios.get(uri).then((response) => {
                  this.items = response.data;
              });
            }
        }
    }
</script>
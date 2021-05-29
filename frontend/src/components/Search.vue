// Edit.vue
<template>
  <div class="container">
        <div class="card">
           <pre> {{item}} </pre>
        </div>
    </div>
</template>
<script>


export default{
        data(){
            return{
                item:{}
            }
        },


        created: function(){
            this.getItem();
        },

        methods: {
            getItem()
            {  //  + this.$route.params.id
              let uri = '/api/ioc/searchbyname/' + this.$route.params.feed +'/'+ this.$route.params.name;
                this.axios.get(uri).then((response) => {
                    this.item = response.data;
                });
            },
        }
    }
</script>
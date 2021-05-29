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
              let uri = '/api/ioc/searchbyid/' + this.$route.params.feed +'/'+ this.$route.params.id;
                this.axios.get(uri).then((response) => {
                    this.item = response.data;
                });
            },
        }
    }
</script>
<template>
  <b-form-checkbox  @click.native.stop @change="changeActiveToday(row)" v-model="row.item.active_today"></b-form-checkbox>
</template>

<script>
  export default {
    name: 'WorkerChangeActive',
    props:["row"],
    data() {
      return {
        formUrl: 'workers',
      }
    },
    created() {
      this.formUrl = "town/"+this.$route.meta.town+"/"+"workers";
    },
    methods: {
      put: function (url, data, callback) {
        return this.$http.put(url, data, null).then(result => {
          callback(result);
        }, error => {
          callback(error);
        });
      },
      changeActiveToday(row) {
        row.item.active_today = !row.item.active_today;
        let self;
        let url = this.formUrl;
        if (row.item.id > 0) {
          self = this;
          url = this.formUrl + '/' + row.item.id;
          self.put(url, row.item, function (result) {
            console.log(result);
            if (result.status === 200) {

            }
          }, error => {
            console.log("ERROR", error);
          });
        }
      },
    },
  }
</script>

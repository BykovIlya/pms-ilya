<template>
  <div>
    <b-table id="workers"
             ref="workers"
             striped
             small
             thead-class="fixed-top sticky-top top-fixed-thead"
             responsive
             show-empty
             :busy.sync="isBusy"
             :items="items"
             :fields="fields"
             :current-page="currentPage"
             :per-page="perPage"
             :empty-text="$t('message.common.emptyTable')"
             :sort-by.sync="sortBy"
             :sort-desc.sync="sortDesc"
             :filter="filter"
             @filtered="onFiltered"
    >
      <template slot="action" slot-scope="row">
        <b-button size="sm" variant="success" @click.stop="showModal(row,true)">
          <font-awesome-icon icon="eye" />
        </b-button>
        <b-btn  size="sm" variant="primary" @click="showModal(row,false)">
          <font-awesome-icon icon="edit" />
        </b-btn>
        <b-btn  size="sm" variant="danger" @click="showModalDelete(row)">
          <font-awesome-icon icon="trash" />
        </b-btn>

      </template>

      <template slot="active_today" slot-scope="row">
        <worker-change-active :row="row"></worker-change-active>
      </template>
      <template slot="phone" slot-scope="row">
        <a :href=telUrl(row.item)>{{row.item.phone}}</a>
      </template>

      <template slot="passport_url" slot-scope="row">
        <a v-if="row.item.passport_url.length>0" :href=imageUrl(row.item) target="_blank">
          <font-awesome-icon  icon="passport" size="lg" :style="{ color: 'purple' }"/>
        </a>

      </template>

    </b-table>

    <b-row>
      <b-col sm="12">
        <div align="right">Всего рабочих: {{totalRows}}</div>
        <b-pagination align="right" :total-rows="totalRows" :per-page="perPage" v-model="currentPage"></b-pagination>
      </b-col>
    </b-row>
  </div>
</template>

<script>

  import WorkerChangeActive from "./WorkerChangeActive";
    export default {
      name: "worker-table",
      props:["town","filter"],
      components: {
        WorkerChangeActive,
      },
      data() {
        return {
          //serverUrl:"http://localhost:5001",
          serverUrl:"http://178.62.14.228:5001",
          formUrl: 'workers',
          sortBy: 'id',
          sortDesc: false,
          // filter: null,
          items:[],
          fields: [
            {
              key: 'active_today',
              label: this.$t('message.workers.table.active_today'),
              sortable: true,
            },
            {
              key: 'name',
              label: this.$t('message.workers.table.name'),
              sortable: true,
            },
            {
              key: 'phone',
              label: this.$t('message.workers.table.phone'),
              sortable: true,
            },
            {
              key: 'address',
              label: this.$t('message.workers.table.address'),
              sortable: true,
            },
            {
              key: 'passport_url',
              label: this.$t('message.workers.table.passport'),
              sortable: true,
              formatter:'existPassport',
            },
            {
              key: 'age',
              label: this.$t('message.workers.table.age'),
              sortable: true,
            },
            {
              key: 'skill',
              label: this.$t('message.workers.table.skill'),
              sortable: true,
            },
            {
              key: 'ready',
              label: this.$t('message.workers.table.ready'),
              sortable: true,
            },
            {
              key: 'action',
              label: this.$t('message.workers.table.actions'),
              class:'actions-width'
            },
          ],
          currentPage: 1,
          perPage: 1000,
          isBusy: false,
          totalRows: 0,
        }
      },
      created() {
        this.reloadView();

        // setInterval(function () {
        //   this.getItems();
        // }.bind(this),30000);
      },
      methods:{
        reloadView(){
          this.formUrl = "town/"+this.$props.town+"/"+"workers";
          this.refreshTable()
        },
        getItems(){
          let url = this.formUrl;
          return this.$http.get(url).then(result => {
            //console.log(result);
            this.isBusy = false;
            if (result.status === 403) {
              //this.showForbiddenAlert = true;
            } else if (result.status === 200 || result.status === 304) {
              //this.showForbiddenAlert = false;
              if (result.body != null) {
                this.items = result.body;
                this.totalRows = this.items.length;
              }else {
                this.items = [];
              }
            }
          }, error => {
            this.isBusy = false;
            console.log("ERROR", error);
          });
        },
        refreshTable() {
          this.getItems();
          if (this.$refs.workers){
            this.$refs.workers.refresh()
          }
        },
        onFiltered (filteredItems) {
          this.totalRows = filteredItems.length;
          this.currentPage = 1
        },
        imageUrl(newItem){
          return this.serverUrl + '/' + newItem.passport_url
        },
        telUrl(newItem){
          return "tel:"+newItem.phone
        },
        existPassport(value){
          return value.length>0
        },
        showModal(row,isRead) {
          this.$emit('showModal',row,isRead);
        },
        showModalDelete(row) {
          console.log("WorkerTable showModalDelete");
          this.$emit('delete',row);
        },
      },
      watch:{
        '$props.town': function (newR, oldR) {
          this.reloadView()
        }
      }
    }
</script>

<style scoped>

</style>

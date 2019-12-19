<template>
  <div>

    <b-table id="сustomers"
             ref="сustomers"
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

      <template slot="phone" slot-scope="row">
        <a :href=telUrl(row.item)>{{row.item.phone}}</a>
      </template>

      <template slot="urflag" slot-scope="row">
        {{urFlagString(row.item.urflag)}}
      </template>

    </b-table>

    <b-row>
      <b-col sm="12">
        <div align="right">Всего заказчиков: {{totalRows}}</div>
        <b-pagination align="right" :total-rows="totalRows" :per-page="perPage" v-model="currentPage"></b-pagination>
      </b-col>
    </b-row>
  </div>
</template>

<script>
    export default {
      name: "customer-table",
      props:["town","filter"],
      data() {
        return {
          formUrl: 'customers',
          sortBy: 'id',
          sortDesc: false,
          // filter: null,
          items:[],
          fields: [
            {
              key: 'name',
              label: this.$t('message.customers.table.name'),
              sortable: true,
            },
            {
              key: 'phone',
              label: this.$t('message.customers.table.phone'),
              sortable: true,
            },
            {
              key: 'address',
              label: this.$t('message.customers.table.address'),
              sortable: true,
            },
            {
              key: 'email',
              label: this.$t('message.customers.table.email'),
              sortable: true,
            },
            {
              key: 'urflag',
              label: this.$t('message.customers.table.urflag'),
              sortable: true,
            },
            {
              key: 'action',
              label: this.$t('message.customers.table.actions'),
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
          this.formUrl = "town/"+this.$props.town+"/"+"customers";
          this.refreshTable()
        },
        getItems(){
          let url = this.formUrl;
          return this.$http.get(url).then(result => {
            console.log(result);
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
          if (this.$refs.customers){
            this.$refs.customers.refresh()
          }
        },
        onFiltered (filteredItems) {
          // Trigger pagination to update the number of buttons/pages due to filtering
          this.totalRows = filteredItems.length;
          this.currentPage = 1
        },
        telUrl(newItem){
          return "tel:"+newItem.phone
        },
        showModal(row,isRead) {
          this.$emit('showModal',row,isRead);
        },
        showModalDelete(row) {
          console.log("CustomerTable showModalDelete");
          this.$emit('delete',row);
        },
        urFlagString(flag){
          if (flag){
            return "Юр лицо"
          }
          return "Физ лицо"
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

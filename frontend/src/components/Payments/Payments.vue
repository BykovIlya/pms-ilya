<template>
  <div>
    <b-navbar toggleable="md" type="dark" :variant="background" fixed="top" sticky>
      <b-navbar-brand>{{title}}</b-navbar-brand>
      <b-navbar-nav>
        <b-form inline>
          <label class="sr-only" for="input_startDate">{{$t('message.payments.startDate')}}</label>
          <b-input class="mb-2 mr-sm-2 mb-sm-0" id="input_startDate" type="date" v-model.date="startDate"></b-input>
          <label class="sr-only" for="input_endDate">{{$t('message.payments.endDate')}}</label>
          <b-input class="mb-2 mr-sm-2 mb-sm-0" id="input_endDate" type="date" v-model.date="endDate"></b-input>
          <b-button variant="success" @click="findWorkers()">{{ $t("message.payments.find") }}</b-button>
        </b-form>
      </b-navbar-nav>
      <b-navbar-nav class="ml-auto">
        <b-nav-form>
          <b-form-input size="md" class="my-1 md-2" v-model="filter" :placeholder="$t('message.common.searchPlaceholder')" ></b-form-input>
          <b-btn size="md" class="my-1 md-0" :disabled="!filter" @click="filter = ''">
            <font-awesome-icon  icon="times-circle" size="lg"/>
          </b-btn>
        </b-nav-form>
      </b-navbar-nav>
    </b-navbar>

    <b-container id="table-container" fluid>
      <b-table id="workers"
               ref="mytable"
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
          <b-button size="sm" variant="success" @click.stop="showDetail(row)">
            <font-awesome-icon icon="eye" />
          </b-button>
        </template>

        <template slot="paid" slot-scope="row">
          <b-form-checkbox :id="workerCheckBoxId(row)" @change="showModalPaid(row)" v-model="row.item.paid"></b-form-checkbox>
        </template>
      </b-table>
      <b-row>
        <b-col sm="12">
          <div align="right">Всего рабочих: {{totalRows}}</div>
          <b-pagination align="right" :total-rows="totalRows" :per-page="perPage" v-model="currentPage"></b-pagination>
        </b-col>
      </b-row>

    </b-container>

    <!-- Modal Delete Component -->
    <b-modal id="modal-paid"
             ref="modalPaid"
             :title="$t('message.payments.paidDesc')"
             @ok="paidItem(paidRow)"
             @cancel="cancelPaid(paidRow)"
             :ok-title="$t('message.actions.ok')"
             :cancel-title="$t('message.actions.cancel')"
             centered>
      <p>{{ $t("message.payments.paidDesc") }} "{{paidRow.name}}"</p>
    </b-modal>
  </div>
</template>


<script>
  export default {
    name: 'Payments',
    data() {
      return {
        title: 'Выплаты',
        background: 'info',
        formUrl: 'payments',
        paidUrl: 'payments/all',
        sortBy: 'worker.name',
        sortDesc: false,
        filter: null,
        startDate: this.formatDate(new Date()),
        endDate: this.formatDate(new Date()),
        paidRow:{
          id:null,
          name:'',
          flag:null
        },
        items:[],
        fields: [
          {
            key: 'worker.name',
            label: this.$t('message.workers.table.name'),
            sortable: true,
          },
          {
            key: 'worker.cart_number',
            label: this.$t('message.workers.table.cart_number'),
            sortable: true,
          },
          {
            key: 'total',
            label: this.$t('message.workers.table.total_price_payment'),
            sortable: true,
          },
          {
            key: 'total_all',
            label: this.$t('message.workers.table.total_price'),
            sortable: true,
          },
          {
            key: 'paid',
            sortable: true,
            label: this.$t('message.workers.table.paid'),
            'class': 'text-center'
          },
          {
            key: 'action',
            label: this.$t('message.workers.table.actions'),
            'class': 'text-center'
          },
        ],
        currentPage: 1,
        perPage: 1000,
        isBusy: false,
        totalRows: 0,
      }
    },
    created() {
      if (!this.$store.getters.isLoggedIn){
        this.$router.push('/login');
      }else{
        //this.showForbiddenAlert = false;
      }
      this.updateView(this.$route.meta.town);

      console.log(this.$route.query.startdate);
      console.log(this.$route.query.enddate);


      if((this.$route.query.startdate != null) && (this.$route.query.enddate != null)){
        this.startDate = this.$route.query.startdate;
        this.endDate = this.$route.query.enddate;
        this.getItems();
      }

    },
    methods: {
      getItems(){

        let start_date = new Date(this.startDate).toISOString();
        let end_date = new Date(this.endDate).toISOString();

        let url = this.formUrl+"?start="+start_date+"&end="+end_date;
        //let url = this.formUrl;
        //this.isBusy = true;
        return this.$http.get(url).then(result => {
          console.log(result);
          this.isBusy = false;
          if (result.status === 403) {
            this.showForbiddenAlert = true;
          } else if (result.status === 200 || result.status === 304) {
            this.showForbiddenAlert = false;
            if (result.body != null) {
              //this.totalRows = result.body.total;
              this.items = result.body;
              this.totalRows = this.items.length;
            }else {
              console.log("body null");
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
        this.$refs.mytable.refresh()
      },
      onFiltered (filteredItems) {
        this.totalRows = filteredItems.length;
        this.currentPage = 1
      },
      updateView(town){
        this.formUrl = "town/1/payments";
        this.paidUrl ="town/1/payments/all";
        //this.getItems();
      },
      formatDate(date) {
        let dd = date.getDate();
        if (dd < 10) dd = '0' + dd;
        let mm = date.getMonth() + 1;
        if (mm < 10) mm = '0' + mm;
        let yy = date.getFullYear();
        return yy + '-' + mm + '-' + dd;
      },
      findWorkers() {
        console.log("Find Workers");
        console.log(this.startDate);
        console.log(this.endDate);

        this.$router.push({ path: 'payments', query: { startdate: this.startDate,enddate:this.endDate }});
        this.getItems();
      },
      workerCheckBoxId(row){
        return "checkbox_paid_"+row.item.worker.id;
      },
      paid: function (url, data, callback) {
        return this.$http.post(url, data, null).then(result => {
          callback(result.body);
        }, error => {
          console.log("ERROR", error);
          if (error.status === 422) {
            callback(error.body);
          }
        });
      },
      cancelPaid(){
        this.refreshTable();
      },
      paidItem(row) {

        let start_date = new Date(this.startDate).toISOString();
        let end_date = new Date(this.endDate).toISOString();
        let worker = row.id;
        console.log(row);
        console.log(row.flag);
        let url = this.paidUrl+"?worker="+worker+"&start="+start_date+"&end="+end_date+"&flag="+row.flag;
        self = this;
        self.paid(url, null, function (result) {
          console.log(result);
          if (result.status === 422) {
          } else {
            self.refreshTable();
            self.hideModalPaid()
          }
        }, error => {
          console.log("ERROR", error);
        });
      },
      showModalPaid(row) {
        if (row.item) {
          this.paidRow.id = row.item.worker.id;
          this.paidRow.name = row.item.worker.name;
          this.paidRow.flag = !row.item.paid;
        }
        this.$refs.modalPaid.show()
      },
      hideModalPaid() {
        this.$refs.modalPaid.hide()
      },
      showDetail(row) {
        if (row != null ){
          console.log("Show Detail payment");
          this.$router.push({ path: 'payments/detail', query: { worker:row.item.worker.id,startdate: this.startDate,enddate:this.endDate }});
          //this.$router.push({ path: 'tasks/edit/'+row.item.id, params: { row:row,isRead: isRead }});
        }else {
          console.log("Show New Task");
          this.$router.push('/tasks/new');
        }

        //this.$refs.task.showModal(row,isRead)
      },
    }
  }
</script>

<style scoped>

</style>

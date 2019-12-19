<template>
  <div>
    <b-navbar toggleable="md" type="dark" :variant="background" fixed="top" sticky>
      <b-navbar-brand>{{title}}</b-navbar-brand>
      <b-navbar-nav>
        <!--<b-form inline>-->
          <!--<label class="sr-only" for="input_startDate">{{$t('message.payments.startDate')}}</label>-->
          <!--<b-input class="mb-2 mr-sm-2 mb-sm-0" id="input_startDate" type="date" v-model.date="startDate"></b-input>-->
          <!--<label class="sr-only" for="input_endDate">{{$t('message.payments.endDate')}}</label>-->
          <!--<b-input class="mb-2 mr-sm-2 mb-sm-0" id="input_endDate" type="date" v-model.date="endDate"></b-input>-->
          <!--<b-button variant="success" @click="findWorkers()">{{ $t("message.payments.find") }}</b-button>-->
        <!--</b-form>-->

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
        <template slot="start_date" slot-scope="row">
          {{formatTableDate(new Date(row.item.start_date))}}
        </template>

        <template slot="start_time" slot-scope="row">
          {{formatTime(new Date(row.item.start_date))}}
        </template>

        <template slot="end_time" slot-scope="row">
          {{formatTime(new Date(row.item.end_date))}}
        </template>

        <template slot="total_hours" slot-scope="row">
          {{getTotalHours(row.item)}}
        </template>

        <template slot="task" slot-scope="row">
          <b-button size="sm" variant="success" @click.stop="showTask(row)">
            {{row.item.task_id}}
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
             :title="$t('message.payments.paid')"
             @ok="paidItem(paidRow)"
             :ok-title="$t('message.actions.ok')"
             :cancel-title="$t('message.actions.cancel')"
             centered>
      <p>{{ $t("message.payments.paidDesc") }} "{{paidRow.id}}"</p>
    </b-modal>
  </div>
</template>


<script>
  export default {
    name: 'PaymentDetail',
    data() {
      return {
        title: 'Детальная информация',
        background: 'info',
        formUrl: 'payments',
        paidUrl: 'payments/task',
        sortBy: 'name',
        sortDesc: false,
        filter: null,
        startDate: this.formatDate(new Date()),
        endDate: this.formatDate(new Date()),
        paidRow:{
          id:null,
          name:''
        },
       items:[],
        fields: [
          {
            key: 'start_date',
            label: this.$t('message.payments.date'),
            sortable: true,
          },
          {
            key: 'price',
            label: this.$t('message.workers.table.price'),
            sortable: true,
          },
          {
            key: 'start_time',
            label: this.$t('message.workers.table.start_time'),
            sortable: true,
          },
          {
            key: 'end_time',
            label: this.$t('message.workers.table.end_time'),
            sortable: true,
          },
          {
            key: 'total_hours',
            label: this.$t('message.workers.table.total_hours'),
            sortable: true,
          },
          {
            key: 'total_price',
            label: this.$t('message.workers.table.total_price'),
            sortable: true,
          },
          {
            key: 'paid',
            label: this.$t('message.workers.table.paid'),
            sortable: true,
            'class': 'text-center'
          },
          {
            key: 'task',
            label: this.$t('message.tasks.table.name')
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
    },
    methods: {
      getItems(){
        let start_date = new Date(this.$route.query.startdate).toISOString();
        let end_date = new Date(this.$route.query.enddate).toISOString();
        let url = this.formUrl+"?worker="+this.$route.query.worker+"&start="+start_date+"&end="+end_date;
        ///let url = this.formUrl;
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
        this.formUrl = "town/1/payments/detail";
        this.paidUrl ="town/1/payments/task";
        this.getItems();
      },
      formatDate(date) {
        let dd = date.getDate();
        if (dd < 10) dd = '0' + dd;
        let mm = date.getMonth() + 1;
        if (mm < 10) mm = '0' + mm;
        let yy = date.getFullYear();
        return yy + '-' + mm + '-' + dd;
      },
      formatTableDate(date) {
        let dd = date.getDate();
        if (dd < 10) dd = '0' + dd;
        let mm = date.getMonth() + 1;
        if (mm < 10) mm = '0' + mm;
        let yy = date.getFullYear();
        return dd + '-' + mm + '-' + yy;
      },
      formatTime(date) {
        let HH = date.getHours();
        if (HH < 10) HH = '0' + HH;
        let MM = date.getMinutes();
        if (MM < 10) MM = '0' + MM;
        return HH + ':' + MM;
      },
      findWorkers() {
        console.log("Find Workers");
        console.log(this.startDate);
        console.log(this.endDate);
      },
      workerCheckBoxId(row){
        return "checkbox_paid_"+row.item.id;
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
      paidItem(row) {
        let url = this.paidUrl + '/' + row.id;
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
          this.paidRow.id = row.item.id;
          //this.paidRow.name = row.item.worker.name;
        }
        this.$refs.modalPaid.show()
      },
      hideModalPaid() {
        this.$refs.modalPaid.hide()
      },
      showTask(row) {
        // if (row != null ){
        //   console.log("Show Detail payment");
        //   this.$router.push({ path: 'payments/detail', query: { startdate: '2018-11-10',enddate:'2018-11-11' }});
        //   //this.$router.push({ path: 'tasks/edit/'+row.item.id, params: { row:row,isRead: isRead }});
        // }else {
        //   console.log("Show New Task");
        //   this.$router.push('/tasks/new');
        // }

        //this.$refs.task.showModal(row,isRead)
      },
      getTotalHours(tw){
        let startTime = new Date(tw.start_date);
        let endTime = new Date(tw.end_date);
        let diff = ((endTime-startTime)/60000/60);
        diff = Math.round(diff * 100) / 100;
        diff = diff - tw.dinner;
        return diff
      }
    }
  }
</script>

<style scoped>

</style>

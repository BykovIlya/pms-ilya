<template>
  <div>
    <b-navbar toggleable="md" type="dark" :variant="background" fixed="top" sticky>
      <b-navbar-brand>{{title}}</b-navbar-brand>
        <b-nav-form>
          <b-form-select v-model="selectedManager">
            <option :value="null" disabled>-- Выберите менеджера --</option>
            <option v-for="managerName in managersNames" v-bind:value="managerName.id">{{managerName.name}}</option>
          </b-form-select>
        </b-nav-form>
        <b-navbar-nav class = "ml-2">
          <b-form inline>
            <label class="sr-only" for="input_startDate">{{$t('message.payments.startDate')}}</label>
            <b-input class="mb-2 mr-sm-2 mb-sm-0" id="input_startDate" type="date" v-model.date="startDate"></b-input>
            <label class="sr-only" for="input_endDate">{{$t('message.payments.endDate')}}</label>
            <b-input class="mb-2 mr-sm-2 mb-sm-0" id="input_endDate" type="date" v-model.date="endDate"></b-input>
            <b-button variant="success" @click="findManagers()">{{ $t("message.payments.find") }}</b-button>
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
      <b-table id="manager-stat"
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
        <template slot="payment_type" slot-scope="row">
          {{paymentTypeString(row.item.payment_type)}}
        </template>

        <template slot="start_date" slot-scope="row" class="widthDate">
          {{formatDate(row.item.start_date)}}
        </template>

        <template slot="end_date" slot-scope="row">
          {{formatDate(row.item.end_date)}}
        </template>

        <template slot="task_workers" slot-scope="row">
          {{len(row.item.task_workers)}}
        </template>

        <template slot="paid_workers" slot-scope="row">
          {{paidWorkers(row)}}
        </template>

        <template slot="total_delta" slot-scope="row">
          {{roundDelta(row.item.total_delta)}}
        </template>

        <template slot="total_manager" slot-scope="row">
          {{roundDelta(row.item.total_manager)}}
        </template>

        <template slot="paid" slot-scope="row">
          {{paid(row)}}
        </template>

      </b-table>
      <b-row>
        <b-col sm="12">
          <div align="left">Всего заявок: {{totalRows}}</div>
          <div align="left">Всего с заказчика: {{totalCustomer}}</div>
          <div align="left">Всего дельта: {{totalDelta}}</div>
          <div align="left">Всего на руки: {{totalManager}}</div>
          <!--<b-pagination align="right" :total-rows="totalRows" :per-page="perPage" v-model="currentPage"></b-pagination>-->
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>


<script>
  export default {
    name: 'ManagerStatistics',
    data() {
      return {
        title: 'Статистика по менеджерам',
        background: 'info',
        formUrl: 'statistics/manager',
        formManagersUrl: 'statistics/managersNames',
        selectedManager: null,
        managersNames: [],
        sortBy: 'id',
        sortDesc: false,
        filter: null,
        startDate: this.formatDate(new Date()),
        endDate: this.formatDate(new Date()),
        items:[],
        fields: [
          {
            key: 'id',
            // label: this.$t('message.tasks.table.name'),
            sortable: true,
          },
          {
            key: 'customer_name',
            label: this.$t('message.tasks.table.customer'),
            sortable: true,
          },
          {
            key: 'start_address',
            label: this.$t('message.tasks.table.start_address'),
            sortable: true,
          },
          {
            key: 'start_date',
            label: this.$t('message.tasks.table.start_date'),
            sortable: true,
            class:'widthDate'
          },
          {
            key: 'payment_type',
            label: this.$t('message.tasks.table.payment_type'),
            sortable: true,
          },
          {
            key: 'task_workers',
            label: this.$t('message.tasks.table.workers_count'),
            sortable: true,
          },
          {
            key: 'total_price',
            label: this.$t('message.tasks.table.total_price'),
            sortable: true,
          },
          {
            key: 'total_delta',
            label: this.$t('message.tasks.table.total_delta'),
            sortable: true,
          },
          {
            key: 'total_manager',
            label: this.$t('message.tasks.modal.total_manager'),
            sortable: true,
          },
          {
            key: 'paid',
            label: this.$t('message.tasks.table.paid'),
            sortable: true,
          },
          {
            key: 'paid_workers',
            label: this.$t('message.tasks.table.paid_workers'),
            sortable: true,
          },
          {
            key: 'status_name',
            label: this.$t('message.tasks.table.status_name'),
            sortable: true,
          },
          {
            key: 'executor_name',
            label: this.$t('message.tasks.table.executor_name'),
            sortable: true,
          },
        ],
        currentPage: 1,
        perPage: 1000,
        isBusy: false,
        totalRows: 0,
        totalDelta: 0,
        totalManager: 0,
        totalCustomer:0,
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
      this.getManagersNames();
      if((this.$route.query.startdate) && (this.$route.query.enddate)){
        this.startDate = this.$route.query.startdate;
        this.endDate = this.$route.query.enddate;
        this.getItems();
      }
    },
    methods: {
      getManagersNames() {
        let url = this.formManagersUrl;
        return this.$http.get(url).then(result => {
          console.log(result);
          this.isBusy = false;
          if (result.status === 403) {
            this.showForbiddenAlert = true;
          } else if (result.status === 200 || result.status === 304) {
            this.showForbiddenAlert = false;
            if (result.body != null) {
              this.managersNames = result.body;
            }else {
              console.log("body of managers names is null");
              this.managersNames = [];
            }
          }
        }, error => {
          this.isBusy = false;
          console.log("ERROR with getting managers names", error);
        });
      },
      getItems(){

        this.totalDelta = 0;
        this.totalManager = 0;
        this.totalCustomer=0;
        let managerId = this.selectedManager;
        let start_date = new Date(this.startDate).toISOString();
        let end_date = new Date(this.endDate).toISOString();

        let url = this.formUrl+"?id="+managerId+ "&start="+start_date+"&end="+end_date;
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

              for (let i=0;i<this.items.length;i++){
                var item = this.items[i];
                this.updateTotalItem(item)
              }

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
        this.currentPage = 1;
        this.totalDelta = 0;
        this.totalManager = 0;
        this.totalCustomer=0;

        for (let index=0;index<filteredItems.length;index++){
          this.updateTotalItem(filteredItems[index]);
        }
      },
      updateView(town){
        //this.formUrl = "town/1/statistics/tasks";
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
      findManagers() {
        console.log("Find Managers");
        this.$router.push({
          path:'/statistics/manager/', query: { id: this.selectedManager,
            startdate: this.startDate, enddate:this.endDate }});
        this.getItems();
      },
      formatDate(date) {
        date = new Date(date);
        let dd = date.getDate();
        if (dd < 10) dd = '0' + dd;
        let mm = date.getMonth() + 1;
        if (mm < 10) mm = '0' + mm;
        let yy = date.getFullYear();

        let HH = date.getHours();
        if (HH < 10) HH = '0' + HH;
        let MM = date.getMinutes();
        if (MM < 10) MM = '0' + MM;

        return dd + '-' + mm + '-' + yy + ' ' + HH + ':' + MM;
      },
      paymentTypeString(flag){
        if (flag){
          return "Нал"
        }
        return "Безнал"
      },
      len(arr){
        if (arr){
          return arr.length;
        }
        return 0;
      },
      paidWorkers(row){
        let result = true;
        if(row.item.task_workers){
          if (row.item.task_workers.length>0){
            row.item.task_workers.forEach(function (item) {
              if(!item.paid){
                result = false;
                return
              }
            });
          }
        }
        if (result){
          return "Да";
        }
        return "Нет";
      },
      paid(row){
        if (row.item.paid){
          return "Да";
        }
        return "Нет";
      },
      updateTotalItem(item) {
        this.totalDelta = this.totalDelta + item.total_delta;
        this.totalDelta= Math.round(this.totalDelta * 100) / 100;
        this.totalManager = this.totalManager + item.total_manager;
        this.totalManager= Math.round(this.totalManager * 100) / 100;
        this.totalCustomer = this.totalCustomer + item.total_price;
        this.totalCustomer= Math.round(this.totalCustomer * 100) / 100;
      },
      roundDelta(item){
        return Math.round(item * 100) / 100;
      }
    },
  }
</script>

<style>
  .widthDate{
    width: 170px;
  }
</style>

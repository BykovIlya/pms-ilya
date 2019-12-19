<template>
  <div ref="task-view">
    <b-navbar toggleable="md" type="dark" :variant="background" fixed="top" sticky>
      <b-navbar-brand>{{title}}</b-navbar-brand>
      <b-navbar-nav>
        <b-btn size="md" variant="success" @click="showModal(null,false)">{{ $t("message.tasks.importHand") }}</b-btn>
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
          <b-button size="sm" variant="success" @click.stop="showModal(row,true)">
            <font-awesome-icon icon="eye" />
          </b-button>
          <b-btn size="sm" variant="primary" @click="showModal(row,false)">
            <font-awesome-icon icon="edit" />
          </b-btn>
          <b-btn size="sm" variant="danger" @click="showModalDelete(row)">
            <font-awesome-icon icon="trash" />
          </b-btn>
        </template>

        <template slot="payment_type" slot-scope="row">
          {{paymentTypeString(row.item.payment_type)}}
          <!--<b-form-checkbox v-model="row.item.payment_type" disabled></b-form-checkbox>-->
        </template>

        <template slot="start_date" slot-scope="row">
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

        <template slot="paid" slot-scope="row">
          {{paid(row)}}
        </template>

      </b-table>
      <b-row>
        <b-col sm="12">
          <div align="right">Всего заявок: {{totalRows}}</div>
          <b-pagination-nav align="right" :link-gen="linkGen" :number-of-pages="numberOfPages()" :per-page="perPage" v-model="currentPage"></b-pagination-nav>
        </b-col>
      </b-row>

    </b-container>

    <!-- Modal Delete Component -->
    <b-modal id="modal-delete"
    ref="modalDelete"
    :title="$t('message.tasks.modal.removeTitle')"
    @ok="removeItem(deleteRow)"
    :ok-title="$t('message.actions.remove')"
    :cancel-title="$t('message.actions.cancel')"
    centered>
    <p>{{ $t("message.tasks.modal.removeDesc") }} "{{deleteRow.start_address}}"</p>
    </b-modal>

  </div>
</template>

<style>
  #tasks td {
    text-align: center;
  }

  #tasks th {
    text-align: center;
  }

  .actions-width{
    width:150px;
  }

  .widthDate{
    width: 120px;
  }

 .top-fixed-thead{
   /*top: 56px !important;*/
   background-color: lightgray;
 }

  #table-container{
    padding: 0px;
  }

  @media screen
  and (min-width: 768px) {
    a[href*="tel:"] {
      pointer-events: none;
      cursor:default;
      text-decoration: none;
    }
  }

  .modal-fullscreen1 .modal {
    padding: 0 !important;
  }
  .modal-fullscreen1 .modal-dialog {
    max-width: 100%;
    min-height: 100%;
    margin: 0;
  }
  .modal-fullscreen1 .modal-content {
    border: 0;
    border-radius: 0;
    min-height: 100%;
    height: auto;
  }

  @media screen
  and (max-width: 768px) {
    .modal-fullscreen1 .modal {
      padding: 0 !important;
    }
    .modal-fullscreen1 .modal-dialog {
      max-width: 100%;
      min-height: 100%;
      margin: 0;
    }
    .modal-fullscreen1 .modal-content {
      border: 0;
      border-radius: 0;
      min-height: 100%;
      height: auto;
    }

    .modal-fullscreen2  .modal {
      padding: 0 !important;
    }
    .modal-fullscreen2 .modal-dialog {
      max-width: 100%;
      height: 100%;
      margin: 0;
    }
    .modal-fullscreen2 .modal-content {
      width: calc(100% - 2rem);
      min-height: 100%;
      height: auto;
      margin: 1rem;
    }
  }
</style>

<script>
  import Task from "./Task";

  export default {
    components: {Task},
    name: 'Tasks',
    data() {
      return {
        title: 'Задачи',
        background: 'info',
        formUrl: 'tasks',
        sortBy: 'id',
        sortDesc: true,
        filter: null,
        items:[],
        deleteRow:{
          id:null,
          start_address:''
        },
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
          // {
          //   key: 'end_address',
          //   label: this.$t('message.tasks.table.end_address'),
          //   sortable: true,
          // },
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
          {
            key: 'action',
            label: this.$t('message.tasks.table.actions'),
            class:'actions-width'
          },
        ],
        currentPage: 1,
        perPage: 25,
        isBusy: false,
        totalRows: 0,
      }
    },
    created() {
      console.log("created tasks");

      if (!this.$store.getters.isLoggedIn){
        this.$router.push('/login');
      }else{
        //this.showForbiddenAlert = false;
      }
      this.currentTown = this.$route.meta.town;
      this.updateView(this.$route.meta.town);

    },
    methods: {
      numberOfPages(){
        var n= Math.ceil(this.totalRows/this.perPage);
        console.log("number of pages = " + n);
        return n
      },
      getItems(){
        //let url = this.formUrl+"?page="+this.currentPage+"&size="+this.perPage;
        let url = this.formUrl;
        //this.isBusy = true;
        return this.$http.get(url).then(result => {
          //console.log(result);
          this.isBusy = false;
          if (result.status === 403) {
            this.showForbiddenAlert = true;
          } else if (result.status === 200 || result.status === 304) {
            this.showForbiddenAlert = false;
            if (result.body != null) {
              //this.totalRows = result.body.total;
              this.items = result.body;
              this.totalRows = this.items.length;
              if (this.$route.hash){
                this.currentPage = this.getPage(this.$route.hash);
                console.log("page = " + this.currentPage);
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
        console.log("refresh table");
        this.getItems();
        this.$refs.mytable.refresh()
      },
      onFiltered (filteredItems) {
        this.totalRows = filteredItems.length;
        //this.currentPage = 1
      },
      telUrl(newItem){
        return "tel:"+newItem.phone
      },
      updateView(town){
        this.formUrl = "town/"+town+"/tasks";
        if (town===0) {
          this.background = "info";
          this.title = "Заявки Москва";
          this.titleModel = "Заявки Москва"
        }
        else if (town===1) {
          this.background = "warning";
          this.title = "Заявки СПБ";
          this.titleModel = "Заявки СПб";
        }
        this.currentTown = town;
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
      showModal(row,isRead) {
        if (this.currentTown===0) {
          if (row != null ){
            this.$router.push({ name: 'TaskEdit', params: { row:row,isRead: isRead }});
          }else {
            this.$router.push('tasks/new');
          }
        }
        else if (this.currentTown===1) {
          if (row != null ){
            this.$router.push({ name: 'TaskEditSpb', params: { row:row,isRead: isRead }});
          }else {
            this.$router.push('tasks/new');
          }
        }
      },
      delete: function (url, data, callback) {
        return this.$http.delete(url, data, null).then(result => {
          callback(result.body);
        }, error => {
          console.log("ERROR", error);
          if (error.status === 422) {
            callback(error.body);
          }
        });
      },
      removeItem(row) {
        let url = this.formUrl + '/' + row.id;
        this.delete(url, null, function (result) {
          if (result.status === 422) {
          } else {
            this.refreshTable();
            this.hideModalDelete()
          }
        }.bind(this));
      },
      showModalDelete(row) {
        if (row.item) {
          this.deleteRow = row.item
        }
        this.$refs.modalDelete.show()
      },
      hideModalDelete() {
        this.$refs.modalDelete.hide()
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
      linkGen (pageNum) {
        return '#page/' + pageNum
      },
      getPage(hash){
        var page = hash.split("/").pop();
        return Number(page)
      }

    },
    watch: {
      $route: function (newR, oldR) {
        this.updateView(newR.meta.town);
      }
    }
  }
</script>

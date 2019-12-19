<template>
  <div>
    <b-navbar type="dark" :variant="background" sticky>
      <b-navbar-brand>{{title}}</b-navbar-brand>
      <b-navbar-nav>
        <b-btn size="md" variant="success" @click="showFreeWorkersTable">{{ $t("message.workers.importHand") }}</b-btn>
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
               ref="TaskWorkers"
               striped
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
          <b-btn  size="sm" variant="danger" @click="removeWorkerFromTask(row)">
          <font-awesome-icon icon="trash" />
          </b-btn>
        </template>

        <template slot="phone" slot-scope="row">
          <a :href=telUrl(row.item)>{{row.item.phone}}</a>
        </template>

        <template slot="passport_url" slot-scope="row">
          <a v-if="row.item.passport_url.length>0" :href=imageUrl(row.item) target="_blank">
            <font-awesome-icon  icon="passport" size="lg" :style="{ color: 'purple' }"/>
          </a>
        </template>

        <template slot="price" slot-scope="row" >
            <b-form-input id="input_price"
                          type="number"
                          class="widthPrice"
                          v-model.number="row.item.price">
            </b-form-input>
        </template>

        <template slot="start_time" slot-scope="row">
          <b-form-input id="input_startTime"
                        type="time"
                        class="widthTime"
                        v-model.date="row.item.start_time"></b-form-input>
        </template>
        <template slot="end_time" slot-scope="row">
          <b-form-input id="input_endTime"
                        type="time"
                        class="widthTime"
                        v-model.date="row.item.end_time"></b-form-input>
        </template>

        <template slot="road_price" slot-scope="row">
          <b-form-input id="input_price"
                        type="number"
                        class="widthPrice"
                        v-model.number="row.item.road_price">
          </b-form-input>
        </template>

        <template slot="dinner" slot-scope="row">
          <b-form-input id="input_dinner"
                        type="number"
                        step="0.5"
                        class="widthPrice"
                        v-model.number="row.item.dinner">
          </b-form-input>
        </template>

        <template slot="prepaid" slot-scope="row">
          <b-form-input id="input_prepaid"
                        type="number"
                        class="widthPrice"
                        v-model.number="row.item.prepaid">
          </b-form-input>
        </template>

        <template slot="description" slot-scope="row">
          <b-form-textarea id="textarea_description"
                           :placeholder="$t('message.tasks.modal.descriptionPlaceholder')"
                           v-model="row.item.description"
                           :rows="2"
                           :max-rows="2">
          </b-form-textarea>
        </template>

      </b-table>
      <b-row>
        <b-col sm="12">
          <div align="right">Всего рабочих: {{totalRows}}</div>
          <b-pagination align="right" :total-rows="totalRows" :per-page="perPage" v-model="currentPage"></b-pagination>
        </b-col>
      </b-row>

    </b-container>
  </div>
</template>

<script>
  import * as moment from "vue";

  export default {
    name: "TaskWorkers",
    props:['workers','task_workers'],
    data() {
      return {
        title: 'Рабочие',
        background: 'info',
        formUrl: 'workers',
        sortBy: 'name',
        sortDesc: false,
        filter: null,
        items:[],
        fields: [
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
            key: 'dinner',
            label: this.$t('message.workers.table.dinner'),
            sortable: true,
          },
          {
            key: 'total_hours',
            label: this.$t('message.workers.table.total_hours'),
            sortable: true,
          },
          {
            key: 'road_price',
            label: this.$t('message.workers.table.road_price'),
            sortable: true,
          },
          {
            key: 'prepaid',
            label: this.$t('message.workers.table.prepaid'),
            sortable: true,
          },

          {
            key: 'total_price',
            label: this.$t('message.workers.table.total_price'),
            sortable: true,
          },
          {
            key: 'description',
            label: this.$t('message.workers.table.description'),
            sortable: true,
          },
          {
            key: 'action',
            label: this.$t('message.workers.table.actions')
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

      this.updateView(0);
    },
    methods: {
      showFreeWorkersTable() {
        this.$emit('showFreeWorkers')
      },
      getWorkerById: function (id,item) {
        let url = this.formUrl + "/id/" + id;
        let self = this;
        var st = item.start_date;
        var et = item.end_date;
        return this.$http.get(url).then(result => {
          console.log(result);
          this.isBusy = false;
          if (result.status === 403) {
            this.showForbiddenAlert = true;
          } else if (result.status === 200 || result.status === 304) {
            this.showForbiddenAlert = false;
            if (result.body != null) {
              let worker= result.body;
              console.log("Worker");
              console.log(worker);

              let taskWorker = {
                worker_id: item.worker_id,
                name: worker.name,
                phone:worker.phone,
                price: item.price,
                road_price:item.road_price,
                start_time: this.formatTime(new Date(st)),
                end_time:this.formatTime(new Date(et)),
                total_price:item.total_price,
                prepaid:item.prepaid,
                dinner:item.dinner,
                description:item.description,
                start_date: this.formatDate(new Date(st)),
                end_date:this.formatDate(new Date(et)),

              };

              taskWorker.total_hours = this.getTotalHours(taskWorker);

              this.items.push(taskWorker)

            } else {
              console.log("body null");
            }
          }
        }, error => {
          this.isBusy = false;
          console.log("ERROR", error);
        });
      },
      onFiltered (filteredItems) {
        // Trigger pagination to update the number of buttons/pages due to filtering
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
      updateView: function (town) {
        this.formUrl = "town/" + town + "/" + "workers";
        if (town === 0) {
          this.background = "info";
          this.title = "Рабочие Москва";
          this.titleModel = "Работник Москва"
        }
        else if (town === 1) {
          this.background = "warning";
          this.title = "Рабочие СПБ";
          this.titleModel = "Работник СПб";
        }

        if (this.$props.task_workers){
          for (var i=0;i<this.$props.task_workers.length;i++){
            let item = this.$props.task_workers[i];
            this.getWorkerById(item.worker_id, item);
          }
        }
        //this.getItems();
      },
      existWorker(id){
        let exist = false;
        for (let i = 0; i < this.items.length; i++) {
          let item = this.items[i];
          if (item.worker_id === id ){
            exist = true;

            break;
          }
        }
        return exist
      },
      removeWorkerFromTask(row){
        let index = this.items.indexOf(row.item);
        this.items.splice(index,1);
        this.$props.workers.splice(index,1);
        //this.$emit('removeWorkerFromTask',row.item);
        console.log("Remove worker ",row.item.name);
      },
      formatDate(date) {
        let dd = date.getDate();
        if (dd < 10) dd = '0' + dd;
        let mm = date.getMonth() + 1;
        if (mm < 10) mm = '0' + mm;
        let yy = date.getFullYear();
        return yy + '-' + mm + '-' + dd;
      },
      formatTime(date) {
        let HH = date.getHours();
        if (HH < 10) HH = '0' + HH;
        let MM = date.getMinutes();
        if (MM < 10) MM = '0' + MM;
        return HH + ':' + MM;
      },
      getTotalHours(tw){
        let startTime = new Date(tw.start_date+" "+tw.start_time);
        let endTime = new Date(tw.start_date+" "+tw.end_time);
        var diff = ((endTime-startTime)/60000/60);
        diff = Math.round(diff * 100) / 100;
        //Вычитаем обед
        diff = diff - tw.dinner;
        return diff
      }
    },
    watch: {
      '$props.workers': {
        handler: function (newVal, oldVal) {
          for (let i = 0; i < newVal.length; i++) {
            let worker = newVal[i];
            if (!this.existWorker(worker.id)){

              let taskWorker = {
                worker_id: worker.id,
                name: worker.name,
                phone:worker.phone,
                price: 0.0,
                road_price:0.0,
                start_time: this.formatTime(new Date()),
                end_time:this.formatTime(new Date()),
                total_price:0.0,
                prepaid:0.0,
                dinner:0.0,
                description:'',
                start_date: this.formatDate(new Date()),
                end_date:this.formatDate(new Date()),
              };
              taskWorker.total_hours = this.getTotalHours(taskWorker);
              this.items.push(taskWorker)
            }
          }
        },
        deep: true
      },
      'items': {
        handler: function (newValues, oldValues) {
          for (let index=0;index<newValues.length;index++){
            let tw = newValues[index];
            let diff = this.getTotalHours(tw);
            tw.total_hours = diff;
            tw.total_price = (diff*tw.price)+tw.road_price-tw.prepaid;
          }
        },
        deep: true
      },
      $route: function (newR, oldR) {
        console.log(newR.meta.town);
        console.log(oldR.meta);
        this.updateView(newR.meta.town);
      }
    }
  }
</script>

<style scoped>
.widthPrice{
  width: 80px;
}
.widthTime{
  width: 105px;
}

</style>

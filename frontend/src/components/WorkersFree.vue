<template>
  <div>
    <!-- Modal Component -->
    <b-modal id="modal-center"
             size="lg"
             class="modal-fullscreen1"
             ref="modalWorkersSelect"
             :title="title"
             :ok-only="true"
             :no-fade="true"
             centered>

      <b-row>
        <b-col>
          <b-form-input size="md" class="my-1 md-2" v-model="filter" :placeholder="$t('message.common.searchPlaceholder')" ></b-form-input>
        </b-col>
        <b-col>
          <b-btn size="md" class="my-1 md-0" :disabled="!filter" @click="filter = ''">
            <font-awesome-icon  icon="times-circle" size="lg"/>
          </b-btn>
        </b-col>
      </b-row>

      <b-container id="table-container" fluid>
        <b-table id="workers"
                 ref="FreeWorkersTable"
                 striped
                 hover
                 thead-class="fixed-top sticky-top top-fixed-thead"
                 tbody-tr-class="pointer-cell"
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
                 @row-clicked="rowClickHandler"
        >
          <template slot="action" slot-scope="row">
            <b-button size="sm" variant="success" @click.stop="showModal(row,true)">
              <font-awesome-icon icon="eye" />
            </b-button>
            <!--<b-btn  size="sm" variant="primary" @click="showModal(row,false)">-->
              <!--<font-awesome-icon icon="edit" />-->
            <!--</b-btn>-->
            <!--<b-btn  size="sm" variant="danger" @click="showModalDelete(row)">-->
              <!--<font-awesome-icon icon="trash" />-->
            <!--</b-btn>-->

          </template>

          <template slot="active_today" slot-scope="row">
            <worker-change-active :row="row"></worker-change-active>
          </template>

          <template slot="phone" slot-scope="row">
            <a :href=telUrl(row.item)>{{row.item.phone}}</a>
          </template>

          <template slot="task_ids" slot-scope="row">
            {{getTaskIdsString(row.item.task_ids)}}
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
      </b-container>
    </b-modal>
  </div>
</template>

<script>
    export default {
      name: "WorkersFree",
      props: ["town","date","end_date"],
      data() {
        return {
          //serverUrl:"http://localhost:5001",
          serverUrl:"http://178.62.14.228:5001",
          title: 'Рабочие',
          background: 'info',
          formUrl: 'workers',
          sortBy: 'last_task_end_date',
          sortDesc: true,
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
              key: 'task_ids',
              label: this.$t('message.workers.table.task_id'),
              sortable: true
            },
            {
              key: 'last_task_end_date',
              label: this.$t('message.workers.table.last_task_end_date'),
              sortable: true
            },
            // {
            //   key: 'action',
            //   label: this.$t('message.workers.table.actions'),
            //   class:'actions-width'
            // },
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

        //this.updateView(this.$props.town);
      },
      methods: {
        showModalSelect(town) {
          this.$refs.modalWorkersSelect.show();
          this.updateView(town);
        },
        hideModal() {
          this.$refs.modalWorkersSelect.hide();
        },
        getItems(){
          //let url = this.formUrl+"?page="+this.currentPage+"&size="+this.perPage;
          let url = this.formUrl+"/date/"+this.$props.date+"/end_date/"+this.$props.end_date;
          //this.isBusy = true;
          let self = this;
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
                this.items.forEach(function (item) {
                  item._rowVariant= self.getRowVariant(item)
                });
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
          this.$refs.tasks.refresh()
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
        updateView(town){
          this.formUrl = "town/"+town+"/"+"workers";
          if (town===0) {
            this.background = "info";
            this.title = "Рабочие Москва";
            this.titleModel = "Работник Москва"
          }
          else if (town===1) {
            this.background = "warning";
            this.title = "Рабочие СПБ";
            this.titleModel = "Работник СПб";
          }

          this.getItems();
        },
        rowClickHandler(record, index) {
          this.$emit('addToTask',record);
          this.hideModal();
        },
        getRowVariant(item){
          if (item.task_ids){
            if (item.task_ids.length>0){
              return 'danger'
            }
            return '';
          }

        },
        getTaskIdsString(ids){
          let res = '';
          if (ids){
            if (ids.length>0){
              ids.forEach(function (item) {
                if (res !== '' ){
                  res = res + ', '+ item
                }else {
                  res = res + item
                }
              });
            }
          }
          return res
        }
      },
      watch: {
        $route: function (newR, oldR) {
          this.updateView(newR.meta.town);
        }
      }
    }
</script>

<style scoped>
  .pointer-cell{
    cursor: pointer;
  }
</style>

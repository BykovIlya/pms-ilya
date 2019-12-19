<template>
  <div>
    <!-- Modal Component -->
    <b-modal id="modal-center"
             size="lg"
             class="modal-fullscreen-cts1"
             ref="modalCustomerSelect"
             :title="titleModel"
             :ok-only="true"
             :no-fade="true"
             centered>
      <b-card no-body>
        <b-tabs card v-model="currentTown">
          <b-tab :title="updateTitle(i)" v-for="i in tabs" :key="i">
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

            <b-table id="сustomers"
                     ref="сustomers"
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
          </b-tab>
        </b-tabs>
      </b-card>


    </b-modal>

  </div>
</template>

<style>
  .pointer-cell{
    cursor: pointer;
  }

  .modal-fullscreen-cts .modal {
    padding: 0 !important;
  }
  .modal-fullscreen-cts .modal-dialog {
    max-width: 100%;
    min-height: 100%;
    margin: 0;
  }
  .modal-fullscreen-cts .modal-content {
    border: 0;
    border-radius: 0;
    min-height: 100%;
    height: auto;
  }

</style>

<script>
  import DropdownDivider from "bootstrap-vue/es/components/dropdown/dropdown-divider";

  export default {
    components: {DropdownDivider},
    name: "CustomerTableSelect",
    props: ["town", "filter"],
    data() {
      return {
        titleModel:'Выбор заказчика',
        formUrl: 'customers',
        sortBy: 'id',
        sortDesc: false,
        // filter: null,
        items: [],
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
            key: 'urflag',
            label: this.$t('message.customers.table.urflag'),
            sortable: true,
          },
          // {
          //   key: 'action',
          //   label: this.$t('message.customers.table.actions'),
          //   class: 'actions-width'
          // },
        ],
        currentPage: 1,
        perPage: 1000,
        isBusy: false,
        totalRows: 0,
        currentTown:0,
        tabs:[0,1]

      }
    },
    created() {
      this.reloadView();
    },
    methods: {
      showModalSelect() {

        this.$refs.modalCustomerSelect.show();
        this.reloadView();
      },
      hideModal() {
        this.$refs.modalCustomerSelect.hide();
      },
      reloadView() {
        this.formUrl = "town/" + this.currentTown + "/" + "customers";
        this.refreshTable()
      },
      getItems() {
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
            } else {
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
        if (this.$refs.customers) {
          this.$refs.customers.refresh()
        }
      },
      onFiltered(filteredItems) {
        // Trigger pagination to update the number of buttons/pages due to filtering
        this.totalRows = filteredItems.length;
        this.currentPage = 1
      },
      telUrl(newItem) {
        return "tel:" + newItem.phone
      },
      showModal(row, isRead) {
        this.$emit('showModal', row, isRead);
      },
      showModalDelete(row) {
        console.log("CustomerTable showModalDelete");
        this.$emit('delete', row);
      },
      urFlagString(flag) {
        if (flag) {
          return "Юр лицо"
        }
        return "Физ лицо"
      },
      rowClickHandler(record, index) {
        this.$emit('selectCustomer',record);
        this.hideModal();
      },
      updateTitle(town){
        if (town === 0){
          return "Москва"
        }else if (town === 1) {
          return "Спб"
        }
        return "Москва"
      }
    },
    watch: {
      '$props.town': function (newR, oldR) {
        this.reloadView()
      },
      'currentTown': function (newR, oldR) {
        this.reloadView()
      }
    }
  }
</script>



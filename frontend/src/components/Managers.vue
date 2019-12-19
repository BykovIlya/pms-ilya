<template>
  <div>
    <b-navbar toggleable="md" type="dark" :variant="background" fixed="top" sticky>
      <b-navbar-brand>{{ $t("message.managers.title") }}</b-navbar-brand>
   <!--   <b-nav-form>
        <b-form-select v-model="selectedManager">
          <option :value="null" disabled>-- Выберите менеджера --</option>
          <option v-for="managerName in managersNames" v-bind:value="managerName.id">{{managerName.name}}</option>
        </b-form-select>
      </b-nav-form>
      <b-navbar-nav class = "ml-2">
          <b-button variant="success" @click="findManagers()">{{ $t("message.managers.see") }}</b-button>
      </b-navbar-nav>
      <b-navbar-nav class = "ml-2">
        <b-button variant="primary" @click="findManagers()">{{ $t("message.managers.edit") }}</b-button>
      </b-navbar-nav>
      <b-navbar-nav class = "ml-2">
        <b-button variant="danger" @click="findManagers()">{{ $t("message.managers.delete") }}</b-button>
      </b-navbar-nav> -->
      <b-navbar-nav class = "ml-auto">
      <b-button variant="success" @click="showModal(false, false)">{{ $t("message.managers.add") }}</b-button>
      </b-navbar-nav>
    </b-navbar>

    <b-container id="table-container" fluid>
      <b-table id="managers"
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
          <b-button size="sm" variant="success" @click.stop="showModal(row,true,true)">
            <font-awesome-icon icon="eye" />
          </b-button>
          <b-btn  size="sm" variant="primary" @click="showModal(row,false,true)">
            <font-awesome-icon icon="edit" />
          </b-btn>
          <b-btn  size="sm" variant="danger" @click="showModalDelete(row)">
            <font-awesome-icon icon="trash" />
          </b-btn>

        </template>
      </b-table>
    </b-container>

    <manager ref="manager" v-on:refresh="refreshTable"></manager>
  </div>
</template>

<script>
  import Manager from "./Manager";

  export default {
      components: {
        Manager,
      },
      name: "Managers",
      data() {
        return {
          selectedManager: null,
          managersNames: [],
          background: 'info',
          formUrl: "managers",
          formManagersUrl: "managers/managersNames",
          sortBy: 'name',
          sortDesc: false,
          filter: null,
          items:[],
          fields: [
            {
              key: 'id',
              // label: this.$t('message.tasks.table.name'),
              sortable: true,
            },
            {
              key: 'username',
              label: this.$t('message.managers.table.username'),
              sortable: true,
            },
            {
              key: 'name',
              label: this.$t('message.managers.table.name'),
              sortable: true
            },
            {
              key: 'email',
              label: this.$t('message.managers.table.email'),
              sortable: true
            },
            {
              key: 'role',
              label: this.$t('message.managers.table.role'),
              sortable: true
            },
            {
              key: 'rate',
              label: this.$t('message.managers.table.rate'),
              sortable: true
            },
            {
              key: 'rate_tax',
              label: this.$t('message.managers.table.rate_tax'),
              sortable: true
            },
            {
              key: 'action',
              label: this.$t('message.managers.table.actions'),
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
        if (!this.$store.getters.isLoggedIn){
          this.$router.push('/login');
        }else{
          //this.showForbiddenAlert = false;
        }
        //this.updateView(this.$route.meta.town);

        this.getManagersNames();
        this.getItems();
      },
      methods: {
        getItems() {
          let url = this.formManagersUrl;
          return this.$http.get(url).then(result => {
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
                  //this.updateTotalItem(item)
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
        onFiltered (filteredItems) {
          this.totalRows = filteredItems.length;
          this.currentPage = 1
        },
        refreshTable() {
          this.getItems();
          this.$refs.mytable.refresh()
        },
        findRole(my) {
          this.$refs.manager.findRole(my)
        },
        showModal(row, isRead, isNew) {
          this.$refs.manager.showModal(row, isRead, isNew)
        },
        showModalDelete(row) {
          this.$refs.manager.showModalDelete(row)
        },
        getManagersNames() {
          let url = this.formManagersUrl;
          return this.$http.get(url).then(result => {
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
      }
    }
</script>

<style scoped>

</style>

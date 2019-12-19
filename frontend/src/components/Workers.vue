<template>
  <div>
    <b-navbar toggleable="md" type="dark" :variant="background" fixed="top" sticky>
      <b-navbar-toggle target="my_nav_collapse"></b-navbar-toggle>

      <b-navbar-brand>{{title}}</b-navbar-brand>
      <b-collapse is-nav id="my_nav_collapse">

        <b-navbar-nav>
          <b-nav-item>
          <b-btn size="md" variant="success" @click="showModal">{{ $t("message.workers.importHand") }}</b-btn>
          </b-nav-item>
        </b-navbar-nav>
        <!-- Right aligned nav items -->
        <b-navbar-nav class="ml-auto">
          <b-nav-form>
            <b-form-input size="md" class="my-1 md-2" v-model="filter"
                          :placeholder="$t('message.common.searchPlaceholder')"></b-form-input>
            <b-btn size="md" class="my-1 md-0" :disabled="!filter" @click="filter = ''">
              <font-awesome-icon icon="times-circle" size="lg"/>
            </b-btn>
          </b-nav-form>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>

    <b-container id="table-container" fluid>
      <worker-table ref="workerTable" :town="currentTown" :filter="filter" v-on:showModal="showModal"
                    v-on:delete="showModalDelete"></worker-table>
    </b-container>

    <worker ref="worker" :titleModel="titleModel" v-on:refresh="refreshTable"></worker>
  </div>
</template>

<style>
  #workers td {
    text-align: center;
  }

  #workers th {
    text-align: center;
  }

  .actions-width {
    width: 150px;
  }

  .top-fixed-thead {
    /*top: 56px !important;*/
    background-color: lightgray;
  }

  #table-container {
    padding: 0px;
  }

  @media screen
  and (min-width: 768px) {
    a[href*="tel:"] {
      pointer-events: none;
      cursor: default;
      text-decoration: none;
    }
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
      height: 100%;
    }

    .modal-fullscreen2 .modal {
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
  import Worker from "./Worker";
  import WorkerTable from "./WorkerTable";

  export default {
    components: {
      WorkerTable,
      Worker
    },
    name: 'Workers',
    data() {
      return {
        currentTown: 0,
        title: 'Рабочие',
        background: 'info',
        filter: null,
      }
    },
    created() {
      if (!this.$store.getters.isLoggedIn) {
        this.$router.push('/login');
      } else {
        //this.showForbiddenAlert = false;
      }

      this.currentTown = this.$route.meta.town;
      this.updateView(this.$route.meta.town);

    },
    methods: {
      updateView(town) {
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
        this.currentTown = town;
      },
      refreshTable() {
        this.$refs.workerTable.refreshTable()
      },
      showModal(row, isRead) {
        this.$refs.worker.showModal(row, isRead)
      },
      showModalDelete(row) {
        console.log("WorkerS showModalDelete");
        this.$refs.worker.showModalDelete(row)
      },
    },
    watch: {
      $route: function (newR, oldR) {
        console.log(newR.meta.town);
        console.log(oldR.meta);
        this.updateView(newR.meta.town);
      }
    }
  }
</script>

<template>
  <div>
    <b-navbar toggleable="md" type="dark" :variant="background" fixed="top" sticky>
      <b-navbar-brand>{{title}}</b-navbar-brand>
      <b-navbar-nav>
        <b-btn size="md" variant="success" @click="showModal">{{ $t("message.customers.importHand") }}</b-btn>
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
      <customer-table ref="customerTable" :town="currentTown" :filter="filter" v-on:showModal="showModal"  v-on:delete="showModalDelete"></customer-table>
    </b-container>

    <customer ref="customer" :titleModel="titleModel" v-on:refresh="refreshTable"></customer>
  </div>
</template>

<style>
  #customers td {
    text-align: center;
  }

  #customers th {
    text-align: center;
  }

  .actions-width{
    width:150px;
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
  import Customer from "./Customer";
  import CustomerTable from "./CustomerTable";
  export default {
    components: {
      CustomerTable,
      Customer
    },
    name: 'Customers',
    data() {
      return {
        currentTown:0,
        title: 'Рабочие',
        background: 'info',
        filter: null,
      }
    },
    created() {
      if (!this.$store.getters.isLoggedIn){
        this.$router.push('/login');
      }else{
        //this.showForbiddenAlert = false;
      }

      this.currentTown = this.$route.meta.town;
      this.updateView(this.$route.meta.town);

    },
    methods: {
      updateView(town){
        this.formUrl = "town/"+town+"/"+"customers";
        if (town===0) {
          this.background = "info";
          this.title = "Заказчики Москва";
          this.titleModel = "Заказчик Москва"
        }
        else if (town===1) {
          this.background = "warning";
          this.title = "Заказчики СПБ";
          this.titleModel = "Заказчик СПб";
        }
        this.currentTown = town;
      },
      refreshTable() {
        this.$refs.customerTable.refreshTable()
      },
      showModal(row,isRead) {
        this.$refs.customer.showModal(row,isRead)
      },
      showModalDelete(row) {
        console.log("CustomerS showModalDelete");
        this.$refs.customer.showModalDelete(row)
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

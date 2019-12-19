<template>
  <div id="app">
    <b-navbar  v-if="isLoggedIn" toggleable="md" type="dark" variant="dark">
      <b-navbar-toggle target="nav_collapse"></b-navbar-toggle>
      <b-navbar-brand :to="{ name: 'Workers'}">PMS</b-navbar-brand>
      <b-collapse is-nav id="nav_collapse">
        <b-navbar-nav>
          <!--<b-nav-item-dropdown :text="$t('message.workers.title')" right>-->
            <!--<b-dropdown-item :to="{ name: 'Workers'}"  exact>{{ $t("message.workers.title") }} Москва</b-dropdown-item>-->
            <!--<b-dropdown-item :to="{ name: 'WorkersSpb'}">{{ $t("message.workers.title") }} СПб</b-dropdown-item>-->
          <!--</b-nav-item-dropdown>-->

          <b-nav-item-dropdown class="topZIndex">
            <template slot="button-content">
              Москва
            </template>
            <b-dropdown-item :to="{ name: 'Workers'}"  exact>{{ $t("message.workers.title") }} Москва</b-dropdown-item>
            <b-dropdown-item :to="{ name: 'Customers'}"  exact>{{ $t("message.customers.title") }} Москва</b-dropdown-item>
            <b-dropdown-item :to="{ name: 'Tasks'}">{{ $t("message.tasks.title") }} Москва</b-dropdown-item>
          </b-nav-item-dropdown>

          <b-nav-item-dropdown class="topZIndex">
            <template slot="button-content">
              СПБ
            </template>
            <b-dropdown-item :to="{ name: 'WorkersSpb'}">{{ $t("message.workers.title") }} СПб</b-dropdown-item>
            <b-dropdown-item :to="{ name: 'CustomersSpb'}"  exact>{{ $t("message.customers.title") }} СПб</b-dropdown-item>
            <b-dropdown-item :to="{ name: 'TasksSpb'}">{{ $t("message.tasks.title") }} СПб</b-dropdown-item>
          </b-nav-item-dropdown>

            <!--<b-nav-item :to="{ name: 'Workers'}"  exact>{{ $t("message.workers.title") }} Москва</b-nav-item>-->
            <!--<b-nav-item :to="{ name: 'Customers'}"  exact>{{ $t("message.customers.title") }} Москва</b-nav-item>-->
            <!--<b-nav-item :to="{ name: 'Tasks'}">{{ $t("message.tasks.title") }} Москва</b-nav-item>-->
            <!---->
            <!--<b-nav-item :to="{ name: 'WorkersSpb'}">{{ $t("message.workers.title") }} СПб</b-nav-item>-->
            <!--<b-nav-item :to="{ name: 'CustomersSpb'}"  exact>{{ $t("message.customers.title") }} СПб</b-nav-item>-->
            <!--<b-nav-item :to="{ name: 'TasksSpb'}">{{ $t("message.tasks.title") }} СПб</b-nav-item>-->

          <b-nav-item :to="{ name: 'Payments'}">{{ $t("message.payments.title") }}</b-nav-item>

          <b-nav-item-dropdown class="topZIndex">
            <template slot="button-content">
              Статистика
            </template>
            <b-dropdown-item :to="{ name: 'ManagerStatistics'}">По менеджерам</b-dropdown-item>
            <b-dropdown-item href="#">По грузчикам</b-dropdown-item>
            <b-dropdown-item href="#">По заказчикам</b-dropdown-item>
          </b-nav-item-dropdown>

          <b-nav-item-dropdown class="topZIndex">
            <template slot="button-content">
              Отчёты
            </template>
            <b-dropdown-item href="#">Общий отчёт</b-dropdown-item>
            <b-dropdown-item href="#">Бух отчёт</b-dropdown-item>
          </b-nav-item-dropdown>

          <b-nav-item :to="{ name: 'Managers'}">{{ $t("message.managers.title") }}</b-nav-item>

        </b-navbar-nav>

        <b-navbar-nav class="ml-auto">
          <b-nav-text>Вы вошли как: {{user.name}}</b-nav-text>
          <b-nav-item @click="logout">{{ $t("message.login.logout") }}</b-nav-item>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>

    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App',
  data () {
    return {
      user: ""
    }
  },
  created(){

  },
  methods: {
    //...mapActions(["logout"])
    logout(){
      this.$store.dispatch("logout").then(res => {
        this.user= "";
        this.$router.push('/login');
      })
    },
    userInfo() {
      this.$http.get("users/info").then(result => {
        console.log(result);
        if (result.status === 200) {
            this.user = result.body
        }
      }, error => {
        console.log(error);
      });
    }
  },
  computed: {
    //...mapGetters(["isLoggedIn"])
    isLoggedIn(){
      this.userInfo();
      return this.$store.getters.isLoggedIn
    },
    isAdmin() {
      return this.$store.getters.isAdmin
    },
    isManager() {
      return this.$store.getters.isManager
    },
    isAccountant() {
      return this.$store.getters.isAccountant
    }
  }
}
</script>

<style>
  .topZIndex{
    z-index: 1050;
  }
</style>

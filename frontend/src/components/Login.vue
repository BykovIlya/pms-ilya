<template>
  <div>
    <b-navbar  v-if="!isLoggedIn" toggleable="md" type="dark" variant="dark">
      <b-navbar-toggle target="nav_collapse"></b-navbar-toggle>
      <b-navbar-brand >PMS</b-navbar-brand>
    </b-navbar>
    <b-container>
      <b-row>
        <div class="Absolute-Center is-Responsive">
          <b-alert variant="danger"
                   dismissible
                   :show="showDismissibleAlert"
                   @dismissed="showDismissibleAlert=false">
            Доступ запрещён
          </b-alert>
          <div class="col-sm-12 col-md-10 col-md-offset-1">
            <b-form id="loginForm"  @submit.stop.prevent="login()">
              <b-form-group id="usernameGroup"
                            :label="$t('message.login.login')"
                            label-for="usernameInput"
              >
                <b-form-input id="usernameInput"
                              type="text"
                              :placeholder="$t('message.login.enterLogin')"
                              v-model="username"
                >
                </b-form-input>

              </b-form-group>

              <b-form-group id="passwordGroup"
                            :label="$t('message.login.password')"
                            label-for="passwordInput"
              >
                <b-form-input id="passwordInput"
                              type="password"
                              :placeholder="$t('message.login.enterPassword')"
                              v-model="password">
                </b-form-input>
              </b-form-group>
              <b-form-group>
                <b-button type="submit" variant="dark" block >{{ $t("message.login.signIn") }}</b-button>
              </b-form-group>
            </b-form>
          </div>
        </div>
      </b-row>
    </b-container>
  </div>


</template>

<script>

export default {
  name: 'Login',
  data () {
    return {
      msg: 'Welcome to PMS',
      username: "",
      password: "",
      showDismissibleAlert: false
    }
  },
  methods: {
    login() {
      this.post("login", {
        username: this.username,
        password: this.password
      },function(result){
        console.log(result)
        if (result.status === 200) {
          this.$store.dispatch("login", {
            token: result.body.token
          }).then(res => {
            this.userInfo(result.body.token);
          })
        }

        if (result.status === 401) {
          this.showDismissibleAlert = true
        }

      }.bind(this));
    },
    post: function (url, data, callback) {
      console.log(data);
      return this.$http.post(url,data,null).then(result => {
        callback(result);
      },error =>{
        callback(error);
      });
    },
    userInfo(token) {
      this.$http.get("users/info").then(result => {
        console.log(result);
        if (result.status === 200) {
          this.$store.dispatch("login", {
            token: token,
            role: result.body.role
          }).then(res => {
            if (result.body.role === "accountant"){
              //this.$router.push({name:"Orders"});
              return
            }
            this.$router.push('/');
          })
        }
      },error =>{
        console.log(error);
        if (error.status === 401) {
          this.showDismissibleAlert = true
        }
        if (error.status === 403) {
          this.showDismissibleAlert = true
        }
      });

    }
  },
  computed: {
    isLoggedIn(){
      return this.$store.getters.isLoggedIn
    }
  }
}


</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.Absolute-Center {
  margin: auto;
  position: absolute;
  top: 0; left: 0; bottom: 0; right: 0;
}

.Absolute-Center.is-Responsive {
  width: 50%;
  height: 50%;
  min-width: 400px;
  max-width: 500px;
  padding: 40px;
}


</style>


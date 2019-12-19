import Vue from 'vue'
import Vuex from 'vuex'

const LOGIN = "LOGIN";
const LOGIN_SUCCESS = "LOGIN_SUCCESS";
const LOGOUT = "LOGOUT";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    isLoggedIn: localStorage.getItem("token"),
    role:localStorage.getItem("role")
  },
  mutations: {
    [LOGIN](state) {
      state.pending = true;
    },
    [LOGIN_SUCCESS](state) {
      state.isLoggedIn = true;
      state.role = localStorage.getItem("role");
      state.pending = false;
    },
    [LOGOUT](state) {
      state.isLoggedIn = false;
      state.role = null;
    }
  },
  actions: {
    login({
            state,
            commit,
            rootState
          }, creds) {
      console.log("login...", creds);
      localStorage.setItem("role", creds.role);
      commit(LOGIN); // show spinner
      return new Promise(resolve => {
        setTimeout(() => {
          localStorage.setItem("token", creds.token);
          localStorage.setItem("role", creds.role);
          commit(LOGIN_SUCCESS);
          resolve();
        }, 10);
      });
    },
    logout({
             commit
           }) {
      localStorage.removeItem("token");
      localStorage.removeItem("role");
      commit(LOGOUT);
    }
  },
  getters: {
    isLoggedIn: state => {
      return state.isLoggedIn;
    },
    isAdmin: state => {
      return state.role === "admin"
    },
    isManager: state => {
      return state.role === "manager"
    },
    isAccountant: state => {
      return state.role === "accountant"
    }

  }
});

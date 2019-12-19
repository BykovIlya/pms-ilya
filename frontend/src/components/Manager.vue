<template>
  <div>
    <!-- Modal Component -->
    <b-modal id="modal-center"
             size="lg"
             class="modal-fullscreen1"
             ref="modal"
             @ok="createProduct"
             :title=title()
             :ok-disabled="isOk"
             centered>
      <b-form @submit.stop.prevent="handleSubmit">
        <b-form-group id="group1"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.managers.modal.username')"
                      label-for="input1">
          <b-form-input id="input1"
                        type="text"
                        :disabled="isDisabled_2"
                        :state="usernameState"
                        aria-describedby="inputUsernameFeedback"
                        :placeholder="$t('message.managers.modal.usernamePlaceholder')"
                        v-model="newItem.username"></b-form-input>
        </b-form-group>
        <b-form-invalid-feedback id="inputUserFeedback">
          Введите не менее пяти символов
        </b-form-invalid-feedback>

        <b-form-group id="group2" v-if ="!isDisabled_2"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.managers.modal.password')"
                      label-for="input3-2">
          <b-form-input id="input2"
                        type="text"
                        :disabled="isDisabled"
                        :state="passwordState"
                        aria-describedby="inputPasswordFeedback"
                        :placeholder="$t('message.managers.modal.passwordPlaceholder')"
                        v-model="newItem.password"></b-form-input>
        </b-form-group>

        <b-form-invalid-feedback id="inputPasswordFeedback">
          Слишком короткий пароль
        </b-form-invalid-feedback>

        <b-form-group id="group3"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.managers.modal.name')"
                      label-for="input3">
          <b-form-input id="input3"
                        type="text"
                        :disabled="isDisabled"
                        :state="nameState"
                        aria-describedby="inputNameFeedback"
                        :placeholder="$t('message.managers.modal.namePlaceholder')"
                        v-model="newItem.name"></b-form-input>
        </b-form-group>

        <b-form-group id="group4"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.managers.modal.email')"
                      label-for="input4">
          <b-form-input id="input4"
                        type="text"
                        :disabled="isDisabled"
                        :state="emailState"
                        aria-describedby="inputEmailFeedback"
                        :placeholder="$t('message.managers.modal.emailPlaceholder')"
                        v-model="newItem.email"></b-form-input>
        <!--  <b-form-invalid-feedback id="inputLiveFeedbackEmail">
            {{ $t("message.managers.modal.errorWithEmail") }}
          </b-form-invalid-feedback> -->
        </b-form-group>

        <b-form-group id="groupOption" v-if ="!isDisabled_2"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.managers.modal.role')"
                      label-for="inputOption">
          <b-form-select id="inputOption"
                        :disabled="isDisabled"
                        v-model="newItem.role"
                         :required="true">
            <template slot="first">
            <option :value="null" disabled>-- Выберите должность --</option>
            </template>
            <option v-for="option in options" v-bind:value="option.value">{{option.text}}</option>
          </b-form-select>
        </b-form-group>

        <b-form-group id="group5"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.managers.modal.rate')"
                      label-for="input5">
          <b-form-input id="input5"
                        type="text"
                        :disabled="isDisabled"
                        :state="rateState"
                        aria-describedby="inputRateFeedback"
                        :placeholder="$t('message.managers.modal.ratePlaceholder')"
                        v-model="newItem.rate"></b-form-input>
        </b-form-group>

        <b-form-group id="group6"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.managers.modal.rate_tax')"
                      label-for="input6">
          <b-form-input id="input6"
                        type="text"
                        :disabled="isDisabled"
                        :state="rate_taxState"
                        aria-describedby="inputRate_taxFeedback"
                        :placeholder="$t('message.managers.modal.rate_taxPlaceholder')"
                        v-model="newItem.rate_tax"></b-form-input>
        </b-form-group>

      </b-form>
    </b-modal>

    <!-- Modal Delete Component -->
    <b-modal id="modal-delete"
             ref="modalDelete"
             class="modal-fullscreen1"
             :title="$t('message.managers.modal.removeTitle')"
             @ok="removeItem(deleteRow)"
             :ok-title="$t('message.actions.remove')"
             :cancel-title="$t('message.actions.cancel')"
             centered>
      <p>{{ $t("message.managers.modal.removeDesc") }} "{{deleteRow.name}}"</p>
    </b-modal>

  </div>
</template>

<script>
    export default {
      name: "Manager",
      computed: {
        usernameState() {
          if (this.isDisabled) {return null;}
          return this.newItem.username.length > 4
        },
        passwordState() {
          if (this.isDisabled) {return null;}
          return this.newItem.password.length > 4
        },
        nameState() {
          if (this.isDisabled) {return null;}
          return this.newItem.name.length > 0
        },
        emailState() {
          if (this.isDisabled) {return null;}
          return this.reg.test(this.newItem.email)
        },
        rateState() {
          if (this.isDisabled) {return null;}
          return this.newItem.rate.length > 0
        },
        rate_taxState() {
          if (this.isDisabled) {return null;}
          return this.newItem.rate_tax.length > 0
        },
        isOk() {
          if (this.isDisabled) {return true;}
          if (!this.isDisabled_2) {
            return !(this.usernameState && this.passwordState &&
            this.nameState && this.emailState &&
            this.rateState && this.rate_taxState);}
          return !(this.usernameState &&
            this.nameState && this.emailState &&
            this.rateState && this.rate_taxState);
        },
      },
      data() {
        return {
          formUrl: 'managers',
          isDisabled:true,
          isDisabled_2: true,
          isDisabled_ok: true,
          reg: /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,24}))$/,
          options : [
            {text: this.$t('message.managers.admin'), value: "admin"},
            {text: this.$t('message.managers.manager'), value: "manager"},
            {text: this.$t('message.managers.accountant'), value: "accountant"},
          ],
          newItem: {
            id: null,
            username: '',
            password: '',
            name: '',
            email: '',
            role: '',
            rate: '',
            rate_tax: '',
          },
          newMan: {
            username: '',
            password: '',
            name: '',
            email: '',
            role: '',
            rate: '',
            rate_tax: '',
          },
          deleteRow: {
            id: null,
            name: '',
          },
        }
      },
      created() {
      },
      methods: {
        showModal(row,isView, isNew) {
          if (row.item) {
            this.setNewItem(row, isNew);
          } else {
            this.clearNewItem(isNew)
          }
          this.$refs.modal.show();
          this.isDisabled = isView;
          this.isDisabled_2 = isNew;
        },
        showModalDelete(row) {
          if (row.item) {
            this.deleteRow = row.item
          }
          this.$refs.modalDelete.show()
        },
        findRole(my) {
          for (let i in this.options) {
            if (i.value === my) {
              return i.text
            }
          }
        },
        setNewItem(row, isNew) {
          this.newItem.id = row.item.id;
          this.newItem.username = row.item.username;
          if (isNew === false) {
            this.newItem.password = row.item.password;
          }
          this.newItem.email = row.item.email;
          this.newItem.role = row.item.role;
          this.newItem.rate = row.item.rate.toString();
          this.newItem.rate_tax = row.item.rate_tax.toString();

          this.newItem.name = row.item.name;
        },
        clearNewItem(isNew) {
          this.newItem.id = null;
          this.newItem.username = '';
          if (isNew === false) {
            this.newItem.password = '';
          }
          this.newItem.email = '';
          this.newItem.role = '';
          this.newItem.rate = '';
          this.newItem.rate_tax = '';
          this.newItem.name = '';
        },
          hideModal() {
          if (this.$refs.fileinput){
            this.$refs.fileinput.reset();
          }
          this.fileProducts = null;
          this.$refs.modal.hide();
        },
        title() {
          if (this.isDisabled_2) {
            return this.$t('message.managers.modal.editTitle')
          }
          return this.$t('message.managers.modal.addTitle')
        },
        createProduct(evt) {
          evt.preventDefault();
          this.handleSubmit()
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
              //this.formErrors = result.errors
            } else {
              this.$emit('refresh');
              this.hideModalDelete()
            }
          }.bind(this));
        },
        hideModalDelete() {
          this.$refs.modalDelete.hide()
        },
        post: function (url, data, callback) {
          console.log(data);
          return this.$http.post(url, data, null).then(result => {
            callback(result);
          }, error => {
            callback(error);
          });
        },
        put: function (url, data, callback) {
          return this.$http.put(url, data, null).then(result => {
            callback(result);
          }, error => {
            callback(error);
          });
        },
        handleSubmit() {
          let self;
          let url = this.formUrl;
          if (this.newItem.id > 0) {
            console.log(this.newItem);
            self = this;
            url = this.formUrl + '/' + this.newItem.id;
            console.log(this.newItem.id);
            self.put(url, self.newItem, function (result) {
              console.log(result);
              if (result.status === 200) {

                self.newItem.id = result.body.id;

                self.hideModal();
                self.$emit('refresh');
              }
            },error => {
              console.log("ERROR", error);
            });

          } else {

            this.newMan.email = this.newItem.email;
            this.newMan.name = this.newItem.name;
            this.newMan.password = this.newItem.password;
            this.newMan.username = this.newItem.username;
            this.newMan.rate = this.newItem.rate;
            this.newMan.rate_tax = this.newItem.rate_tax;
            this.newMan.role = "manager";
            let self = this;
            self.post(url, this.newMan, function (result) {
                self.hideModal();
                self.$emit('refresh');
            },error => {
              console.log("ERROR", error);

            });
          }
        },

      }
    }
</script>

<style scoped>

</style>

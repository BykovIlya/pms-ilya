<template>
  <div>
    <!-- Modal Component -->
    <b-modal id="modal-center"
             size="lg"
             class="modal-fullscreen1"
             ref="modal"
             :title="titleModel"
             @ok="createProduct"
             :ok-disabled="isDisabled"
             centered>
      <b-form @submit.stop.prevent="handleSubmit">
        <b-form-group id="group1"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.customers.modal.nameInput')"
                      label-for="input1">
          <b-form-input id="input1"
                        type="text"
                        :disabled="isDisabled"
                        :placeholder="$t('message.customers.modal.namePlaceholder')"
                        v-model="newItem.name"></b-form-input>
        </b-form-group>

        <b-form-group id="group2"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.customers.modal.phoneInput')"
                      label-for="input2">
          <b-form-input id="input2"
                        type="text"
                        :disabled="isDisabled"
                        :placeholder="$t('message.customers.modal.phonePlaceholder')"
                        aria-describedby="inputLiveFeedbackPhone"
                        v-model="newItem.phone"></b-form-input>
          <b-form-invalid-feedback id="inputLiveFeedbackPhone">
            {{ $t("message.customers.modal.phoneError") }}
          </b-form-invalid-feedback>
        </b-form-group>

        <b-form-group id="group6"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.customers.modal.addressInput')"
                      label-for="input6">
          <b-form-input id="input6"
                        type="text"
                        :disabled="isDisabled"
                        :placeholder="$t('message.customers.modal.addressPlaceholder')"
                        v-model="newItem.address"></b-form-input>
        </b-form-group>

        <b-form-group id="groupEmail"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.customers.modal.email')"
                      label-for="inputEmail">
          <b-form-input id="inputEmail"
                        type="text"
                        :disabled="isDisabled"
                        :placeholder="$t('message.customers.modal.emailPlaceholder')"
                        v-model="newItem.email"></b-form-input>
        </b-form-group>

        <b-form-group id="groupCompanyName"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.customers.modal.company_name')"
                      label-for="inputCompanyName">
          <b-form-input id="inputCompanyName"
                        type="text"
                        :disabled="isDisabled"
                        :placeholder="$t('message.customers.modal.company_namePlaceholder')"
                        v-model="newItem.company_name"></b-form-input>
        </b-form-group>

        <b-form-group id="groupRef"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.customers.modal.referee')"
                      label-for="inputRef">
          <b-form-input id="inputRef"
                        type="text"
                        :disabled="isDisabled"
                        :placeholder="$t('message.customers.modal.refereePlaceholder')"
                        v-model="newItem.referee"></b-form-input>
        </b-form-group>

        <b-form-group id="group_urflag"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.customers.modal.urflagInput')"
                      label-for="input_payment_type"
        >
          <b-form-radio-group id="input_payment_type"
                              v-model="newItem.urflag"
                              :options="options"
                              :disabled="isDisabled"
                              name="radioOpetions">
          </b-form-radio-group>
        </b-form-group>

        <b-form-group id="group8"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.customers.modal.descriptionInput')"
                      label-for="input7">
          <b-form-textarea id="textarea1"
                           :disabled="isDisabled"
                           :placeholder="$t('message.customers.modal.descriptionPlaceholder')"
                           v-model="newItem.description"
                           :rows="3"
                           :max-rows="6">
          </b-form-textarea>
        </b-form-group>
        <b-form-group id="group_rate"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.customers.modal.rateInput')"
                      label-for="input_rate">
          <b-form-input id="input_rate"
                        type="number"
                        :disabled="isDisabled"
                        v-model.number="newItem.rate"></b-form-input>
        </b-form-group>
        <b-form-group id="group_sale"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.customers.modal.saleInput')"
                      label-for="input_sale">
          <b-form-input id="input_sale"
                        type="number"
                        :disabled="isDisabled"
                        v-model.number="newItem.sale"></b-form-input>
        </b-form-group>

      </b-form>
    </b-modal>

    <!-- Modal Delete Component -->
    <b-modal id="modal-delete"
             ref="modalDelete"
             class="modal-fullscreen1"
             :title="$t('message.customers.modal.removeTitle')"
             @ok="removeItem(deleteRow)"
             :ok-title="$t('message.actions.remove')"
             :cancel-title="$t('message.actions.cancel')"
             centered>
      <p>{{ $t("message.customers.modal.removeDesc") }} "{{deleteRow.name}}"</p>
    </b-modal>

  </div>
</template>

<script>
  export default {
    name: 'Customer',
    props:["titleModel"],
    data() {
      return {
        formUrl: 'customers',
        newItem: {
          id: null,
          name: '',
          phone: '',
          address: '',
          email:'',
          company_name:'',
          active: true,
          description: '',
          town_id: 0,
          urflag: false,
          rate: 0,
          sale: 0,
          referee:'',
        },
        deleteRow: {
          id: null,
          name: '',
        },
        options: [
          {text: 'Юр лицо', value: true},
          {text: 'Физ лицо', value: false},
        ],
        isDisabled:true,
      }
    },
    created() {
      this.formUrl = "town/" + this.$route.meta.town + "/" + "customers";
    },
    methods: {
      showModal(row,isView) {
        if (row.item) {
          this.setNewItem(row);
        } else {
          this.clearNewItem()
        }
        this.isDisabled = isView;
        this.$refs.modal.show()
      },
      hideModal() {
        if (this.$refs.fileinput){
          this.$refs.fileinput.reset();
        }
        this.fileProducts = null;
        this.$refs.modal.hide();
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
      clearNewItem() {
        this.newItem.id = null;
        this.newItem.name = '';
        this.newItem.phone = '';
        this.newItem.address = '';
        this.newItem.email = '';
        this.newItem.company_name = '';
        this.newItem.active = true;
        this.newItem.description = '';
        this.newItem.town_id = 0;
        this.newItem.urflag = false;
        this.newItem.rate = 0;
        this.newItem.sale = 0;
        this.newItem.referee = '';
      },
      setNewItem(row){
        this.newItem = row.item;
      },
      createProduct(evt) {
        evt.preventDefault();
        this.handleSubmit()
      },
      handleSubmit() {
        let self;
        let url = this.formUrl;
        this.newItem.town_id = this.$route.meta.town;
        if (this.newItem.id > 0) {
          console.log(this.newItem);
          self = this;
          url = this.formUrl + '/' + this.newItem.id;
          self.put(url, self.newItem, function (result) {
            console.log(result);
            if (result.status === 422) {
              self.isValidPhone = false
            } else if (result.status === 200) {
              self.newItem.id = result.body.id;
              self.hideModal();
              self.$emit('refresh');
            }
          },error => {
            console.log("ERROR", error);
          });
        } else {
          let self = this;
          self.post(url, this.newItem, function (result) {
            console.log(result);
            if (result.status === 422) {
              self.isValidPhone = false
            } else if (result.status === 200) {

              self.newItem.id = result.body.id;

              self.hideModal();
              self.$emit('refresh');
            }
          },error => {
            console.log("ERROR", error);
          });
        }
      },
      telUrl(newItem){
        return "tel:"+newItem.phone
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
      showModalDelete(row) {
        console.log("Customer showModalDelete");
        if (row.item) {
          this.deleteRow = row.item
        }
        this.$refs.modalDelete.show()
      },
      hideModalDelete() {
        this.$refs.modalDelete.hide()
      }
    },
    watch: {
    }
  }
</script>

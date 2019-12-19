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
                      :label="$t('message.workers.modal.nameInput')"
                      label-for="input1">
          <b-form-input id="input1"
                        type="text"
                        :disabled="isDisabled"
                        :placeholder="$t('message.workers.modal.namePlaceholder')"
                        v-model="newItem.name"></b-form-input>
        </b-form-group>

        <b-form-group id="group2"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.workers.modal.phoneInput')"
                      label-for="input2">
          <b-form-input id="input2"
                        type="text"
                        :disabled="isDisabled"
                        :placeholder="$t('message.workers.modal.phonePlaceholder')"
                        aria-describedby="inputLiveFeedbackPhone"
                        v-model="newItem.phone"></b-form-input>
          <b-form-invalid-feedback id="inputLiveFeedbackPhone">
            {{ $t("message.workers.modal.phoneError") }}
          </b-form-invalid-feedback>
        </b-form-group>

        <b-form-group id="group3"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.workers.modal.ageInput')"
                      label-for="input3">
          <b-form-input id="input3"
                        type="text"
                        :disabled="isDisabled"
                        :placeholder="$t('message.workers.modal.agePlaceholder')"
                        v-model="newItem.age"></b-form-input>
        </b-form-group>
        <b-form-group id="group4"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.workers.modal.passport_typeInput')"
                      label-for="input4">
          <b-form-input id="input4"
                        type="text"
                        :disabled="isDisabled"
                        :placeholder="$t('message.workers.modal.passport_typePlaceholder')"
                        v-model="newItem.passport_type"></b-form-input>
        </b-form-group>
        <b-form-group id="group5"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.workers.modal.passport_urlInput')"
                      label-for="input5"
        >
          <b-form-file id="input5"
                       ref="fileinput"
                       :placeholder="$t('message.workers.modal.passport_urlPlaceholder')"
                       :disabled="isDisabled"
                       v-if = !isShowFile(newItem)
                       v-model="fileProducts"></b-form-file>

          <div id="input5" v-if = isShowFile(newItem) style="display: flex;align-items: center;justify-content: left;">
            <a :href=imageUrl(newItem) target="_blank">
              <b-img  width="75" height="75" alt="File" left :src=imageUrl(newItem) class="m-1" ></b-img>
            </a>
            <b-btn v-if="!isDisabled" variant="danger" size="sm" class="mx-2 md-0" @click="removeFile(newItem)">{{ $t("message.common.searchClear") }}</b-btn>
          </div>

        </b-form-group>
        <b-form-group id="group6"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.workers.modal.addressInput')"
                      label-for="input6">
          <b-form-input id="input6"
                        type="text"
                        :disabled="isDisabled"
                        :placeholder="$t('message.workers.modal.addressPlaceholder')"
                        v-model="newItem.address"></b-form-input>
        </b-form-group>
        <b-form-group id="group7"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.workers.modal.cart_numberInput')"
                      label-for="input7">
          <b-form-input id="input7"
                        type="text"
                        :disabled="isDisabled"
                        :placeholder="$t('message.workers.modal.cart_numberPlaceholder')"
                        v-model="newItem.cart_number"></b-form-input>
        </b-form-group>

        <b-form-group id="group8"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.workers.modal.descriptionInput')"
                      label-for="input7">
          <b-form-textarea id="textarea1"
                           :disabled="isDisabled"
                           :placeholder="$t('message.workers.modal.descriptionPlaceholder')"
                           v-model="newItem.description"
                           :rows="3"
                           :max-rows="6">
          </b-form-textarea>
        </b-form-group>
        <b-form-group id="group10"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.workers.modal.skillInput')"
                      label-for="input10">
          <b-form-input id="input10"
                        type="text"
                        :disabled="isDisabled"
                        :placeholder="$t('message.workers.modal.skillPlaceholder')"
                        v-model="newItem.skill"></b-form-input>
        </b-form-group>

        <b-form-group id="group11"
                      horizontal
                      :label-cols='2'
                      label-size='sm'
                      label='Small'
                      :label="$t('message.workers.modal.readyInput')"
                      label-for="input11">
          <b-form-input id="input11"
                        type="text"
                        :disabled="isDisabled"
                        :placeholder="$t('message.workers.modal.readyPlaceholder')"
                        v-model="newItem.ready"></b-form-input>
        </b-form-group>

      </b-form>
    </b-modal>

    <!-- Modal Delete Component -->
    <b-modal id="modal-delete"
             ref="modalDelete"
             class="modal-fullscreen1"
             :title="$t('message.workers.modal.removeTitle')"
             @ok="removeItem(deleteRow)"
             :ok-title="$t('message.actions.remove')"
             :cancel-title="$t('message.actions.cancel')"
             centered>
      <p>{{ $t("message.workers.modal.removeDesc") }} "{{deleteRow.name}}"</p>
    </b-modal>

  </div>
</template>

<script>
  export default {
    name: 'Worker',
    props:["titleModel"],
    data() {
      return {
        serverUrl:"http://178.62.14.228:5001",
        formUrl: 'workers',
        newItem: {
          id: null,
          name: '',
          phone: '',
          address: '',
          age: '',
          passport_type: '',
          passport_url: '',
          cart_number: '',
          rating: 0.0,
          active_today: false,
          active: true,
          description: '',
          group_id: 1,
          skill: '',
          ready: '',
          town_id: 0,
        },
        deleteRow: {
          id: null,
          name: '',
        },
        fileProducts: null,
        isDisabled:true,
      }
    },
    created() {
      this.formUrl = "town/" + this.$route.meta.town + "/" + "workers";
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
        this.newItem.age = '';
        this.newItem.address = '';
        this.newItem.passport_type = '';
        this.newItem.passport_url = '';
        this.newItem.cart_number = '';
        this.newItem.rating = 0.0;
        this.newItem.active_today = true;
        this.newItem.active = true;
        this.newItem.description = '';
        this.newItem.group_id = 1;
        this.newItem.skill = '';
        this.newItem.ready = '';
        this.newItem.town_id = 0;

        this.fileProducts = null;
      },
      setNewItem(row){
        this.newItem.id = row.item.id;
        this.newItem.name = row.item.name;
        this.newItem.phone = row.item.phone;
        this.newItem.age = row.item.age;
        this.newItem.address = row.item.address;
        this.newItem.passport_type = row.item.passport_type;
        this.newItem.passport_url = row.item.passport_url;
        this.newItem.cart_number = row.item.cart_number;
        this.newItem.rating = row.item.rating;
        this.newItem.active_today = row.item.active_today;
        this.newItem.active = row.item.active;
        this.newItem.description = row.item.description;
        this.newItem.group_id = row.item.group_id;
        this.newItem.skill = row.item.skill;
        this.newItem.ready = row.item.ready;
        this.newItem.town_id = row.item.town_id;
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

              if (self.fileProducts != null){
                let url = self.formUrl + "/import/"+self.newItem.id;
                let formData = new FormData();
                formData.append('file', self.fileProducts);

                self.$http.post(url, formData, null).then(result => {
                  console.log(result);
                  if (result.status === 200) {
                    self.hideModal();
                    self.$emit('refresh');
                  }
                },error => {
                  console.log("ERROR", error);
                  self.hideModal();
                  self.$emit('refresh');

                })
              }else {
                console.log("Not file");
                self.hideModal();
                self.$emit('refresh');

              }
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

              if (self.fileProducts != null){
                let url = self.formUrl + "/import/"+self.newItem.id;
                let formData = new FormData();
                formData.append('file', self.fileProducts);

                self.$http.post(url, formData, null).then(result => {
                  console.log(result);
                  if (result.status === 200) {
                    self.hideModal();
                    self.$emit('refresh');

                  }
                },error => {
                  console.log("ERROR", error);
                  self.hideModal();
                  self.$emit('refresh');

                })
              }else {
                console.log("Not file");
                self.hideModal();
                self.$emit('refresh');
              }
            }
          },error => {
            console.log("ERROR", error);

          });
        }
      },
      imageUrl(newItem){
        return this.serverUrl + '/' + newItem.passport_url
      },
      telUrl(newItem){
        return "tel:"+newItem.phone
      },
      isShowFile(newItem){
        return newItem.passport_url.length > 0
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
        console.log("Worker showModalDelete");
        if (row.item) {
          this.deleteRow = row.item
        }
        this.$refs.modalDelete.show()
      },
      hideModalDelete() {
        this.$refs.modalDelete.hide()
      },
      removeFile(row) {
        row.passport_url = ''
      },
      isShowFile(newItem){
        return newItem.passport_url.length > 0
      },
    },
    watch: {
    }
  }
</script>

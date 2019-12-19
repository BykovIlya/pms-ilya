<template>
  <div>
    <b-navbar toggleable="md" type="dark" :variant="background" >
      <b-navbar-brand>{{titleModel}}</b-navbar-brand>
      <b-navbar-nav class="ml-auto">
        <b-nav-form>
          <!--<b-form-group id="group16_1"-->
                        <!--horizontal-->
                        <!--:label-cols='3'-->
                        <!--label-size='sm'-->
                        <!--label='Small'-->
                        <!--:label="$t('message.tasks.modal.status_id')"-->
                        <!--label-for="input16_1"-->
                        <!--label-class="topLabelColor"-->
          <!--&gt;-->
            <!--<b-form-select id="input16_1" v-model="newItem.status_id" :options="status_options" :disabled="isDisabled"-->
                           <!--style="width:160px;">-->
            <!--</b-form-select>-->
          <!--</b-form-group>-->
        </b-nav-form>
        <b-nav-item>
          <b-button variant="success" :disabled="isDisabled" @click="handleSubmit()">Сохранить</b-button>
          <b-button variant="danger" @click.stop="hideModal()">Отменить</b-button>
        </b-nav-item>
      </b-navbar-nav>
    </b-navbar>

    <!--<b-breadcrumb :items="$route.meta.breadcrumbs"/>-->
    <b-form @submit.stop.prevent="handleSubmit">
      <b-container fluid class="mt-2">
        <b-row>
          <b-col md="6">
            <b-card bg-variant="light">
              <b-row>
                <b-col>
                  <b-form-group id="group_customer"
                                :label-cols='2'
                                label-size='md'
                                label='Small'
                                :label="$t('message.tasks.modal.customer')"
                                label-for="input_customer">

                    <b-input-group>
                      <b-form-input id="input_customer"
                                    type="text"
                                    readonly
                                    :disabled="isDisabled"
                                    :placeholder="$t('message.tasks.modal.customerPlaceholder')"
                                    v-model="newItem.customer_name"></b-form-input>
                      <b-input-group-append>
                        <b-btn variant="outline-success" :disabled="isDisabled" @click.stop="showCustomerTable()">Выбрать</b-btn>
                        <b-btn variant="danger" :disabled="isDisabled" @click.stop="removeCustomer()">Очистить</b-btn>
                      </b-input-group-append>
                    </b-input-group>

                  </b-form-group>
                </b-col>
              </b-row>
              <b-row>
                <b-col>
                  <b-form-group id="group_contact"
                                :label-cols='2'
                                label-size='md'
                                label='Small'
                                :label="$t('message.tasks.modal.contact')"
                                label-for="input_contact">
                    <b-form-input id="input_contact"
                                  type="text"
                                  :disabled="isDisabled"
                                  :placeholder="$t('message.tasks.modal.contactPlaceholder')"
                                  v-model="newItem.contact"></b-form-input>
                  </b-form-group>
                </b-col>
              </b-row>
              <b-row>
                <b-col>
                  <b-form-group id="group_start_address"
                                :label-cols='2'
                                label-size='md'
                                label='Small'
                                :label="$t('message.tasks.modal.start_address')"
                                label-for="input_start_address">
                    <b-form-input id="input_start_address"
                                  type="text"
                                  :disabled="isDisabled"
                                  :placeholder="$t('message.tasks.modal.start_addressPlaceholder')"
                                  v-model="newItem.start_address"></b-form-input>
                  </b-form-group>
                </b-col>
                <b-col>
                  <b-form-group id="group_end_address"
                                :label-cols='2'
                                label-size='sm'
                                label='Small'
                                :label="$t('message.tasks.modal.end_address')"
                                label-for="input_end_address">
                    <b-form-input id="input_end_address"
                                  type="text"
                                  :disabled="isDisabled"
                                  :placeholder="$t('message.tasks.modal.end_addressPlaceholder')"
                                  aria-describedby="inputLiveFeedbackPhone"
                                  v-model="newItem.end_address"></b-form-input>
                  </b-form-group>
                </b-col>
              </b-row>
              <b-row>
                <b-col>
                  <b-form-group id="group_startDate"
                                :label-cols='2'
                                label-size='sm'
                                label='Small'
                                :label="$t('message.tasks.modal.start_date')"
                                label-for="input_startDate">
                    <b-form-input id="input_startDate"
                                  type="date"
                                  :disabled="disabledStartDate"
                                  :placeholder="$t('message.tasks.modal.start_datePlaceholder')"
                                  v-model.date="startDate"></b-form-input>
                  </b-form-group>
                </b-col>
                <b-col>
                  <b-form-group id="group_startTime"
                                :label-cols='2'
                                label-size='sm'
                                label='Small'
                                :label="$t('message.tasks.modal.start_time')"
                                label-for="input_startTime">
                    <b-form-input id="input_startTime"
                                  type="time"
                                  :disabled="isDisabled"
                                  v-model.time="startTime"></b-form-input>
                  </b-form-group>
                </b-col>
              </b-row>
              <b-row>
                <b-col>
                  <b-form-group id="group_endDate"
                                :label-cols='2'
                                label-size='sm'
                                label='Small'
                                :label="$t('message.tasks.modal.end_date')"
                                label-for="input_startDate">
                    <b-form-input id="input_startDate"
                                  type="date"
                                  :disabled="disabledStartDate"
                                  v-model.date="endDate"></b-form-input>
                  </b-form-group>
                </b-col>
                <b-col>
                  <b-form-group id="group_endTime"
                                :label-cols='2'
                                label-size='sm'
                                label='Small'
                                :label="$t('message.tasks.modal.end_time')"
                                label-for="input_endTime">
                    <b-form-input id="input_endTime"
                                  type="time"
                                  :disabled="isDisabled"
                                  v-model.time="endTime"></b-form-input>
                  </b-form-group>
                </b-col>
              </b-row>
              <b-row>
                <b-col cols="2">
                  <b-form-group id="group_payment_first_count"
                                :label-cols='2'
                                label-size='sm'
                                label='Small'
                                :label="$t('message.tasks.modal.payment_first_count')"
                                label-for="input_payment_first_count">
                    <b-form-input id="input_payment_first_count"
                                  type="number"
                                  min="0"
                                  :disabled="isDisabled"
                                  v-model.number="newItem.payment_first_count"></b-form-input>
                  </b-form-group>
                </b-col>
                <b-col>
                  <b-form-group id="group6"
                                :label-cols='2'
                                label-size='sm'
                                label='Small'
                                :label="$t('message.tasks.modal.payment_first')"
                                label-for="input6">
                    <b-form-input id="input6"
                                  type="number"
                                  min="0"
                                  :disabled="isDisabled"
                                  :placeholder="$t('message.tasks.modal.payment_firstPlaceholder')"
                                  v-model.number="newItem.payment_first"></b-form-input>
                  </b-form-group>
                </b-col>


                <b-col>
                  <b-form-group id="group_payment_second"
                                :label-cols='2'
                                label-size='sm'
                                label='Small'
                                :label="$t('message.tasks.modal.payment_second')"
                                label-for="input_payment_second">
                    <b-form-input id="input_payment_second"
                                  type="number"
                                  min="0"
                                  :disabled="isDisabled"
                                  :placeholder="$t('message.tasks.modal.payment_secondPlaceholder')"
                                  v-model.number="newItem.payment_second"></b-form-input>
                  </b-form-group>
                </b-col>

                <b-col cols="2">
                  <!--<b-form-group id="group_payment_second_count"-->
                  <!--:label-cols='2'-->
                  <!--label-size='sm'-->
                  <!--label='Small'-->
                  <!--:label="$t('message.tasks.modal.payment_second_count')"-->
                  <!--label-for="input_payment_second_count">-->
                  <!--<b-form-input id="input_payment_second_count"-->
                  <!--type="number"-->
                  <!--min="0"-->
                  <!--:disabled="isDisabled"-->
                  <!--v-model.number="newItem.payment_second_count"></b-form-input>-->
                  <!--</b-form-group>-->
                </b-col>

              </b-row>
              <b-row>
                <b-col>
                  <b-form-group id="group_cargo_description"
                                :label-cols='2'
                                label-size='sm'
                                label='Small'
                                :label="$t('message.tasks.modal.cargo_description')"
                                label-for="textarea_cargo_description">
                    <b-form-textarea id="textarea_cargo_description"
                                     :disabled="isDisabled"
                                     :placeholder="$t('message.tasks.modal.cargo_description')"
                                     v-model="newItem.cargo_description"
                                     :rows="4"
                                     :max-rows="4">
                    </b-form-textarea>
                  </b-form-group>
                </b-col>
              </b-row>
              <b-row>
                <b-col>
                  <b-form-group id="group_payment_type"
                                :label-cols='3'
                                label-size='sm'
                                label='Small'
                                :label="$t('message.tasks.modal.payment_type')"
                                label-for="input_payment_type"
                  >
                    <b-form-radio-group id="input_payment_type"
                                        v-model="newItem.payment_type"
                                        :disabled="isDisabled"
                                        :options="options"
                                        name="radioOpetions">
                    </b-form-radio-group>
                  </b-form-group>
                </b-col>
              </b-row>
            </b-card>
          </b-col>
          <b-col md="6">
            <b-card bg-variant="light">
              <b-row>
                <b-col>
                  <b-form-group id="group_additional_cost"
                                :label-cols='2'
                                label-size='sm'
                                label='Small'
                                :label="$t('message.tasks.modal.additional_cost')"
                                label-for="input_additional_cost">
                    <b-form-input id="input_additional_cost"
                                  type="number"
                                  :disabled="isDisabled"
                                  v-model.number="newItem.additional_cost"></b-form-input>
                  </b-form-group>
                  <b-form-group id="group_additional_income"
                                :label-cols='2'
                                label-size='sm'
                                label='Small'
                                :label="$t('message.tasks.modal.additional_income')"
                                label-for="input_additional_income">
                    <b-form-input id="input_additional_income"
                                  type="number"
                                  :disabled="isDisabled"
                                  v-model.number="newItem.additional_income"></b-form-input>
                  </b-form-group>
                </b-col>
                <b-col>
                  <b-form-group id="group_description"
                                :label-cols='2'
                                label-size='sm'
                                label='Small'
                                :label="$t('message.tasks.modal.description')"
                                label-for="textarea_description">
                    <b-form-textarea id="textarea_description"
                                     :disabled="isDisabled"
                                     :placeholder="$t('message.tasks.modal.descriptionPlaceholder')"
                                     v-model="newItem.description"
                                     :rows="4"
                                     :max-rows="6">
                    </b-form-textarea>
                  </b-form-group>
                </b-col>
              </b-row>

              <b-row>
                <b-col>
                  <b-form-group id="group_total_hours"
                                :label-cols='2'
                                label-size='sm'
                                label='Small'
                                :label="$t('message.tasks.modal.total_hours')"
                                label-for="input_total_hours">
                    <b-form-input id="input_total_hours"
                                  type="number"
                                  :disabled="isDisabled"
                                  readonly
                                  :placeholder="$t('message.tasks.modal.total_hoursPlaceholder')"
                                  v-model.number="newItem.total_hours"></b-form-input>
                  </b-form-group>
                </b-col>
                <b-col>
                  <b-form-group id="group_total_worker_price"
                                :label-cols='2'
                                label-size='sm'
                                label='Small'
                                :label="$t('message.tasks.modal.total_worker_price')"
                                label-for="input_total_worker_price">
                    <b-form-input id="input_total_worker_price"
                                  type="number"
                                  :disabled="isDisabled"
                                  readonly
                                  :placeholder="$t('message.tasks.modal.total_worker_pricePlaceholder')"
                                  v-model.number="newItem.total_worker_price"></b-form-input>
                  </b-form-group>
                </b-col>
                <b-col>
                  <b-form-group id="group_total_price"
                                :label-cols='2'
                                label-size='sm'
                                label='Small'
                                :label="$t('message.tasks.modal.total_price')"
                                label-for="input_total_price">
                    <b-form-input id="input_total_price"
                                  type="number"
                                  readonly
                                  :disabled="isDisabled"
                                  :placeholder="$t('message.tasks.modal.total_pricePlaceholder')"
                                  v-model.number="newItem.total_price"></b-form-input>
                  </b-form-group>
                </b-col>
                <b-col>
                  <b-form-group id="group_paid"
                                :label-cols='1'
                                label-size='sm'
                                label='Small'
                                :label="$t('message.tasks.modal.paid')"
                                label-for="input_paid">
                    <b-form-checkbox id="input_paid"
                                     :disabled="isDisabled || !this.$store.getters.isAdmin"
                                     v-model="newItem.paid">
                    </b-form-checkbox>
                  </b-form-group>

                </b-col>
              </b-row>
              <b-row>
                <b-col>
                  <b-form-group id="group_declarated_cost"
                                :label-cols='2'
                                label-size='sm'
                                label='Small'
                                :label="$t('message.tasks.modal.declarated_cost')"
                                label-for="input_declarated_cost">
                    <b-form-input id="input_declarated_cost"
                                  type="number"
                                  :disabled="isDisabled"
                                  v-model.number="newItem.declarated_cost"></b-form-input>
                  </b-form-group>
                </b-col>
                <b-col>

                </b-col>
              </b-row>
              <b-row>
                <b-col>
                  <b-form-group id="group_total_delta"
                                :label-cols='2'
                                label-size='sm'
                                label='Small'
                                :label="$t('message.tasks.modal.total_delta')"
                                label-for="input_total_delta">
                    <b-form-input id="input_total_delta"
                                  type="number"
                                  readonly
                                  :disabled="isDisabled"
                                  :placeholder="$t('message.tasks.modal.total_deltaPlaceholder')"
                                  v-model.number="newItem.total_delta"></b-form-input>
                  </b-form-group>
                </b-col>
                <b-col>
                  <b-form-group id="group_total_manager"
                                :label-cols='2'
                                label-size='sm'
                                label='Small'
                                :label="$t('message.tasks.modal.total_manager')"
                                label-for="input_total_manager">
                    <b-form-input id="input_total_manager"
                                  type="number"
                                  readonly
                                  :disabled="isDisabled"
                                  :placeholder="$t('message.tasks.modal.total_managerPlaceholder')"
                                  v-model.number="newItem.total_manager"></b-form-input>
                  </b-form-group>
                </b-col>
              </b-row>

              <b-row>
                <b-col>
                  <b-form-group id="group_autor"
                                :label-cols='2'
                                label-size='sm'
                                label='Small'
                                :label="$t('message.tasks.modal.author')"
                                label-for="author">
                    {{author.name}}
                  </b-form-group>
                </b-col>
                <b-col>
                  <b-form-group id="group_executor"
                                :label-cols='2'
                                label-size='sm'
                                label='Small'
                                :label="$t('message.tasks.modal.executor')"
                                label-for="input_executor">
                    {{executor.name}}
                  </b-form-group>
                </b-col>
                <b-col>

                </b-col>
              </b-row>
            </b-card>
          </b-col>
        </b-row>
      </b-container>
    </b-form>

    <b-container fluid class="mt-2">
      <b-row>
        <b-col cols="12">
          <div>
            <b-navbar type="dark" :variant="backgroundTW" sticky>
              <b-navbar-brand>{{titleTW}}</b-navbar-brand>
              <b-navbar-nav>
                <b-btn size="md" variant="success" :disabled="isDisabled" @click="showFreeWorkers">{{ $t("message.workers.importHand") }}
                </b-btn>
              </b-navbar-nav>
              <b-navbar-nav class="ml-auto">
                <b-nav-form>
                  <b-form-input size="md" class="my-1 md-2" v-model="filter"
                                :placeholder="$t('message.common.searchPlaceholder')"></b-form-input>
                  <b-btn size="md" class="my-1 md-0" :disabled="!filter" @click="filter = ''">
                    <font-awesome-icon icon="times-circle" size="lg"/>
                  </b-btn>
                </b-nav-form>
              </b-navbar-nav>
            </b-navbar>

            <b-container id="table-container" fluid>
              <b-table id="workers"
                       ref="TaskWorkers"
                       striped
                       thead-class="fixed-top sticky-top top-fixed-thead"
                       responsive
                       small
                       show-empty
                       :busy.sync="isBusy"
                       :items="newItem.task_workers"
                       :fields="fieldsTW"
                       :current-page="currentPage"
                       :per-page="perPage"
                       :empty-text="$t('message.common.emptyTable')"
                       :sort-by.sync="sortBy"
                       :sort-desc.sync="sortDesc"
                       :filter="filter"
                       @filtered="onFiltered"
              >

                <template slot="action" slot-scope="row">
                  <b-button size="sm" variant="success" @click.stop="showWorkerInfo(row.item.worker,true)">
                    <font-awesome-icon icon="eye" />
                  </b-button>
                  <b-btn size="sm" variant="danger" :disabled="isDisabled"
                         @click="removeWorkerFromTask(row)">
                    <font-awesome-icon icon="trash"/>
                  </b-btn>
                </template>

                <template slot="phone" slot-scope="row">
                  <a :href=telUrl(row.item)>{{row.item.phone}}</a>
                </template>

                <template slot="passport_url" slot-scope="row">
                  <a v-if="row.item.passport_url.length>0" :href=imageUrl(row.item) target="_blank">
                    <font-awesome-icon icon="passport" size="lg" :style="{ color: 'purple' }"/>
                  </a>
                </template>

                <template slot="price" slot-scope="row">
                  <b-form-input id="input_price"
                                type="number"
                                :disabled="isDisabled"
                                class="widthPrice"
                                v-model.number="row.item.price">
                  </b-form-input>
                </template>

                <template slot="start_date" slot-scope="row">

                  <b-form-input id="input_startDate"
                                type="date"
                                class="widthDate"
                                :disabled="isDisabled"
                                v-model.date="row.item.string_start_date"></b-form-input>
                </template>

                <template slot="start_time" slot-scope="row">
                  <b-form-input id="input_startTime"
                                type="time"
                                :disabled="isDisabled"
                                class="widthTime"
                                v-model.date="row.item.start_time"></b-form-input>
                </template>

                <template slot="end_date" slot-scope="row">

                  <b-form-input id="input_endDate"
                                type="date"
                                class="widthDate"
                                :disabled="isDisabled"
                                v-model.date="row.item.string_end_date"></b-form-input>
                </template>

                <template slot="end_time" slot-scope="row">
                  <b-form-input id="input_endTime"
                                type="time"
                                :disabled="isDisabled"
                                class="widthTime"
                                v-model.date="row.item.end_time"></b-form-input>
                </template>

                <template slot="road_price" slot-scope="row">
                  <b-form-input id="input_price"
                                type="number"
                                :disabled="isDisabled"
                                class="widthPrice"
                                v-model.number="row.item.road_price">
                  </b-form-input>
                </template>

                <template slot="dinner" slot-scope="row">
                  <b-form-input id="input_dinner"
                                type="number"
                                :disabled="isDisabled"
                                step="0.5"
                                class="width50"
                                v-model.number="row.item.dinner">
                  </b-form-input>
                </template>

                <template slot="prepaid" slot-scope="row">
                  <b-form-input id="input_prepaid"
                                type="number"
                                :disabled="isDisabled"
                                class="widthPrice"
                                v-model.number="row.item.prepaid">
                  </b-form-input>
                </template>

                <template slot="description" slot-scope="row">
                  <b-form-textarea id="textarea_description"
                                   :disabled="isDisabled"
                                   :placeholder="$t('message.tasks.modal.descriptionPlaceholder')"
                                   v-model="row.item.description"
                                   :rows="2"
                                   :max-rows="2">
                  </b-form-textarea>
                </template>

                <template slot="paid" slot-scope="row">
                  <b-button type="button" :variant="getPaidVariant(row)" :disabled="isDisabled" @click.stop="showModalAccept(row)">
                    {{getPaidTitle(row)}}
                  </b-button>

                </template>
              </b-table>
              <b-row>
                <b-col sm="12">
                  <div align="right">Всего рабочих: {{totalRows}}</div>
                  <b-pagination align="right" :total-rows="totalRows" :per-page="perPage"
                                v-model="currentPage"></b-pagination>
                </b-col>
              </b-row>

            </b-container>
          </div>

          <!--<task-workers ref="taskWorkers" :workers="task_workers" :task_workers="newItem.task_workers"-->
          <!--v-on:showFreeWorkers="showFreeWorkers"></task-workers>-->
        </b-col>
      </b-row>
    </b-container>

    <workers-free ref="freeWorkersTable" :town="currentTown" :date="getStartDateTime()" :end_date="getEndDateTime()" v-on:addToTask="addToTask"/>

    <customer-table-select ref="customerTable" :town="currentTown" v-on:selectCustomer="updateCustomer"/>

    <worker ref="worker" :titleModel="titleTW"></worker>

    <b-modal id="modal-accept"
             ref="modalAccept"
             :title="$t('message.tasks.modal.acceptWorkerTitle')"
             @ok="acceptItem(acceptWorker)"
             :ok-title="$t('message.actions.ok')"
             :cancel-title="$t('message.actions.cancel')"
             centered>
      <p>{{ $t("message.tasks.modal.acceptWorkerDesc") }} "{{acceptWorker.name}}"</p>
    </b-modal>
  </div>
</template>


<script>
  import WorkersFree from "./WorkersFree";
  import CustomerTableSelect from "./Customers/CustomerTableSelect";
  import TaskWorkers from "./TaskWorkers";
  import Worker from "./Worker";

  // ISO8601 in local time zone
  function localISOString(d) {

    //var d = new Date()
    var pad = function (n) {
      return n < 10 ? '0' + n : n;
    }
    //   , tz = d.getTimezoneOffset() // mins
    //   , tzs = (tz>0?"-":"+") + pad(parseInt(Math.abs(tz/60)));
    //
    // if (tz%60 != 0)
    //   tzs += pad(Math.abs(tz%60));
    //
    // if (tz === 0) // Zulu time == UTC
    //   tzs = 'Z';

    return d.getFullYear() + '-'
      + pad(d.getMonth() + 1) + '-'
      + pad(d.getDate()) + 'T'
      + pad(d.getHours()) + ':'
      + pad(d.getMinutes()) + ':'
      //+ pad(d.getSeconds()) + tzs+":00";
      + pad(d.getSeconds()) + "+03:00";
  };

  export default {
    components: {
      TaskWorkers,
      CustomerTableSelect,
      WorkersFree,
      Worker
    },
    name: 'Task',
    props: ['isRead', 'row'],
    data() {
      return this.defaultData()
    },
    created() {
      if (!this.$store.getters.isLoggedIn) {
        this.$router.push('/login');
      } else {
        //this.showForbiddenAlert = false;
      }

      this.showModal(this.$props.row, this.$props.isRead);

      this.userInfo();

      this.statusUrl = "taskstatus";
      this.getStatuses();

    },
    methods: {
      defaultData() {
        return {
          currentTown: 0,
          formUrl: 'tasks',
          //workersUrl: 'workers',
          statusUrl: 'status',
          customersUrl: 'customers',
          titleModel: "Добавление заявки",
          background: 'info',
          disabledStartDate: false,
          startDate: this.formatDate(new Date()),
          startTime: '',
          endDate: this.formatDate(new Date()),
          endTime: '',
          options: [
            {text: 'Нал', value: true},
            {text: 'Без нал', value: false},
          ],
          status_options: [],
          user: '',
          author: '',
          executor: '',
          newItem: {
            id: null,
            start_address: '',
            end_address: '',
            start_date: '',
            end_date: '',
            payment_type: true,
            payment_first_count: 0,
            payment_first: 250,
            payment_second_count: 0,
            payment_second: 0,
            workers_count: 0,
            total_hours: 0,
            total_price: 0,
            total_worker_price: 0,
            total_delta: 0,
            total_manager: 0,
            additional_cost: 0,
            additional_income: 0,
            description: '',
            cargo_description: '',
            status_id: 1,
            customer_id: 0,
            customer_name: '',
            author: 0,
            executor: 0,
            paid: false,
            contact: '',
            declarated_cost: 0,
            task_workers: [],
            customer:{},
          },
          task_workers: [],
          fields: [
            {
              key: 'active_today',
              label: this.$t('message.workers.table.active_today'),
              sortable: true,
            },
            {
              key: 'name',
              label: this.$t('message.workers.table.name'),
              sortable: true,
            },
            {
              key: 'phone',
              label: this.$t('message.workers.table.phone'),
              sortable: true,
            },
            {
              key: 'action',
              label: this.$t('message.workers.table.actions'),
              class: 'actions-width'
            },
          ],
          isDisabled: false,

          // Task Workers Default Data
          titleTW: 'Рабочие',
          backgroundTW: 'info',
          formUrlTW: 'workers',
          sortBy: 'name',
          sortDesc: false,
          filter: null,
          // items:[],
          fieldsTW: [
            {
              key: 'name',
              label: this.$t('message.workers.table.name'),
              sortable: true,
            },
            {
              key: 'phone',
              label: this.$t('message.workers.table.phone'),
              sortable: true,
            },
            {
              key: 'price',
              label: this.$t('message.workers.table.price'),
              sortable: true,
            },
            {
              key: 'start_date',
              label: this.$t('message.workers.table.start_date'),
              sortable: true,
            },
            {
              key: 'start_time',
              label: this.$t('message.workers.table.start_time'),
              sortable: true,
            },
            {
              key: 'end_date',
              label: this.$t('message.workers.table.end_date'),
              sortable: true,
            },
            {
              key: 'end_time',
              label: this.$t('message.workers.table.end_time'),
              sortable: true,
            },
            {
              key: 'dinner',
              label: this.$t('message.workers.table.dinner'),
              sortable: true,
            },
            {
              key: 'total_hours',
              label: this.$t('message.workers.table.total_hours'),
              sortable: true,
            },
            {
              key: 'prepaid',
              label: this.$t('message.workers.table.prepaid'),
              sortable: true,
            },
            {
              key: 'road_price',
              label: this.$t('message.workers.table.road_price'),
              sortable: true,
            },

            {
              key: 'total_price',
              label: this.$t('message.workers.table.total_price'),
              sortable: true,
            },
            {
              key: 'description',
              label: this.$t('message.workers.table.description'),
              sortable: true,
            },
            {
              key: 'paid',
              label: this.$t('message.workers.table.paid'),
              sortable: true,
              'class': 'text-center'
            },
            {
              key: 'action',
              label: this.$t('message.workers.table.actions')
            },
          ],
          currentPage: 1,
          perPage: 1000,
          isBusy: false,
          totalRows: 0,
          acceptWorker:{
            worker_id:null,
            name:''
          },
        }
      },
      showModal(row, isRead) {
        if (isRead) {
          this.isDisabled = isRead;
        }
        if (row) {
          if (row.item) {
            this.currentTown = this.$route.meta.town;
            this.updateView(this.$route.meta.town);
            this.setNewItem(row);
          } else {
            this.clearNewItem();
          }
        } else {
          this.clearNewItem();
        }

      },
      hideModal() {
        this.goBack();
        // if (this.currentTown === 0){
        //   this.$router.push({name: 'Tasks'});
        // }else if (this.currentTown === 1)
        // {
        //   this.$router.push({name: 'TasksSpb'})
        // }
      },
      goBack() {
        if (window.history.length > 1) {
          console.log("-1");
          this.$router.go(-1)
        } else {
          console.log("/");
          this.$router.push('/')
        }
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
        Object.assign(this.$data, this.defaultData());
        this.currentTown = this.$route.meta.town;
        this.updateView(this.$route.meta.town);
      },
      setNewItem(row) {

        this.newItem = row.item;
        let tws = this.newItem.task_workers;
        this.newItem.task_workers = [];

        let startD = new Date(row.item.start_date);
        let endD = new Date(row.item.end_date);
        this.startDate = this.formatDate(startD);
        this.startTime = this.formatTime(startD);
        this.endDate = this.formatDate(endD);
        this.endTime = this.formatTime(endD);

        this.customersUrl = "town/0/customers";
        if (this.newItem.customer_id > 0) {
          this.getCustomerById(this.newItem.customer_id)
        }

        let self = this;
        if (tws) {
          if (tws.length > 0) {
            tws.forEach(function (item, i, arr) {
              self.getWorkerById(item.worker_id, item);
            });

            this.disabledChangeStartDate(tws)
          }
        }
      },

      handleSubmit() {
        let url = this.formUrl;
        this.newItem.town_id = this.$route.meta.town;
        this.newItem.author = Number(this.newItem.author);
        this.newItem.executor = Number(this.newItem.executor);

        // let dateStart = new Date(this.startDate + " " + this.startTime);
        // let dateEnd = new Date(this.endDate + " " + this.endTime);
        // this.newItem.start_date = dateStart.toISOString();
        // this.newItem.end_date = dateEnd.toISOString();

        this.newItem.start_date = this.formatServerDateTime(this.startDate, this.startTime);
        this.newItem.end_date = this.formatServerDateTime(this.endDate, this.endTime);

        this.parseNulls();
        if (this.newItem.id > 0) {
          console.log(this.newItem);
          let self = this;
          url = this.formUrl + '/' + this.newItem.id;
          self.put(url, self.newItem, function (result) {
            console.log(result);
            if (result.status === 422) {

            } else if (result.status === 200) {
              self.newItem.id = result.body.id;
              self.hideModal();
              self.$emit('refresh');
            }
          }, error => {
            console.log("ERROR", error);
          });
        }
        else {
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
          }, error => {
            console.log("ERROR", error);
          });
        }
      },
      parseNulls() {
        //Проверки на нулевые значения в number
        if (!this.newItem.payment_first_count) {
          this.newItem.payment_first_count = 0;
        }
        if (!this.newItem.payment_first) {
          this.newItem.payment_first = 0;
        }
        if (!this.newItem.payment_second_count) {
          this.newItem.payment_second_count = 0;
        }
        if (!this.newItem.payment_second) {
          this.newItem.payment_second = 0;
        }

        if (!this.newItem.additional_cost) {
          this.newItem.additional_cost = 0;
        }
        if (!this.newItem.additional_income) {
          this.newItem.additional_income = 0;
        }
        if (!this.newItem.declarated_cost) {
          this.newItem.declarated_cost = 0;
        }

        this.newItem.task_workers.forEach(function (item, i, arr) {
          if (!item.price) {
            item.price = 0.0;
          }
          if (!item.road_price) {
            item.road_price = 0.0;
          }
          if (!item.total_price) {
            item.total_price = 0.0;
          }
          if (!item.prepaid) {
            item.prepaid = 0.0;
          }
          if (!item.dinner) {
            item.dinner = 0.0;
          }
        });
      },
      paymentTypeString(flag) {
        if (flag) {
          return "Нал"
        }
        return "Безнал"
      },
      formatDate(date) {
        let dd = date.getDate();
        if (dd < 10) dd = '0' + dd;
        let mm = date.getMonth() + 1;
        if (mm < 10) mm = '0' + mm;
        let yy = date.getFullYear();
        return yy + '-' + mm + '-' + dd;
      },
      formatTime(date) {
        // console.log(date);
        // let date1 = new Date(localISOString(date));
        // console.log(localISOString(date));
        // console.log(date1);
        let HH = date.getHours();
        if (HH < 10) HH = '0' + HH;
        let MM = date.getMinutes();
        if (MM < 10) MM = '0' + MM;
        return HH + ':' + MM;
      },
      formatServerDateTime(date, timeStr) {
        let dateItem = new Date(date + " " + timeStr);
        return localISOString(dateItem);
      },
      getStatuses() {
        let url = this.statusUrl;
        //this.isBusy = true;
        let self = this;
        return this.$http.get(url).then(result => {
          console.log(result);
          if (result.status === 403) {
            this.showForbiddenAlert = true;
          } else if (result.status === 200 || result.status === 304) {
            this.showForbiddenAlert = false;
            if (result.body != null) {
              self.status_options = [];
              result.body.forEach(function (e) {
                var disabled = false;
                // if (self.newItem.status_id>=5){
                //   if(e.id<5){
                //     disabled = true;
                //   }
                // }
                //
                // if(e.id === 7 ){
                //   disabled = false;
                // }

                self.status_options.push({value: e.id, text: e.name, disabled: disabled})
              });
            } else {
              this.status_options = [];
            }
          }
        }, error => {
          //this.isBusy = false;
          console.log("ERROR", error);
        });
      },
      showCustomerTable() {
        this.$refs.customerTable.showModalSelect();
      },
      updateCustomer(customer) {
        this.newItem.customer = customer;
        this.newItem.customer_id = customer.id;
        this.newItem.customer_name = customer.name;
        this.updateTotals();
      },
      getCustomerById(customer_id) {
        let url = this.customersUrl + '/id/' + customer_id;
        let self = this;
        return this.$http.get(url).then(result => {
          console.log(result);
          if (result.status === 403) {
            this.showForbiddenAlert = true;
          } else if (result.status === 200 || result.status === 304) {
            this.showForbiddenAlert = false;
            if (result.body != null) {
              self.updateCustomer(result.body);
            } else {
              self.newItem.customer = {};
              self.newItem.customer_name = '';
            }
          }
        }, error => {
          //this.isBusy = false;
          console.log("ERROR", error);
        });
      },
      removeCustomer() {
        this.newItem.customer_name = '';
        this.newItem.customer_id = 0;
        this.newItem.customer = {};
      },
      updateTotals() {
        //if (this.newItem.status_id < 5) {
          this.newItem.total_worker_price = this.calculateWorkersPrice();
          this.newItem.total_hours = this.calculateWorkersHours();
          let transfer = this.calculateWorkersTransfer();

        if (this.newItem.payment_first != null) {
          this.newItem.total_price = (this.newItem.payment_first * this.newItem.total_hours) + transfer - this.newItem.additional_cost + this.newItem.additional_income;
          if ((this.newItem.payment_second != null) && (this.newItem.payment_second > 0) &&(this.newItem.payment_first_count>0)) {
            let firstCountHours = this.newItem.payment_first_count*this.newItem.task_workers.length;
            if (firstCountHours>0){
              this.newItem.total_price = this.newItem.payment_first * firstCountHours + (this.newItem.payment_second * (this.newItem.total_hours-firstCountHours)) + transfer - this.newItem.additional_cost + this.newItem.additional_income;
            }
          }
        }

        //}
      },
      updateTotalDelta() {
        //if (this.newItem.status_id < 5) {
          if ((this.newItem.total_price != null) && (this.newItem.total_worker_price != null)) {


            if (this.newItem.payment_type) {
              //физ лицо
              this.newItem.total_delta = this.newItem.total_price - this.newItem.total_worker_price;
            } else {
              //юр лицо
              this.newItem.total_delta = (this.newItem.total_price - this.newItem.total_worker_price) - ((this.newItem.total_price / 1.20) * 0.20);
            }

            //Аванс не должен идти в дельту
            let prepaid = this.calculateWorkersPrepaid();
            this.newItem.total_delta = this.newItem.total_delta - prepaid;

            if (this.newItem.customer){
              if(this.newItem.customer.sale>0){
                this.newItem.total_delta = this.newItem.total_delta - (this.newItem.total_hours * this.newItem.customer.sale)
              }
            }


            this.newItem.total_delta = Math.round(this.newItem.total_delta * 100) / 100;
            this.newItem.total_delta = this.newItem.total_delta - this.newItem.declarated_cost;

            this.updateTotalManager();
          }
        //}
      },
      updateTotalManager() {
        let rate = 0;
        if (this.newItem.customer!=null){
          if (this.newItem.customer.rate>0){
            rate = this.newItem.customer.rate
          }else {
            if (this.newItem.payment_type) {
              rate = this.executor.rate
            }else {
              rate = this.executor.rate_tax
            }
          }
        }else {
          if (this.newItem.payment_type) {
            rate = this.executor.rate
          }else {
            rate = this.executor.rate_tax
          }
        }

        this.newItem.total_manager = this.newItem.total_delta * rate / 100;
        this.newItem.total_manager = Math.round(this.newItem.total_manager * 100) / 100;
      },


      //Task Workers Methods
      showFreeWorkers() {
        this.$refs.freeWorkersTable.showModalSelect(this.currentTown);
      },
      addToTask(worker) {
        if (!this.existWorker(worker.id)) {
          let st = this.formatTime(new Date());
          let et = this.formatTime(new Date());
          if (this.startTime) {
            st = this.startTime;
          }
          if (this.endTime) {
            et = this.endTime;
          }
          let taskWorker = {
            worker_id: worker.id,
            name: worker.name,
            phone: worker.phone,
            price: 0.0,
            road_price: 0.0,
            start_time: st,
            end_time: et,
            total_price: 0.0,
            prepaid: 0.0,
            dinner: 0.0,
            description: '',
            start_date: this.formatServerDateTime(this.startDate, st),
            end_date: this.formatServerDateTime(this.endDate, et),
            string_start_date: this.startDate,
            string_end_date: this.endDate,
            paid: false
          };
          taskWorker.total_hours = this.getTotalHours(taskWorker);
          this.newItem.task_workers.push(taskWorker)

        } else {
          console.log("Worker not Added, because Worker exist")
        }
      },
      getStartDateTime() {
        let st = this.formatTime(new Date());
        if (this.startTime) {
          st = this.startTime;
        }
        return this.formatServerDateTime(this.startDate, st)
      },
      getEndDateTime() {
        let st = this.formatTime(new Date());
        if (this.endTime) {
          st = this.endTime;
        }
        return this.formatServerDateTime(this.endDate, st)
      },
      existWorker(id) {
        let exist = false;
        for (let i = 0; i < this.newItem.task_workers.length; i++) {
          let item = this.newItem.task_workers[i];
          if (item.worker_id === id) {
            exist = true;
            break;
          }
        }
        return exist
      },
      removeWorkerFromTask(row) {
        let index = this.newItem.task_workers.indexOf(row.item);
        this.newItem.task_workers.splice(index, 1);
      },
      getTotalHours(tw) {
        let startTime = new Date(tw.start_date);
        let endTime = new Date(tw.end_date);
        let diff = ((endTime - startTime) / 60000 / 60);
        diff = Math.round(diff * 100) / 100;
        //Вычитаем обед
        diff = diff - tw.dinner;
        return diff
      },
      calculateWorkersHours() {
        let sum = 0.0;
        for (let index = 0; index < this.newItem.task_workers.length; index++) {
          let tw = this.newItem.task_workers[index];
          //let diff = this.getTotalHours(tw);
          //sum = sum + diff;
          sum = sum + tw.total_hours;
        }
        return sum;
      },
      calculateWorkersPrice() {
        var sum = 0.0;
        for (let index = 0; index < this.newItem.task_workers.length; index++) {
          let tw = this.newItem.task_workers[index];
          sum = sum + tw.total_price;
        }
        return sum;
      },
      calculateWorkersTransfer() {
        var sum = 0.0;
        for (let index = 0; index < this.newItem.task_workers.length; index++) {
          let tw = this.newItem.task_workers[index];
          sum = sum + tw.road_price;
        }
        return sum;
      },
      calculateWorkersPrepaid() {
        var sum = 0.0;
        for (let index = 0; index < this.newItem.task_workers.length; index++) {
          let tw = this.newItem.task_workers[index];
          sum = sum + tw.prepaid;
        }
        return sum;
      },

      //TW
      getWorkerById: function (id, item) {
        let url = this.formUrlTW + "/id/" + id;
        let self = this;
        let st = item.start_date;
        let et = item.end_date;
        return this.$http.get(url).then(result => {
          console.log(result);
          this.isBusy = false;
          if (result.status === 403) {
            this.showForbiddenAlert = true;
          } else if (result.status === 200 || result.status === 304) {
            this.showForbiddenAlert = false;
            if (result.body != null) {
              let worker = result.body;
              console.log("Worker");
              console.log(worker);

              let taskWorker = {
                worker_id: item.worker_id,
                name: worker.name,
                phone: worker.phone,
                price: item.price,
                road_price: item.road_price,
                start_time: this.formatTime(new Date(st)),
                end_time: this.formatTime(new Date(et)),
                total_price: item.total_price,
                prepaid: item.prepaid,
                dinner: item.dinner,
                description: item.description,
                start_date: this.formatServerDateTime(this.formatDate(new Date(st)), this.formatTime(new Date(st))),
                end_date: this.formatServerDateTime(this.formatDate(new Date(et)), this.formatTime(new Date(et))),
                string_start_date: this.formatDate(new Date(st)),
                string_end_date: this.formatDate(new Date(et)),
                paid: item.paid,
                worker:worker
              };

              taskWorker.total_hours = this.getTotalHours(taskWorker);
              self.newItem.task_workers.push(taskWorker)

            } else {
              console.log("body null");
            }
          }
        }, error => {
          this.isBusy = false;
          console.log("ERROR", error);
        });
      },
      onFiltered(filteredItems) {
        this.totalRows = filteredItems.length;
        this.currentPage = 1
      },
      imageUrl(newItem) {
        return this.serverUrl + '/' + newItem.passport_url
      },
      telUrl(newItem) {
        return "tel:" + newItem.phone
      },
      existPassport(value) {
        return value.length > 0
      },

      disabledChangeStartDate(newValues) {
        if (newValues.length > 0) {
          this.disabledStartDate = true
        } else {
          this.disabledStartDate = false
        }
      },
      workerCheckBoxId(row) {
        return "checkbox_paid_" + row.item.worker_id;
      },
      userInfo() {
        this.$http.get("users/info").then(result => {
          if (result.status === 200) {
            this.user = result.body
          }
        }, error => {
          console.log(error);
        });
      },
      authorInfo(userId) {
        this.$http.get("users/id/" + userId).then(result => {
          console.log(result);
          if (result.status === 200) {
            this.author = result.body
          }
        }, error => {
          console.log(error);
        });
      },
      executorInfo(userId) {
        this.$http.get("users/id/" + userId).then(result => {
          console.log(result);
          if (result.status === 200) {
            this.executor = result.body;
            this.updateTotalDelta();
          }
        }, error => {
          console.log(error);
        });
      },
      updateView: function (town) {
        this.formUrl = "town/" + town + "/tasks";
        this.formUrlTW = "town/" + town + "/workers";
        if (town === 0) {
          this.background = "info";
          this.title = "Заявки Москва";
          this.titleModel = "Заявки Москва"
        }
        else if (town === 1) {
          this.background = "warning";
          this.title = "Заявки СПБ";
          this.titleModel = "Заявки СПб";
        }
      },
      showWorkerInfo(row, isRead) {
        this.$refs.worker.showModal({item:row}, isRead)
      },
      acceptItem(row) {
        this.acceptWorker.paid = !this.acceptWorker.paid;
        this.newItem.task_workers.find(x => x.worker_id === this.acceptWorker.worker_id).paid = this.acceptWorker.paid;

        this.hideModalAccept()
      },
      showModalAccept(row) {
        //console.log(this.$store.getters.isManager);
        if (this.$store.getters.isManager){
          if (!row.item.paid){
            if (row.item) {
              this.acceptWorker = row.item
            }
            this.$refs.modalAccept.show();
          }
        }else {
          this.acceptWorker = row.item;
          this.acceptWorker.paid = !this.acceptWorker.paid;
          this.newItem.task_workers.find(x => x.worker_id === this.acceptWorker.worker_id).paid = this.acceptWorker.paid;
        }

      },
      hideModalAccept() {
        this.$refs.modalAccept.hide()
      },
      getPaidVariant(row){
        if (row.item.paid){
          return 'success'
        }else {
          return 'danger'
        }
      },
      getPaidTitle(row){
        if (row.item.paid){
          return 'Оплачено'
        }else {
          return 'Не оплачено'
        }
      }

    },
    watch: {
      'newItem.payment_first': {
        handler: function (newVal, oldVal) {
          if (newVal !== oldVal) {
            this.updateTotals()
          }
        },
        deep: true
      },
      'newItem.payment_first_count': {
        handler: function (newVal, oldVal) {
          if (newVal !== oldVal) {
            this.updateTotals()
          }
        },
        deep: true
      },
      'newItem.payment_second': {
        handler: function (newVal, oldVal) {
          if (newVal !== oldVal) {
            this.updateTotals()
          }
        },
        deep: true
      },
      'newItem.payment_second_count': {
        handler: function (newVal, oldVal) {
          if (newVal !== oldVal) {
            this.updateTotals()
          }
        },
        deep: true
      },
      'newItem.total_hours': {
        handler: function (newVal, oldVal) {
          if (newVal !== oldVal) {
            this.updateTotals()
          }
        },
        deep: true
      },
      'newItem.additional_cost': {
        handler: function (newVal, oldVal) {
          this.updateTotals()
        },
        deep: true
      },
      'newItem.additional_income': {
        handler: function (newVal, oldVal) {
          this.updateTotals()
        },
        deep: true
      },
      'newItem.declarated_cost': {
        handler: function (newVal, oldVal) {
          this.updateTotalDelta()
        },
        deep: true
      },
      'newItem.workers_count': {
        handler: function (newVal, oldVal) {
          if (newVal !== oldVal) {
            this.updateTotals()
          }
        },
        deep: true
      },
      'newItem.total_worker_price': {
        handler: function (newVal, oldVal) {
          this.updateTotalDelta()
        },
        deep: true
      },
      'newItem.total_price': {
        handler: function (newVal, oldVal) {
          this.updateTotalDelta()
        },
        deep: true
      },
      'newItem.payment_type': {
        handler: function (newVal, oldVal) {
          this.updateTotalDelta()
        },
        deep: true
      },
      'startDate': {
        handler: function (newVal, oldVal) {
          this.updateTotals()
        },
        deep: true
      },
      'startTime': {
        handler: function (newVal, oldVal) {
          this.updateTotals()
        },
        deep: true
      },
      'endDate': {
        handler: function (newVal, oldVal) {
          this.updateTotals()
        },
        deep: true
      },
      'endTime': {
        handler: function (newVal, oldVal) {
          this.updateTotals()
        },
        deep: true
      },
      'user': {
        handler: function (newVal, oldVal) {
          if (this.newItem.author === 0) {
            this.newItem.author = this.user.id
          }
          if (this.newItem.executor === 0) {
            this.newItem.executor = this.user.id
          }
        },
        deep: true
      },
      'newItem.author': {
        handler: function (newVal, oldVal) {
          if (newVal > 0) {
            this.authorInfo(newVal);
          }

        },
        deep: true
      },
      'newItem.executor': {
        handler: function (newVal, oldVal) {
          if (newVal > 0) {
            this.executorInfo(newVal);
          }
        },
        deep: true
      },
      'newItem.paid': {
        handler: function (newVal, oldVal) {
          // if(newVal === true){
          //   this.newItem.status_id = 6;
          // }else {
          //   this.newItem.status_id = 5;
          // }
        },
        deep: true
      },

      //Task Workers Watch
      'newItem.task_workers': {
        handler: function (newValues, oldValues) {
          //if (newValues !== oldValues) {
          for (let index = 0; index < newValues.length; index++) {
            let tw = newValues[index];
            if (tw.start_time != null) {
              if (tw.end_time != null) {
                tw.start_date = this.formatServerDateTime(tw.string_start_date, tw.start_time);
                tw.end_date = this.formatServerDateTime(tw.string_end_date, tw.end_time);
                let diff = this.getTotalHours(tw);
                if (diff < 4) {
                  tw.total_hours = 4;
                } else {
                  tw.total_hours = diff;
                }
                tw.total_price = (tw.total_hours * tw.price) + tw.road_price - tw.prepaid;
              }
            }
          }
          this.updateTotals();

          this.disabledChangeStartDate(newValues)
          //}
        },
        deep: true
      },
      $route: function (newR, oldR) {
        this.updateView(newR.meta.town);
      }
    }
  }
</script>

<style>
  .topLabelColor {
    color: white;
    font-size: 20px;
  }

  @media screen
  and (min-width: 768px) {
    a[href*="tel:"] {
      pointer-events: none;
      cursor: default;
      text-decoration: none;
    }
  }

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
  }

  /*TW*/
  .widthPrice {
    width: 70px;
  }

  .widthTime {
    width: 105px;
  }

  .widthDate {
    width: 170px;
  }

  .width50 {
    width: 65px;
  }
</style>

import Vue from 'vue'
import Router from 'vue-router'
//import HelloWorld from '@/components/HelloWorld'
import Workers from '@/components/Workers'
import Login from '@/components/Login'
import Forbidden from '@/components/Forbidden'
import Tasks from '@/components/Tasks'
import Task from '@/components/Task'
import Customers from '@/components/Customers/Customers'
import Payments from '@/components/Payments/Payments'
import PaymentDetail from '@/components/Payments/PaymentDetail'
import ManagerStatistics from '@/components/Statistics/ManagerStatistic'
import Managers from '@/components/Managers'

Vue.use(Router);

const routes = [
    {
      path: '/',
      name: 'Workers',
      component: Workers,
      meta: {town:0, requiresAuth: true , admin:true , manager : true, accountant : false}
    },
    {
      path: '/forbidden',
      name: 'Forbidden',
      component: Forbidden
    },
    {
      path: '/login',
      name: 'Login',
      component: Login
    },
    {
      path: '/spb/workers',
      name: 'WorkersSpb',
      component: Workers,
      meta: {town:1, requiresAuth: true , admin:true , manager : true, accountant : false}
    },
    {
      path: '/tasks',
      name: 'Tasks',
      component: Tasks,
      meta: {town:0,requiresAuth: true , admin:true , manager : true, accountant : false}
    },
    {
      path: '/spb/tasks',
      name: 'TasksSpb',
      component: Tasks,
      meta: {town:1,requiresAuth: true , admin:true , manager : true, accountant : false}
    },
    {
      path: '/tasks/new',
      name: 'Task',
      component: Task,
      meta: {town:0,breadcrumbs:[{'text':'Заявки', to: { name: 'Tasks' }},{'text':'Заявка', to: { name: 'Task' }}],requiresAuth: true , admin:true , manager : true, accountant : false},
      props: true
    },
    {
      path: '/tasks/edit',
      name: 'TaskEdit',
      component: Task,
      meta: {town:0,requiresAuth: true , admin:true , manager : true, accountant : false},
      props: true
    },
  {
    path: '/spb/tasks/new',
    name: 'TaskSpb',
    component: Task,
    meta: {town:1,breadcrumbs:[{'text':'Заявки', to: { name: 'Tasks' }},{'text':'Заявка', to: { name: 'Task' }}],requiresAuth: true , admin:true , manager : true, accountant : false},
    props: true
  },
  {
    path: '/spb/tasks/edit',
    name: 'TaskEditSpb',
    component: Task,
    meta: {town:1,requiresAuth: true , admin:true , manager : true, accountant : false},
    props: true
  },
    {
      path: '/customers',
      name: 'Customers',
      component: Customers,
      meta: {town:0, requiresAuth: true , admin:true , manager : true, accountant : false}
    },
    {
      path: '/spb/customers',
      name: 'CustomersSpb',
      component: Customers,
      meta: {town:1, requiresAuth: true , admin:true , manager : true, accountant : false}
    },
    {
      path: '/payments',
      name: 'Payments',
      component: Payments,
      meta: {requiresAuth: true , admin:true , manager : true, accountant : false}
    },
    {
      path: '/payments/detail',
      name: 'PaymentDetail',
      component: PaymentDetail,
      meta: {requiresAuth: true , admin:true , manager : true, accountant : false}
    },
    {
      path: '/statistics/manager',
      name: 'ManagerStatistics',
      component: ManagerStatistics,
      meta: {requiresAuth: true , admin:true , manager : false, accountant : false}
    },
    {
      path: '/managers',
      name: 'Managers',
      component: Managers,
      meta: {requiresAuth: true , admin:true , manager : false, accountant : false}
    }
  ];

const router = new Router({routes,mode:'history'});

router.beforeEach((to, from, next) => {
  if(to.meta.requiresAuth) {
    const role = localStorage.getItem('role');
    const token = localStorage.getItem('token');

    if(!token || !role) {
      next({name:'Login'})
      return
    }

    if(to.meta.admin) {
      if(role === 'admin') {
        next();
        return
      }
    }
    if(to.meta.manager) {

      if(role === 'manager') {
        next();
        return
      }
    }
    if(to.meta.security) {
      if(role === 'accountant') {
        next();
        return
      }
    }

    next({name:'Forbidden'})
  }else {
    next()
  }
});

export default router;

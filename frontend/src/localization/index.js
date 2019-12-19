import Vue from 'vue'
import VueI18n from 'vue-i18n';
Vue.use(VueI18n);

import VueCookie from 'vue-cookie'
Vue.use(VueCookie)

const messages = {
  ru: {
    message: {
      actions:{
        remove:"Удалить",
        cancel:"Отмена",
        download:"Скачать",
        upload:"Загрузить",
        export:"Выгрузить",
        corrected:"Исправлено",
        ok:"Да",
      },
      common:{
        emptyTable:"В таблице нет записей",
        search:"Поиск",
        searchPlaceholder:"Начните набор",
        searchClear:"Очистить",
      },
      login:{
        title:'Добро пожаловать!',
        login:"Логин",
        password:"Пароль",
        enterLogin:"Введите логин",
        enterPassword:"Введите пароль",
        signIn:"Войти",
        logout:"Выйти"
      },
      workers:{
        title:'Рабочие',
        description:'',
        importExcel:'Импорт работников из EXCEL',
        importHand:'Добавить работника',
        table:{
          active_today:"Занят",
          name:"Имя",
          phone:"Телефон",
          rating:" Рейтинг",
          age:"Возраст",
          passport:"Паспорт",
          address:"Адрес",
          skill:" Умение",
          ready:" Готовность",
          description:" Примечание",
          actions:"Действия",
          view:"Просмотр",
          edit:"Ред.",
          remove:"Удалить",
          add_in_task:"В заявку",
          price:"Ставка",
          start_date:"Дата начала",
          end_date:"Дата конца",
          start_time:"Время начала",
          end_time:"Время конца",
          road_price:"Трансфер",
          prepaid:"Аванс",
          dinner:"Обед",
          total_price_payment:"Осталось оплатить",
          total_price:"Итого",
          total_hours:"Часы",
          task_id:"Задача",
          last_task_end_date:"Конец последней заявки",
          paid:"Оплачено",
          cart_number:"Карта"

        },
        modal:{
          title:"Работник",
          nameInput:"Имя:",
          namePlaceholder:"Например Иванов Иван",
          phoneInput:"Телефон:",
          phonePlaceholder:"Без пробелов и +7 Например 91112223344",
          phoneError:"Работник с таким Телефоном уже есть в системе",
          ageInput:"Возраст:",
          agePlaceholder:"Например 20",
          ageError:"Возраст указан неверно минимум 14, максимум 100",
          addressInput:"Адрес:",
          addressPlaceholder:"Район, метро, адрес если есть",
          passport_typeInput:"Тип паспорта:",
          passport_typePlaceholder:"Например РФ",
          passport_urlInput:"Паспорт:",
          passport_urlPlaceholder:"Добавьте документ",
          cart_numberInput:"Номер карты:",
          cart_numberPlaceholder:"Например 4276 0022 2121 3211",
          skillInput:"Умения",
          skillPlaceholder:"Например: Шуруповерт",
          readyInput:"Готовнсть",
          readyPlaceholder:"Например: работает по выходным",
          descriptionInput:"Примечание:",
          descriptionPlaceholder:"Введите нужную информацию которой нет в полях выше",
          removeTitle:"Удаление работника",
          removeDesc:"Вы уверены, что хотите удалить работника?",
          importTitle:"Импорт товаров",
          template:"Шаблон EXCEL",
          importDesc:"Скачать шаблон можно по ссылке",
        }
      },
      tasks:{
        title:'Заявки',
        description:'',
        importExcel:'Импорт заявок из EXCEL',
        importHand:'Добавить заявку',
        table:{
          name:"Заявка",
          customer:"Заказчик",
          start_address:"Начальный адрес",
          end_address:" Конечный адрес",
          start_date:"Дата начала",
          end_date:"Дата окончания",
          payment_type:"Тип оплаты",
          payment_first:"Стоимость часа",
          payment_second:"Стоимость последующего часа",
          workers_count:"Кол-во работников",
          total_hours:"Итого часов",
          total_price:"Итого с заказчика",
          total_worker_price:"Итого работникам",
          total_delta:"Дельта",
          additional_cost:"Доп. расходы",
          additional_income:" Доп. доходы",
          description:"Примечание",
          cargo_description:"Описание груза",
          status_name:"Статус",
          paid_workers:"Оплачено грузчикам",
          executor_name:"Исполнитель",
          paid:"Оплачено",
          contact:"Контакт",
          declarated_cost:"Расходы",
          actions:"Действия",
          view:"Просмотр",
          edit:"Ред.",
          remove:"Удалить",
        },
        modal:{
          title:"Работник",
          customer:"Заказчик",
          customerPlaceholder:"Выберите заказчика",
          start_address:"Начальный адрес",
          start_addressPlaceholder:"Введите начальный адрес",
          end_address:" Конечный адрес",
          end_addressPlaceholder:"Введите конечный адрес",
          start_date:"Дата начала",
          start_datePlaceholder:"Введите дату начала",
          start_time:"Время начала",
          end_date:"Дата окончания",
          end_datePlaceholder:"Введите дату окончания",
          end_time:"Время окончания",
          payment_type:"Тип оплаты",
          payment_typePlaceholder:"Введите тип оплаты",
          payment_first_count:"Часы",
          payment_first:"Стоимость",
          payment_firstPlaceholder:"Стоимость",
          payment_second_count:"Часы 2",
          payment_second:"Стоимость 2",
          payment_secondPlaceholder:"Стоимость 2",
          workers_count:"Количество работников",
          workers_countPlaceholder:"Введите кол-во работников",
          total_hours:"Итого часов",
          total_hoursPlaceholder:"Введите итого часов",
          total_price:"Итого с заказчика",
          total_pricePlaceholder:"Введите итого заказчикам",
          total_worker_price:" Итого работникам",
          total_worker_pricePlaceholder:"Введите итого работникам",
          total_delta:"Дельта",
          total_deltaPlaceholder:"Введите дельта",
          additional_cost:" Доп. расходы",
          additional_income:" Доп. доходы",
          additional_costPlaceholder:"Введите доп расходы",
          description:"Примечание",
          descriptionPlaceholder:"Введите примечание",
          cargo_description:"Описание груза",
          cargo_descriptionPlaceholder:"Введите описание груза",
          status_id:"Статус",
          status_idPlaceholder:"Введите статус",
          total_manager:"Менеджеру на руки",
          total_managerPlaceholder:"Менеджеру на руки",
          removeTitle:"Удаление заявку",
          removeDesc:"Вы уверены, что хотите удалить заявку",
          author:"Автор",
          executor:"Исполнитель",
          paid:"Оплачено",
          contact:"Контакт",
          contactPlaceholder:"Контактное лицо на объекте",
          declarated_cost:"Расходы",
          acceptWorkerTitle:"Подтвержление оплаты работнику",
          acceptWorkerDesc:"Отключить оплату будет невозможно. Применить для ",
        }
      },
      customers:{
        title:'Заказчики',
        description:'',
        importHand:'Добавить заказчика',
        table:{
          name:"Имя/Название",
          phone:"Телефон",
          address:"Адрес",
          description:" Примечание",
          urflag:"Тип",
          actions:"Действия",
          view:"Просмотр",
          edit:"Ред.",
          remove:"Удалить",
          email:"Email",
        },
        modal:{
          title:"Заказчик",
          nameInput:"Имя:",
          namePlaceholder:"Например, Иванов Иван",
          phoneInput:"Телефон:",
          phonePlaceholder:"Без пробелов и +7. Например, 91112223344",
          phoneError:"Работник с таким Телефоном уже есть в системе",
          addressInput:"Адрес:",
          addressPlaceholder:"Район, метро, адрес, если есть",
          descriptionInput:"Примечание:",
          descriptionPlaceholder:"Введите нужную информацию, которой нет в полях выше",
          urflagInput:"Тип",
          rateInput:"Ставка менеджеру, %",
          saleInput:"Скидка за час, руб.",
          removeTitle:"Удаление заказчика",
          removeDesc:"Вы уверены, что хотите удалить заказчика?",
          email:"Email:",
          company_name:"Название компании:",
          emailPlaceholder: "ex: smth@smth.com",
          company_namePlaceholder: "Полное название компании",
          referee: "Рефери:",
          refereePlaceholder: "Название фирмы, сайта, откуда клиент узнал о нас",
        }
      },
      payments:{
        title:'Выплаты',
        description:'',
        find:'Найти',
        date:'Дата',
        startDate:'С',
        endDate:'По',
        paid:'Подтверждение действия',
        paidDesc:'Подтвердить изменение оплаты для',
        table:{
          name:"Имя/Название",
          phone:"Телефон",
          address:"Адрес",
          description:" Примечание",
          urflag:"Тип",
          actions:"Действия",
          view:"Просмотр",
          edit:"Ред.",
          remove:"Удалить",
        },
      },
      statistics:{
        title:'Статистика',
        description:'',
        find:'Показать',
        date:'Дата',
        startDate:'С',
        endDate:'По',
      },
      users:{
        title:'Персональные предложения',
        description:'Список персональных предложений для каждого покупателя',
        personal:{
          title:"Персональные предложения",
          description:'',
          conditionTitle:"Условие:",
          conditionDesc:"Количество покупок > 0 и Акция на товар > 0",
          customer:"Покупатель:",
          login:"Логин:",
          email:"Email:",
          table:{
            name:"Наименование",
            barcode:"Штрихкод",
            price:"Цена",
            sale:"Акция",
            count:"Количество покупок",
          },
        },
        table:{
          login:"Логин",
          name:"Имя",
          email:"Email",
          actions:"Действия",
        },
      },
      orders:{
        title:'Транзации',
        description:'Список успешных транзакция',
        exportTitle:"Экспорт данных о покупках",
        exportFileMessage:"Файл можно скачать по ссылке",
        exportExcel:"Выгрузка Excel",
        table:{
          human_date:"Дата создания",
          sum:"Сумма покупки",
          username:"Логин покупателя",
          uuid:"Id транзации",
        },
      },
      oos:{
        title:'Out Of Shelf',
        description:'Список отсутствующих товаров на полке по мнениям пользователей',
        table:{
          human_date:"Дата создания",
          name:"Наименование",
          barcode:"Штрихкод",
          username:"Логин покупателя",
          actions:"Действия",
        },
        removeTitle:"Удаление записи",
        removeDesc:"Вы уверены, что хотите удалить запись",
      },
      managers: {
        title: "Менеджеры",
        see: "Посмотреть",
        edit: "Редактировать",
        delete: "Удалить",
        add: "Добавить менеджера",
        admin: "Администратор",
        manager: "Менеджер",
        accountant: "Бухгалтер",
        table: {
          username: "Логин",
          password: "Пароль",
          role: "Должность",
          name: "Имя",
          email: "Почта",
          rate: "Ставка",
          rate_tax: "Такса",
          actions: "Действия",
        },
        modal: {
          username: "Логин:",
          password: "Пароль:",
          name: "Имя:",
          email: "Почта:",
          rate: "Ставка:",
          rate_tax: "Такса:",
          usernamePlaceholder: "Введите логин",
          passwordPlaceholder: "Введите пароль",
          rolePlaceholder: "Администратор, менеджер, и т.п.",
          namePlaceholder: "Фамилия и имя",
          emailPlaceholder: "smth@smth.com",
          ratePlaceholder: "",
          rate_taxPlaceholder: "",
          errorWithEmail: "Пользователь с таким адресом почты уже существует!",
          removeTitle:"Удаление менеджера",
          removeDesc:"Вы уверены, что хотите удалить менеджера?",
          editTitle: "Изменение данных менеджера",
          addTitle: "Добавление нового менеджера",
          role:"Должность",
        }
      }
    }
  }
};

Vue.config.lang = VueCookie.get('locale') || 'ru';

const i18n = new VueI18n({
  locale:  Vue.config.lang, // set locale
  messages, // set locale messages
});

export default i18n;
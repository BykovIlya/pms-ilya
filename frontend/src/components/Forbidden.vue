<template>
  <b-container>
    <b-alert variant="danger" center show>
      Доступ запрещён. Обратитесь к администратору.
    </b-alert>
  </b-container>
</template>

<script>

export default {
  name: 'Forbidden',
  data () {
    return {
      msg: 'Forbidden'
    }
  }
}
</script>

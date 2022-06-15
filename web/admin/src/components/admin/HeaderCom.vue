<template>
  <a-button type="danger" @click="logout">Logout</a-button>
</template>

<script>
export default {
  methods: {
    logout() {
      window.sessionStorage.clear('token')
      this.$router.push('/login')
    }
  }
}
</script>

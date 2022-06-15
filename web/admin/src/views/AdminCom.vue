<template>
  <a-layout class="container">
    <a-layout-sider>
      <div class="logo">
        <span>My Logo</span>
      </div>
      <NavCom></NavCom>
    </a-layout-sider>
    <a-layout>
      <a-layout-header class="headerBtn">
        <a-button type="danger" @click="logout">Logout</a-button>
      </a-layout-header>
      <a-layout-content>
        <router-view></router-view>
      </a-layout-content>
      <a-layout-footer>
        <FooterCom></FooterCom>
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script>
import NavCom from '../components/admin/NavCom.vue'
import FooterCom from '../components/admin/FooterCom.vue'
export default {
  components: { NavCom, FooterCom },
  methods: {
    logout() {
      window.sessionStorage.clear('token')
      this.$router.push('/login')
    }
  }
}
</script>

<style scoped>
.container {
  height: 100%;
}
.headerBtn {
  display: flex;
  justify-content: right;
  align-items: center;
}
.logo {
  height: 32px;
  margin: 16px;
  background-color: white;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 17px;
}
</style>

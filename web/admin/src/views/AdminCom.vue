<template>
  <a-layout class="container">
    <NavCom></NavCom>
    <a-layout>
      <a-layout-header class="headerBtn">
        <HeaderCom></HeaderCom>
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
import HeaderCom from '@/components/admin/HeaderCom.vue'

export default {
  components: { NavCom, FooterCom, HeaderCom }
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
</style>

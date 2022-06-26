<template>
  <a-layout-sider breakpoint="lg" v-model="collapsed">
    <div class="logo">
      <span>{{ collapsed ? 'Blog' : 'My Blog' }}</span>
    </div>
    <a-menu theme="dark" mode="inline" @click="goToPage">
      <a-menu-item key="index">
        <a-icon type="dashboard"> </a-icon>
        <span>Dashboard</span>
      </a-menu-item>
      <a-sub-menu>
        <span slot="title"
          ><a-icon type="file-text"></a-icon
          ><span>Article Management</span></span
        >
        <a-menu-item key="addart"
          ><a-icon type="edit"></a-icon><span>Write Article</span></a-menu-item
        >
        <a-menu-item key="artlist"
          ><a-icon type="snippets"></a-icon
          ><span>Article List</span></a-menu-item
        >
      </a-sub-menu>
      <a-menu-item key="catelist"
        ><a-icon type="tags"></a-icon><span>Category List</span></a-menu-item
      >
      <a-menu-item key="userlist"
        ><a-icon type="user"></a-icon><span>User List</span></a-menu-item
      >
    </a-menu>
  </a-layout-sider>
</template>

<script>
export default {
  data() {
    return {
      collapsed: false
    }
  },
  methods: {
    goToPage(item) {
      this.$router.push('/admin/' + item.key).catch((err) => err)
    }
  }
}
</script>

<style>
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

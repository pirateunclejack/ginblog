<template>
  <div>
    <h3>User List</h3>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search placeholder="input username to search" enter-button />
        </a-col>
        <a-col :span="4">
          <a-button type="primary">Add</a-button>
        </a-col>
      </a-row>
      <a-table
        rowKey="username"
        :columns="columns"
        :pagination="paginationOption"
        :dataSource="userlist"
        bordered
      >
        <span slot="role" slot-scope="role">{{
          role === 1 ? 'Administrator' : 'Register'
        }}</span>
        <template slot="action">
          <div class="actionSlot">
            <a-button type="primary" style="margin-right: 15px">Edit</a-button>
            <a-button type="danger">Delete</a-button>
          </div>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '10%',
    key: 'id',
    align: 'center'
  },
  {
    title: 'Username',
    dataIndex: 'username',
    width: '20%',
    key: 'username',
    align: 'center'
  },
  {
    title: 'Role',
    dataIndex: 'role',
    width: '20%',
    key: 'role',
    scopedSlots: { customRender: 'role' },
    align: 'center'
  },
  {
    title: 'Action',
    dataIndex: 'action',
    width: '30%',
    key: 'action',
    scopedSlots: { customRender: 'action' },
    align: 'center'
  }
]
export default {
  data() {
    return {
      paginationOption: {
        pageSizeOptions: ['2', '5', '10'],
        defaultCurrent: 1,
        defaultPageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `Total ${total}`,
        onChange: (current, pageSize) => {
          this.paginationOption.defaultCurrent = current
          this.paginationOption.defaultPageSize = pageSize
          this.getUserList()
        },
        onshowSizeChange: (current, size) => {
          this.paginationOption.defaultCurrent = current
          this.paginationOption.defaultPageSize = size
          this.getUserList()
        }
      },
      userlist: [],
      columns
    }
  },
  created() {
    this.getUserList()
  },
  methods: {
    async getUserList() {
      const { data: res } = await this.$http.get('users', {
        params: {
          // username: this.queryParam.username,
          // pagesize: this.queryParam.pagesize,
          // pagenum: this.queryParam.pagenum
          pagesize: this.paginationOption.defaultPageSize,
          pagenum: this.paginationOption.defaultCurrent
        }
      })
      if (res.status !== 200) return this.$message.error(res.message)
      this.userlist = res.data
      this.paginationOption.total = res.total
    }
  }
}
</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}
</style>

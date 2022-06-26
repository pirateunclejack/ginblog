<template>
  <div>
    <h3>User List</h3>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search
            v-model="queryParam.username"
            placeholder="input username to search"
            enter-button
            allowClear
            @search="getUserList"
          />
        </a-col>
        <a-col :span="4">
          <a-button type="primary">Add</a-button>
        </a-col>
      </a-row>
      <a-table
        rowKey="username"
        :columns="columns"
        :pagination="pagination"
        :dataSource="userlist"
        bordered
        @change="handleTableChange"
      >
        <span slot="role" slot-scope="role">{{
          role === 1 ? 'Administrator' : 'Subscriber'
        }}</span>
        <template slot="action" slot-scope="text, data">
          <div class="actionSlot">
            <a-button type="primary" style="margin-right: 15px">Edit</a-button>
            <a-button type="danger" @click="deleteUser(data.ID)"
              >Delete</a-button
            >
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
      pagination: {
        pageSizeOptions: ['2', '4', '8'],
        pageSize: 2,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `Total ${total}`
      },
      userlist: [],
      columns,
      queryParam: {
        username: '',
        pagesize: 2,
        pagenum: 1
      },
      visible: false
    }
  },
  created() {
    this.getUserList()
  },
  methods: {
    async getUserList() {
      const { data: res } = await this.$http.get('admin/users', {
        params: {
          username: this.queryParam.username,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      this.userlist = res.data
      this.pagination.total = res.total
    },
    handleTableChange(pagination, filters, sorter) {
      const pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.queryParam.pagesize = pagination.pageSize
      this.queryParam.pagenum = pagination.current

      if (pagination.pageSize !== this.pagination.pageSize) {
        this.queryParam.pagenum = 1
        pager.current = 1
      }
      this.pagination = pager
      this.getUserList()
    },
    deleteUser(id) {
      this.$confirm({
        title: 'Do you want to delete this user?',
        content:
          'Are you sure you want to delete this user? Once it is deleted , it will be impossible to recover it!',
        onOk: async () => {
          const { data: res } = await this.$http.delete(`user/${id}`)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$message.success('Delete success')
          this.getUserList()
        },
        onCancel: () => {
          this.$message.info('Cancel delete')
        }
      })
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

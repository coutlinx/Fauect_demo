// 登录组件
const Login = {
  data() {
    return {
      todoTitle: '', // 待办标题
      todos: []
    }
  },
  template: `
    <div class="login">
      <el-row>
        <el-input v-model="todoTitle" placeholder="请输入内容"></el-input>
        <el-button @click="handleAddTodo">新增待办</el-button>
      </el-row>


      <el-divider content-position="left">待办列表</el-divider>

      <el-row v-for="item in todos" :key="item.title">
        <el-checkbox @change="finishAddTodo(item)">{{item.title}}</el-checkbox>
      </el-row>
    </div>
  `,
  mounted: function () {
    this.queryTodo() // 加载时，查询所有待办项目
  },
  methods: {
    // 查询待办列表
    queryTodo() {
      // 发起请求
      axios({
        method: 'get',
        url: `/todo`
      })
        .then(ret => {
          console.log(JSON.stringify(ret))

          if (ret.data.code == 0) {
            console.log("ok")
            this.todos = ret.data.data
          } else {
            console.log("msg" + ret.data.message)
            this.$message(ret.data.message)
          }
        })
        .catch(err => {
          console.log(err)
          this.$message(err)
        })
    },

    // 增加待办
    handleAddTodo() {
      // 检查参数
      if (this.todoTitle === '') {
        this.$message('请输入待办标题')
        return
      }

      // 发起请求
      axios({
        method: 'post',
        url: `/todo`,
        data: {title: this.todoTitle}
      })
        .then(ret => {
          console.log(JSON.stringify(ret))

          if (ret.data.code == 0) {
            console.log("ok")
            this.$message('添加成功')
            this.queryTodo()
          } else {
            console.log("msg" + ret.data.message)
            this.$message(ret.data.message)
          }
        })
        .catch(err => {
          console.log(err)
          this.$message(err)
        })

      this.todoTitle = ''
    },

    // 完成待办
    finishAddTodo(item) {
      console.log("finishAddTodo:" + JSON.stringify(item))

      // 反转状态
      item.state = !item.state

      let todo = {}
      todo.id = item.id
      todo.title = item.title
      todo.state = item.state

      // 发起请求
      axios({
        method: 'put',
        url: `/todo`,
        data: todo
      })
        .then(ret => {
          console.log(JSON.stringify(ret))

          if (ret.data.code == 0) {
            this.queryTodo()
            console.log("ok")
          } else {
            console.log("msg" + ret.data.message)
            this.$message(ret.data.message)
          }
        })
        .catch(err => {
          console.log(err)
          this.$message(err)
        })
    },
  }
}
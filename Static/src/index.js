const vue = new Vue({
  // 注册 Header、Login 子组件
  components: {
    Header,
    Login,
  },
  // 实例绑定的元素
  el: '#app',
  data: {
    login: false, // 登录状态
    user: '', // 当前用户
  },
  template: `
    <el-container>

      <!-- 头部导航组件 -->
      <el-header class="header"><Header :user="user" :login="login" @logout="(val) => {this.login = val;}" /></el-header>

      <!-- 主要内容组件 -->
      <el-container>
        <el-main>

          <!-- 登录组件 -->
          <Login v-if="!login" @login="(val) => {}" />

        </el-main>
      </el-container>
    </el-container>
    `,
  methods: {
    
  }
})
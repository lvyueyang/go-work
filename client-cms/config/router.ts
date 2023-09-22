const router = {
  routes: [
    { path: '/login', component: 'login' },
    { path: '/nopassword', component: 'nopassword' },
    { path: '/init-root-user', component: 'initRootUser' },
    { path: '/', redirect: '/admin-user/list' },
    {
      path: '/',
      component: '@/layouts/main',
      meta: {
        isMenuRoot: true,
      },
      routes: [
        {
          path: '/userinfo',
          component: 'userinfo',
          title: '用户信息',
          menuHide: true,
        },
        {
          path: '/admin-user',
          title: '后台账户管理',
          // icon: UserOutlined,
          routes: [
            {
              path: '/admin-user/list',
              component: 'adminUser',
              title: '用户列表',
            },
            {
              path: '/admin-user/role',
              component: 'adminRole',
              title: '角色管理',
            },
          ],
        },
        // {
        //   path: '/logger',
        //   component: 'logger',
        //   title: '系统日志',
        // },

        {
          path: '/news',
          title: '新闻管理',
          routes: [
            {
              path: '/news/list',
              title: '新闻列表',
              component: 'news',
            },
            {
              path: '/news/create',
              component: 'news/form',
              title: '新增新闻',
            },
            {
              path: '/news/update/:id',
              component: 'news/form',
              title: '修改新闻',
              menuHide: true,
            },
          ],
        },

        // {
        //   path: '/banner',
        //   title: '广告位管理',
        //   routes: [
        //     {
        //       path: '/banner/list',
        //       title: '广告列表',
        //       component: 'banner',
        //     },
        //     {
        //       path: '/banner/create',
        //       title: '新增广告',
        //       component: 'banner/form',
        //     },
        //     {
        //       path: '/banner/update/:id',
        //       component: 'banner/form',
        //       title: '修改广告',
        //       menuHide: true,
        //     },
        //   ],
        // },

        // {
        //   path: '/cominfo',
        //   title: '信息管理',
        //   routes: [
        //     {
        //       path: '/cominfo/info',
        //       component: 'cominfo/info',
        //       title: '公司介绍',
        //     },
        //   ],
        // },
      ],
    },
    { path: '*', component: '404' },
  ],
};
export default router;

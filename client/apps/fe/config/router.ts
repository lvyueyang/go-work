const router = {
  routes: [
    { path: '/login', component: 'login' },
    { path: '/home', component: 'home' },
    { path: '*', component: '404' },
  ],
};
export default router;

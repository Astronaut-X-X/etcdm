import {createApp} from 'vue'
import App from './app.vue'
import './style.css';
import Router from './router';

createApp(App)
.use(Router)
.mount('#app')

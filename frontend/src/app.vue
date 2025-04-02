<template>
  <n-config-provider :theme="theme">
    <n-layout has-sider position="absolute">
      <n-layout-sider :width="64" :bordered="true">
        <n-menu :collapsed="true" :collapsed-width="64" :collapsed-icon-size="22" :options="menuOptions"
          v-model:value="activeKey" @update:value="handleUpdateValue" />
      </n-layout-sider>
      <n-layout-content>
        <n-message-provider>
          <RouterView />
        </n-message-provider>
      </n-layout-content>
    </n-layout>
  </n-config-provider>
</template>

<script setup lang="ts">
import { NLayout, NLayoutSider, NLayoutContent, NMessageProvider, NConfigProvider } from 'naive-ui';
import { NMenu, type MenuOption } from 'naive-ui';
import { NIcon } from 'naive-ui';
import {
  AlbumsOutline as CubeIcon,
  ServerOutline as ServerIcon,
  SettingsOutline as SettingsIcon,
} from '@vicons/ionicons5';
import type { Component } from 'vue';
import { h, ref, computed, watch } from 'vue';
import { useRouter, useRoute } from 'vue-router'
import { getConfig } from './config';

// 从配置中获取主题设置并创建响应式引用
const theme = computed(() => {
  const config = getConfig();
  console.log('[config]', config);
  return config?.theme;
});

const menuOptions = ref<MenuOption[]>([
  {
    label: '数据',
    key: 'data',
    icon: renderIcon(CubeIcon)
  },
  {
    label: '连接',
    key: 'connection',
    icon: renderIcon(ServerIcon)
  },
  {
    label: '设置',
    key: 'setting',
    icon: renderIcon(SettingsIcon)
  },
]);

function renderIcon(icon: Component) {
  return () => h(NIcon, null, { default: () => h(icon) })
}

const router = useRouter();

const activeKey = ref('data');
function handleUpdateValue(key: string, item: MenuOption) {
  router.push({ path: "/" + key });
}
const route = useRoute();

// 监听路由变化，同步更新菜单选中状态
watch(
  () => route.path,
  (newPath) => {
    // 去除路径开头的斜杠，获取一级路由名称
    const routeName = newPath.split('/')[1] || 'data';
    activeKey.value = routeName;
  },
  { immediate: true }
);

</script>

<style></style>

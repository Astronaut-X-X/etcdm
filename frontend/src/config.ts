import { darkTheme, lightTheme } from 'naive-ui';
import type { BuiltInGlobalTheme } from 'naive-ui/es/themes/interface';
import { ref } from 'vue';

interface Config {
  theme?: BuiltInGlobalTheme;
}

// 使用 ref 创建响应式配置
const globalConfig = ref<Config>({
  theme: lightTheme,
});

export function getConfig(): Config {
  return globalConfig.value;
}

export function setConfig(config: Config): void {
  globalConfig.value.theme = config.theme;
}

<template>
  <n-layout has-sider position="absolute">
    <n-layout-content>
      <n-card title="Setting" :bordered="false" size="small" style="height: 100%;">
        <n-form ref="formRef" label-placement="left" label-width="auto" size="small">
          <n-form-item label="Theme" path="theme">
            <n-radio-group v-model:value="theme" name="radiogroup" @update:value="changeTheme">
              <n-space>
                <n-radio key="light" value="light">Light</n-radio>
                <n-radio key="light" value="dark">Dark</n-radio>
              </n-space>
            </n-radio-group>
          </n-form-item>
        </n-form>
      </n-card>
    </n-layout-content>
  </n-layout>
</template>

<script lang="ts" setup>
import { NLayout, NLayoutContent, NForm, NFormItem, NRadio, NRadioGroup, NSpace, NCard } from 'naive-ui';
import { darkTheme } from 'naive-ui';
import { ref } from 'vue';
import { getConfig, setConfig } from '../../config';

// region theme
const theme = ref('light');
function initTheme() {
  const config = getConfig();
  console.log(config.theme?.name);
  if (config.theme?.name == "dark") {
    theme.value = 'dark';
  }
}
initTheme();

function changeTheme(val: string) {
  if (val == 'light') {
    setConfig({ theme: undefined });
  } else {
    setConfig({ theme: darkTheme });
  }
}
// endregion


</script>

<style scoped></style>

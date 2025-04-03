<template>
  <n-layout has-sider position="absolute">
    <n-layout-sider bordered width="240">
      <n-card :bordered="false" size="small" title="Connections">
        <template #header-extra>
          <n-button size="tiny" @click="handleAdd">
            <n-icon>
              <AddIcon />
            </n-icon>
          </n-button>
        </template>

        <!-- <n-input-group style="margin-bottom: 20px;">
          <n-input size="small" />
          <n-button size="small">
            Srearch
          </n-button>
        </n-input-group> -->

        <n-flex vertical justify="center">
          <n-flex justify="space-between" v-for="item in list" :key="item.id" class="connction-line">
            <n-button text tag="a" size="small" @click="() => { loadItem(item.id) }">
              {{ item.name }}
            </n-button>
            <n-button quaternary size="tiny" @click="() => { handleDelete(item.id) }">
              <template #icon>
                <n-icon>
                  <TrashIcon />
                </n-icon>
              </template>
            </n-button>
          </n-flex>
        </n-flex>
      </n-card>
    </n-layout-sider>
    <n-layout-content>
      <n-card :bordered="false" size="small" :title="formTitle" style="height: 100%;">
        <n-form ref="formRef" size="small" label-placement="left" label-width="90" label-align="right"
          :model="formValue" :rules="rules">
          <n-form-item label="连接名称" path="name">
            <n-input v-model:value="formValue.name" placeholder="输入连接名称" />
          </n-form-item>
          <n-form-item label="主机地址" path="host">
            <n-input v-model:value="formValue.host" placeholder="输入主机地址" />
          </n-form-item>
          <n-form-item label="端口号" path="port">
            <n-input-number v-model:value="formValue.port" placeholder="输入端口号" />
          </n-form-item>
          <n-form-item label="启用认证" path="enable_auth">
            <n-switch v-model:value="formValue.enable_auth" />
          </n-form-item>
          <n-form-item v-if="formValue.enable_auth" label="用户名" path="username">
            <n-input v-model:value="formValue.username" placeholder="输入用户名" />
          </n-form-item>
          <n-form-item v-if="formValue.enable_auth" label="密码" path="password">
            <n-input v-model:value="formValue.password" type="password" placeholder="输入密码" />
          </n-form-item>
          <n-form-item label="启用TLS" path="enable_tls">
            <n-switch v-model:value="formValue.enable_tls" />
          </n-form-item>
          <n-form-item v-if="formValue.enable_tls" label="CA证书" path="ca">
            <n-upload ref="caFile" action="" :max="1" :multiple="false" :default-upload="false" @change="uploadCAFile">
              <n-button>选择文件</n-button>
            </n-upload>
          </n-form-item>
          <n-form-item v-if="formValue.enable_tls" label="客户端证书" path="cert">
            <n-upload ref="certFile" action="" :max="1" :multiple="false" :default-upload="false"
              @change="uploadCertFile">
              <n-button>选择文件</n-button>
            </n-upload>
          </n-form-item>
          <n-form-item v-if="formValue.enable_tls" label="客户端密钥" path="key">
            <n-upload ref="keyFile" action="" :max="1" :multiple="false" :default-upload="false"
              @change="uploadKeyFile">
              <n-button>选择文件</n-button>
            </n-upload>
          </n-form-item>
        </n-form>
        <template #footer>
          <n-flex justify="end" style="width: 100%;">
            <n-button size="small" type="tertiary" ghost attr-type="button" @click="" v-if="isEdit">
              删除
            </n-button>
            <n-button size="small" type="primary" ghost attr-type="button" @click="handleTest" :loading="testing">
              测试
            </n-button>
            <n-button size="small" type="primary" attr-type="button" @click="handleUpdate" v-if="isEdit"
              :loading="updating">
              更新
            </n-button>
            <n-button size="small" type="primary" attr-type="button" @click="handleSave" v-if="!isEdit"
              :loading="saving">
              保存
            </n-button>
          </n-flex>
        </template>
      </n-card>
    </n-layout-content>
  </n-layout>
</template>

<script lang="ts" setup>
import { NLayout, NLayoutSider, NLayoutContent } from 'naive-ui';
import { NButton, NIcon, NFlex, NCard, NForm, NFormItem, NInput, NInputNumber, NSwitch, NUpload } from 'naive-ui';
import { NInputGroup } from 'naive-ui';
import {
  Add as AddIcon,
  TrashBinOutline as TrashIcon,
} from '@vicons/ionicons5';
import type { FormInst, FormRules, UploadFileInfo } from 'naive-ui';
import { useMessage } from 'naive-ui'
import { ref } from 'vue';
import type { Connection } from './connection_model';
import { useConnection } from './use-connection';
import { computed } from 'vue';

const $message = useMessage();
const $connection = useConnection();

// region list
const list = ref<Connection[]>([]);
async function loadData() {
  const data = {
    page: 1,
    size: 10,
    total: 0,
    keyword: "",
    params: {},
  } as any;
  const r = await $connection.api.list(data);
  if (r.success) { list.value = r.data; }
}
loadData();
// endregion list


// region model
const model = ref<Connection>();
async function loadItem(id: string) {
  const r = await $connection.api.get(id);
  if (r.success) {
    model.value = r.data;
    formValue.value = r.data;
    isEdit.value = true;
  } else {
    $message.error(r.message);
  }
}

// endregion model

const isEdit = ref(false);
const formTitle = computed(() => { return isEdit.value ? "编辑连接" : "创建连接"; });
const formRef = ref<FormInst | null>(null);
const formValue = ref<Connection>({
  id: '',
  name: 'localhost',
  host: 'localhost',
  port: 2379,
  enable_auth: false,
  username: '',
  password: '',
  enable_tls: false,
  ca: '',
  cert: '',
  key: ''
});
const rules = ref<FormRules>({
  id: [{ required: false, message: '请输入连接ID', trigger: 'blur' }],
  name: [{ required: true, message: '请输入连接名称', trigger: 'blur' }],
  host: [{ required: true, message: '请输入主机地址', trigger: 'blur' }],
  port: [{ required: true, message: '请输入端口号', trigger: 'blur', type: 'number', min: 1, max: 65535 }],
  username: [{ required: false, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: false, message: '请输入密码', trigger: 'blur' }],
  ca: [{ required: false, message: '请输入CA证书', trigger: 'blur' }],
  cert: [{ required: false, message: '请输入客户端证书', trigger: 'blur' }],
  key: [{ required: false, message: '请输入客户端密钥', trigger: 'blur' }],
});

function uploadCAFile(options: { file: UploadFileInfo }) {
  if (options.file) {
    // 读取文件内容
    const reader = new FileReader();
    reader.readAsText(options.file.file as File);
    reader.onload = () => {
      formValue.value.ca = reader.result as string;
    };
  }
}

function uploadCertFile(options: { file: UploadFileInfo }) {
  if (options.file) {
    // 读取文件内容
    const reader = new FileReader();
    reader.readAsText(options.file.file as File);
    reader.onload = () => {
      formValue.value.cert = reader.result as string;
    };
  }
}

function uploadKeyFile(options: { file: UploadFileInfo }) {
  if (options.file) {
    // 读取文件内容
    const reader = new FileReader();
    reader.readAsText(options.file.file as File);
    reader.onload = () => {
      formValue.value.key = reader.result as string;
    };
  }
}

function handleValidateClick(e: MouseEvent) {
  e.preventDefault()
  formRef.value?.validate((errors) => {
    if (!errors) {
    }
    else {
      console.log(errors)
    }
  })
}

function handleAdd(e: MouseEvent) {
  e.preventDefault();
  isEdit.value = false;
  formValue.value = {
    id: '',
    name: 'localhost',
    host: 'localhost',
    port: 2379,
    enable_auth: false,
    username: '',
    password: '',
    enable_tls: false,
    ca: '',
    cert: '',
    key: ''
  };
}

const testing = ref(false);
async function handleTest(e: MouseEvent) {
  e.preventDefault();
  const vr = await formRef.value?.validate();
  console.log(vr);
  if (vr?.warnings) { return; }
  testing.value = true;
  const r = await $connection.api.test(formValue.value);
  console.log(r);
  if (r.success) {
    $message.success('连接成功');
  } else {
    $message.error(r.message);
  }
  testing.value = false;
}

const updating = ref(false);
async function handleUpdate(e: MouseEvent) {
  e.preventDefault()
  const vr = await formRef.value?.validate();
  if (vr?.warnings) { return; }
  updating.value = true;
  const r = await $connection.api.update(formValue.value);
  if (r.success) {
    $message.success('修改成功');
    loadData();
  } else {
    $message.error(r.message);
  }
  updating.value = false;
}

const saving = ref(false);
async function handleSave(e: MouseEvent) {
  e.preventDefault()
  const vr = await formRef.value?.validate();
  if (vr?.warnings) { return; }
  saving.value = true;
  const r = await $connection.api.create(formValue.value);
  if (r.success) {
    $message.success(isEdit.value ? '修改成功' : '保存成功');
    loadData();
  } else {
    $message.error(r.message);
  }
  saving.value = false;
}

async function handleDelete(id: string) {
  await $connection.api.delete(id);
  loadData();
}

</script>

<style scoped>
.connction-line{
  background:rgba(250, 250, 252, 1);
  padding: 4px 10px;
  border-radius: 4px;
  border: 1px solid #e8e8e8;
  transition: all 0.3s;
}

.connction-line:hover{
  background:rgba(240, 240, 242, 1);
}
</style>

<template>
  <n-layout has-sider position="absolute">
    <n-layout-sider bordered width="240">
      <n-card :bordered="false" size="small" title="Data Manager">
        <n-select v-if="showSelect" v-model:value="connctionId" :options="options" size="small"
          style="margin-bottom: 8px;" @update:value="loadTree" />
        <!-- <n-input-group style="margin-bottom: 20px;">
          <n-input size="small" />
          <n-button size="small">
            搜索
          </n-button>
        </n-input-group> -->
        <n-flex v-else vertical justify="center" style="margin: 8px 0 24px;">
          <n-button size="small" @click="createConneion">
            Create Connection
          </n-button>
        </n-flex>

        <n-flex vertical justify="center">
          <n-tree block-line :data="tree"
            :override-default-node-click-behavior="clickItem as TreeOverrideNodeClickBehavior" />
        </n-flex>
      </n-card>
    </n-layout-sider>
    <n-layout-content>
      <n-card :bordered="false" size="small" style="height: 100%;">
        <n-data-table size="small" :columns="columns" :data="tabel" :pagination="pagination" :bordered="false" />
      </n-card>

      <!--  -->
      <n-modal v-model:show="showModal" preset="card" style="width: 60%;">
        <n-form size="small" label-placement="left" label-width="auto">
          <n-form-item label="Key">
            <n-input v-model:value="currentData.key" />
          </n-form-item>
          <n-form-item label="Value">
            <n-input type="textarea" :autosize="{ minRows: 20, maxRows: 20 }" v-model:value="formattedValue" />
          </n-form-item>
        </n-form>
        <template #footer>
          <n-flex justify="end">
            <n-button type="tertiary" @click="showModal = false">
              Cancel
            </n-button>
            <n-button type="primary" @click="saveData">
              Save
            </n-button>
          </n-flex>
        </template>
      </n-modal>

      <!--  -->
      <n-modal v-model:show="showDeleteModal" preset="card" style="width: 60%;">
        <n-form size="small" label-placement="left" label-width="auto">
          <n-form-item label="Key">
            <n-input v-model:value="currentData.key" :disabled="true" />
          </n-form-item>
          <n-form-item label="Value">
            <n-input type="textarea" :autosize="{ minRows: 20, maxRows: 20 }" v-model:value="formattedValue"
              :disabled="true" />
          </n-form-item>
        </n-form>
        <template #footer>
          <n-flex justify="end">
            <n-button type="tertiary" @click="showDeleteModal = false">
              Cancel
            </n-button>
            <n-button type="error" @click="deleteData">
              Delete
            </n-button>
          </n-flex>
        </template>
      </n-modal>
    </n-layout-content>
  </n-layout>
</template>

<script lang="ts" setup>
import { NLayout, NLayoutSider, NLayoutContent, NModal, NForm, NFormItem, NInput, c } from 'naive-ui';
import { NSelect, NFlex, NCard, NTree, NDataTable, NButton } from 'naive-ui';
import type { TreeOption, TreeOverrideNodeClickBehavior } from 'naive-ui';
import { ref, h, computed } from 'vue';
import { Connection } from '../connection/connection_model';
import { useConnection } from '../connection/use-connection';
import { useData } from './use-data';
import router from '../../router';

const $connection = useConnection();
const $data = useData();

// region list
const showSelect = ref(false);
const connctionId = ref("");
const options = ref<{ label: string, value: string }[]>();
async function loadData() {
  options.value = undefined;
  const req = {
    page: 1,
    size: 10,
    total: 0,
    keyword: '',
    params: {},
  }
  const r = await $connection.api.list(req);
  if (r.success) {
    console.log(r.data);
    options.value = r.data.map((item: Connection) => {
      return {
        label: item.name,
        value: item.id,
      };
    });
    if (options.value && options.value.length > 0) {
      connctionId.value = options.value[0].value;
      showSelect.value = true;
      loadTree();
    }
  }
}
loadData();
// endregion list

// region tree
const tree = ref<TreeOption[]>();
async function loadTree() {
  const req = {
    id: connctionId.value,
    keyword: '',
  }
  console.log(req);
  const r = await $data.api.getKeyTree(req);
  if (r.success) {
    console.log(r.data);
    tree.value = r.data;
  }
}
loadTree();

const clickItem = (info: { option: TreeOption }) => {
  console.log(info);
  const prefix = info.option.key;
  loadTabel(prefix as string);
}
// endregion tree

// region page
const columns = [
  {
    title: 'Key',
    key: 'key',
  },
  {
    title: 'Value',
    key: 'value',
    ellipsis: true,
  },
  {
    title: 'Action',
    key: 'actions',
    width: 120,
    render(row: any) {
      return [
        h(NButton,
          {
            strong: true,
            tertiary: true,
            size: 'small',
            onClick: () => {
              console.log(row);
              currentData.value = row;
              showModal.value = true;
            }
          },
          { default: () => 'Edit' }
        ),
        h(NButton,
          {
            strong: true,
            tertiary: true,
            size: 'small',
            style: {
              marginLeft: '8px',
            },
            onClick: () => {
              console.log(row);
              currentData.value = row;
              deleteKey.value = row.key;
              showDeleteModal.value = true;
            }
          },
          { default: () => 'Del' }
        ),
      ]
    }
  }
];
const tabel = ref();
const pagination = ref({
  page: 1,
  pageSize: 10,
  total: 0,
  pageSizes: [10, 20, 50, 100],
  showSizePicker: true,
  showQuickJumper: true,
})
async function loadTabel(prefix: string) {
  const req = {
    page: pagination.value.page,
    size: pagination.value.pageSize,
    total: 0,
    keyword: "",
    params: {
      id: connctionId.value,
      prefix: prefix,
    }
  };
  console.log(req);
  const r = await $data.api.listData(req);
  if (r.success) {
    console.log(r.data);
    const { data, total } = r.data;
    tabel.value = data;
    pagination.value.total = total;
  }
}
// endregion page

// region action
function createConneion() {
  router.push('/connection');
}
// endregion action

// region edit
const currentData = ref({
  key: '',
  value: '',
});
const showModal = ref(false);

// 格式化JSON显示
const formattedValue = computed({
  get() {
    try {
      const parsed = JSON.parse(currentData.value.value);
      return JSON.stringify(parsed, null, 2);
    } catch {
      return currentData.value.value;
    }
  },
  set(val: string) {
    currentData.value.value = val;
  }
});

// 保存时尝试压缩JSON
const saveData = async () => {
  try {
    const parsed = JSON.parse(formattedValue.value);
    currentData.value.value = JSON.stringify(parsed);
  } catch {
    // 如果不是有效的JSON，保持原值
  }
  const req = {
    id: connctionId.value,
    key: currentData.value.key,
    value: currentData.value.value,
  };
  const r = await $data.api.putData(req);
  if (r.success) {
    loadTabel(currentData.value.key);
    // todo tip
  }

  showModal.value = false;
}
// endregion edit

// region delete
const showDeleteModal = ref(false);
const deleteKey = ref('');
const deleteData = async () => {
  const req = {
    id: connctionId.value,
    key: deleteKey.value,
  };
  const r = await $data.api.deleteData(req);
  if (r.success) {
    loadTabel(currentData.value.key);
    showDeleteModal.value = false;
    // todo tip 
  }
}
// endregion delete

</script>

<style scoped></style>

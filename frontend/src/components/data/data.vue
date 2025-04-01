<template>
  <n-layout has-sider position="absolute">
    <n-layout-sider bordered width="240">
      <n-card :bordered="false" size="small" title="Data Manager">
        <n-select v-model:value="connctionId" :options="options" size="small" style="margin-bottom: 8px;"
          @update:value="loadTree" />
        <!-- <n-input-group style="margin-bottom: 20px;">
          <n-input size="small" />
          <n-button size="small">
            搜索
          </n-button>
        </n-input-group> -->

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
    </n-layout-content>
  </n-layout>
</template>

<script lang="ts" setup>
import { NLayout, NLayoutSider, NLayoutContent } from 'naive-ui';
import { NSelect, NFlex, NCard, NTree, NDataTable, NButton } from 'naive-ui';
import type { TreeOption, TreeOverrideNodeClickBehavior } from 'naive-ui';
import { ref, h } from 'vue';
import { Connection } from '../connection/connection_model';
import { useConnection } from '../connection/use-connection';
import { KeyItem } from './data_model';
import { useData } from './use-data';


const $connection = useConnection();
const $data = useData();

// region list
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
            }
          },
          { default: () => 'Edit' }
        ),
        h(NButton,
          {
            strong: true,
            tertiary: true,
            size:'small',
            style: {
              marginLeft: '8px',
            },
            onClick: () => {
              console.log(row);
            }
          },
          { default: () => 'Del' }
        ),
      ]
    }
  }

  // {
  //   title: 'CreRevision',
  //   key: 'create_revision',
  //   width: 90,
  // },
  // {
  //   title: 'ModRevision',
  //   key: 'mod_revision', 
  //   width: 90,
  // },
  // {
  //   title: 'Version',
  //   key: 'version',
  //   width: 90,
  // },
  // {
  //   title: 'Lease',
  //   key: 'lease',
  //   width: 90,
  // }
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

</script>

<style scoped></style>
